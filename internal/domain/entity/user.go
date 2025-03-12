package entity

type User struct {
	ID       string `json:"id"`
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=user admin"`
}

// NewUser creates a new user with default role.
func NewUser(username, email, password string) *User {
	return &User{
		ID:       "", // Will be set by repository
		Username: username,
		Email:    email,
		Password: password,
		Role:     "user", // Default role
	}
}
