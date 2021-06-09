package mw

import (
	"net/http"

	"github.com/balboeng/sergeitter/db"
)

/* CheckDbStatus mw that check db status*/
func CheckDbStatus(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
