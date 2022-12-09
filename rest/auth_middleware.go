package rest

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"golang-rest/contextkeys"

	"firebase.google.com/go/auth"
)

// AuthMiddleware is the first step in handling API requests used to authenticate calls.
func AuthMiddleware(authClient *auth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			token := strings.Replace(authHeader, "Bearer ", "", 1)
			t, err := authClient.VerifyIDToken(context.Background(), token)
			if err != nil {
				// client.RespondError(w, http.StatusUnauthorized, fmt.Sprintf("Auth not Ok: %v", err.Error()))
				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextkeys.UsernName, "userID")))
				return
			}

			var userID string
			if username, found := t.Claims["name"]; found {
				userID = fmt.Sprintf("%v", username)
			} else {
				userID = t.UID
			}

			ctx := context.WithValue(r.Context(), contextkeys.UsernName, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
