package schemas

type LoginType struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterType struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}
