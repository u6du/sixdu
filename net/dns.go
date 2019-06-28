package net

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/rs/zerolog/log"
)

/*
import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"sixdu/config"
)

var DnsDot []string

func init() {
	DnsDot = config.Li(
		"dns/dot",
		`dns.rubyfish.cn
dot-jp.blahdns.com
dns.google
security-filter-dns.cleanbrowsing.org
dot.securedns.eu
sdns.233py.com
edns.233py.com
ndns.233py.com
dns.quad9.net
wdns.233py.com
dot-de.blahdns.com
1dot1dot1dot1.cloudflare-dns.com
dns.brahma.world`)
}

*/

func NewResolver(addr string) *net.Resolver {
	var dialer net.Dialer
	tlsConfig := &tls.Config{
		ServerName:         addr,
		ClientSessionCache: tls.NewLRUClientSessionCache(32),
		InsecureSkipVerify: false,
	}

	return &net.Resolver{
		PreferGo: true,
		Dial: func(context context.Context, _, address string) (net.Conn, error) {
			conn, err := dialer.DialContext(context, "tcp", addr+":853")
			if err != nil {
				return nil, err
			}

			_ = conn.(*net.TCPConn).SetKeepAlive(true)
			_ = conn.(*net.TCPConn).SetKeepAlivePeriod(10 * time.Minute)
			return tls.Client(conn, tlsConfig), nil
		},
	}

}
func Txt(name, nameserver string) *string {
	resolve := NewResolver(nameserver)
	n := 0
	for {
		li, err := resolve.LookupTXT(context.Background(), name)
		if err != nil {
			log.Warn().Err(err).Msg(nameserver)

			if n > 1 {
				return nil
			}

			n += 1
		}

		for _, i := range li {
			return &i
		}
	}
}

/*
func init() {
	net.DefaultResolver = newResolver("1.0.0.1")
}
*/

/*
func DialNew(nameserver string) func(context.Context, string, string) (net.Conn, error) {
	return func(ctx context.Context, network, address string) (net.Conn, error) {
		d := net.Dialer{}
		return d.DialContext(ctx, "udp", nameserver+":53")
	}
}

var NAMESERVER = []string{
	"180.76.76.76",
	"119.29.29.29",
	"223.5.5.5",
	"1.1.1.1",
	"208.67.220.220",
	"8.8.8.8",
	"9.9.9.9",
	"114.114.114.114",
	"223.6.6.6",
}

func main() {
	for _, nameserver := range NAMESERVER {


	}
}
*/
