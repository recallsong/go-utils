package httpfs

import (
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/recallsong/go-utils/net/echox"
)

type HttpServer struct {
	*echox.EchoServer
	fs        http.FileSystem
	templates *template.Template
}

func NewHttpServer(fs http.FileSystem) *HttpServer {
	s := &HttpServer{
		EchoServer: echox.New(),
		fs:         fs,
	}
	return s
}

func (s *HttpServer) Static(path, dir string) {
	var handler http.Handler
	if len(dir) <= 0 || dir == "/" {
		handler = http.FileServer(s.fs)
	} else {
		handler = http.FileServer(&subFS{
			routePath:  path,
			dirPrefix:  dir,
			FileSystem: s.fs,
		})
	}
	s.Echo.GET(filepath.Join(path, "**"), func(c echo.Context) error {
		handler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}

func (s *HttpServer) LoadTempletes(dir string) error {
	s.templates = template.New("")
	s.Echo.Renderer = s
	return s.loadTemps(dir, dir)
}

func (s *HttpServer) loadTemps(prefix, dir string) error {
	file, err := s.fs.Open(dir)
	if err != nil {
		return err
	}
	dirs, err := file.Readdir(0)
	if err != nil {
		file.Close()
		return err
	}
	file.Close()
	for _, file := range dirs {
		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			err = s.loadTemps(prefix, path)
			if err != nil {
				return err
			}
		} else {
			err = s.loadTemp(path, path[len(prefix)+1:])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *HttpServer) loadTemp(path, name string) error {
	file, err := s.fs.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	tmp := s.templates.New(name)
	tmp.Parse(string(bytes))
	return nil
}

func (s *HttpServer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return s.templates.ExecuteTemplate(w, name, data)
}

type subFS struct {
	routePath string
	dirPrefix string
	http.FileSystem
}

func (fs *subFS) Open(name string) (http.File, error) {
	return fs.FileSystem.Open(filepath.Join(fs.dirPrefix, name[len(fs.routePath):]))
}
