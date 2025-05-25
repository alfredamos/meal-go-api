package authenticate

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetCookieHandler(c *gin.Context, token string) {
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.String(http.StatusOK, "Cookie has been set")
}

func GetCookieHandler(c *gin.Context) {
	cookie, err := c.Cookie("token")
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