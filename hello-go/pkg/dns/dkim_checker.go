package dns

import (
	"errors"
	"strings"
)

// ValidateDKIMRecord checks if a given DKIM record is present in the DKIM records for a domain
func ValidateDKIMRecord(domain string, selector string, records []string) (bool, error) {
	rec, err := GetDKIMRecords(domain, selector)
	if err != nil {
		return false, err
	}

	for _, record := range records {
		if strings.Contains(rec, record) {
			return true, nil
		}
	}

	return false, errors.New("no valid DKIM record found")
}

// GetDKIMRecords returns the DKIM records for a given domain and selector
func GetDKIMRecords(domain string, selector string) (string, error) {
	recs, err := ResolveDNS(selector+"._domainkey."+domain, "CNAME")
	if err != nil {
		return "", err
	}

	if len(recs) == 0 {
		return "", errors.New("no DKIM records found")
	}

	return recs[0].Data, nil
}
