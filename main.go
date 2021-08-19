package main

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"log"
	"strconv"
	"time"
)

func main() {
	//var a = make(map[string]bool)
	//
	//a["ss"] = true
	MainKafka()
}

func getSessionNo(payType string, userId int64) string {
	var s bytes.Buffer
	if payType == "wx_lite" {
		s.WriteString("11")
	} else {
		s.WriteString("10")
	}
	s.WriteString(fmt.Sprintf("%d00000", time.Now().Unix()))

	var uid = fmt.Sprintf("%08v", userId)
	if len(uid) > 8 {
		uid = uid[len(uid)-8:]
	}
	s.WriteString(uid)
	return s.String()
}

func getTableName(orderNo string) string {
	uintNumber := crc32.ChecksumIEEE([]byte(orderNo))
	log.Print(uintNumber)
	tableNameFix := uintNumber % 32
	return fmt.Sprintf("%s_%d", "pay_order", tableNameFix)
	//return P.Consistent().Get(orderNo)
}

func GetTableNameBySessionNo(sessionNo string) string {
	uid := fmt.Sprintf("%08v", sessionNo[len(sessionNo)-8:])
	uidNumber, _ := strconv.ParseInt(uid, 10, 64)
	tableNameFix := uidNumber % 32

	return fmt.Sprintf("%s_%d", "pay_order", tableNameFix)
}
