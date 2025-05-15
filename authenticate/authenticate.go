package authenticate

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("JWT_TOKEN_SECRET")

func GenerateToken(name string, email string, userId uint, role string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"name": name, "email": email, "userId": userId, "role": role, "expiresAt": time.Now().Add(time.Hour * 2)})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(c *gin.Context){
	//----> Get the token from cookie.
	token, err := GetTokenFromCookie(c)

	//----> Check for error.
	if err != nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail","message": "Invalid credential!", "statusCode": http.StatusUnauthorized})
		return 
	}

	//----> Check for valid token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		//----> Return the secret key for signing
    return []byte(secretKey), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail","message": "Invalid credential!", "statusCode": http.StatusUnauthorized})
		return 
	}

	isValidToken := parsedToken.Valid

	if !isValidToken {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail","message": "Invalid credential!", "statusCode": http.StatusUnauthorized})
		return 
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		//----> Access claims
		name := claims["name"].(string)
		email := claims["email"].(string)
		role := claims["role"].(string)

		//----> Set the claims on gin context
		c.Set("name", name)
		c.Set("email", email)
		c.Set("role", role)
		fmt.Println("I'm authenticated!")
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail","message": "Invalid credential!", "statusCode": http.StatusUnauthorized})
		return 
	}

}

func GetTokenFromCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie("token")
	if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			
			return string(""), err
	}

	return token, nil

}

func GetRole(role string) string {
	if role == "Admin" {
		return "senior"
	}
	if role == "Customer" {
		return "Customer"
	}
	return "Staff"
}