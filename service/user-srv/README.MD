# Api Doc

## UserService.Signin 登陆

```micro cli
micro call go.micro.srv.usersrv UserService.Signin '{"loginname":"18310497688","password":"test"}'
```

## UserService.Signup 注册

req

```micro cli
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
