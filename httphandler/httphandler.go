package httphandler

import (
	"net/http"

	"go.uber.org/zap"
)

type Handler struct {
	Mux    *http.ServeMux
	logger *zap.SugaredLogger
}

func New() *Handler {
	mux := http.NewServeMux()
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	h := Handler{Mux: mux, logger: sugar}
	h.registerRoutes()

	return &h
}

func (h *Handler) registerRoutes() {
	h.Mux.HandleFunc("/", h.hello)
	h.Mux.HandleFunc("/gk", h.kenobi)
}

func (h *Handler) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello there!"))
}

func (h *Handler) kenobi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Generak Kenobi!"))
}
