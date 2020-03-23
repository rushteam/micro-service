# Api Doc

## UserService.Signup 注册

req

```
#micro cli
micro call go.micro.srv.usersrv UserService.Signup '{"nickname":"菜饼子","phone":"183","password":"test"}'
```

rsp
```
{
	"uid": "27",
	"nickname": "菜饼子",
	"status": 1,
	"updatedAt": "2020-03-06 04:59:29",
	"createdAt": "2020-03-06 04:59:29"
}
```

## UserService.Signin 登陆

req

```
#micro cli
micro call go.micro.srv.usersrv UserService.Signin '{"loginname":"18310497688","password":"test"}'
```

rsp
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLXNydiIsInN1YiI6IjEiLCJhdWQiOiIiLCJleHAiOjE1ODQxMzAxNDQsIm5iZiI6MTU4MzUyNTM0NCwiaWF0IjoxNTgzNTI1MzQ0fQ.rwnFiE2DrEukl97V_ozRsaSOH0WZkNFkasAhnLaGHYo",
	"uid": "1"
}
```

## UserService.OAuthAuthorize 第三方登陆

req

```
#micro cli
micro call go.micro.srv.usersrv UserService.OAuthAuthorize '{"platform":"wx","appid":"test"}'
```

rsp
```
```