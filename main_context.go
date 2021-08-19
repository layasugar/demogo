package main

import (
	"context"
	"fmt"
	"log"
)

type ctxTest struct {
	UserId int64
	Sid    int64
}

func (c ctxTest) String() string {
	return fmt.Sprintf("%d,%d", c.UserId, c.Sid)
}

func MainContext() {
	var ctx = context.WithValue(context.Background(), "sdasdasdasdasda", ctxTest{UserId: 1, Sid: 2})

	a := fmt.Sprintf("%s", ctx)
	log.Print(a)
}
