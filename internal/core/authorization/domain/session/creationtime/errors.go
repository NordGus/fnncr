package creationtime

import "errors"

var (
	ErrCreationTimeExceedMaxAge = errors.New("creationtime: creation time exceed allowed max age")
)
