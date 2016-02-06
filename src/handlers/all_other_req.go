package handlers

import (
	"github.com/garyburd/redigo/redis"
	"github.com/zenazn/goji/web"
	"handlers/getAll"
	"handlers/getOne"
	"handlers/robots"
	"handlers/sitemap"
	"net/http"
	"notjsbots"
	"startones"
	"strconv"
	//		"fmt"
	"github.com/hiteshmodha/goDevice"
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

		notjsbot := false

		if strings.Contains(user_agent, "msnbot") || strings.Contains(user_agent, "bingbot") || strings.Contains(user_agent, "yahoo") {

			notjsbot = true
		}

		deviceType := goDevice.GetType(r)

		if deviceType == "Mobile" {

			golog.Info("Mobile")
		} else if deviceType == "Web" {
			golog.Info("Web")

		} else if deviceType == "Tab" {

			golog.Info("Tablet")
		}

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

			//						fmt.Println("id_arr", len(id_arr))

			golog.Info("id_arr " + strconv.Itoa(len(id_arr)))

			if len(id_arr) == 2 {

				golog.Info("path index.html?? " + path)
				if deviceType == "Mobile" {
					http.ServeFile(w, r, "/home/juno/git/fi_FI_mobile_react/version_00/dist/index.html")

				} else {

					http.ServeFile(w, r, "/home/juno/git/fi_FI_desk_mobile/version_23/dist/index.html")

				}

			} else if len(id_arr) > 2 {

				character, exist := getOne.GetById(golog, rds, site, id_arr[1])

				//				existbyId := true

				if exist {

					golog.Info(user_agent)

					if notjsbot {

						notjsbots.CreateNotJsPage(golog, c, w, r, variant, character, site)

					} else {

						if deviceType == "Mobile" {
							http.ServeFile(w, r, "/home/juno/git/fi_FI_mobile_react/version_00/dist/index.html")

						} else {

							http.ServeFile(w, r, "/home/juno/git/fi_FI_desk_mobile/version_23/dist/index.html")

						}

					}

				} else {

					http.NotFound(w, r)

				}

			} else {

				characters, exist := getAll.GetAll(golog, rds, site)

				if !exist {

					http.NotFound(w, r)
				} else {

					notjsbots.CreateNotJsPageIndex(golog, c, w, r, variant, characters, site)

				}

			}
			//						http.NotFound(w, r)

		}

	} else {

		//		golog.Err("variant NOT found!!!")

		if strings.HasPrefix(path, "/robots.txt") {

			robots.Generate(golog, w, r, site)

		}

	}

}
