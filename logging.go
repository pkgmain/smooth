package smooth

import (
	"fmt"

	"github.com/gobuffalo/pop/logging"
	"github.com/rs/zerolog/log"
)

func ZeroPopLogger(lvl logging.Level, s string, args ...interface{}) {
	if lvl == logging.SQL {
		s = "[SQL] - " + s

		if len(args) > 0 {
			fmtArgs := make([]string, len(args))
			for i, a := range args {
				switch a.(type) {
				case string:
					fmtArgs[i] = fmt.Sprintf("%q", a)
				default:
					fmtArgs[i] = fmt.Sprintf("%v", a)
				}
			}
			s = fmt.Sprintf("%s | %s", s, fmtArgs)
		} else {
			s = fmt.Sprintf("%s", s)
		}
	} else {
		s = fmt.Sprintf(s, args...)
		s = fmt.Sprintf("%s", s)
	}

	switch lvl {
	case logging.Debug:
		log.Debug().Msg(s)
	case logging.Info:
		log.Info().Msg(s)
	case logging.Warn:
		log.Warn().Msg(s)
	case logging.Error:
		log.Error().Msg(s)
	default:
		log.Debug().Msg(s)
	}
}
