package auth

type Auth struct {
	accessToken  string
	refreshToken string
}

func (a *Auth) SetAccessToken(accessToken string) {
	a.accessToken = accessToken
}

func (a *Auth) SetRefreshToken(refreshToken string) {
	a.refreshToken = refreshToken
}
