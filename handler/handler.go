package handler

import (
	"net/http"

	"github.com/shoaibahmed997/automata/mouse"
)

func MainHandler() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/api/test", mouse.MouseTest)

	http.HandleFunc("/api/recordMouse", mouse.RecordMouse)

	http.HandleFunc("/api/allMacros", mouse.ReturnAllMacros)

}
