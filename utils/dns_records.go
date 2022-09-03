package utils

import (
	"net"
	"regexp"
	"strings"
)

// MxRecord is a method to get mx-record for a domain
func MxRecord(domain string) []string {
	var mxRecord []string

	mxRecords, _ := net.LookupMX(domain)
	for _, mx := range mxRecords {
		mxRecord = append(mxRecord, mx.Host)
	}

	return mxRecord
}

// SpfRecord is a method to check if it contains spf record in its txt-record
func SpfRecord(domain string) (bool, string) {
	regExp := regexp.MustCompile(`\s*v=spf1.*`)

	txtRecords, _ := net.LookupTXT(domain)
	for _, txt := range txtRecords {
		if regExp.MatchString(txt) {
			return true, txt
		}
	}

	return false, ""
}

// DmarcRecord is a method to check the DMARC record
func DmarcRecord(domain string) (bool, string) {
	dmarcRecords, _ := net.LookupTXT("_dmarc." + domain)
	for _, dmarc := range dmarcRecords {
		if strings.HasPrefix(dmarc, "v=DMARC1") {
			return true, dmarc
		}
	}
	return false, ""
}
