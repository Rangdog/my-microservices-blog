package middleware

import (
	"net/http"
	"strings"
	response "user-service/internal/pkg/Response"

	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(jwtSecret string) func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			authHeader:= r.Header.Get("Authorization")
			if authHeader == ""{
				response.Error(w, http.StatusUnauthorized, jwt.ErrTokenMalformed)
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
				if _, ok:=token.Method.(*jwt.SigningMethodHMAC); !ok{
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(jwtSecret), nil
			})
			if err != nil || !token.Valid{
				response.Error(w, http.StatusUnauthorized, jwt.ErrTokenInvalidId)
				return
			}

			next.ServeHTTP(w,r)
		})
	}
}