package artist

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/internal/handlers"
)

const (
	root      = "/"
	artistURL = "/artist/"

	static = "/static/"

	OK                  = "OK"
	MethodNotAllowed    = "Method Not Allowed"
	BadRequest          = "Bad Request"
	InternalServerError = "Internal Server Error"
	NotFound            = "Not Found"

	PostFormError = "Post Form Error"
)

var _ handlers.Handler = &Handler{}

type Handler struct {
	Mux *http.ServeMux
}

func NewHandler() *Handler {
	router := *http.NewServeMux()
	h := Handler{&router}
	h.Register(h.Mux)
	return &h
}

func middleware(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"%s %s\n",
			r.Method,
			r.URL,
		)
		h.ServeHTTP(w, r)
	})
}

func (h *Handler) Register(router *http.ServeMux) {
	files := http.FileServer(http.Dir("./static"))
	router.Handle(static, middleware(http.StripPrefix(static, files).ServeHTTP))

	router.Handle(artistURL, middleware(h.artist))
	router.Handle(root, middleware(h.artists))
}

func (h *Handler) artists(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.error(w, NotFound, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		h.error(w, MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	artists, err := GetArtists()
	if err != nil {
		h.error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("").ParseFiles("templates/index.html", "templates/artists.html", "templates/header.html")
	if err != nil {
		h.error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index", artists)
	if err != nil {
		h.error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	log.Printf("%d %s\n", http.StatusOK, "OK")
}

func (h *Handler) artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.error(w, MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/artist/"))
	if err != nil {
		h.error(w, NotFound, http.StatusNotFound)
		return
	}
	artist, err := GetArtistById(strconv.Itoa(id))
	if err != nil {
		h.error(w, NotFound, http.StatusNotFound)
		return
	}
	log.Println(http.StatusOK)

	tmpl, err := template.New("").ParseFiles("templates/index.html", "templates/artist.html", "templates/header.html")
	if err != nil {
		h.error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index", artist)
	if err != nil {
		h.error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	log.Printf("%d %s\n", http.StatusOK, "OK")
}

func URLParam(r *http.Request, name string) string {
	ctx := r.Context()
	params := ctx.Value("params").(map[string]string)
	return params[name]
}

func (h *Handler) error(w http.ResponseWriter, err string, code int) {
	log.Println(err)
	tmpl, err1 := template.New("").ParseFiles("templates/error_page.html", "templates/index.html", "templates/header.html")
	if err1 != nil {
		log.Println(err1)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	err = fmt.Sprintf("%d %s", code, err)
	err2 := tmpl.ExecuteTemplate(w, "index", err)
	if err2 != nil {
		log.Println(err2)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	log.Println(err)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}
