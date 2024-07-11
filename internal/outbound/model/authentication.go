package model

type AuthRequest struct {
	Username string
	Password string
	AppName  string
}

type AuthResponse struct {
	FullName string
	RoleList []string
}
