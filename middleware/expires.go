package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Parse the expires query arg and return true if the expiration date is valid
// Will always return true if query arg is not set
// Will return true if the expiration time is after the curnnet time
// Will return false otherwise
func VerifyExpires(qs map[string]interface{}) bool {
	_, expiresSet := qs["expires"]
	if !expiresSet {
		return true
	}

	expires, _ := strconv.ParseInt(qs["expires"].(string), 10, 64)
	if time.Now().Unix() < expires {
		return true
	}

	return false
}

func Expiry() gin.HandlerFunc {
	return func(c *gin.Context) {
		if VerifyExpires(c.MustGet("parameters").(map[string]interface{})) {
			c.Next()
			return
		}
		c.String(http.StatusUnauthorized, "Request expired")
		c.Abort()
	}
}
