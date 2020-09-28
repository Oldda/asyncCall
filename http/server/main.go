package server

import(
	"net/http"
	"log"
	"time"
)

type ServerMux struct{
	router *router
}

func NewServerMux()*ServerMux{
	return &ServerMux{
		router: NewRouter(),
	}
}

func (mux *ServerMux)Run(addr string){
	server := &http.Server{
		Addr : addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("server running and listening at "+addr)
	log.Fatal(server.ListenAndServe())
}

func(mux *ServerMux)GET(pattern string,handler handlerFunc){
	path := "GET"+":"+pattern
	mux.router.Registe(path,handler)
}

func(mux *ServerMux)POST(pattern string,handler handlerFunc){
	path := "POST"+":"+pattern
	mux.router.Registe(path,handler)
}

func(mux *ServerMux)PUT(pattern string,handler handlerFunc){
	path := "PUT"+":"+pattern
	mux.router.Registe(path,handler)
}

func(mux *ServerMux)DELETE(pattern string,handler handlerFunc){
	path := "DELETE"+":"+pattern
	mux.router.Registe(path,handler)
}

//handler无参数，无返回值
func(mux *ServerMux)Ticker(space string, handler func()){
	duration,err := time.ParseDuration(space)
	if err != nil{
		log.Println("ticker duration parsed error")
		return
	}
	ticker := time.NewTicker(duration)
	for range ticker.C{
		handler()
	}
}

func (mux *ServerMux)ServeHTTP(writer http.ResponseWriter, req *http.Request){
	mux.router.Handle(writer,req)
}