package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/badoux/checkmail"
)

func extractDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		fmt.Println("Email ", email, " is not valid")
		return ""
	}
	return parts[1]
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please Enter Email Address:")
	for scanner.Scan() {

		email := scanner.Text()
		err := checkmail.ValidateFormat(email)
		if err != nil {
			fmt.Println("Invalid email format:", err)
		} else {
			domain := extractDomain(email)
			if len(domain) > 0 {
				checkDomain(domain)
			}

		}
		fmt.Printf("Please Enter Email Address:")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error: couldn't read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {

		fmt.Println("MX Records for ", domain, ":")
		for _, mx := range mxRecords {
			fmt.Println(mx.Host, ":", mx.Pref)
		}
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Erro: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}

	}
	if hasSPF == true {
		fmt.Println("hasSPF=True,spfRecrd: ", spfRecord)
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	if hasDMARC == true {
		fmt.Println("hasDMARC=true, DMARCRecord: ", dmarcRecord)
	}

}
