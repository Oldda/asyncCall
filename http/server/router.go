package server

import(
	"net/http"
	"log"
)

type handlerFunc func(http.ResponseWriter,*http.Request)

type router struct{
	routes map[string]handlerFunc
}

func NewRouter()*router{
	return &router{
		routes : make(map[string]handlerFunc),
	}
}

func(this *router)Registe(pattern string,handler handlerFunc){
	log.Println("registe route "+pattern)
	this.routes[pattern] = handler
}

func (this *router)Handle(w http.ResponseWriter,r *http.Request){
	path := r.Method + ":" + r.URL.Path
	if handler,ok := this.routes[path]; ok{
		handler(w,r)
		return
	}
	http.NotFound(w,r)
}