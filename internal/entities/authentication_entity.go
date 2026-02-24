package entities

type (
	LoginRequest struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}
	TokenResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}
)
