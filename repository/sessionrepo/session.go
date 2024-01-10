package sessionrepo

type (
	Session struct {
		ID     string `redis:"id"`
		UserID int64  `redis:"userID"`
	}
)

func (s Session) UserId() int64 {
	return s.UserID
}
