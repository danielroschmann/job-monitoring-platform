package redis

import (
	"context"
	"log"
	"os"

	goredis "github.com/redis/go-redis/v9"
)

var (
	Client *goredis.Client
	Ctx    = context.Background()
)

func Connect() {
	Client = goredis.NewClient(&goredis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})

	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Redis...")
}
