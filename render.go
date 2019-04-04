package smooth

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gobuffalo/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

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

// RenderJSONError logs an error from err. If the err's cause is sql.ErrNoRows,
// it will render a status code of http.StatusNotFound and the corresponding
// http#StatusText. Otherwise, it will render the cause of the error with a
// status code of http.StatusInternalServerError.
func RenderJSONError(w http.ResponseWriter, r *http.Request, err error) {
	log.Error().Msgf("%+v", err)

	if errors.Cause(err) == sql.ErrNoRows {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, http.StatusText(http.StatusNotFound))
		return
	}

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
