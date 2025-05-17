package models

import (
	"errors"
	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/initializers"
	"golang.org/x/crypto/bcrypt"
)

type LoginModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (loginModel *LoginModel) Login() (string, error) {
	var user User //----> Declare user variable.
	
	//----> Check if the user email is attached to a genuine user.
	email := loginModel.Email
	result := initializers.DB.Where("email = ?", email).First(&user)

	//----> Record exist.
	if result.RowsAffected == 0{		
		return string(""), errors.New("invalid credentials")
	}

	//----> Check if the user-password is correct.
	password := loginModel.Password
	isValidPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	
	//----> Check validity of password.
	if isValidPassword != nil{
		return string(""), errors.New("invalid credential")
	}

	//----> Get token.
	token, err := authenticate.GenerateToken(user.Name, user.Email,user.ID, string(user.Role))
	
	//----> Check for errors.
	if err != nil {
		return string(""), errors.New("invalid credential")
	}
	
	//----> send back the response
	return token, nil
}

type ChangePasswordModel struct {
	Email    string  `json:"email" binding:"required"`
	OldPassword string  `json:"oldPassword" binding:"required"`
	NewPassword string  `json:"newPassword" binding:"required"`
	ConfirmPassword string  `json:"confirmPassword" binding:"required"`
}

func (changePasswordModel *ChangePasswordModel) ChangePassword() error {
	var user User
	//----> Extract the email, oldPassword, newPassword and confirmPassword 
	email := changePasswordModel.Email
	oldPassword := changePasswordModel.OldPassword
	newPassword := changePasswordModel.NewPassword
	confirmPassword := changePasswordModel.ConfirmPassword

	//----> Compare the newPassword with confirmPassword, they must be equal.
	isMatchPassword := newPassword == confirmPassword

	if !isMatchPassword{
		return errors.New("passwords must match")
	}

	//----> Check for existence of user by getting first marched record.
	result := initializers.DB.Where("email = ?", email).First(&user)

	//----> Record does not exist.
	if result.RowsAffected == 0{		
		return errors.New("invalid credentials")
	}

	//----> Compare the old password with the one in the database, they must match.
	isValidPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	
	//----> Check validity of password.
	if isValidPassword != nil{
		return errors.New("invalid credential")
	}

	//----> Hatch password
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 14)

	//----> Check for error.
	if err != nil{
		return errors.New("invalid credentials")
	}

	//----> Update the password.
	user.Password = string(hashedPassword)
	
	//----> Save the change in database.
	initializers.DB.Save(&user)

	return nil
}

type Role string

const (
	Admin   Role = "Admin"
	Staff Role = "Staff"
	customer Role = "User"
)

type Gender string

const (
	Male Gender = "Male"
	Female Gender = "Female"
)

type EditProfileModel struct {
	Name string  `json:"name"`
	Email    string  `json:"email" binding:"required"`
	Phone string  `json:"phone"`
	Address string  `json:"address"`
	Image string  `json:"image"`
	Gender Gender  `json:"gender"`
	Password string  `json:"password" binding:"required"`
	Role Role  `json:"role"`
}

func (editProfileModel *EditProfileModel) EditProfile() error {
	//----> Declare the user variable.
	var user User

	//----> Check for the availability of user.
	email := editProfileModel.Email
	result := initializers.DB.Where("email = ?", email).First(&user)

	//----> Record does not exist.
	if result.RowsAffected == 0{		
		return errors.New("invalid credentials")
	}

	//----> Check for password validity.
	password := editProfileModel.Password
	isValidPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	
	//----> Check validity of password.
	if isValidPassword != nil{
		return errors.New("invalid credential")
	} 

	//----> Update the user profile.
	user.Name =editProfileModel.Name
	user.Address = editProfileModel.Address
	user.Image =  editProfileModel.Image
	user.Gender = editProfileModel.Gender
	user.Phone = editProfileModel.Phone

	//----> Save the change in database.
	initializers.DB.Model(&user).Updates(&user)

	return nil
}

type SignupModel struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Gender Gender `json:"gender" binding:"required"`
	Password string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	Role Role `json:"role"`
	Image string `json:"image" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func (signup *SignupModel) Signup() error{
	var user User
	//----> Check the matches of password and confirmPassword.
	isMatchPassword := signup.Password == signup.ConfirmPassword
	
	//----> Check for error
	if !isMatchPassword{
		return errors.New("passwords must match")
	}
	
	//----> Check for existence of user by getting first marched record.
	result := initializers.DB.Where("email = ?", signup.Email).First(&user)

	//----> Record exist.
	if result.RowsAffected > 0{		
		return errors.New("invalid credentials")
	}

	//----> Hatch password
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup.Password), 14)

	//----> Check for error.
	if err != nil{
		return errors.New("invalid credentials")
	}

	//----> Populate user.
	user = User{
		Name: signup.Name,
		Email: signup.Email,
		Phone: signup.Phone,
		Gender: signup.Gender,
		Image: signup.Image,
		Address: signup.Address,
		Role: "Customer",
		//Role: signup.Role,
		Password: string(hashedPassword),
	}

	//----> Save the new user in the database.
	initializers.DB.Create(&user)

	//----> Send back the response.
	return nil
}