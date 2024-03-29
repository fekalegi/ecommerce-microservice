package dto

type RequestUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	RoleID   int    `json:"role_id"`
}

type ResponseLogin struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ResponseRefreshToken struct {
	AccessToken string `json:"access_token"`
}

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseAddUser struct {
	UserID int `json:"user_id"`
}
