// 052 url parsing

package main

import "fmt"
import "net"
import "net/url"

func main() {
	p := fmt.Println
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	p(u.Scheme)
	p(u.User)
	p(u.User.Username())

	pw, _ := u.User.Password()
	p(pw)

	p(u.Host)
	h, pt, _ := net.SplitHostPort(u.Host)
	p(h)
	p(pt)

	p(u.Path)
	p(u.Fragment)

	p(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"][0])
}
