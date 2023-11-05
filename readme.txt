# Email Verifier Tool

## Introduction

Email Verifier Tool is a practical Go application designed to help users verify email addresses, extract domain information, and check essential email-related records such as MX, SPF, and DMARC. Whether you need to validate email addresses or explore domain configurations, this tool simplifies the process.

## Features

- Validate email addresses for correct format.
- Extract domain names from email addresses.
- Perform MX (Mail Exchange) lookups using `net.LookupMX`.
- Print MX records for the given domain.
- Check SPF (Sender Policy Framework) records using `net.LookupTXT`.
- Print SPF records for the domain.
- Check DMARC (Domain-based Message Authentication, Reporting, and Conformance) records using `net.LookupTXT`.
- Print DMARC records for the domain.

## Prerequisites

Make sure you have Go installed. If not, you can download it https://golang.org/dl/.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/email-verifier-tool.git
   cd email-verifier-tool

2. Build the application:
    go build email_verifier.go

3.Run the tool:
   ./email_verifier

Usage
-Enter an email address when prompted.
-The tool will validate the email format.
-If the email is valid, it will extract the domain name.
-It will perform MX, SPF, and DMARC lookups and display the results.

Dependencies
This project relies on the following Go packages:

-log
-bufio
-os
-strings
-badoux/checkmail for email validation

You can install external dependencies using go get.

License
This project is licensed under the MIT License.

Acknowledgments
The Go programming language and its net package for simplifying network-related tasks.
badoux/checkmail for email validation.
Contact
If you have questions or suggestions, feel free to contact us at amarjitaa@gmail.com.
