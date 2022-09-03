package utils

import "net"

// MxRecord is a method to get mx-record for a domain
func MxRecord(domain string) []string {
	var mxRecord []string

	mxRecords, _ := net.LookupMX(domain)
	for _, mx := range mxRecords {
		mxRecord = append(mxRecord, mx.Host)
	}

	return mxRecord
}
