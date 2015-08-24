package find_all_keywords

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"log/syslog"
	"testing"
)

func TestFindAllKeywords(t *testing.T) {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	redisprotocol := "tcp"
	redishost := ":6379"

	c, err := redis.Dial(redisprotocol, redishost)
	if err != nil {

		golog.Crit(err.Error())

	}
	defer c.Close()

	keywords := FindAllKeywords(*golog, c, "fi_FI", "porno")

	

	for i, keyword := range keywords {

		fmt.Println(i, keyword)

	}
	
	fmt.Println("len", len(keywords))

}
