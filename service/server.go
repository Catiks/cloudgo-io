package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

/*服务器生成*/
func NewServer() *negroni.Negroni {

	formatter := render.New()

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			fmt.Println(root)
			fmt.Println("Root above")
		}
	}
	mx.HandleFunc("/", homeHandler).Methods("GET")
	mx.HandleFunc("/", CheckIn).Methods("POST")
	mx.HandleFunc("/unknown", InvalidRequest)
	//保证能收到html的其他数据，比如图片，css等
	mx.PathPrefix("/assets/static").Handler(http.StripPrefix("/assets/static/", http.FileServer(http.Dir(webRoot+"/assets/static/"))))
}
