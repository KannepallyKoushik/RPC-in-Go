package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"time"
)

func main() {

	for true {
		var reply string
		var response string
		var otp string
		status := "active"
		client, err := rpc.DialHTTP("tcp", "192.168.0.114:4040")

		if err != nil {
			log.Fatal("Connection error: ", err)
		}
		fmt.Println("Welcome to OTP Application")
		fmt.Println("Options :- \n 1.Request for OTP 2.End Transaction")
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			fmt.Println("Your Request for Otp has been sent, Waiting for Response")
			client.Call("API.GetOTP", "", &reply)

			fmt.Println("Your otp is: ", reply)

			afterFuncTimer := time.AfterFunc(time.Second*20, func() {
				//fmt.Println("Your time is completed")
				status = "expired"
			})
			defer afterFuncTimer.Stop()
			fmt.Println("You now have 20 seconds to enter otp")

			fmt.Println("\n Enter the otp you recieved")
			fmt.Scanln(&otp)
			rotp := otp + "-" + status

			client.Call("API.ValidateOTP", rotp, &response)
			fmt.Println("Validating your OTP")

			fmt.Println(response)
			fmt.Println()
		case 2:
			fmt.Println()
			fmt.Println("Thankyou for using our Application")
			os.Exit(3)
		default:
			fmt.Println("Choose from the options given")

		}
	}
}
