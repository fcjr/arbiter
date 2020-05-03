package arbiter

import (
	"fmt"
	"net"
)

type NewProxyParams struct{}

// Proxy is an MITM-Capable HTTP Proxy
type Proxy struct{}

// NewProxy returns a new MITM-Capable HTTP Proxy
func NewProxy(params *NewProxyParams) *Proxy {
	return &Proxy{}
}

// Serve accepts incoming HTTP connections on the listener l,
// creating a new service goroutine for each and handles the request.
//
// HTTP/2 support is only enabled if the Listener returns *tls.Conn
// connections and they were configured with "h2" in the TLS
// Config.NextProtos.
//
// Serve always returns a non-nil error.
func (p *Proxy) Serve(l net.Listener) error {
	defer l.Close()

	return fmt.Errorf("unimplemented!")
}
