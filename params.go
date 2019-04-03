package smooth

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

func HasParam(r *http.Request, key string) bool {
	return chi.URLParam(r, key) != ""
}

func GetUUIDFromPath(r *http.Request, key string) (uuid.UUID, *validate.Errors) {
	errs := validate.NewErrors()

	param := chi.URLParam(r, key)
	id, err := uuid.FromString(param)
	if err != nil {
		errs.Add(key, err.Error())
	}

	return id, errs
}
