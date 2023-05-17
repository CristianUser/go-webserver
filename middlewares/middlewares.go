package middlewares

import (
	"encoding/json"
	"net/http"

	"pronesoft/server/model"
	"pronesoft/server/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		tokenString := token.ExtractToken(c)
		session, err := model.GetSessionByToken(tokenString)

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		b, _ := json.Marshal(&session)
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)

		c.Set("userId", session.UserId)
		c.Set("session", m)
		c.Next()
	}
}
