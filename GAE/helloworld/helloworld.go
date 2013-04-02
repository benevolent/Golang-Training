package helloworld

import(
	"fmt"
	"net/http"
)

func init(){
	http.HandleFunc("/",handle)
}

func handle(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<html><body>HelloWorld!! こんちゃ!!</body></html>")
}
