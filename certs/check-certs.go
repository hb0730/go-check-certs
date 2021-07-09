package certs

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func Check(addr string) ([]Cert, error) {
	dialer := new(net.Dialer)
	dialer.Timeout = 3 * time.Second
	conn, err := tls.DialWithDialer(dialer, "tcp", fmt.Sprintf("%s:443", addr), nil)
	if err != nil {
		return nil, err
	}
	err = conn.VerifyHostname(addr)
	if err != nil {
		return nil, err
	}
	pc := conn.ConnectionState().PeerCertificates
	certs := make([]Cert, 0, len(pc))
	for _, cert := range pc {
		c := Cert{}
		c.NotAfter = cert.NotAfter.Format(time.RFC3339)
		c.NotBefore = cert.NotBefore.Format(time.RFC3339)
		c.CommName = cert.Subject.CommonName
		certs = append(certs, c)

	}
	return certs, nil
}

type Cert struct {
	NotBefore string `json:"notBefore"`
	NotAfter  string `json:"notAfter"`
	CommName  string `json:"commName"`
}
