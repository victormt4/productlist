package middleware

import (
	"errors"
	"net/http"
	"productlist/services/responseservice"
)

//RecoverMiddleware Return a 500 response on panic()
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			r := recover()
			if r != nil {

				var err error

				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("unknown error")
				}

				resService := responseservice.GetResponseServiceWriter(w)
				resService.SendServerError(struct {
					Message string
				}{Message: err.Error()})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
