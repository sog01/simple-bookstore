package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/sog01/simple-bookstore/internal/model"
	"github.com/sog01/simple-bookstore/internal/service"
)

// BooksHandlerWeb represent a handler for books web
type BooksHandlerWeb struct {
	booksService service.Books
	t            *template.Template
}

// NewBooksHandlerWeb construct handler for books web
func NewBooksHandlerWeb(booksService service.Books) *BooksHandlerWeb {
	files := []string{
		"./web/templates/index.html",
	}

	tt, _ := template.ParseFiles(files...)
	return &BooksHandlerWeb{
		booksService: booksService,
		t:            tt,
	}
}

// HandleStaticFS handle for static file server
func (bhw *BooksHandlerWeb) HandleStaticFS() http.Handler {
	return http.FileServer(http.Dir("./web/static"))
}

// HandleGetBooksPage handle for get books web page
func (bhw *BooksHandlerWeb) HandleGetBooksPage(w http.ResponseWriter, r *http.Request) {
	req := &model.GetBooksRequest{}
	req.Page, _ = strconv.Atoi(r.FormValue("page"))

	booksResp, err := bhw.booksService.GetBooks(r.Context(), req)
	if err != nil {
		writeResponseInternalError(w, err)
		return
	}

	pageList := []map[string]interface{}{}
	for page := 1; page <= booksResp.MaxPage; page++ {
		pm := map[string]interface{}{
			"page": page,
		}
		if req.Page == page {
			pm["active"] = "active"
		}
		pageList = append(pageList, pm)
	}
	bhw.t.ExecuteTemplate(w, "index.html", map[string]interface{}{
		"bookList": booksResp.BookList,
		"pageList": pageList,
	})
}

// Router used to define the books web handler router
func (bhw *BooksHandlerWeb) Router(mux *http.ServeMux) {
	mux.Handle("/static/", http.StripPrefix("/static/", bhw.HandleStaticFS()))
	mux.HandleFunc("/", bhw.HandleGetBooksPage)
}
