package insert_local_redis

import (
	"domains"
	"encoding/json"
	//	"fmt"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	//	"time"
)

func InsertCharacter(golog syslog.Writer, c redis.Conn, key string, character domains.CharacterRedis) {

	update_insert := false

	field := character.Id

	if exist, err := redis.Int(c.Do("HEXISTS", key, field)); err != nil {

		golog.Crit(err.Error())

	} else {

		if exist == 0 {

			update_insert = true

		} else {

			if bcharacter, err := redis.Bytes(c.Do("HGET", key, field)); err != nil {

				golog.Crit(err.Error())

			} else {

				var characterold domains.CharacterRedis

				err := json.Unmarshal(bcharacter, &characterold)
				if err != nil {

					golog.Crit(err.Error())

				}

				character.Name = characterold.Name
				character.Age = characterold.Age
				character.Phone = characterold.Phone
				character.City = characterold.City

			}

		}

	}

	if update_insert == true {

		bcharacter, _ := json.Marshal(character)

		if _, err := c.Do("HSET", key, field, bcharacter); err != nil {

			golog.Crit(err.Error())

		}

	}

}
