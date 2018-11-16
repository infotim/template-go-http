package main

import (
	"fmt"
	"net/http"
)

func (s *server) handlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, s.config.PONG_MESSAGE)
	}
}
