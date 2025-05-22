package authenticate

import (
	"errors"
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
	token, err := getTokenFromCookie(c)

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
		os.Exit(1)
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
		userId := claims["userId"]

		//----> Set the claims on gin context
		c.Set("name", name)
		c.Set("email", email)
		c.Set("role", role)
		c.Set("userId", userId)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail","message": "Invalid credential!", "statusCode": http.StatusUnauthorized})
		return 
	}

}
func VerifyTokenJwt(c *gin.Context){
	//----> Get the token from cookie.
	token, err := getTokenFromCookie(c)

	//----> Check for error.
	if err != nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail","message": "Invalid credential!", "statusCode": http.StatusUnauthorized})
		return 
	}

	//----> Check for valid token
	 parsedToken := validateToken(c, token)

	 //----> Get user claims.
	 err = getUserClaims(c, parsedToken)

	 //----> Check for error.
	 if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail","message": "Invalid credential!", "statusCode": http.StatusUnauthorized})
		return
	}
	c.Next()
	/* if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		//----> Access claims
		name := claims["name"].(string)
		email := claims["email"].(string)
		role := claims["role"].(string)
		userId := claims["userId"]

		//----> Set the claims on gin context
		c.Set("name", name)
		c.Set("email", email)
		c.Set("role", role)
		c.Set("userId", userId)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail","message": "Invalid credential!", "statusCode": http.StatusUnauthorized})
		return 
	} */

}

func getUserClaims(c *gin.Context, parsedToken jToken) error{
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		//----> Access claims
		name := claims["name"].(string)
		email := claims["email"].(string)
		role := claims["role"].(string)
		userId := claims["userId"]

		//----> Set the claims on gin context
		c.Set("name", name)
		c.Set("email", email)
		c.Set("role", role)
		c.Set("userId", userId)
		c.Next()
		return nil
	} else {
		return errors.New("invalid credentials")
	}

}

func getTokenFromCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie("token")
	if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			
			return string(""), err
	}

	return token, nil

}

type jToken *jwt.Token

func validateToken(c *gin.Context, token string) jToken{
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
		os.Exit(1)
	}

	isValidToken := parsedToken.Valid

	if !isValidToken {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail","message": "Invalid credential!", "statusCode": http.StatusUnauthorized})
		os.Exit(1) 
	}

	return parsedToken
}

func getRole(role string) string {
	if role == "Admin" {
		return "senior"
	}
	if role == "Customer" {
		return "Customer"
	}
	return "Staff"
}