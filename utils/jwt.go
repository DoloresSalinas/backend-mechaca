package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
    "fmt"
    
)

var claveSecreta = []byte("clave_secreta")

func CrearToken(id int, nombre string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": id,  
        "nombre":   nombre,
        "exp": time.Now().Add(10 * time.Minute).Unix(), 
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(claveSecreta)
}

func ValidarToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
        }
        return claveSecreta, nil
    })

    if err != nil {
        return nil, err
    }

    return token, nil
}
