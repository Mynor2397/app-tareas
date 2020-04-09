package helpers

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"

	mod "github.com/Mynor2397/api-simple/src/models"
)

var (
	// PrivateKey es la clave privada
	PrivateKey *rsa.PrivateKey

	// PublicKey es la clave publica
	PublicKey *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./src/certificates/private.rsa")
	if err != nil {
		log.Fatal("Error al leer la llave privada: ", err)
	}

	publicBytes, err := ioutil.ReadFile("./src/certificates/public.rsa.pub")
	if err != nil {
		log.Fatal("Error al leer la llave privada: ", err)
	}

	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("Error al parsear la llave privada: ", err)
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("Error al parsear la llave publica: ", err)
	}
}

// GenerateJWT genera un token jwt
func GenerateJWT(user mod.User) string {
	claims := mod.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "marold97@outlook.com",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(PrivateKey)
	if err != nil {
		log.Fatal("El token no se pudo generar: ", err)
	}

	return result
}
