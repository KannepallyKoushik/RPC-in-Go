# RPC-in-Go
1. Demonstrate client sending remote procedure request and server returning the result to the client stub. 
2. Implement one of the variations and demonstrate the same. Demonstrate the differences between the chosen variation and basic RPC.



1.	File Structure :
Assignment_3 rpc
	Transcient-Synchronious Communication(Variation 1)
		Client
			sync_client.go
		Server.go

	Transcient-Asynchrounous Communication (Variation 2)
		Client
			async_client.go
		Server.go
README.docx
	-------------------------------------------------------------------------------------------------------------------
Server :-
 It is a simple server which handles Http request and serves on port No- 4040
Packages used – net/http , net/rpc , crypto/rand ,io
Procedures Defined – GetOTP()	-> Gets empty string and pointer to reply message as Parameters. Returns out the otp generated randomly
– ValidateOTP() -> Gets the otp for validation along with its status(expired or   active). Returns Validation result to the Client

Built-In Methods Used from RPC Package- 
i.	Register() -> Accepts an Interface as Parameters - > Stores all Suitable methods into the interface depending on type of interface we pass



It adds up another interface which stores the sent interface type,value,name, suitable methods to the client which they can call remotely to access procedures that are written inside the server
Ex:- Call() method




Clients :-
Clients make rpc calls to the procedures running on server by sending http requests and access them remotely.
Pakages used – net/rpc , time, fmt
Built-In Methods Used from RPC Package-
i.	DailHTTP(“network” , IP_of_Server: port_NO) - >
This function in-turn calls – Dail() function from “net” package-> which returns out connection object
This connection object has all the suitable methods required to invoke the procedures that are written at server side which is running at some port number as a process


Now we can send HTTP requests to the server to which we are connected by this connection object
client, err := rpc.DialHTTP("tcp", "192.168.0.114:4040")

Since we are storing that connection object inside the client variable, We are now able to call the procedures that are written as server side using Call() method and store responses.
Ex:-
client.Call("API.GetOTP", "", &reply)
client.Call("API.ValidateOTP", rotp, &response)


Instructions to Execute:-
	Go to the directory where server.go is present in Terminal and type the command go run server.go
	Open the Client program in any editor.
	At line number 18, Change the IP addr with the IP of the system on which you are running the server.go
Ex:- client, err := rpc.DialHTTP("tcp", "192.168.0.114:4040")

	Now open another the terminal and go to directory where client program is present and type command
Go run {client_program_name}.go
Here, client_program_name is -> sync_client.go / async_client.go








2.	Variations Implemented:-

i.	Transcient- Synchronous communication:-

-> Here the client waits for the response(Which is otp) that needs to come from the server . And then starts the timer . We need to send the otp to the server within the time limit given or it expires.
Which shows synchrounous behaviour i.e., the client is waiting untill the response comes and blocking all other code at client side

-> Here we don’t have any message queuing system so it cant be persistent

	Here is the timer function which runs in background

 afterFuncTimer := 
                time.AfterFunc(time.Second*20, func() {
                //fmt.Println("Your time is completed")  -> uncomment this to debug when the timer is finishing
                status = "expired"
                
            })



 

ii.	Transicent – Asynchronous Communication:-

-> Here the client doesnot wait for the response and starts running the timer as soon as request is sent.
-> Since there is no message queuing system this can be transcient way of communication.


 


Note :-
If we are using asynchronous type of communication the timer starts before recieveing the response.
If the is any delay in the network and we have less time to enter otp , the timer expires before the otp arrives to the client. For which we need to initiate the transcation again.
This can be a disadvantage in banking applications since the functions happen asynchronously there.
