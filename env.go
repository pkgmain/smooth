package smooth

import "os"

func GetEnvOrDefault(key, def string) string {
	val, found := os.LookupEnv(key)
	if !found {
		return def
	}

	return val
}
