package authentication

type FormLogin struct {
	ActionURL string
	Username  string
	Password  string
	Failed    bool
}
