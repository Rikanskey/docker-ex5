package v1

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
)

var Count = 0
var Form = "<h3> Hello , Волобуев Игорь P42081</h3>"

func (h handler) GetCount(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, Count)
}

func (h handler) GetCountAndInc(w http.ResponseWriter, r *http.Request) {
	oldCount := Count
	Count++
	render.Respond(w, r, oldCount)
}

func (h handler) GetAbout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, Form)
	return
}
