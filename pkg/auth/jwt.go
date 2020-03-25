package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// CheckPassword accept a password hash as first arg and raw password as second arg
// 系统暂无使用Password
//func CheckPassword(curPasswordHash, password string) error {
//	if len(password) == 0 {
//		return fmt.Errorf("password should not be empty")
//	}
//	bytePassword := []byte(password)
//	// Make sure the second param `bcrypt generator cost` between [4, 32)
//	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
//	byteHashedPassword := []byte(curPasswordHash)
//	return bcrypt.CompareHashAndPassword(passwordHash, byteHashedPassword)
//}

// Strips 'BEARER ' prefix from token string
func stripBearerPrefixFromTokenString(token string) (string, error) {
	// Should be a bearer token
	if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER " {
		return token[7:], nil
	}
	return token, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "BEARER " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var TokenExtractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// JWTMiddleware解析HTTP Header中的"Authorization"值或url中的access_token(OAuth2)
// 并将解析出来的值写入user
func JWTMiddleware(credentialsRequired bool, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, TokenExtractor, func(token *jwt.Token) (interface{}, error) {
			// validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil && credentialsRequired {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		if token != nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				for k := range claims {
					value := claims[k]
					c.Set(k, value)
				}
			}
		}
		return
	}
}

func JWTGen(secret, accountId string, customerId string) (string, error) {
	if secret == "" || (accountId == "" && customerId == "") {
		fmt.Errorf("secret %v or ids %v must be not empty", secret, accountId)
	}

	type MyCustomClaims struct {
		AccountId  string `json:"accountId"`
		CustomerId string `json:"customerId"`
		jwt.StandardClaims
	}

	iat := time.Now()
	//expireTime := iat.Add(duration).Unix()
	claims := MyCustomClaims{
		accountId,
		customerId,
		jwt.StandardClaims{
			Issuer:   "graphQL",
			IssuedAt: iat.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(secret))
}
