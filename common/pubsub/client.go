package pubsub

import (
	"context"

	"github.com/dungvan/mailstation/common"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	*redis.Client
}

type Message = redis.Message

var subscribedTopics = make(map[common.PubSubTopic]*redis.PubSub)

func New(rdb *redis.Client) *Client {
	return &Client{rdb}
}

func (c *Client) Publish(ctx context.Context, topic common.PubSubTopic, message string) error {
	return c.Client.Publish(ctx, topic.String(), message).Err()
}

func (c *Client) Subscribe(ctx context.Context, topic common.PubSubTopic) <-chan *Message {
	ps := c.Client.Subscribe(ctx, topic.String())
	subscribedTopics[topic] = ps
	return ps.Channel()
}

func (c *Client) Unsubscribe(ctx context.Context, topic common.PubSubTopic) error {
	if ps, ok := subscribedTopics[topic]; ok {
		delete(subscribedTopics, topic)
		if err := ps.Unsubscribe(ctx, topic.String()); err != nil {
			return err
		}
	}
	return nil
}
