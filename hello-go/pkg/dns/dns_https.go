package dns

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"masasdani.com/hello/pkg/http"
)

// Record represents a DNS record
type Record struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	TTL      int    `json:"TTL"`
	Data     string `json:"data"`
	Priority int    `json:"priority"`
}

const (
	// ARecordType is the type for A records
	ARecordType = "A"
	// TXTRecordType is the type for TXT records
	TXTRecordType = "TXT"
	// MXRecordType is the type for MX records
	MXRecordType = "MX"
	// googleDNSURL is the URL for the Google DNS API
	googleDNSURL = "https://dns.google/resolve"
)

// ResolveDNS resolves a DNS record for a given domain and type
func ResolveDNS(domain string, types string) ([]Record, error) {
	// Create an HTTP client
	client := http.CreateHTTPClient()
	res, err := client.Get(googleDNSURL + "?name=" + domain + "&type=" + types)
	if err != nil {
		// Handle error
		log.Fatalf("Error: %v", err)
		return nil, err
	}

	result, err := http.ReadResponseBody(res)
	if err != nil {
		// Handle error
		log.Fatalf("Error: %v", err)
		return nil, err
	}

	if result["Answer"] == nil {
		// No records found
		return []Record{}, nil
	}

	var records []Record
	recs := result["Answer"].([]interface{})
	for _, rec := range recs {
		rec := rec.(map[string]interface{})
		data := rec["data"].(string)
		priority := 0

		// If the record type is an MX record, parse the priority and data
		if types == MXRecordType {
			dataArr := strings.Split(data, " ")
			priority, _ = strconv.Atoi(dataArr[0])
			data = dataArr[1]
		}
		records = append(records, Record{
			Name:     rec["name"].(string),
			Type:     types,
			TTL:      int(rec["TTL"].(float64)),
			Data:     data,
			Priority: priority,
		})
	}

	// Sort the records by priority if the type is an MX record
	if types == MXRecordType {
		sort.Slice(records, func(i, j int) bool {
			return records[i].Priority < records[j].Priority
		})
	}
	return records, nil

}

func ResolveARecord(domain string) ([]Record, error) {
	return ResolveDNS(domain, ARecordType)
}

func ResolveTXTRecord(domain string) ([]Record, error) {
	return ResolveDNS(domain, TXTRecordType)
}

func ResolveMXRecord(domain string) ([]Record, error) {
	return ResolveDNS(domain, MXRecordType)
}
