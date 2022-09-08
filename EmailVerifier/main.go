package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)  //to get the input of the domain
	fmt.Println("domain,hasMX,hasSPF,spfRecords,hasDMARC,dmarcRecords")

	for scanner.Scan() {  // multiple inputs but not at a time
		checkDomain(scanner.Text())  //for multiple items to scan
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from the input: %v \n", err)
	}

}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord string
	var dmarcRecord string
	/*
A DNS 'mail exchange' (MX) record directs email to a mail server.
The MX record indicates how email messages should be routed in accordance with the Simple Mail Transfer Protocol
	*/

	mxRecords, err := net.LookupMX(domain)// to return mail server names are validated to be properly formatted presentation-format domain names. If

	if err != nil {
		log.Printf("Error: %v \n", err)
	}
	if len(mxRecords) > 0 { // if length is greater than 0 then te domain has mx records
		hasMX = true
	}
/*The DNS ‘text’ (TXT) record lets a domain administrator enter text into the Domain Name System (DNS).
 The TXT record was originally intended as a place for human-readable notes.
*/
	txtRecords, err := net.LookupTXT(domain)//LookupTXT returns the DNS TXT records for the given domain name.
	if err != nil {
		log.Printf("Error : %v \n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {		// to compare the string as the spfrecords is usually in the given format
			hasSPF = true
			spfRecord = record
			break
		}
	}
	/*DMARC, which stands for Domain-based Message Authentication, Reporting, and Conformance, 
	is a DNS TXT Record that can be published for a domain to control what happens if a message fails authentication
	*/
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error : %v\n ", err)
	}


	for _,record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {  // to compare the string as the dmarcrecords is usually in the given format
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v ,%v, %v, %v, %v, %v \n", domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord)
	return
}
