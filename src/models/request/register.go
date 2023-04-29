package request

type RegisterBody struct {
	Email    string `json="email"`
	Password string `json="password"`
	Name     string `json="name"`
}
