package handler

import (
	"html/template"
	"log"
	"net/http"

	"git/ssengerb/ascii-art-web/pkg"
)

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	input := r.FormValue("convert")
	if len(input) > 300 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	banner := r.FormValue("fonts")
	result, number, err := pkg.Fonts(banner, input)
	if err != nil {
		if number == 2 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		return
	}

	err = tmp.Execute(w, result)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
