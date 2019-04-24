package smooth

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func RenderNoContent(w http.ResponseWriter, r *http.Request) {
	render.NoContent(w, r)
}

// RenderJSON renders v as JSON with a status code of http.StatusOK.
func RenderJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, v)
}

// RenderJSONErrorWithStatus renders v as JSON with as status code of status.
func RenderJSONWithStatus(w http.ResponseWriter, r *http.Request, status int, v interface{}) {
	render.Status(r, status)
	render.JSON(w, r, v)
}

// RenderBadRequest logs a warning with errs and then renders errs as JSON with
// a status code of http.StatusBadRequest.
func RenderBadRequest(w http.ResponseWriter, r *http.Request, errs *validate.Errors) {
	log.Warn().Msg(errs.String())

	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, errs)
}

// RenderJSONError, by default, verbosely error logs err and renders the
// cause of the error with a status code of http.StatusInternalServerError.
// However, if the err's cause is sql.ErrNoRows, it will render a status code
// of http.StatusNotFound and the corresponding http#StatusText.
func RenderJSONError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Cause(err) == sql.ErrNoRows {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, http.StatusText(http.StatusNotFound))
		return
	}

	log.Error().Msgf("%+v", err)
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, errors.Cause(err).Error())
}

// RenderJSONErrorWithStatus logs a warning if status is 4xx or logs an error if
// status is 5xx. It then renders err's cause as JSON with a status code of status.
func RenderJSONErrorWithStatus(w http.ResponseWriter, r *http.Request, status int, err error) {
	if status >= http.StatusBadRequest && status < http.StatusInternalServerError {
		log.Warn().Msgf("%+v", err)
	} else if status >= http.StatusInternalServerError {
		log.Error().Msgf("%+v", err)
	}

	render.Status(r, status)
	render.JSON(w, r, errors.Cause(err).Error())
}

// RenderHTML will find the template in the the templates *packr.Box and render
// it with the given (optional) model. If multiple models are provided, they
// will be merged.
func RenderHTML(w http.ResponseWriter, r *http.Request, template string, model ...map[string]interface{}) {
	t, err := templates.FindString(template)
	if err != nil {
		//todo: do something better like render and error page
		log.Error().Msgf("%+v", err)
		return
	}

	// Merge the models
	data := make(map[string]interface{})
	for _, m := range model {
		for k, v := range m {
			data[k] = v
		}
	}

	c := plush.NewContextWith(data)
	html, err := plush.Render(t, c)
	if err != nil {
		//todo: do something better like render and error page
		log.Error().Msgf("%+v", err)
		return
	}

	render.Status(r, http.StatusOK)
	render.HTML(w, r, html)
}
