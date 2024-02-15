package config

import (
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var TokenAuth *jwtauth.JWTAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil, jwt.WithAcceptableSkew(72*time.Hour))
