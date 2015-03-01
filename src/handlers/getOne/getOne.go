package getOne

import (
	//	"database/sql"
	"domains"
	"github.com/garyburd/redigo/redis"
	//	"fmt"
	//	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"log/syslog"
)

func GetById(golog syslog.Writer, c redis.Conn, site string, id string) domains.Character {

	var character domains.Character

	if bcharacterRedis, err := redis.Bytes(c.Do("HGET", site, id)); err != nil {

		golog.Crit(err.Error())

	} else {

		if err := json.Unmarshal(bcharacterRedis, &character); err != nil {
			golog.Crit(err.Error())
		}

	}

	return character

}
