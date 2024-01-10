package usersrepo

type User struct {
	ID             int64
	AccessName     string
	PasswordDigest string
}

func (u User) Id() int64 {
	return u.ID
}

func (u User) Username() string {
	return u.AccessName
}

func (u User) PasswordHash() []byte {
	return []byte(u.PasswordDigest)
}
