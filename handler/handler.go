package handler

import (
	"net/http"
	"sync"
)

type CustomFile struct {
	m    sync.Mutex
	file string
}

func (h *CustomFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.m.Lock()
	defer h.m.Unlock()

  println("REQUESTURI        :" , r.RequestURI)
	h.file = "/home/onadrog/.config/swaypowerbtn/custom.css"
  w.Write([]byte(h.file))
}

func Handler() {
	http.Handle("/", &CustomFile{})
}
