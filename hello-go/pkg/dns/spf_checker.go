package dns

import (
	"errors"
	"strings"
)

// GetSPFRecords returns the SPF records for a given domain
func GetSPFRecords(domain string) ([]string, error) {
	recs, err := ResolveTXTRecord(domain)
	if err != nil {
		return nil, err
	}

	var spfRecords []string
	for _, rec := range recs {
		if strings.HasPrefix(rec.Data, "v=spf1") {
			spfRecords = append(spfRecords, rec.Data)
		}
	}
	if len(spfRecords) == 0 {
		return nil, errors.New("no SPF records found")
	}
	if len(spfRecords) > 1 {
		return spfRecords, errors.New("multiple SPF records found")
	}
	return spfRecords, nil
}

// ValidateSPFRecord checks if a given SPF record is present in the SPF records for a domain
func ValidateSPFRecord(domain string, record string) (bool, error) {
	recs, err := GetSPFRecords(domain)
	if err != nil {
		return false, err
	}

	if strings.Contains(recs[0], record) {
		return true, nil
	}

	return false, errors.New("Need to add " + record + " to the SPF record for " + domain + "!")
}
