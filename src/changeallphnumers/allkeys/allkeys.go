package allkeys

import (
//	"fmt"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
)

func GetAllKeys(golog syslog.Writer, c redis.Conn)[]string {

	var retallkeys []string

	if allkeys, err := redis.Strings(c.Do("KEYS", "*")); err != nil {

		golog.Crit(err.Error())

	} else {

		retallkeys = allkeys
	}

return retallkeys

}
