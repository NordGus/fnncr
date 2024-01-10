package sessionrepo

type (
	Session struct {
		ID     string
		UserID int64
	}
)

func (s Session) UserId() int64 {
	return s.UserID
}
