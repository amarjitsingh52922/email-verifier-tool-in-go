# Email Verifier Tool

![Go Version](https://img.shields.io/badge/Go-v1.17-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

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
