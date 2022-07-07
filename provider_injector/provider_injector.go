package provider_injector

import (
	"context"
	"errors"
)

type Student struct {
	ClassNo int
}

func NewStudent() Student {
	return Student{ClassNo: 10}
}

type Class struct {
	ClassNo int
}

func NewClass(stu Student) Class {
	return Class{ClassNo: stu.ClassNo}
}

type School struct {
	ClassNo int
}

func NewSchool(ctx context.Context, class Class) (School, error) {
	if class.ClassNo == 0 {
		return School{}, errors.New("cannot provider school when class is 0")
	}
	return School{ClassNo: class.ClassNo}, nil
}
