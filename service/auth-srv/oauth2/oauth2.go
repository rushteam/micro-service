package oauth2

import(
	"github.com/gin-gonic/gin"
	"github.com/RangelReale/osin"
	"github.com/RangelReale/osin/example"
	"log"
)

//InitOAuthServer ..
func InitOAuthServer(r *gin.RouterGroup) *osin.Server {
	conf := osin.NewServerConfig()
	conf.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	conf.AllowedAccessTypes = osin.AllowedAccessType{
		osin.AUTHORIZATION_CODE,
		osin.REFRESH_TOKEN,
		osin.PASSWORD,
		osin.CLIENT_CREDENTIALS,
		osin.ASSERTION,
	}
	conf.AllowGetAccessRequest = true
	conf.AllowClientSecretInParams = true
	//todo 这里要实现落地存储 最好是用micro service方式
	server := osin.NewServer(conf, example.NewTestStorage())
	return server
	//return osin.NewServer(serverConfig, repo.NewStorage(dbconn.DB.DB()))
}

func Routers() {
	resp := r.server.NewResponse()
	defer resp.Close()
	if ar := r.server.HandleAuthorizeRequest(resp, c.Request); ar != nil {
		userId, ok := handleLoginPage(r, ar, c)
		if !ok {
			return
		}
		ar.UserData = userId
		ar.Authorized = true
		r.server.FinishAuthorizeRequest(resp, c.Request, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		log.Printf("ERROR: %s\n", resp.InternalError)
	}
	if !resp.IsError {
		//resp.Output["custom_parameter"] = 42
	}
	osin.OutputJSON(resp, c.Writer, c.Request)
}
