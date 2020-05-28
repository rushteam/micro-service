module github.com/rushteam/micro-service

go 1.13

require (
	github.com/RangelReale/osin v1.0.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gbrlsnchs/jwt/v3 v3.0.0-rc.2
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/google/go-querystring v1.0.0
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.7.0
	github.com/pborman/uuid v1.2.0
	github.com/rushteam/gosql v0.0.0-00010101000000-000000000000
	github.com/yuin/goldmark v1.1.31 // indirect
	golang.org/x/mod v0.3.0 // indirect
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2 // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/tools v0.0.0-20200526224456-8b020aee10d2 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/rushteam/gosql => ../gosql
