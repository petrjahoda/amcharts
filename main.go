package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/html/*filepath", http.Dir("html"))
	router.ServeFiles("/css/*filepath", http.Dir("css"))
	router.ServeFiles("/data/*filepath", http.Dir("data"))
	router.GET("/", home)
	router.POST("/get_downtime_raw_data", getDowntimes)
	router.POST("/get_downtime_users_data", getDowntimeUsers)
	router.POST("/get_users_downtimes", getUsersDowntimes)
	_ = http.ListenAndServe(":80", router)
}




func home(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
http.ServeFile(writer, request, "./html/index.html")
}