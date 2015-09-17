package handlers

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/zenazn/goji/web"
	"handlers/getAll"
	"handlers/getOne"
	"log/syslog"
	"net"
	"net/http"
	"startones"
	"strings"
	"sync"
)

var startOnce sync.Once
var golog syslog.Writer

var config domains.Config

func MhandleAll(c web.C, w http.ResponseWriter, r *http.Request) {

	startOnce.Do(func() {
		golog, config = startones.Start()

	})

	//	var variant string
	//	var user_agent string

	//	for k, v := range r.Header {
	//
	//		golog.Info("key: " + k + " value: " + v[0])
	//
	//
	//		if k == "X-Variant" {
	//			variant = r.Header["X-Variant"][0]
	//
	//		}
	//		if k == "User-Agent" {
	//			user_agent = r.Header["User-Agent"][0]
	//
	//		}
	//	}

	//	golog.Info("User-Agent ->"+user_agent+" variant->" +variant)

	//	golog.Info("UserAgent " + r.UserAgent() + " Host " + r.Host + " RequestURI " + r.RequestURI + " r.RemoteAddr " + r.RemoteAddr + " referer " + r.Referer())

	site, _, _ := net.SplitHostPort(r.Host)

	if site == "localhost" {

		site = "www.test.com"
	}

	if strings.HasPrefix(site, "192.168.") {

		site = "www.test.com"
	}

	rds, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}
	defer rds.Close()

	var bytes []byte
	var e error

	id := c.URLParams["id"]

	//	if strings.HasPrefix(user_agent,"msnbot") {
	//
	//		golog.Info("msnbot!!!")
	//
	//
	//	}

	if id == "" {

		characters, exist := getAll.GetAll(golog, rds, site)

		if !exist {

			http.NotFound(w, r)
		} else {

			bytes, e = json.Marshal(characters)
			if e != nil {

				golog.Err(e.Error())

			}

		}

	} else {

		character, _ := getOne.GetById(golog, rds, site, id)

		bytes, e = json.Marshal(character)
		if e != nil {

			golog.Err(e.Error())

		}

	}

	w.Write(bytes)

}
