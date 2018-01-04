package service

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//检测是否正确Must
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	//ParseFiles是在文件形式上进行模板的具体化
	tmpl.Execute(w, nil)
}

/*deal with the unknown problem*/
//Handler
func InvalidRequestHandler() http.Handler {
	return http.HandlerFunc(InvalidRequest)
}

//implementation
func InvalidRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid Request!", 514)
}

//template output
func CheckIn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["name"][0]
	pn := r.Form["phonenumber"][0]
	id := r.Form["idnumber"][0]
	tmpl := template.Must(template.ParseFiles("template/info.html"))
	tmpl.Execute(w, map[string]string{
		"Name": name,
		"Pn":   pn,
		"ID":   id,
	})
}
