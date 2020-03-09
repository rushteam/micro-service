package session

import (
	"time"

	jwt "github.com/gbrlsnchs/jwt/v3"
	//github.com/dgrijalva/jwt-go
)

var _secret = "861b1508cb81764b65fa90e460dbf1c1"

//Token ..
type Token struct {
	jwt.Payload
}

//Encode ..
func Encode(claims *Token, secret string) (string, error) {
	if secret == "" {
		secret = _secret
	}
	hs := jwt.NewHS256([]byte(secret))
	token, err := jwt.Sign(claims, hs)
	return string(token), err
}

//Decode ..
func Decode(token, secret string) (Token, error) {
	if secret == "" {
		secret = _secret
	}
	var claims Token
	hs := jwt.NewHS256([]byte(secret))
	_, err := jwt.Verify([]byte(token), hs, &claims)
	return claims, err
}

/*
New ...
	iss(Issuser)：代表这个JWT的签发主体；
	sub(Subject)：代表这个JWT的主体，即它的所有人；
	aud(Audience)：代表这个JWT的接收对象；
	exp(Expiration time)：是一个时间戳，代表这个JWT的过期时间；
	nbf(Not Before)：是一个时间戳，代表这个JWT生效的开始时间，意味着在这个时间之前验证JWT是会失败的；
	iat(Issued at)：是一个时间戳，代表这个JWT的签发时间；
	jti(JWT ID)：是JWT的唯一标识
*/
func New(Issuer, Subject string, Audience string) *Token {
	now := time.Now()
	claims := &Token{
		Payload: jwt.Payload{
			Issuer:         Issuer,
			Subject:        Subject,
			Audience:       jwt.Audience{Audience},
			ExpirationTime: jwt.NumericDate(now.Add(time.Hour * 24 * 7)),
			NotBefore:      jwt.NumericDate(now),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "",
		},
	}
	return claims
}
