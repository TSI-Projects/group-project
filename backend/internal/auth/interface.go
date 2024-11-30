package auth

type IAuthClient interface {
	Login(username, password string) (string, error)
	Logout(accessToken string)
}
