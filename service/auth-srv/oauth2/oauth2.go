package oauth2

import (
	"log"

	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
)

//NewDefaultOsinServer
func NewDefaultOsinServer() *osin.Server {
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
	server := osin.NewServer(conf, NewTestStorage())
	return server
	//return osin.NewServer(serverConfig, repo.NewStorage(dbconn.DB.DB()))
}

//LoginPageHandler login page handle
type LoginPageHandler func(ar *osin.AuthorizeRequest, c *gin.Context) bool

//OAuth2 oauth engine
type OAuth2 struct {
	server *osin.Server
	//handerLoginPage func(ar *osin.AuthorizeRequest, w http.ResponseWriter, r *http.Request) bool
	loginPageHandler LoginPageHandler
}

//New ..
func New(s *osin.Server) *OAuth2 {
	o := &OAuth2{server: s}
	return o
}

//SetLoginPageHandler set login page handle
func (s *OAuth2) SetLoginPageHandler(h LoginPageHandler) {
	s.loginPageHandler = h
}

//InitRouter init oauth2 routes
func (s *OAuth2) InitRouter(r *gin.Engine) *OAuth2 {
	o2 := r.Group("/oauth2")
	o2.GET("/authorize", s.Authorize)
	o2.POST("/authorize", s.Authorize)
	o2.GET("/token", s.Token)

	//r.GET("/authorize",s.rest.authorize)
	return s
}

//Authorize authorize process
func (s *OAuth2) Authorize(c *gin.Context) {
	resp := s.server.NewResponse()
	defer resp.Close()
	if ar := s.server.HandleAuthorizeRequest(resp, c.Request); ar != nil {
		// HANDLE LOGIN PAGE HERE
		if !s.loginPageHandler(ar, c) {
			return
		}
		//ar.UserData = userId
		ar.Authorized = true
		s.server.FinishAuthorizeRequest(resp, c.Request, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		log.Printf("ERROR: %s\n", resp.InternalError)
	}
	if !resp.IsError {
		log.Printf("ERROR: %s\n", resp.InternalError)
		//resp.Output["custom_parameter"] = 42
	}
	osin.OutputJSON(resp, c.Writer, c.Request)
}

//Token token process
func (s *OAuth2) Token(c *gin.Context) {
	resp := s.server.NewResponse()
	defer resp.Close()

	if ar := s.server.HandleAccessRequest(resp, c.Request); ar != nil {
		ar.Authorized = true
		s.server.FinishAccessRequest(resp, c.Request, ar)
	}
	osin.OutputJSON(resp, c.Writer, c.Request)
}
