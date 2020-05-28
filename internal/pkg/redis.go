package pkg

import "github.com/go-redis/redis"

type RedisClient struct {
	client *redis.Client
}

func New(address string) RedisClient {
	c := redis.NewClient(&redis.Options{
		Addr: address,
	})
	_, err := c.Ping().Result()
	if err != nil {
		panic(err)
	}
	return RedisClient{
		client: c,
	}
}

func (r RedisClient) SetPair(key, value string) error {
	return r.client.Set(key, value, 0).Err()
}

func (r RedisClient) GetPair(key string) (string, error) {
	return r.client.Get(key).Result()
}
