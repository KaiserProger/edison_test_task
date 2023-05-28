package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		w.Header().Set("Access-Control-Allow-Origin", "http://80.243.141.49:3000") // SET FRONTEND ADDRESS WHEN COMES TO PRODUCTION!!!
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
