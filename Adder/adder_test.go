package Adder

import (
	"errors"
	"fmt"
	"testing"
)

func testSub() error {
	v1 := 10
	v2 := -10
	if res := Adder(v1, -v2); res != v1-v2 {
		return errors.New(fmt.Sprintf("%d + %d ==> %v", v1, v2, res))
	}
	v1 = 9
	v2 = 60
	if res := Adder(v1, -v2); res != v1-v2 {
		return errors.New(fmt.Sprintf("%d + %d ==> %v", v1, v2, res))
	}
	v1 = 400
	v2 = 20
	if res := Adder(v1, -v2); res != v1-v2 {
		return errors.New(fmt.Sprintf("%d + %d ==> %v", v1, v2, res))
	}
	v1 = 2100
	v2 = 37
	if res := Adder(v1, -v2); res != v1-v2 {
		return errors.New(fmt.Sprintf("%d + %d ==> %v", v1, v2, res))
	}
	return nil
}

func testAdding() error {
	v1 := 10
	v2 := 10
	if res := Adder(v1, v2); res != v1+v2 {
		return errors.New(fmt.Sprintf("%d + %d ==> %v", v1, v2, res))
	}
	v1 = 9
	v2 = 60
	if res := Adder(v1, v2); res != v1+v2 {
		return errors.New(fmt.Sprintf("%d + %d ==> %v", v1, v2, res))
	}
	v1 = 400
	v2 = 20
	if res := Adder(v1, v2); res != v1+v2 {
		return errors.New(fmt.Sprintf("%d + %d ==> %v", v1, v2, res))
	}
	v1 = 2100
	v2 = 37
	if res := Adder(v1, v2); res != v1+v2 {
		return errors.New(fmt.Sprintf("%d + %d ==> %v", v1, v2, res))
	}
	return nil
}

func TestAdder(t *testing.T) {

	if err := testAdding(); err != nil {
		t.Errorf("error in adding %v", err)
	}
	if err := testSub(); err != nil {
		t.Errorf("error in adding %v", err)
	}
}
