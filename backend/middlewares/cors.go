package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // SET FRONEND ADDRESS WHEN COMES TO PRODUCTION!!!
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
