package main

import (
	"domains"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"handlers/getOne"
	"io/ioutil"
	"log"
	"log/syslog"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"handlers/sitemap/createmapfile"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	rds, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())

	}
	defer rds.Close()

	files, _ := filepath.Glob("upload/img/*")

	set_all_id := make(map[int]struct{})

	for _, f := range files {

		if intid, err := strconv.Atoi(strings.Split(f, "/")[2]); err != nil {

			fmt.Println(err.Error())

		} else {

			set_all_id[intid] = struct{}{}

		}

	}

	fmt.Println("img num", len(files))

	mapfiles, _ := filepath.Glob("maps/*")

	for _, f := range mapfiles {

		fmt.Println(f)

		site := strings.Split(f, "_")[1]

		site = strings.TrimSuffix(site, ".xml")
		fmt.Println(site)

		xmlFile, err := os.Open(f)

		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer xmlFile.Close()

		XMLdata, _ := ioutil.ReadAll(xmlFile)

		var pages domains.Pages
		xml.Unmarshal(XMLdata, &pages)

		set_id_insitemap := make(map[int]struct{})

		for _, page := range pages.Pages {

			if intid, err := strconv.Atoi(strings.Split(page.Loc, "/")[3]); err != nil {

				fmt.Println(err.Error())

			} else {

				set_id_insitemap[intid] = struct{}{}

			}

		}

		fmt.Println("pages num", len(pages.Pages))

		set_free_id := make(map[int]struct{})

		for k := range set_all_id {

			if _, ok := set_id_insitemap[k]; !ok {

				set_free_id[k] = struct{}{}

			}

		}

		fmt.Println("freeid", len(set_free_id))

		var id_to_add []int

		if len(set_free_id) > 10 {

			i := 0

			for k := range set_free_id {

				i++

				id_to_add = append(id_to_add, k)

				if i >= 10 {
					break
				}

			}

			var Url *url.URL

			for _, k := range id_to_add {

				Url, err = url.Parse("http://" + site)
				if err != nil {
					golog.Crit(err.Error())
				}

				character, _ := getOne.GetById(*golog, rds, site, strconv.Itoa(k))

				permlink := strings.Split(character.Moto, " ")

				Url.Path += "/" + strconv.Itoa(character.Id) + "/" + permlink[0] + "-" + permlink[1] + ".html"

				fmt.Println(Url.String())
				doc := new(domains.Page)
				doc.Loc = Url.String()
				doc.Lastmod = time.Now().Format(time.RFC3339)
				doc.Changefreq = "weekly"

				pages.Pages = append(pages.Pages, doc)

			}

			resultXml, err := xml.MarshalIndent(pages, "", "  ")
			if err != nil {

				golog.Crit(err.Error())
			}
			
			createmapfile.Createmap(*golog,f,resultXml)

		}

	}

}
