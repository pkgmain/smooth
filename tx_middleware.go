package smooth

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

func PopTXMiddleware(conn *pop.Connection) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			_ = conn.Transaction(func(tx *pop.Connection) error {
				ctx := context.WithValue(r.Context(), "tx", tx)
				ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

				next.ServeHTTP(ww, r.WithContext(ctx))

				if ww.Status() >= http.StatusBadRequest {
					return errors.New("rolling back tx")
				}

				return nil
			})
		}

		return http.HandlerFunc(fn)
	}
}

func GetTX(r *http.Request) *pop.Connection {
	return r.Context().Value("tx").(*pop.Connection)
}
