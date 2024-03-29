package main

import (
	"fmt"
	"log"
	"os"

	"github.com/atphinn/Blackhatgo/metasploit-minimal/rpc"
	//"github.com/blackhat-go/bhg/ch-3/metasploit-minimal/rpc"
)

func main()  {
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := "msf"

	if host == "" || pass == "" {
		log.Fatalln("Missing required enviorment varible MSFHOST or MSFPASS")
	}
	msf, err := rpc.New(host, user, pass)
	if err != nil {
		log.Panicln(err)
	}
	defer msf.Logout()

	sessions, err := msf.SessionList()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Sessions:")

	for _, session := range sessions {
		fmt.Printf("%e5d %s\n", session.ID, session.Info)
	}
		
	
}