package find_all_keywords

import (
	"github.com/garyburd/redigo/redis"
	"log/syslog"
//	"fmt"
)

func FindAllKeywords(golog syslog.Writer, c redis.Conn, locale string, themes string) map[int]string {
	
	
	queuename := locale + ":" + themes+ ":keywords"
	
	mapkeywords := make(map[int]string)

	if keywords, err := redis.Strings(c.Do("ZRANGEBYSCORE", queuename,"-inf","+inf")); err != nil {

		golog.Crit("err "+err.Error())

	} else {
		
	for i, keyword := range keywords {

		mapkeywords[i] = keyword 
				
	}
		
//		return 	keywords	
	}

	return mapkeywords
}
