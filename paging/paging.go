package paging

import (
	"net/http"

	"github.com/gobuffalo/pop"
)

type Request struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type Params struct {
	r *http.Request
}

func NewParams(r *http.Request) *Params {
	return &Params{r}
}

func (p *Params) Get(key string) string {
	return p.r.URL.Query().Get(key)
}

type Info struct {
	// Current page you're on
	Page int `json:"page"`
	// Number of results you want per page
	PerPage int `json:"perPage"`
	// Page * PerPage (ex: 2 * 20, Offset == 40)
	Offset int `json:"offset"`
	// Total potential records matching the query
	TotalEntriesSize int `json:"totalEntriesSize"`
	// Total records returns, will be <= PerPage
	CurrentEntriesSize int `json:"currentEntriesSize"`
	// Total pages
	TotalPages int `json:"totalPages"`
}

type Page struct {
	Info
	Content []interface{} `json:"content"`
}

func NewPage(paginator *pop.Paginator, content []interface{}) *Page {
	return &Page{Info(*paginator), content}
}
