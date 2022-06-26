package google

// User is a retrieved and authentiacted user.
type User struct {
	Sub     string `json:"sub"`
	Picture string `json:"picture"`
	Email   string `json:"email"`
}
