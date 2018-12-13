package utils

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
	"reflect"
	"sort"
	"strconv"
)

//SignSHA1 sha1签名
func SignSHA1(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		io.WriteString(h, s)
	}
	return hex.EncodeToString(h.Sum(nil))
	// return fmt.Sprintf("%x")
}

//SignMD5 md5签名
func SignMD5(frist, laster string, params ...string) string {
	sort.Strings(params)
	h := md5.New()
	io.WriteString(h, frist)
	for _, s := range params {
		io.WriteString(h, s)
	}
	io.WriteString(h, laster)
	return hex.EncodeToString(h.Sum(nil))
	// return fmt.Sprintf("%x")
}

// Sign 微信支付签名.
//  params: 待签名的参数集合
//  apiKey: api密钥
//  h:      hash.Hash, 如果为 nil 则默认用 md5.New(), 特别注意 h 必须是 initial state.
func Sign(params map[string]string, first, laster string, h hash.Hash) string {
	if h == nil {
		h = md5.New()
	}
	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var s string
	bufw := bufio.NewWriterSize(h, 128)
	s += first
	bufw.WriteString(first)
	for i, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		bufw.WriteString(k)
		bufw.WriteByte('=')
		bufw.WriteString(v)
		s += k + "=" + v
		if i < len(keys)-1 {
			s += "&"
			bufw.WriteByte('&')
		}
	}
	s += laster
	//fmt.Println(s)
	bufw.WriteString(laster)
	bufw.Flush()
	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}

//Struct2Map ...
func Struct2Map(obj interface{}, tagName string) map[string]string {
	//t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	//if t.Kind() == reflect.Ptr {
		//t = t.Elem()
		v = reflect.Indirect(v)
	//}
	t := v.Type()
	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		var k string
		if tagName != "" {
			k = t.Field(i).Tag.Get(tagName)
		} else {
			k = t.Field(i).Name
		}
		vv := v.Field(i).Interface()
		switch vv.(type) {
		case string:
			// data[t.Field(i).Name] = vv.(string)
			data[k] = vv.(string)
		case int:
			data[k] = strconv.Itoa(vv.(int))
		case int64:
			data[k] = strconv.FormatInt(vv.(int64), 10)
		}
	}
	return data
}
