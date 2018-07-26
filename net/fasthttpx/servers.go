package fasthttpx

import (
	"fmt"
	"net"
	"os"

	"github.com/valyala/fasthttp"
)

type baseServer struct {
	svr fasthttp.Server
	ln  net.Listener
}

func (s *baseServer) Close() error {
	return s.ln.Close()
}

func (s *baseServer) Listen(addr string) error {
	if s.ln != nil {
		err := s.ln.Close()
		if err != nil {
			return err
		}
		s.ln = nil
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.ln = ln
	return err
}

func (s *baseServer) Serve() error {
	return s.svr.Serve(s.ln)
}

type TcpServer struct {
	baseServer
}

func NewTcpServer(h fasthttp.RequestHandler) *TcpServer {
	return &TcpServer{baseServer: baseServer{svr: fasthttp.Server{Handler: h}}}
}

type TLSTcpServer struct {
	baseServer
	certFile, keyFile string
}

func NewTLSTcpServer(h fasthttp.RequestHandler, certFile, keyFile string) *TLSTcpServer {
	return &TLSTcpServer{
		baseServer: baseServer{svr: fasthttp.Server{Handler: h}},
		certFile:   certFile, keyFile: keyFile,
	}
}

func (s *TLSTcpServer) Serve() error {
	return s.svr.ServeTLS(s.ln, s.certFile, s.keyFile)
}

type UnixServer struct {
	baseServer
}

func NewUnixServer(h fasthttp.RequestHandler) *UnixServer {
	return &UnixServer{baseServer: baseServer{svr: fasthttp.Server{Handler: h}}}
}

func (s *UnixServer) Listen(addr string) error {
	if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("unexpected error when trying to remove unix socket file %q: %s", addr, err)
	}
	if s.ln != nil {
		err := s.ln.Close()
		if err != nil {
			return err
		}
		s.ln = nil
	}
	ln, err := net.Listen("unix", addr)
	if err != nil {
		return err
	}
	s.ln = ln
	return err
}
