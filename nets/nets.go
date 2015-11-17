// Simple HTTP server with recovery from panic attacks
//
// Author: Dmitri Krasnenko

package nets

import (
	"net/http"
	"log"
	"fmt"
)


type ServeMux struct {
	*http.ServeMux
}

func NewServeMux() *ServeMux {
	return &ServeMux{http.NewServeMux()}
}

func (m *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request){
	//503 recovery
	defer R503(
		w,r,
	)

	m.ServeMux.ServeHTTP(
		w,r,
	)
}

func R503(w http.ResponseWriter, req *http.Request) {
	if err := recover(); err != nil {
		http.NotFound(w, req)
	}
}


func Pass(){
}

func DeployHTTP() {
	http_mux := NewServeMux()

	http_mux.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {

		if name := r.URL.Query().Get("who"); name != "" {
			w.Write([]byte("Hello " + name + "!"))
		}else {
			w.Write([]byte("Hello Anonymous!"))
		}
	})

	http_mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("It's panic attack!")
	})

	fmt.Println("Listening on port 8000.")
	if err := http.ListenAndServe(":8000", http_mux); err != nil {
		log.Fatal("Unable to start server: ",err)
	}
}