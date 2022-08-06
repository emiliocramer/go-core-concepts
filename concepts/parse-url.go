package concepts

import (
	"fmt"
	"net"
	"net/url"
)

func Parser() {
	h := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse to check for errors
	u, err := url.Parse(h)
	if err != nil {
		panic(err)
	}

	// Print different parts of the parsed url
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// We can call net's SplitHostPort to split it (returns host, port, err)
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// k=v
	fmt.Println(u.RawQuery)

	// map[k: [v]]
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)

	// v
	fmt.Println(m["k"][0])

}
