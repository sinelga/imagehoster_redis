package handlers

import (
	"github.com/garyburd/redigo/redis"
	"github.com/zenazn/goji/web"
	"net/http"
	//	"domains"

	"handlers/getOne"
	"net"
	"startones"
	"strconv"
	"strings"
	//	"sync"
	//	"log/syslog"
)

func Elaborate(c web.C, w http.ResponseWriter, r *http.Request) {

	startOnce.Do(func() {
		golog, config = startones.Start()

	})
	
	golog.Info("UserAgent " + r.UserAgent() + " Host " + r.Host + " RequestURI " + r.RequestURI + " r.RemoteAddr " + r.RemoteAddr + " referer " + r.Referer())

	site, _, _ := net.SplitHostPort(r.Host)
	
	golog.Info("Elaborate other ->site "+site+" host "+r.Host)

	if site == "localhost" {

		site = "www.test.com"
	}

	if strings.HasPrefix(site, "192.168.") {

		site = "www.test.com"
	}

	if strings.HasPrefix(site, "127.0.0.1") {

		site = "www.test.com"
	}

	rds, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}
	defer rds.Close()

	path := r.URL.Path

	id_arr := strings.Split(path, "/")

	if len(id_arr) > 0 {

		if _, err := strconv.Atoi(id_arr[1]); err == nil {

			_, exist := getOne.GetById(golog, rds, site, id_arr[1])

			if exist {

				http.ServeFile(w, r, "/home/juno/git/8_fi_FIporno_desk/dist/index.html")

			}

		}

	}

	http.ServeFile(w, r, "/home/juno/git/8_fi_FIporno_desk/dist/404.html")

}
