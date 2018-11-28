package session

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gbrlsnchs/jwt/v2"
)

//Token ..
type Token struct {
	jwt.JWT
}

//Encode ..
func Encode(secret string, claims *Token) (string, error) {
	hs256 := jwt.NewHS256(secret)
	claims.SetAlgorithm(hs256)
	rand.Seed(time.Now().UnixNano())
	claims.SetKeyID(strconv.Itoa(rand.Intn(10000)))
	payload, err := jwt.Marshal(claims)
	if err != nil {
		return "", err
	}
	token, err := hs256.Sign(payload)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

//Decode ..
func Decode(secret string, token string) (Token, error) {
	var claims Token
	hs256 := jwt.NewHS256(secret)
	payload, sig, err := jwt.Parse(token)
	if err != nil {
		return claims, err
	}
	if err = hs256.Verify(payload, sig); err != nil {
		return claims, err
	}
	if err = jwt.Unmarshal(payload, &claims); err != nil {
		return claims, err
	}
	return claims, nil
}

//New ...
func New(Issuer, Subject, Audience string) *Token {
	now := time.Now()
	claims := &Token{
		jwt.JWT{
			Issuer:         Issuer,
			Subject:        Subject,
			Audience:       Audience,
			ExpirationTime: now.Add(time.Hour * 24 * 7).Unix(),
			NotBefore:      now.Add(time.Minute * 30).Unix(),
			IssuedAt:       now.Unix(),
			ID:             "",
		},
	}
	return claims
}
