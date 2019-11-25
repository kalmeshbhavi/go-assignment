package adapter

import "net/http"

type ServerAdapter func(http.Handler) http.Handler

func ServerApply(h http.Handler, adapters ...ServerAdapter) http.Handler {

	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
