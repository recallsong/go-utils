package httpx

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/recallsong/go-utils/net/netutil"
)

type HttpServer struct {
	*http.Server
	Listener net.Listener
}

func NewHttpServer(handler http.Handler) *HttpServer {
	s := &HttpServer{
		Server: new(http.Server),
	}
	s.Server.Handler = handler
	return s
}

func (s *HttpServer) Listen(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	ln = netutil.TcpKeepAliveListener(ln)
	if s.TLSConfig != nil {
		ln = tls.NewListener(ln, s.TLSConfig)
	}
	s.Listener = ln
	s.Server.Addr = addr
	return nil
}

func (s *HttpServer) Serve() error {
	err := s.Server.Serve(s.Listener)
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}

func (s *HttpServer) Close() error {
	return s.Server.Close()
}
