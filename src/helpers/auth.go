package helpers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	mod "github.com/Mynor2397/api-simple/src/models"
)

// AuthUserID autenticar al usuario
func AuthUserID(r *http.Request) (int64, string, error) {
	claims := mod.Claim{}

	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, claims, func(token *jwt.Token) (interface{}, error) {
		return PublicKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			typeError := err.(*jwt.ValidationError)
			switch typeError.Errors {
			case jwt.ValidationErrorExpired:
				return 0, "", err
			case jwt.ValidationErrorSignatureInvalid:
				return 0, "", err
			}
		default:
			return 0, "", err

		}
	}

	var (
		id      int64
		rolname string
	)

	if token.Valid {
		id = claims.ID
		rolname = claims.NameRol
	}

	return id, rolname, nil
}
