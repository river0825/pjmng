package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"river0825/cleanarchitecture/core/arch/port/gin/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AuthorizationMiddleware struct {
	authService service.JWTTokenAuthor
}

func NewAuthorizationMiddleware(
	authService service.JWTTokenAuthor,
) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{
		authService: authService,
	}
}

func (m *AuthorizationMiddleware) DevBindUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		setIfNotEmpty := func(key, value string) {
			if value != "" {
				c.Set(key, value)
			}
		}

		setIfNotEmpty("nk_user_id", c.Request.Header.Get("X-USER-ID"))
		setIfNotEmpty("nk_is_subscribed", c.Request.Header.Get("X-IS-SUBSCRIBE"))
		setIfNotEmpty("nk_organization", c.Request.Header.Get("X-ORGANIZATION-ID"))

		c.Next()
	}
}
func (m *AuthorizationMiddleware) ValidateAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := m.getAccessTokenFromHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("[ValidateAccessToken] invalid authorization header: %s", err.Error()),
			})
			c.Abort()
			return
		}

		result, err := m.authService.ValidateAccessToken(c, accessToken)

		log.Logger.Info().Any("user info", result).Msg("user info")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("[ValidateAccessToken] invalid token or expired: %s", err.Error()),
			})
			c.Abort()
			return
		}

		c.Set("user-meta", result.Meta)
		c.Next()
	}
}

func (m *AuthorizationMiddleware) getAccessTokenFromHeader(c *gin.Context) (string, error) {
	authToken := c.Request.Header.Get("Authorization")
	authComps := strings.Split(authToken, " ")
	if len(authComps) != 2 || strings.ToLower(authComps[0]) != "bearer" {
		return "", errors.New("[getAccessTokenFromHeader] invalid authorization header")
	}

	return authComps[1], nil
}

func (m *AuthorizationMiddleware) CmsAllowUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get("nk_user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "[CmsAllowUsers] user id not found",
			})
			c.Abort()
			return
		}

		// in array
		allowUsers := []int{24}
		for _, v := range allowUsers {
			if id == v {
				c.Next()
				return
			}
		}
	}
}
