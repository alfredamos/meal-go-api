package models

import (
	"errors"
	"fmt"

	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/initializers"
	"golang.org/x/crypto/bcrypt"
)

type LoginModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (loginModel *LoginModel) Login() (LoginResp, error) {
	user := User{} //----> Declare user variable.
	
	//----> Check if the user email is attached to a genuine user.
	email := loginModel.Email
	err := initializers.DB.Where("email = ?", email).First(&user).Error

	//----> Record does not exist.
	if err != nil{		
		return LoginResp{}, errors.New("invalid credentials")
	}

	//----> Check if the user-password is correct.
	password := loginModel.Password
	isValidPasswordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	
	//----> Check validity of password.
	if isValidPasswordError != nil{
		return LoginResp{}, errors.New("invalid credential")
	}

  userId := user.ID //----> User-id.
	userRole := string(user.Role) //----> User-role.
	userName := user.Name //----> User-name.
	userEmail := user.Email //----> User-email

	//----> Get token.
	token, err := authenticate.GenerateToken(userName, userEmail,userId, userRole)
	
	//----> Check for errors.
	if err != nil {
		return LoginResp{}, errors.New("invalid credential")
	}

	loginResp := makeLoginResp(token, user)

	//----> send back the response
	return loginResp, nil
}

type ChangePasswordModel struct {
	Email    string  `json:"email" binding:"required"`
	OldPassword string  `json:"oldPassword" binding:"required"`
	NewPassword string  `json:"newPassword" binding:"required"`
	ConfirmPassword string  `json:"confirmPassword" binding:"required"`
}

func (changePasswordModel *ChangePasswordModel) ChangePassword() error {
	user := User{} //----> Declare user variable.
	
	//----> Extract the email, oldPassword, newPassword and confirmPassword 
	email := changePasswordModel.Email
	oldPassword := changePasswordModel.OldPassword
	newPassword := changePasswordModel.NewPassword
	confirmPassword := changePasswordModel.ConfirmPassword

	//----> Compare the newPassword with confirmPassword, they must be equal.
	isMatchPassword := matchPassword(newPassword, confirmPassword)
	fmt.Println("In change-password-model, isMatchPassword : ", isMatchPassword)
	fmt.Println("In change-password-model, email : ", email)
	if !isMatchPassword{
		return errors.New("passwords must match")
	}

	//----> Check for existence of user by getting first marched record.
	err := initializers.DB.Where("email = ?", email).First(&user).Error
	fmt.Println("In change-password-model, error after save in db : ", err)
	//----> Record does not exist.
	if err != nil{		
		return errors.New("invalid credentials")
	}

	//----> Compare the old password with the one in the database, they must match.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	fmt.Println("In change-password-model, isValidPasswordError after comparing with old password : ", err)
	//----> Check validity of password.
	if err != nil{
		return errors.New("invalid credential")
	}

	//----> Hatch password
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 14)
	fmt.Println("In change-password-model, error hatching password : ", err)
	//----> Check for error.
	if err != nil{
		return errors.New("invalid credentials")
	}

	//----> Update the password.
	user.Password = string(hashedPassword)
	
	//----> Save the change in database.
	err = initializers.DB.Save(&user).Error

	fmt.Println("In change-password-model, error after saving new password : ", err)

	//----> Check for error.
	if err != nil{
		return errors.New("updated password cannot be saved")
	}

	return nil
}

func GetCurrentUser(userId string) (User, error){
	//----> Get user with the given id.
	user, err := userGetById(userId)
	
	//----> Check for error.
	if err != nil {
		return User{}, errors.New("user cannot be retrieved")
	}

	//----> Send back the response.
	 return user, nil
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
	user := User{} //----> Declare user variable.

	//----> Check for the availability of user.
	email := editProfileModel.Email
	err := initializers.DB.Where("email = ?", email).First(&user).Error
	
	//----> Record does not exist.
	if err != nil{		
		return errors.New("invalid credentials")
	}
	
	//----> Check for password validity.
	password := editProfileModel.Password
	isValidPasswordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	
	//----> Check validity of password.
	if isValidPasswordError != nil{
		return errors.New("invalid credential")
	} 

	//----> Update the user profile.
	user.Name = editProfileModel.Name
	user.Address = editProfileModel.Address
	user.Image =  editProfileModel.Image
	user.Gender = editProfileModel.Gender
	user.Phone = editProfileModel.Phone

	//----> Save the change in database.
	err = initializers.DB.Model(&user).Updates(&user).Error

	//----> Check for error.
	if err != nil{
		return errors.New("updated profile cannot be saved")
	}

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
	user := User{} //----> Declare user variable.
	
	//----> Check the matches of password and confirmPassword.
	isMatchPassword := matchPassword(signup.Password,signup.ConfirmPassword)
	
	//----> Check for error.
	if !isMatchPassword{
		return errors.New("passwords must match")
	}
	
	//----> Check for existence of user by getting first marched record.
	err := initializers.DB.Where("email = ?", signup.Email).First(&user).Error

	//----> Record exist.
	if err == nil{		
		return errors.New("invalid credentials")
	}

	//----> Hatch password
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup.Password), 14)

	//----> Check for error.
	if err != nil{
		return errors.New("invalid credentials")
	}

	//----> Populate user.
	user = populateUser(*signup, string(hashedPassword))
	
	//----> Save the new user in the database.
	initializers.DB.Create(&user)

	//----> Send back the response.
	return nil
}

func matchPassword(password, confirmPassword string) bool {
	return password == confirmPassword
}

func populateUser(signup SignupModel, hashedPassword string) User{
	//----> Populate user.
	user := User{
		Name: signup.Name,
		Email: signup.Email,
		Phone: signup.Phone,
		Gender: signup.Gender,
		Image: signup.Image,
		Address: signup.Address,
		Role: "Customer",
		//Role: signup.Role,
		Password: hashedPassword,
	}

	return user
}