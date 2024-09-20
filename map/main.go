package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type serviceErr struct {
	service ErrorIdentity
	err     error
}

type ChangeFundnameProcessor func(old, new, scheduleRef string, ch chan serviceErr) (bool, error)

func CallChangeFundcodePortService(old, new, scheduleRef string, ch chan serviceErr) (bool, error) {
	err := errors.New("this error")
	fmt.Println("Port Service Error")
	if ch != nil {
		ch <- serviceErr{
			service: PortService,
			err:     err,
		}
	}
	return true, err
}
func CallChangeFundcodeFundService(old, new, scheduleRef string, ch chan serviceErr) (bool, error) {
	// err := errors.New("this Fundservice error")
	if ch != nil {
		ch <- serviceErr{
			service: FundService,
			err:     nil,
		}
	}

	return true, nil
}
func CallChangeFundcodeRoboService(old, new, scheduleRef string, ch chan serviceErr) (bool, error) {
	if ch != nil {
		ch <- serviceErr{
			service: RoboService,
			err:     nil,
		}
	}
	return true, nil

}
func CallChangeFundcodePlanService(old, new, scheduleRef string, ch chan serviceErr) (bool, error) {
	if ch != nil {
		ch <- serviceErr{
			service: PlanService,
			err:     nil,
		}
	}
	return true, nil

}

type ErrorServices struct {
	PortService *string `json:"port_service,omitempty"`
	FundService *string `json:"fund_service,omitempty"`
}

func (e *ErrorServices) addError(service ErrorIdentity, err error) *ErrorServices {
	if err == nil {
		return e
	}
	if e == nil {
		e = &ErrorServices{}
	}
	_err := err.Error()
	switch service {
	case PortService:
		e.PortService = &_err
	case FundService:
		e.FundService = &_err
	}
	fmt.Printf("Added error: %#v\n", e)
	return e
}

type ErrorIdentity int

const (
	PortService ErrorIdentity = iota
	FundService
	RoboService
	PlanService
)

const a = "test"

// const c = [2]int{1,2}
var CFPs = []ChangeFundnameProcessor{
	CallChangeFundcodePortService,
	CallChangeFundcodeFundService,
	CallChangeFundcodeRoboService,
	CallChangeFundcodePlanService,
}

func main() {
	// u8max := utf8.UTFMax
	// uniReplace := unicode.ReplacementChar
	// var a []interface{}
	// var c interface{}
	// err := errors.New("prefix" + "error_msg")
	// b := []int{1, 2, 3}

	// c = b

	// a = []interface{}{b}

	// fmt.Printf("C: %v\n", c)
	// fmt.Printf("A: %v\n", a)

	// CFPs = append(CFPs, CallChangeFundcodeTransferService())
	// CFPs[1] = CallChangeFundcodeFundService
	// wg.Add(len(process))
	// for _, process := range CFPs {
	// 	go process("A", "B", "134")
	// }

	var err *ErrorServices

	// err.FundService = "test"
	// if b, _ := json.Marshal(err); len(b) > 0 {
	// 	fmt.Println(string(b))
	// }

	errCh := make(chan serviceErr, len(CFPs))

	for _, fn := range CFPs {
		go fn("a", "b", "1234", errCh)
	}
	for i := range CFPs {
		fmt.Println(i)
		e := <-errCh
		err = err.addError(e.service, e.err)
	}
	fmt.Printf("After Added: %#v\n", err)
	if b, _ := json.Marshal(err); len(b) > 0 {
		fmt.Println(string(b))
	}
	if err.PortService != nil {

	}

	testErr := ErrorServices{}
	json.Unmarshal([]byte(`{"port_service": "port error"}`), &testErr)
	fmt.Printf("Result: %#v", testErr)
}
