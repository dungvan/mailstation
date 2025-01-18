package pubsub

import (
	"context"

	"github.com/dungvan/mailstation/common"
	"github.com/redis/go-redis/v9"
)

type PubSubClient struct {
	*redis.Client
}

type Message = redis.Message

var subscribedTopics = make(map[common.PubSubTopic]*redis.PubSub)

func NewPubSubClient(rdb *redis.Client) *PubSubClient {
	return &PubSubClient{rdb}
}

func (c *PubSubClient) Publish(ctx context.Context, topic common.PubSubTopic, message string) error {
	return c.Client.Publish(ctx, topic.String(), message).Err()
}

func (c *PubSubClient) Subscribe(ctx context.Context, topic common.PubSubTopic) <-chan *Message {
	ps := c.Client.Subscribe(ctx, topic.String())
	subscribedTopics[topic] = ps
	return ps.Channel()
}

func (c *PubSubClient) Unsubscribe(ctx context.Context, topic common.PubSubTopic) error {
	if ps, ok := subscribedTopics[topic]; ok {
		delete(subscribedTopics, topic)
		if err := ps.Unsubscribe(ctx, topic.String()); err != nil {
			return err
		}
	}
	return nil
}
