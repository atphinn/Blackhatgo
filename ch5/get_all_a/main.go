package main

import (
  "fmt"

  "github.com/miekg/dns"
)

func main()  {
  var msg dns.Msg                    //Create a new message
  fqdn := dns.Fqdn("stacktitan.com") //transform the domain into a FQSDN to be exchanged with a DNS server
  msg.SetQuestion(fqdn, dns.TypeA) //Sends the message to the supplied server address
  in, err := dns.Exchange(&msg, "8.8.8.8:53")   //Sends the message to the supplied server address
  if err != nil {
    panic(err) //If there is an error then stop the program
  }
  //Validate the answer is 1 and if it isnt that means no record
  if len(in.Answer) < 1 {
    fmt.Println("No Records")
    return
  }
  for_, answer := range in.Answer{// loop throught the answers and confirm that it was successful
    if a, ok := answer.(*dns.A); ok {
			fmt.Println(a.A)
    }
  }
