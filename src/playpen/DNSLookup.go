package main

import (
	"fmt"
	"net"
)

func listIPs(url string) {
	iprecords, _ := net.LookupIP(url)
	for _, ip := range iprecords {
		fmt.Println(ip)
	}
}

func listCNAME(url string) {
	cname, _ := net.LookupCNAME(url)
	fmt.Println(cname)
}

func listNames(ipAddress string) {
	ptr, _ := net.LookupAddr(ipAddress)
	for _, ptrvalue := range ptr {
		fmt.Println(ipAddress, ptrvalue)
	}
}

func listNameServers(url string) {
	nameservers, _ := net.LookupNS(url)
	for _, ns := range nameservers {
		fmt.Println(ns)
	}
}

func listMailServers(url string) {
	mxRecords, _ := net.LookupMX(url)
	for _, mx := range mxRecords {
		fmt.Println(mx.Host, mx.Pref)
	}
}

func listSRV(url string) {
	services := []string{"xmpp-server"}
	for _, s := range services {
		cname, srvs, _ := net.LookupSRV(s, "tcp", url)

		fmt.Printf("\ncname: %s \n\n", cname)

		for _, srv := range srvs {
			fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
		}
	}

}

func listTXT(url string) {
	txtrecords, _ := net.LookupTXT(url)

	for _, txt := range txtrecords {
		fmt.Println(txt)
	}
}

func main() {
	// casa := "casa.uh.edu"
	// fb := "facebook.com"
	// uh := "uh.edu"
	goog := "google.com"
	// math := "math.uh.edu"
	// cs := "cs.uh.edu"
	// listIPs(casa)
	// listCNAME(casa)

	// for i := 30; i <= 70; i++ {
	// 	ip := fmt.Sprintf("%s%d", "129.7.181.", i)
	// 	// fmt.Println(ip)
	// 	listNames(ip)
	// }

	// listNameServers(cs)
	// listMailServers(goog)
	// listTXT(goog)
	listSRV(goog)
}
