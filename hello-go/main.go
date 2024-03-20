package main

import (
	"fmt"

	"masasdani.com/hello/pkg/dns"
)

func main() {
	dnsTest()
	spfTest()
}

func dnsTest() {
	domain := "masasdani.com"
	as, _ := dns.ResolveARecord(domain)
	for _, a := range as {
		fmt.Println(a)
	}
	ts, _ := dns.ResolveTXTRecord(domain)
	for _, t := range ts {
		fmt.Println(t)
	}
	ms, _ := dns.ResolveMXRecord(domain)
	for _, m := range ms {
		fmt.Println(m)
	}
}

func spfTest() {
	domain := "mtarget.co"
	spf, err := dns.ValidateSPFRecord(domain, "include:_spf.google.com")
	fmt.Println(spf, err)
}
