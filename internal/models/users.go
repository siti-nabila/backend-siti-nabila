package models

type (
	RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleId   *int   `json:"role_id,omitempty"`
	}

	LoginReqeust struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AuthResponse struct {
		ErrorMessage string `json:"error,omitempty"`
		Token        string `json:"token,omitempty"`
	}
)
