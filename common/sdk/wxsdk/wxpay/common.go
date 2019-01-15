package wxpay

import (
	"bytes"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"hash"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const mchURL = "https://api.mch.weixin.qq.com"
const (
	//Success 成功
	Success = "SUCCESS"
	//Fail 失败
	Fail = "FAIL"
)

//IRequest ..
type IRequest interface {
	URL() string
}

//IResponse ..
type IResponse interface {
	Error() error
}

//Request ..
func Request(req IRequest, rsp IResponse) error {
	data, err := xml.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := http.Post(req.URL(), "application/xml", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = xml.NewDecoder(resp.Body).Decode(rsp)
	if err != nil {
		return err
	}
	return rsp.Error()
}

//Response ..
func Response(raw []byte, rsp IResponse) error {
	err := xml.Unmarshal(raw, rsp)
	if err != nil {
		return err
	}
	return rsp.Error()
}

//Sign ..
func Sign(args interface{}, secret string, h hash.Hash) string {
	h.Reset()
	params, err := query.Values(args)
	if err != nil {
		return ""
	}
	querystr, err := url.QueryUnescape(params.Encode())
	if err != nil {
		return ""
	}
	h.Write([]byte(querystr))
	h.Write([]byte(fmt.Sprintf("&key=%s", secret)))
	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}

//NotifyReplySuccess ..
func NotifyReplySuccess() string {
	return `<xml>
	<return_code><![CDATA[SUCCESS]]></return_code>
	<return_msg><![CDATA[OK]]></return_msg>
	</xml>`
}

//NotifyReplyFail ..
func NotifyReplyFail(msg string) string {
	return fmt.Sprintf(`<xml>
	<return_code><![CDATA[FAIL]]></return_code>
	<return_msg><![CDATA[%s]]></return_msg>
	</xml>`, msg)
}
