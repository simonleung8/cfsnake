package redis

import (
	"encoding/json"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/simonleung8/cfsnake/parser"
	. "github.com/simonleung8/cfsnake/server"
)

const (
	REDIS_CHANNEL = "CFSNAKE"
)

type Redis struct {
	err  error
	Conn redis.Conn
}

func NewRedis(services parser.Services) (*Redis, error) {
	//Parse the service configurations
	services := parser.Services{}

	if os.Getenv("VCAP_SERVICES") != "" {
		//setup our redis connections using redis
		services.Parse(os.Getenv("VCAP_SERVICES"))
	} else {
		return nil, error.New("No VCAP_SERVICES")
	}

	var red Redis

	red.Conn, red.err = redis.Dial("tcp", services.Redis.Hostname+":"+services.Redis.Port)
	if red.err != nil {
		return nil, err
	}

	if services.Redis.Password != "" {
		red.Conn.Do("AUTH", services.Redis.Password)
	}

	red.Subscribe(REDIS_CHANNEL)

	return &red, nil
}

/* function that should constantly be polling the redis service */
func (red *Redis) Read(data chan Player) {
	var person Person

	for {
		switch v := sub.Receive().(type) {
		case redis.Message:
			json.Unmarshal([]byte(v.Data), &person)
			Player <- person
		}
	}
}

func (red *Redis) Push(person Person) error {
	data, err := json.Marshal(&Person)
	if err != nil {
		return err
	}

	red.Do("PUBLISH", REDIS_CHANNEL, data)

	return nil
}
