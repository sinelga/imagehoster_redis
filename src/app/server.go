package main

import (
	//	"net/http"
	"github.com/rs/cors"
	"github.com/zenazn/goji"
	"handlers"
	"log"
	"log/syslog"
//	"net/http"
)

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	goji.Use(c.Handler)

	goji.Get("/api", handlers.MhandleAll)
	goji.Get("/api/:id", handlers.MhandleAll)
	//	goji.Get("/img/:id/:imgfile/:img/:mime/:width/:height", handlers.ImageShow)
	goji.Get("/img/:id/:imgfile/:width/:height", handlers.ImageShow)
	goji.Get("/chat/:uuid/:phone/:ask", handlers.GetChatAnswer)
//	goji.Get("/interceptor", handlers.MhandleAll)
	//	goji.Get("/uploadapi", handlers.ShowAll)
	//	goji.Get("/fullimage/", http.FileServer(http.Dir("upload/img/")))
	goji.Get("/fullimage/:id/original/:imgfile", handlers.ImageFullShow)
	//	goji.Get("/fullimage/:id/original/:imgfile",http.FileServer(http.Dir("upload/img")))
	//	goji.Get("/fullimage",handlers.ImageFullShow)

	//	goji.Options("/upload",handlers.MakeUpload)
	//	goji.Post("/upload",handlers.MakeUpload)

//	staticFilesLocation := "/home/juno/git/8_fi_FIporno_desk/dist"
//	goji.Handle("/*", http.FileServer(http.Dir(staticFilesLocation)))
	goji.Get("/*",handlers.Elaborate)

	goji.Serve()
}
