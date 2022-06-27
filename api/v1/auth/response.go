package auth

type Response struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Primary bool   `json:"primary"`
}
