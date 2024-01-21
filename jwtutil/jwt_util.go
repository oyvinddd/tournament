package jwtutil

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

const (
	// secret is used for signing the newly generated jwt token
	secretJwtKey string = "Nw8c3ctiJlgEYb56AqyqRhmH8a5KSfsX" // FIXME: don't hardcode

	// the ttl (in minutes) of a jwt access token
	accessTokenExpiration time.Duration = time.Minute * 15

	// the ttl (in hours) for a jwt refresh token
	refreshTokenExpiration time.Duration = time.Hour * 72

	// expiration date of the created jwt token (unix timestamp)
	jwtClaimKeyExp string = "exp"

	// subject claim stored in the jwt token (in this case account id)
	jwtClaimKeySub string = "sub"

	// type claim stores the account type (admin or regular)
	jwtClaimKeyTyp string = "typ"

	// authorization http header key
	authHeader string = "Authorization"
)

var (
	errUnauthorized   = errors.New("unauthorized")
	errMissingToken   = errors.New("missing access token")
	errMalformedToken = errors.New("malformed token")
	// TODO: we need a way of distinguishing this error on the client, since we need to take special action
	errExpiredToken = errors.New("expired token")
)

func GenerateAccessAndRefreshTokens(sub string, typ uint) (string, string, error) {
	accessToken, err := GenerateToken(sub, typ, accessTokenExpiration)
	// TODO: should not add claims to refresh token?
	refreshToken, err := GenerateToken(sub, typ, refreshTokenExpiration)
	return accessToken, refreshToken, err
}

func GenerateToken(sub string, typ uint, exp time.Duration) (string, error) {
	// a map of the claims we want to include in the token
	claims := jwt.MapClaims{
		jwtClaimKeyExp: time.Now().Add(exp).Unix(),
		jwtClaimKeySub: sub,
		jwtClaimKeyTyp: typ,
	}
	// create a new token object, specifying the signing method
	// and the claims you would like it to contain
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretJwtKey))
}

func ValidateToken(token string) (string, uint, error) {
	sub, typ, err := parseAndValidateToken(token)
	if err != nil {
		return "", 0, err
	}
	// cast subject to string and return it
	subStr, _ := sub.(string)
	// cast type to int and return it
	typInt, _ := typ.(uint)
	return subStr, typInt, nil
}

func ValidateTokenFromRequest(r *http.Request) (interface{}, interface{}, error) {
	tokenStr, err := extractTokenFromRequestHeader(r.Header)
	if err != nil {
		return nil, nil, err
	}
	return parseAndValidateToken(*tokenStr)
}

func extractTokenFromRequestHeader(header http.Header) (*string, error) {
	authHeader := strings.Split(header.Get(authHeader), "Bearer ")
	if len(authHeader) != 2 {
		return nil, errMissingToken
	}
	return &authHeader[1], nil
}

func parseAndValidateToken(tokenStr string) (interface{}, interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errUnauthorized
		}
		return []byte(secretJwtKey), nil
	})
	if !token.Valid && err != nil {
		return nil, nil, err //mappedError(err)
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims[jwtClaimKeySub], claims[jwtClaimKeyTyp], nil
}

/*
func mappedError(err error) error {
	if ve, ok := err.(*jwt.ErrTokenMalformed); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return errMalformedToken
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return errExpiredToken
		}
	}
	return errUnauthorized
}
*/
