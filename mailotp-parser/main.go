package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/dungvan/mailstation/common"
	"github.com/dungvan/mailstation/common/db"
	"github.com/dungvan/mailstation/common/memcache"
	"github.com/dungvan/mailstation/common/pubsub"
	"github.com/dungvan/mailstation/mailotp-parser/otpservice"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	redisURL     = "redis://localhost:6379/0"
	pubsubClient *pubsub.Client
	publisher    eventPublisher
	subscriber   eventSubscriber
	cacheClient  memcache.Client
	dbClient     *gorm.DB
)

func init() {
	if os.Getenv("REDIS_URL") != "" {
		redisURL = os.Getenv("REDIS_URL")
	}
	redisOpts, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf(`failed connect to redis: %v
		make sure you have set the redis connection environment variables
			REDIS_URL	Optional	default 'redis://localhost:6379/0'
		`, err)
	}
	rdb := redis.NewClient(redisOpts)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf(`failed connect to redis: %v
		make sure you have set the redis connection environment variables
			REDIS_URL	Optional	default 'redis://localhost:6379/0'
		`, err)
	}
	pubsubClient = pubsub.New(rdb)
	publisher = pubsubClient
	subscriber = pubsubClient

	cacheClient = memcache.New(rdb)
	dbClient, err = db.New("", false)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	otpservice.Init(dbClient, cacheClient)
}

type eventPublisher interface {
	Publish(ctx context.Context, topic common.PubSubTopic, message string) error
}

type eventSubscriber interface {
	Subscribe(ctx context.Context, topic common.PubSubTopic) <-chan *pubsub.Message
}

func main() {
	defer pubsubClient.Close()

	// Create a channel to handle parse mail events -> publish otp event if parse succeeded or drop the event.
	newParsedEvents := make(chan *common.MailOTP)
	// parse mail otp events

	incomingMsgCh := subscriber.Subscribe(context.Background(), common.INCOMING_EMAIL_TOPIC)

	for {
		select {
		case msg := <-incomingMsgCh:
			go parseMailOTP(msg, newParsedEvents)
		case email := <-newParsedEvents:
			go publishNewOPTEvent(email)
		}
	}
}

func parseMailOTP(msg *pubsub.Message, mailEvent chan<- *common.MailOTP) {
	var email *common.IncomingEmail
	if err := json.Unmarshal([]byte(msg.Payload), &email); err != nil {
		log.Println("failed to unmarshal email message:", err)
		return
	}

	// cache hit mail sender -> service_name, subject_regexp, content_regexp, start_delimiter, end_delimiter
	// or cache miss -> SELECT service_name, subject_regexp, content_regexp, start_delimiter, end_delimiter FROM otp_services WHERE sender = ${email.From}
	// caching service_name, subject_regexp, content_regexp, start_delimiter, end_delimiter by mail sender

	templates, err := otpservice.GetTemplateBySender(email.From)
	if err != nil {
		log.Println("templates could not be found for sender", email.From, ":", err)
		return
	}

	for _, template := range templates {
		if template.SubjectRegexp.MatchString(email.Subject) {
			otp := template.ExtractOTP(email.TextBody)
			if otp != "" {
				mailEvent <- &common.MailOTP{
					Email:   email.To[0],
					OTP:     otp,
					Service: template.ServiceName,
				}
				return
			}
		}
	}
	log.Println("no OTP template matches for email", email.Subject, "from", email.From)
}

func publishNewOPTEvent(email *common.MailOTP) {
	emailBytes, err := json.Marshal(email)
	if err != nil {
		log.Println(err)
		return
	}

	if err := publisher.Publish(context.Background(), common.NEW_OTP_TOPIC, string(emailBytes)); err != nil {
		log.Println(err)
		return
	}
}
