package mock

import (
	"api-gateway/app/domain/dao"
	"api-gateway/app/domain/dco"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateMockToken(user dao.User) (string, error) {
	convertedPermissions := make([]string, len(user.Permissions))
	for i, v := range user.Permissions {
		convertedPermissions[i] = v.ConvertToID()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dco.JWTClaim{
		Username:    user.Username,
		Department:  user.Department.ID.String(),
		Permissions: convertedPermissions,
		IsAdmin:     user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(dco.JWTExpirationTime)),
		},
	})

	return token.SignedString([]byte("secret"))
}
