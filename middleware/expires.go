package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Expiry() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.MustGet("parameters").(map[string]interface{})

		_, exists := params["expires"]
		if !exists {
			c.Next()
		}

		expires, err := strconv.ParseInt(params["expires"].(string), 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "Expires format must be integer seconds since (unix) epoch")
			c.Abort()
		}

		requestExpired := time.Now().Unix() > expires

		if requestExpired {
			c.String(http.StatusUnauthorized, "Request expired")
			c.Abort()
			return
		}

		c.Next()
	}
}
