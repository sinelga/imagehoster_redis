package handlers

import (
	"github.com/garyburd/redigo/redis"
	"github.com/zenazn/goji/web"
	"handlers/getOne"
	"handlers/robots"
	"handlers/sitemap"
	"net/http"
	"startones"
	"strconv"
	"strings"
)

func Elaborate(c web.C, w http.ResponseWriter, r *http.Request) {

	startOnce.Do(func() {
		golog, config = startones.Start()

	})

	var variant string
	for k, v := range r.Header {

		golog.Info("key: " + k + " value: " + v[0])
	}

	variant = r.Header["X-Variant"][0]

	if variant != "" {

		//		golog.Info("UserAgent " + r.UserAgent() + " Host " + r.Host + " RequestURI " + r.RequestURI + " r.RemoteAddr " + r.RemoteAddr + " referer " + r.Referer())

		site := r.Host

		golog.Info("Elaborate other ->site " + site + " host " + r.Host)

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

		if strings.HasPrefix(path, "/robots.txt") {

			robots.Generate(golog, w, site)

		} else if strings.HasPrefix(path, "/sitemap.xml") {

			sitemap.CheckGenerate(golog, w, site)

		} else {

			id_arr := strings.Split(path, "/")

			if len(id_arr) > 0 {

				if _, err := strconv.Atoi(id_arr[1]); err == nil {

					golog.Info("Only control on exist ")
					_, exist := getOne.GetById(golog, rds, site, id_arr[1])

					if exist {

						http.ServeFile(w, r, "/home/juno/git/fi_FI_desk_mobile/version_"+variant+"/dist/index.html")

					}

				}

			}

			http.NotFound(w, r)

		}

	} else {

		golog.Err("variant NOT found!!!")

	}

}
