package main

import (
	"crypto/rand"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
	"time"
)

type API int

var otp string

func main() {
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}
}

func (a *API) GetOTP(empty string, reply *string) error {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, 6)
	n, err := io.ReadAtLeast(rand.Reader, b, 6)
	if n != 6 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	otp = string(b)

	// Giving little delay
	time.Sleep(time.Second * 5)

	*reply = otp
	return nil
}

func (a *API) ValidateOTP(rotp string, reply *string) error {
	var response string

	req := strings.Split(rotp, "-")

	if req[1] == "active" {
		if req[0] == otp {
			response = "Transaction Successfull"
		} else {
			response = "OTP Not Valid"
		}
	} else {
		response = "Sorry your otp has expired"
	}
	time.Sleep(time.Second * 5)
	*reply = response
	return nil
}
