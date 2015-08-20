package find_all_keywords

import (
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	"fmt"
)

func FundAllKeywords(golog syslog.Writer, c redis.Conn, locale string, themes string) {
	
	
	queuename := locale + ":" + themes+ ":keywords"
	
	if keywords, err := redis.Values(c.Do("ZRANGEBYSCORE","-inf","+inf","withscores", queuename)); err != nil {

		golog.Crit(err.Error())

	} else {
		
		fmt.Println(len(keywords))
		
	}

}
