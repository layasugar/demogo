package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

const (
	pubKeyPrefix = "-----BEGIN PUBLIC KEY-----"      //公钥前缀
	pubKeySuffix = "-----END PUBLIC KEY-----"        //公钥后缀
	prvKeyPrefix = "-----BEGIN RSA PRIVATE KEY-----" //私钥前缀
	prvKeySuffix = "-----END RSA PRIVATE KEY-----"   //私钥后缀
	PubKeyType   = 1                                 //公钥类型
	PrvKeyType   = 2                                 //私钥类型
	InputBase64  = 1                                 //base64输入
	InputOrigin  = 2                                 //原样输入
	OutPutBase64 = 1                                 //base64输出
	OutPutOrigin = 2                                 //原样输出
)

//标准格式化公钥私钥
func RsaParseKey(key string, keyType int, input int, output int) (string, error) {
	var strPure string
	switch input {
	case InputBase64:
		strBytes, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			return "", err
		}
		strPure = string(strBytes)
	case InputOrigin:
		strPure = key
	default:
		return "", errors.New("not support input")
	}
	if strPure == "" {
		return "", errors.New("key is empty")
	}

	//1.去掉换行符"\n"或者"\r\n"或者"\r"
	strPure = strings.Replace(strPure, "\n", "", -1)
	strPure = strings.Replace(strPure, "\r", "", -1)
	//2.匹配首尾符,去掉首尾符
	if keyType == PubKeyType {
		strPure = strings.TrimPrefix(strPure, pubKeyPrefix)
		strPure = strings.TrimSuffix(strPure, pubKeySuffix)
	}
	if keyType == PrvKeyType {
		strPure = strings.TrimPrefix(strPure, prvKeyPrefix)
		strPure = strings.TrimSuffix(strPure, prvKeySuffix)
	}

	//3.去掉空格
	strPure = strings.Replace(strPure, " ", "", -1)
	//4.按照指定的长度,对字符串进行换行
	str := WarpWord(strPure, "\n", 64)
	//5.加上首尾符
	if keyType == PubKeyType {
		str = fmt.Sprintf("%s\n%s%s", pubKeyPrefix, str, pubKeySuffix)
	} else {
		str = fmt.Sprintf("%s\n%s%s", prvKeyPrefix, str, prvKeySuffix)
	}
	//6.返回
	switch output {
	case OutPutBase64:
		return base64.StdEncoding.EncodeToString([]byte(str)), nil
	case OutPutOrigin:
		return str, nil
	default:
		return "", errors.New("not support output")
	}
}

//按指定位数拆分支付串,并填充指定字符串
func WarpWord(s string, sep string, len int) string {
	var returnStr string
	var str string
	for i, r := range s {
		str = str + string(r)
		fmt.Printf("i%d r %c\n", i, r)
		if i > 0 && (i+1)%len == 0 {
			fmt.Printf("=>(%d) '%v'\n", i, str)
			returnStr += str + sep
			str = ""
		}
	}
	returnStr += str + sep
	//fmt.Println(str)
	return returnStr
}
