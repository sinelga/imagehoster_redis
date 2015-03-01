package getAll

import (
	//	"database/sql"
	"domains"
	"fmt"
	"github.com/garyburd/redigo/redis"
	//	_ "github.com/go-sql-driver/mysql"
	"log/syslog"

	"encoding/json"
	"math/rand"
	"time"
	//		"strings"
)

func shuffleSlice(slice []interface{}) []interface{} {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func GetAll(golog syslog.Writer, c redis.Conn, site string) []domains.CharacterRedis {

	var charactersRedis []domains.CharacterRedis

	if allfilds, err := redis.Strings(c.Do("HKEYS", site)); err != nil {

		golog.Crit(err.Error())

	} else {

		//				set := make(map[int]string)
		list := make([]interface{}, len(allfilds))

		for i, fild := range allfilds {

			list[i] = fild

		}

		shuffleSlice(list)

		var args []interface{}

		args = append(args, site)

		for i := 0; i < 10; i++ {

			args = append(args, list[i].(string))

		}

		if bcharactersRedis, err := redis.MultiBulk(c.Do("HMGET", args...)); err != nil {

			golog.Crit(err.Error())

		} else {

			fmt.Println(len(bcharactersRedis))

			for _, x := range bcharactersRedis {
				var v, ok = x.([]byte)
				if ok {
					fmt.Println(string(v))

					var character domains.CharacterRedis

					if err := json.Unmarshal(v, &character); err != nil {
						golog.Crit(err.Error())
					} else {
						
						charactersRedis =append(charactersRedis,character)
						
					}

				}
			}

		}

	}

	return charactersRedis

}
