package sessionrepo

type (
	Session struct {
		ID     string `redis:"id" json:"id"`
		UserID int64  `redis:"userID" json:"userID"`
	}
)

func (s Session) UserId() int64 {
	return s.UserID
}
