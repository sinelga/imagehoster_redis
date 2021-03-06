package sitemap

import (
	"bytes"
	"domains"
	"encoding/xml"
	"github.com/garyburd/redigo/redis"
	"handlers/sitemap/createmapfile"
	"io"
	"log/syslog"
//	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sitemap_maker/getLinks"
//	"strconv"
	"strings"
	"time"
	"fmt"
	"github.com/gosimple/slug"
)

func CheckGenerateSlug(golog syslog.Writer, w http.ResponseWriter, site string) {

	if _, err := os.Stat("maps"); os.IsNotExist(err) {

		os.Mkdir("maps", 0777)

	}

	filestr := "maps/sitemap_" + site + ".xml"

	if _, err := os.Stat(filestr); os.IsNotExist(err) {

		golog.Info(filestr + "  NOT exist")

		limit := random(600, 1000)

		c, err := redis.Dial("tcp", ":6379")
		if err != nil {

			golog.Crit(err.Error())

		}
		defer c.Close()

		characters := getLinks.GetAllLinks(golog, c, site)

		var Url *url.URL

		docList := new(domains.Pages)
		docList.XmlNS = "http://www.sitemaps.org/schemas/sitemap/0.9"
		
		for i := 0; i < limit; i++ {

			Url, err = url.Parse("http://" + site)
			if err != nil {
				golog.Crit(err.Error())
			}
			slug.MaxLength=20
			slugresalt := slug.Make(characters[i].Description)
			fmt.Println(slugresalt)
			
			permlink := strings.Split(characters[i].Moto, " ")
//			Url.Path += "/" + strconv.Itoa(characters[i].Id) + "/" + permlink[0] + "-" + permlink[1] + "-" +slugresalt+".html"
			Url.Path += "/" + characters[i].Id + "/" + permlink[0] + "-" + permlink[1] + "-" +slugresalt+".html"

			doc := new(domains.Page)
			doc.Loc = Url.String()
			doc.Lastmod = characters[i].Created_at.Format(time.RFC3339)
			doc.Changefreq = "weekly"

			docList.Pages = append(docList.Pages, doc)

		}

		resultXml, err := xml.MarshalIndent(docList, "", "  ")
		if err != nil {

			golog.Crit(err.Error())
		}

		go createmapfile.Createmap(golog, filestr, resultXml)
		w.Header().Add("Content-type", "application/xml")
		w.Write(resultXml)

	} else {
		f, _ := os.Open(filestr)
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, f)

		w.Header().Add("Content-type", "application/xml")
		w.Write(buf.Bytes())

	}

}


