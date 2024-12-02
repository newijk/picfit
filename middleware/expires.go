package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thoas/picfit/constants"
)

// Parse the expires query arg and return true if the expiration date is valid
// Will always return true if query arg is not set
// Will return true if the expiration time is after the curnnet time
// Will return false otherwise
func VerifyExpires(qs map[string]interface{}) bool {
	_, expiresSet := qs[constants.ExpiresParamName]
	if !expiresSet {
		return true
	}

	expires, _ := strconv.ParseInt(qs[constants.ExpiresParamName].(string), 10, 64)
	return time.Now().Unix() < expires
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
