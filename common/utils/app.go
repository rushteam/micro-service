package utils

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//CurrentRoot 当前程序所在目录
func CurrentRoot() string {
	//获取路径的方法 此种方法获取的是程序所在路径 而不是程序运行路径
	execpath, err := os.Executable() // 获得程序路径
	if err != nil {
		log.Panic("获取执行路径错误")
	}
	return filepath.Dir(execpath)
}

//SnakeString  转 snake_string
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

//CamelString  camel_string to CamelString
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

//RandomStr 随机生成字符串
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//FormatDate 格式化时间
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
