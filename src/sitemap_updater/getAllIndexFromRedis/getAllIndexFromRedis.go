package getAllIndexFromRedis

import (
	"github.com/garyburd/redigo/redis"
	//	"encoding/json"
	"log/syslog"
)

func GetAllIndexForSite(golog syslog.Writer, c redis.Conn, site string) []string {

	if keywordindexs, err := redis.Strings(c.Do("HKEYS", site)); err != nil {

		golog.Crit("keywordindexs " + err.Error())

	} else {

		return keywordindexs

	}

	return nil

}
