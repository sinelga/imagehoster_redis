package notjsbots

import (
"log/syslog"
	"net/http"
	"github.com/zenazn/goji/web"
	"domains"
	"fmt"
)
func CreateNotJsPage(golog syslog.Writer,c web.C, w http.ResponseWriter, r *http.Request,variant string,character domains.Character){
	
	golog.Info("CreateNotJsPage msnbot!!! "+character.Name)
//	http.ServeFile(w, r, "/home/juno/git/fi_FI_desk_mobile/version_"+variant+"/dist/index.html")

	fmt.Fprintf(w, "Hello")
	
}
