package jwt

import (
	"time"

	"github.com/MiguelAngelderobles/microblog/models"
	jwt "github.com/dgrijalva/jwt-go" //Creo un alias para manejarlo más fácil
)

func GeneroJWT(usu models.Usuario) (string, error) {
	miClave := []byte("SkillFactoryGo_Avalith")

	payload := jwt.MapClaims{
		"email":            usu.Email,
		"nombre":           usu.Nombre,
		"apellidos":        usu.Apellidos,
		"fecha_nacimiento": usu.FechaNacimiento,
		"biografia":        usu.Biografia,
		"ubicacion":        usu.Ubicacion,
		"sitioWeb":         usu.SitioWeb,
		"_id":              usu.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
