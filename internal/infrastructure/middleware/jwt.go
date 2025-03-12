package middleware

import (
"errors"
"net/http"
"strings"
"time"

"github.com/gin-gonic/gin"
"github.com/golang-jwt/jwt/v5"
"web-server/internal/domain/constants"
"web-server/internal/infrastructure/config"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, role string) (string, error) {
	cfg := config.GetConfig()
	
	claims := JWTClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWTExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(cfg.JWTSecret)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
		c.JSON(http.StatusUnauthorized, constants.ErrAuthHeaderRequired)
		c.Abort()
		return
		}
		
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, constants.ErrInvalidAuthFormat)
		c.Abort()
		return
		}
		
		tokenString := bearerToken[1]
		claims := &JWTClaims{}
		
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected signing method")
		}
		return config.GetConfig().JWTSecret, nil
		})
		
		if err != nil {
		c.JSON(http.StatusUnauthorized, constants.ErrInvalidToken)
		c.Abort()
		return
		}
		
		if !token.Valid {
		c.JSON(http.StatusUnauthorized, constants.ErrInvalidToken)
		c.Abort()
		return
		}

		// Add claims to context
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
		c.JSON(http.StatusUnauthorized, constants.ErrRoleNotFound)
		c.Abort()
		return
		}
		
		roleStr := role.(string)
		allowed := false
		for _, r := range allowedRoles {
		if r == roleStr {
		allowed = true
		break
		}
		}
		
		if !allowed {
		c.JSON(http.StatusForbidden, constants.ErrInsufficientPermissions)
		c.Abort()
		return
		}

		c.Next()
	}
}