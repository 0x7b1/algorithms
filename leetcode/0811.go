package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	subdomains := make(map[string]int)
	var visits []string

	for _, cpdomain := range cpdomains {
		pair := strings.Split(cpdomain, " ")
		counter, _ := strconv.Atoi(pair[0])
		domain := pair[1]

		mapSubdomains(subdomains, domain, counter)
	}

	for subdomain, counter := range subdomains {
		visits = append(visits, strconv.Itoa(counter)+" "+subdomain)
	}

	return visits
}

func mapSubdomains(subdomains map[string]int, url string, counter int) {
	domains := strings.Split(url, ".")
	subdomain := domains[len(domains)-1]
	subdomains[subdomain] += counter

	for i := len(domains) - 2; i >= 0; i-- {
		subdomain = domains[i] + "." + subdomain
		subdomains[subdomain] += counter
	}
}

func main() {
	//fmt.Println(mapSubdomains("ax.bx.cx"))
	//fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
