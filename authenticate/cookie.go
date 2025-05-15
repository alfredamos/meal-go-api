package authenticate

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetCookieHandler(c *gin.Context, token string) {
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.String(http.StatusOK, "Cookie has been set")
}

func GetCookieHandler(c *gin.Context) {
	cookie, err := c.Cookie("user")
	if err != nil {
			c.String(http.StatusNotFound, "Cookie not found")
			return
	}
	c.String(http.StatusOK, "Cookie value: %s", cookie)
}

func DeleteCookieHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.String(http.StatusOK, "Cookie has been deleted")
}