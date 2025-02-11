package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hanks177/podman/v4/pkg/api/types"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// referenceIDHandler adds X-Reference-Id Header allowing event correlation
// and Apache style request logging
func referenceIDHandler() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		// Only log Apache access_log-like entries at Info level or below
		out := ioutil.Discard
		if logrus.IsLevelEnabled(logrus.InfoLevel) {
			out = logrus.StandardLogger().Out
		}

		return handlers.CombinedLoggingHandler(out,
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				rid := r.Header.Get("X-Reference-Id")
				if rid == "" {
					if c := r.Context().Value(types.ConnKey); c == nil {
						rid = uuid.New().String()
					} else {
						rid = fmt.Sprintf("%p", c)
					}
				}

				r.Header.Set("X-Reference-Id", rid)
				w.Header().Set("X-Reference-Id", rid)
				h.ServeHTTP(w, r)
			}))
	}
}
