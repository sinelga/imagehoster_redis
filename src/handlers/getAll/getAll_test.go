package getAll

import (
    "testing"
    "github.com/garyburd/redigo/redis"
    "startones"
    
)

func TestGetAll(t *testing.T) {
	
	golog, _ := startones.Start()
	
		rds, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}
	defer rds.Close()
	
	GetAll(golog, rds,"www.test.com")
	

}

