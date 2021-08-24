package rediskeyspacenotif

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func KeyspaceNotif() {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDISADDR"),
		DB:   0, // index DB
	})
	// set config keyspace notif -> expire
	var ctx = context.Background()
	var statusCmd = rdb.ConfigSet(ctx, "notify-keyspace-events", "KEAx")
	log.Println(statusCmd)

	// test connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println(err)
	}
	log.Println(pong)

	// example set expire
	var status = rdb.SetEX(ctx, "test", "value", 1*time.Minute)
	log.Println(status)

	// 0 -> index DB
	pubSub := rdb.PSubscribe(ctx, "__keyevent@0__:expired")
	log.Printf("pubSub:%+v\n", pubSub)
	for {
		msgi, err := pubSub.Receive(ctx)
		if err != nil {
			panic(err)
		}
		switch msg := msgi.(type) {
		case *redis.Message:
			log.Printf("Message: %s %s\n", msg.Channel, msg.Payload) //msg.Payload == id in redis
			continue
		case *redis.Subscription:
			log.Printf("Subscription: %s %s %d\n", msg.Kind, msg.Channel, msg.Count)
			if msg.Count == 0 {
				continue
			}
			continue
		case error:
			log.Printf("error: %v\n", msg)
			continue
		}
		log.Println("Go routine exit")
	}
}
