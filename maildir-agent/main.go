package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/mail"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/dungvan/mailstation/common"
	"github.com/dungvan/mailstation/common/pubsub"
	"github.com/jhillyerd/enmime"
	"github.com/redis/go-redis/v9"
)

var (
	// regexStr is a regular expression to match the maildir file path format <maildomain>/<mailuser>/new/<mailfile>
	regexStr     = regexp.MustCompile(`^[^/]*/[^/]*/new/\d{10}\.[a-zA-Z0-9]+\.(.+),[A-Z]=\d+,[A-Z]=\d+:\d,[A-Z]$`)
	maildirPath  = "."
	redisURL     = "redis://localhost:6379/0"
	pubsubClient *pubsub.Client
	publisher    eventPublisher
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
}

type eventPublisher interface {
	Publish(ctx context.Context, topic common.PubSubTopic, message string) error
}

func main() {
	defer pubsubClient.Close()
	// Specify the path to the maildir directory

	flag.StringVar(&maildirPath, "maildir", maildirPath, "Path to the maildir directory")
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %v [OPTION]", os.Args[0])
		flag.PrintDefaults()
	}

	fileInfo, err := os.Stat(maildirPath)
	if err != nil || !fileInfo.IsDir() {
		flag.Usage()
		log.Fatal("maildir must be a directory")
	}

	// Create a channel to handle new file events -> read mailfile
	newFileEvents := make(chan string)
	// Create a channel to handle read mailfile events -> publish mail event
	readMailEvents := make(chan *common.IncomingEmail)
	// Create a channel to handle published mail events -> remove mail file
	publishedMailEvents := make(chan string)

	// Start watching the maildir directory for file events
	ticker := time.NewTicker(time.Second * 10)
	lastTime := time.Time{}

	log.Printf("Start Watching maildir: %s\n", maildirPath)

	// Process file events
	for {
		select {
		case tickTime := <-ticker.C:
			go watchMaildir(maildirPath, lastTime, newFileEvents)
			lastTime = tickTime
		case path := <-newFileEvents:
			go readMailFile(path, readMailEvents)
		case email := <-readMailEvents:
			go publishMailEvent(email, publishedMailEvents)
		case path := <-publishedMailEvents:
			go removeMailFile(path)
		}
	}
}

func watchMaildir(path string, lastTime time.Time, fileEvents chan<- string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a regular file and not a directory.
		// Check if file path are in correct format.
		// Check if the file was created within the last minute.
		if info.Mode().IsRegular() && regexStr.MatchString(path) && info.ModTime().After(lastTime) {
			fileEvents <- path
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func readMailFile(path string, emailEvents chan<- *common.IncomingEmail) {
	// Open the mail file
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// Read the contents of the mail file and process it
	envelop, err := enmime.ReadEnvelope(file)
	if err != nil {
		log.Println(err)
		return
	}
	from, _ := envelop.AddressList("From")
	toList, _ := envelop.AddressList("To")
	subject := envelop.GetHeader("Subject")
	textBody := envelop.Text

	emailEvents <- &common.IncomingEmail{
		From: from[0].Address,
		To: func(list []*mail.Address) []string {
			var str []string
			for _, a := range list {
				str = append(str, a.Address)
			}
			return str
		}(toList),
		Subject:      subject,
		TextBody:     textBody,
		MailFilePath: path,
	}
}

func publishMailEvent(email *common.IncomingEmail, emailEvents chan<- string) {
	emailBytes, err := json.Marshal(email)
	if err != nil {
		log.Println(err)
		return
	}

	if err := publisher.Publish(context.Background(), common.INCOMING_EMAIL_TOPIC, string(emailBytes)); err != nil {
		log.Println(err)
		return
	}
	log.Println("published incoming email event:", email.MailFilePath)

	emailEvents <- email.MailFilePath
}

func removeMailFile(path string) {
	if err := os.Remove(path); err != nil {
		log.Println(err)
	}
	log.Println("removed mailfile:", path)
}
