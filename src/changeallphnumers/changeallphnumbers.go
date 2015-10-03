package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"startones"
	"changeallphnumers/allkeys"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

var fromFlag = flag.String("from", "", "from must be 0700.. 0600..")
var fromTo = flag.String("to", "", "from must be 0700.. 0600..")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)

	}

	golog, _ := startones.Start()

	redisprotocol := "tcp"
	redishost := ":6379"

	c, err := redis.Dial(redisprotocol, redishost)
	if err != nil {

		golog.Crit(err.Error())

	}
	defer c.Close()
	
	allkeys := allkeys.GetAllKeys(golog,c)
	
	for _,key :=range allkeys{
		
		fmt.Println(key)
	}
	
	

}
