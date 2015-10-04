package elaboratekey

import (
	"domains"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
)

func ScanKey(golog syslog.Writer, c redis.Conn, key string, fromphnumer string, tophnumber string) {

	if allfields, err := redis.Strings(c.Do("HKEYS", key)); err != nil {

		golog.Crit(err.Error())

	} else {

		for _, field := range allfields {

			if val, err := redis.Bytes(c.Do("HGET", key, field)); err != nil {
				golog.Crit(err.Error())

			} else {

				var characterRedis domains.CharacterRedis

				if err := json.Unmarshal(val, &characterRedis); err != nil {

					golog.Crit(err.Error())

				} else {

					if characterRedis.Phone == fromphnumer {

						fmt.Println(characterRedis.Phone)

						characterRedis.Phone = tophnumber

						bcharacter, _ := json.Marshal(characterRedis)

						if _, err := c.Do("HSET", key, field, bcharacter); err != nil {

							golog.Crit(err.Error())

						}

					}

				}

			}

		}

	}


}
