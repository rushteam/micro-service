module github.com/rushteam/micro-service

go 1.13

require (
	github.com/RangelReale/osin v1.0.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gbrlsnchs/jwt/v3 v3.0.0-rc.1
	github.com/gin-gonic/gin v1.5.0
	github.com/go-log/log v0.2.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/google/go-querystring v1.0.0
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro/v2 v2.6.0
	github.com/micro/micro/v2 v2.2.0 // indirect
	github.com/pborman/uuid v1.2.0
	github.com/rushteam/gosql v0.0.0-20200305224957-f3ad7954ff89
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/rushteam/gosql => ../gosql
