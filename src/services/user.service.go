package services

import (
	"context"
	"database/sql"
	"strings"

	"golang.org/x/crypto/bcrypt"

	help "github.com/Mynor2397/api-simple/src/helpers"
	mod "github.com/Mynor2397/api-simple/src/models"
	uti "github.com/Mynor2397/api-simple/src/utils"
)

// CreateUser nos permite crear un usario
func CreateUser(ctx context.Context, user mod.User) error {
	user.Email = strings.TrimSpace(user.Email)
	user.Name = strings.TrimSpace(user.Name)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := "insert into users (name, email, password, rol_idrol) values (?, ?, ?, ?);"
	_, err = db.QueryContext(ctx, query, user.Name, user.Email, string(hashedPassword), 2)
	if err != nil {
		return err
	}

	return uti.ErrUserOk
}

// Login servicio de consulta a la base de datos
func Login(ctx context.Context, user mod.User) (mod.ResponseToken, error) {
	var out mod.ResponseToken

	user.Email = strings.TrimSpace(user.Email)
	user.Name = strings.TrimSpace(user.Name)

	var passwordClient string
	query := "SELECT r.idrol, r.name as rol, u.idusers, u.name, u.email,  u.password FROM users u"
	query += " INNER JOIN rol r ON r.idrol = u.rol_idrol WHERE email = ?"

	row := db.QueryRowContext(ctx, query, user.Email).Scan(&user.IDRol, &user.NameRol, &user.ID, &user.Name, &user.Email, &passwordClient)

	if row == sql.ErrNoRows {
		return out, uti.ErrUserNotFound
	}

	if row != nil {
		return out, row
	}

	hashedPasswordDatabase := []byte(passwordClient)
	valuePassword := bcrypt.CompareHashAndPassword(hashedPasswordDatabase, []byte(user.Password))
	if valuePassword != nil {
		return out, uti.ErrUserNotFound
	}

	user.Password = ""
	out.ID = user.ID
	out.IDRol = user.IDRol
	out.Name = user.Name

	token := help.GenerateJWT(user)
	out.Token = token

	return out, nil

}
