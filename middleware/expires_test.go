package middleware

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestVerifyExpires(t *testing.T) {
	plusOneMine := strconv.FormatInt(time.Now().Add(time.Duration(time.Minute)).Unix(), 10)
	qs := map[string]interface{}{
		"expires": plusOneMine,
	}
	assert.True(t, VerifyExpires(qs))
}

func TestVerifyExpiresInvalid(t *testing.T) {
	qs := map[string]interface{}{
		"expires": "notvalid",
	}
	assert.False(t, VerifyExpires(qs))
}

func TestVerifyExpiresNotSet(t *testing.T) {
	qs := map[string]interface{}{}
	assert.True(t, VerifyExpires(qs))
}

func TestVerifyExpiresExpired(t *testing.T) {
	oneMinuteAgo := strconv.FormatInt(time.Now().Add(-time.Duration(time.Minute)).Unix(), 10)
	qs := map[string]interface{}{
		"expires": oneMinuteAgo,
	}
	assert.False(t, VerifyExpires(qs))
}
