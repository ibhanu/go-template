package middleware

import (
	"net/http"
	"strings"
	"time"

	"web-server/internal/domain/constants"
	"web-server/internal/infrastructure/config"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	*jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"` // Access token expiration in seconds
}

func GenerateTokenPair(userID, role string) (*TokenPair, error) {
	cfg := config.GetConfig()

	// Generate access token
	accessClaims := JWTClaims{
		UserID:    userID,
		Role:      role,
		TokenType: "access",
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWTExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(cfg.JWTSecret)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshClaims := JWTClaims{
		UserID:    userID,
		Role:      role,
		TokenType: "refresh",
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWTRefreshExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(cfg.JWTRefreshSecret)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    int64(cfg.JWTExpiration.Seconds()),
	}, nil
}

func RefreshToken(refreshTokenString string) (*TokenPair, error) {
	cfg := config.GetConfig()
	claims := &JWTClaims{}

	// Parse and validate refresh token
	token, err := jwt.ParseWithClaims(refreshTokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, constants.ErrUnexpectedSigningMethod
		}
		return cfg.JWTRefreshSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, constants.ErrInvalidRefreshToken
	}

	if claims.TokenType != "refresh" {
		return nil, constants.ErrInvalidTokenType
	}

	// Generate new token pair
	return GenerateTokenPair(claims.UserID, claims.Role)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, constants.ErrAuthHeaderRequired())
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, constants.ErrInvalidAuthFormat())
			c.Abort()
			return
		}

		tokenString := bearerToken[1]
		claims := &JWTClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, constants.ErrUnexpectedSigningMethod
			}
			return config.GetConfig().JWTSecret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, constants.ErrInvalidToken())
			c.Abort()
			return
		}

		if !token.Valid || claims.TokenType != "access" {
			c.JSON(http.StatusUnauthorized, constants.ErrInvalidToken())
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
			c.JSON(http.StatusUnauthorized, constants.ErrRoleNotFound())
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, constants.ErrInternalServer())
			c.Abort()
			return
		}

		allowed := false
		for _, r := range allowedRoles {
			if r == roleStr {
				allowed = true
				break
			}
		}

		if !allowed {
			c.JSON(http.StatusForbidden, constants.ErrInsufficientPermissions())
			c.Abort()
			return
		}

		c.Next()
	}
}
