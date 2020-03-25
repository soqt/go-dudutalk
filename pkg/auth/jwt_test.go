package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TestAuth(t *testing.T) {
	secret := "APP_SECRET"
	accountId := "testId"
	gin.SetMode(gin.TestMode)

	t.Run("Test JWT Gen function", func(t *testing.T) {
		tokenString, err := JWTGen(secret, accountId, "")
		if err != nil {
			t.Errorf("expected a jwt string, got error %v", err)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(secret), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			accountId := claims["accountId"]
			if accountId != accountId {
				t.Errorf("token expected to have accountId claim 'testId', got %v", claims)
			}
		} else {
			t.Errorf("token expected, got %v", claims)
		}
	})

	t.Run("Test JWT Middleware if Authorization header is provided", func(t *testing.T) {
		resp := httptest.NewRecorder()
		c, r := gin.CreateTestContext(resp)
		r.Use(JWTMiddleware(false, secret))

		r.GET("/test", func(c *gin.Context) {
			accountId, found := c.Get("accountId")
			if accountId != accountId || !found {
				t.Errorf("accountId `testId` expected, got %v", accountId)
			}
			c.Status(200)
		})

		token, _ := JWTGen(secret, accountId, "")
		c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
		c.Request.Header.Add("Authorization", "Bearer "+token)
		r.ServeHTTP(resp, c.Request)
	})

	t.Run("Test JWT Middleware, fall through if Authorization header is missing", func(t *testing.T) {
		resp := httptest.NewRecorder()
		c, r := gin.CreateTestContext(resp)

		r.Use(JWTMiddleware(false, secret))

		r.GET("/test", func(c *gin.Context) {
			accountId, found := c.Get("account")
			if found {
				t.Errorf("accounId is not expected to be found, but got %v", accountId)
			}
			c.Status(200)
		})

		c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(resp, c.Request)
	})

}
