package main

import (
	"fmt"
	"hash/crc32"
)

func main_order()  {
	fmt.Print(GetTableName("232807136219970"))
}

func  GetTableName(orderNo string) string {
	uintNumber := crc32.ChecksumIEEE([]byte(orderNo))
	tableNameFix := uintNumber % 32
	return fmt.Sprintf("%s_%d", "pay_order", tableNameFix)
	//return p.Consistent().Get(orderNo)
}