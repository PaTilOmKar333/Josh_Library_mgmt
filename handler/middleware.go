package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"project/models"
	"project/service"
	"strings"
)

func IsAuthorized(authorizedRole []string, authToken service.AuthTokenInterface, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authenticationResponse models.AuthenticationResponse

		var err error

		if r.Header["Authorization"] == nil {
			authenticationResponse.Message = "Please Login First"
			authenticationResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(authenticationResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}

		tokenString := r.Header["Authorization"][0]
		splitToken := strings.Split(tokenString, "Bearer ")
		reqToken := splitToken[1]

		claimss, err := authToken.ValidateToken(reqToken, authorizedRole)
		if err != nil {
			if strings.Contains(err.Error(), "token is expired") {
				authenticationResponse.Message = "Authentication Token is expired.Please Login again"
				authenticationResponse.StatusCode = http.StatusBadRequest
				res, _ := json.Marshal(authenticationResponse)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(res)
				return
			} else if strings.Contains(err.Error(), "illegal") {
				authenticationResponse.Message = "Token is illegal."
				authenticationResponse.StatusCode = http.StatusBadRequest
				res, _ := json.Marshal(authenticationResponse)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(res)
				return
			} else {
				log.Println(err)
				return
			}
		}

		roleFromClaims := claimss.Role
		for _, val := range authorizedRole {
			if val == roleFromClaims {
				handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "ClaimsToVerify", claimss)))
				return
			}
		}
		authenticationResponse.Message = "You are Unauthorized person"
		authenticationResponse.StatusCode = http.StatusBadRequest
		res, _ := json.Marshal(authenticationResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)

	}
}
