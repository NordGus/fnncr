package creationtime

import "time"

func isTooOld(when time.Time) bool {
	return time.Since(when) > maxAge
}
