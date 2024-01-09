package authentication

import (
	"errors"
)

// authenticate receives a username and a password and authenticates the user and returns the session key. If it can't
// find the user in the system or it can't authenticated returns an error.
func (s *Service) authenticate(username string, password string) (string, error) {
	// TODO: Implement a data storage to retrieve users and authenticate them

	return "test", errors.New("authentication: authenticate not implemented yet")
}
