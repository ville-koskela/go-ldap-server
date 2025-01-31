package domain

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	UID      int    `json:"uid"`
	GID      int    `json:"gid"`
}
