package verifier

import (
	"log"
	"net"
	"strings"
)

func VerifyEmail(domain string) DomainPropertiesData {
	mx := verifyMX(domain)
	spf, spfRecord := verifySPF(domain)
	dmarc, dmarcRecords := verifyDMARC(domain)

	return DomainPropertiesData{
		Mx:           mx,
		Spf:          spf,
		Dmarc:        dmarc,
		SpfRecords:   spfRecord,
		DmarcRecords: dmarcRecords,
	}
}

func verifyMX(domain string) bool {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	return len(mxRecords) > 0
}

func verifySPF(domain string) (bool, string) {
	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			return true, record
		}
	}

	return false, ""
}

func verifyDMARC(domain string) (bool, string) {
	dmarcRecords, err := net.LookupTXT("_dmarc" + domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			return true, record
		}
	}

	return false, ""
}
