package main

import (
	"github.com/miekg/dns"
)

//main function
func main() {
	var msg dns.Msg                    //Create a new message
	fqdn := dns.Fqdn("stacktitan.com") //transform the domain into a FQSDN to be exchanged with a DNS server
	msg.SetQuestion(fqdn, dns.TypeA)   //Modifies the internal state of the message with the intent to look up an A record
	dns.Exchange(&msg, "8.8.8.8:53")   //Sends the message to the supplied server address
}
