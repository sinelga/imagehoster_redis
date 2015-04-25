package robots

import (
	"bytes"
	"log/syslog"
	"net/http"
)

func Generate(golog syslog.Writer, w http.ResponseWriter, site string) {

	var buffer bytes.Buffer
	buffer.WriteString("User-agent: *\nAllow: /\nSitemap: http://" + site + "/sitemap.xml\n")

	w.Header().Add("Content-type", "text/plain")
	w.Write(buffer.Bytes())

}
