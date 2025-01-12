package handlers

import "net/http"

func (h *Handler) Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
