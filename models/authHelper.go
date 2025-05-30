package models

type UserAuthResp struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Image   string `json:"image"`
	Gender  Gender `json:"gender"`
	Role    Role   `json:"role"`
	Address string `json:"address"`
}

type LoginResp struct {
	CurrentUser UserAuthResp `json:"user"`
	Token       string       `json:"token"`
	IsAdmin     bool         `json:"isAdmin"`
	IsLoggedIn  bool         `json:"isLoggedIn"`
}

func makeLoginResp(token string, user User) LoginResp {
	return LoginResp{
		CurrentUser: UserAuthResp{
			ID:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Phone:   user.Phone,
			Image:   user.Image,
			Gender:  user.Gender,
			Role:    user.Role,
			Address: user.Address,
		},
		IsAdmin:    user.Role == "Admin",
		IsLoggedIn: true,
		Token:      token,
	}
}