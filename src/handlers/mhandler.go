package handlers

import (
	"domains"
	"encoding/json"
	"github.com/zenazn/goji/web"
	"github.com/garyburd/redigo/redis"
	"handlers/getAll"
	"handlers/getOne"
	"log/syslog"
	"net/http"
	"startones"
	"sync"
)

var startOnce sync.Once
var golog syslog.Writer

var config domains.Config

func MhandleAll(c web.C, w http.ResponseWriter, r *http.Request) {

	startOnce.Do(func() {
		golog, config = startones.Start()

	})
	
	site := r.Host
	
	rds, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}
	defer rds.Close()

	golog.Info("id " + c.URLParams["id"])
	golog.Info(r.Method)
	id := c.URLParams["id"]

	var bytes []byte
	var e error

	if id == "" {

		characters := getAll.GetAll(golog, rds,site)

		bytes, e = json.Marshal(characters)
		if e != nil {

			golog.Err(e.Error())

		}

	} else {

		golog.Info("id " + id)

		character := getOne.GetById(golog,rds,site, id)

		bytes, e = json.Marshal(character)
		if e != nil {

			golog.Err(e.Error())

		}

	}

	w.Write(bytes)

}
