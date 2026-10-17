package dto

type (
	RegisterRequest struct {
		Email           string `json:"email" validate:"required,email"`
		Username        string `json:"username" validate:"required,min=3,max=50"`
		Password        string `json:"password" validate:"required,min=4,max=50"`
		PasswordConfirm string `json:"password_confirm" validate:"required,min=4,max=50,eqfield=Password"`
	}

	RegisterResponse struct {
		ID string `json:"id"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=4,max=50"`
	}

	LoginResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)

type (
	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}
	RefreshTokenResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)
