package response

// objeto com os campos do user (response)
type UserResponse struct {
	ID    string `json: "id"`
	Email string `json: "email"`
	Name  string `json: "name"`
	Age   int8   `json: "age"`
}
