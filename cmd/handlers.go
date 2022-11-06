package cmd

import (
	"ascii-art-web/packages"
	"net/http"
	"text/template"
)

var (
	word   string
	banner string
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// w.Header().Set("Allow", http.MethodGet)
		CheckerErrors(w, 405)
		// http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		CheckerErrors(w, 404)
		// http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		CheckerErrors(w, 500)
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		CheckerErrors(w, 500)
		return
	}

	// http.ServeFile(w, r, "index.html")
}

func Create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	// er, err := template.ParseFiles("error.html")
	if err != nil {
		CheckerErrors(w, 500)
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		// w.Header().Set("Allow", http.MethodPost)
		CheckerErrors(w, 405)
		// http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	word = r.FormValue("word")
	banner = r.FormValue("banner")
	if len(banner) == 0 {
		banner = "standard"
	}
	// check, err1 := pkg.HashCheck("./bannerFiles/" + banner + ".txt")
	// if err1 != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusInternalServerError)
	// 	return
	// }
	// if !check {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	return

	// }
	result, StatusCode := packages.AsciiArt(word, banner)
	if StatusCode == 400 {
		// err = tmpl.Execute(w, http.StatusText(500))
		CheckerErrors(w, 400)
		// http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if StatusCode == 500 {
		CheckerErrors(w, 500)
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
	err = tmpl.Execute(w, result)
	if err != nil {
		CheckerErrors(w, 500)
		// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
