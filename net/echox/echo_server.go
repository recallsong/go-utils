package echox

import (
	"crypto/tls"
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/recallsong/go-utils/net/servegrp"
)

type EchoServer struct {
	*echo.Echo
}

func New() *EchoServer {
	s := &EchoServer{
		Echo: echo.New(),
	}
	s.Echo.HideBanner = true
	return s
}

func (s *EchoServer) GetHttpServer(addr string) (error, servegrp.ServeItem) {
	s.Echo.Server.Addr = addr
	return nil, &httpServer{s.Echo, false}
}

func (s *EchoServer) GetHttpsServer(addr, certFile, keyFile string) (error, servegrp.ServeItem) {
	if certFile == "" || keyFile == "" {
		return errors.New("invalid tls configuration"), nil
	}
	var err error
	svr := s.Echo.TLSServer
	svr.TLSConfig = new(tls.Config)
	svr.TLSConfig.Certificates = make([]tls.Certificate, 1)
	svr.TLSConfig.Certificates[0], err = tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err, nil
	}
	if !s.Echo.DisableHTTP2 {
		svr.TLSConfig.NextProtos = append(svr.TLSConfig.NextProtos, "h2")
	}
	s.Echo.TLSServer.Addr = addr
	return nil, &httpServer{s.Echo, true}
}

type httpServer struct {
	*echo.Echo
	IsTLS bool
}

func (s *httpServer) Serve() error {
	var err error
	if s.IsTLS {
		err = s.StartServer(s.TLSServer)
	} else {
		err = s.StartServer(s.Server)
	}
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}

func (s *httpServer) Close() error {
	if s.IsTLS {
		return s.TLSServer.Close()
	} else {
		return s.Server.Close()
	}
}
