package handlers

import (
	"github.com/garyburd/redigo/redis"
	"github.com/zenazn/goji/web"
	"handlers/getOne"
	"handlers/robots"
	"handlers/sitemap"
	"net/http"
	"notjsbots"
	"startones"
	//	"strconv"
	"strings"
)

func Elaborate(c web.C, w http.ResponseWriter, r *http.Request) {

	startOnce.Do(func() {
		golog, config = startones.Start()

	})

	var variant string
	var user_agent string
	//	var path string
	path := r.URL.Path
	site := r.Host

	golog.Info("all_other_req.Elaborate")

	for k, v := range r.Header {

		//		golog.Info("all_other_req key: " + k + " value: " + v[0])

		if k == "X-Variant" {
			//			variant = r.Header["X-Variant"][0]
			variant = v[0]

		}

		if k == "User-Agent" {
			user_agent = r.Header["User-Agent"][0]

		}

	}

	if variant != "" {

		//		golog.Info("UserAgent " + r.UserAgent() + " Host " + r.Host + " RequestURI " + r.RequestURI + " r.RemoteAddr " + r.RemoteAddr + " referer " + r.Referer())

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

			robots.Generate(golog, w, r, site)

		} else if strings.HasPrefix(path, "/sitemap.xml") {

			sitemap.CheckGenerate(golog, w, site)

		} else {

			id_arr := strings.Split(path, "/")

			if len(id_arr) > 0 {

				character, exist := getOne.GetById(golog, rds, site, id_arr[1])

				if exist {

					if strings.HasPrefix(user_agent, "msnbot") {

						notjsbots.CreateNotJsPage(golog, c, w, r, variant, character)

					} else {

						http.ServeFile(w, r, "/home/juno/git/fi_FI_desk_mobile/version_"+variant+"/dist/index.html")
						//						http.ServeFile(w, r, "/home/juno/git/"+variant+"_fi_FI_desk_mobile/dist/index.html")
					}

				}

			} else {

				if strings.HasPrefix(user_agent, "msnbot") {

					golog.Info("msnbot!!! not exist")

				} else {
					http.NotFound(w, r)
				}

			}
			//			http.NotFound(w, r)

		}

	} else {

		//		golog.Err("variant NOT found!!!")

		if strings.HasPrefix(path, "/robots.txt") {

			robots.Generate(golog, w, r, site)

		}

	}

}
