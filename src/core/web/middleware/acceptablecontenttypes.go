package core

import (
	"net/http"
)

func AcceptableContentTypes(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedTypes := [2]string{"application/json", "application/xml"}

		m := make(map[string]bool)
		for i := 0; i < len(allowedTypes); i++ {
			m[allowedTypes[i]] = true
		}

		if _, ok := m[r.Header.Get("Accept")]; ok {
			h.ServeHTTP(w, r)
		}

		w.WriteHeader(http.StatusNotAcceptable)
	})
}
