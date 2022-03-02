package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"text/template"

	"github.com/Karkhana-Studio/Backend-server/common/httpservice"
	log "github.com/Karkhana-Studio/Backend-server/common/logingservice"
	"github.com/Karkhana-Studio/Backend-server/lib/config"
	"github.com/Karkhana-Studio/Backend-server/lib/handler"
	"github.com/Karkhana-Studio/Backend-server/lib/service"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Clutch-n-Gearz-Backend
// @version 1.4.1
// @description Backend support for the hottest place to get your favorite superbike
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /

// Server - Structure for server
type Server struct {
	*config.ServerConfiguration
	router     *echo.Echo
	templates  *template.Template
	dbClient   *gorm.DB
	httpClient httpservice.HttpService
}

// NewServer - Constructor function for server
func NewServer(cfg *config.Configuration, db *gorm.DB) *Server {
	server := &Server{
		ServerConfiguration: cfg.Server,
		router:              echo.New(),
		templates:           loadTemplates(cfg.Server.WebDir),
		dbClient:            db,
		httpClient:          *httpservice.NewHttpService(),
	}

	server.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	server.router.Use(middleware.Logger())
	server.router.Use(middleware.Recover())
	server.router.Use(middleware.CORS())

	server.registerCommonAPI(cfg.Server)
	server.registerListingAPI(cfg.Server)
	server.router.Static(path.Join(cfg.Server.APIPath, "public"), path.Join(cfg.Server.WebDir, "public"))

	server.router.Renderer = server
	server.router.HideBanner = true
	server.router.HidePort = true
	return server
}

func (s *Server) registerCommonAPI(cfg *config.ServerConfiguration) {

	commonHandler := &handler.CommonHandler{
		ServerConfiguration: cfg,
	}
	s.router.GET(path.Join(cfg.APIPath, "/api/v1/health"), commonHandler.Health)
	s.router.GET(path.Join(cfg.APIPath, "/api/v1/swagger"), commonHandler.Swagger)
}

func (s *Server) registerListingAPI(cfg *config.ServerConfiguration) {
	listingHandler := &handler.ListingHandler{
		ServerConfiguration: cfg,
		ListingService:      service.NewListingService(),
	}

	s.router.GET(path.Join(cfg.APIPath, "/api/v1/listing/women"), listingHandler.ListWomenPageDetails)
}

// Start - This function will start the echo server
func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%s", s.Host, s.Port)
	log.Infof("Listening on %s", address)
	return s.router.Start(address)
}

// Stop - This function will stop the echo server
func (s *Server) Stop(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func loadTemplates(webDir string) *template.Template {
	templatePath := path.Join(webDir, "templates", "*.html")
	return template.Must(template.New("").Delims("[[", "]]").ParseGlob(templatePath))
}

// Render - Render the HTML to echo server
func (s *Server) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return s.templates.ExecuteTemplate(w, name, data)
}
