package server

import (
	"gosolve/internal/index"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Webserver represents a webserver
type Webserver struct {
	Srv *echo.Echo
	Im  *index.IndexManager
}

// NewWebServer Return an instance of WebServer
// This function creates a new instance of WebServer with the provided
// IndexManager and LogLevel.
func NewWebServer(im *index.IndexManager, logLevel *string) *Webserver {
	// Create a new Echo server
	srv := echo.New()

	// Root level middleware
	// This middleware recovers from panics and logs to the console.
	srv.Use(middleware.Recover())

	// Common Log Format
	// This is the log format used by the middleware.
	logFormat := `[${time_custom}] host=${host} uri=${uri} method=${method} path=${path} server_protocol=${protocol} remote_addr=${remote_ip} status=${status} bytes_out=${bytes_out} referer=${referer} http_user_agent=${user_agent} error=${error} latency=${latency} latency_human=${latency_human}`
	// This is the custom time format used by the middleware.
	customTimeFormat := "2/Jan/2006:15:04:05 -0300"
	// Append a newline to the log format.
	logFormat += "\n"
	// Set the log format and output for the middleware.
	srv.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           logFormat,
		CustomTimeFormat: customTimeFormat,
		Output:           os.Stdout,
	}))

	// Set the log level for the server
	// This sets the log level for the server based on the provided LogLevel.
	switch *logLevel {
	case "ERROR":
		srv.Logger.SetLevel(8)
		log.SetLevel(log.ErrorLevel)
	case "WARNING":
		srv.Logger.SetLevel(3)
		log.SetLevel(log.WarnLevel)
	case "INFO":
		srv.Logger.SetLevel(2)
		log.SetLevel(log.InfoLevel)
	case "DEBUG":
		srv.Logger.SetLevel(1)
		log.SetLevel(log.DebugLevel)
	default:
		srv.Logger.SetLevel(2)
		log.SetLevel(log.InfoLevel)
	}

	// Init and return web server
	// This creates a new instance of WebServer with the provided
	// IndexManager and LogLevel.
	return &Webserver{
		Srv: srv,
		Im:  im,
	}
}

// Start Method to start server http and register handlers
// This method starts the Echo server and registers handlers for it.
// The server is started with the provided host and port.
// The server starts a goroutine to start the metrics server on port 8380.
// The server logs to the console the message "Metrics server started on port 8380"
func (w *Webserver) Start(host *string, port *int) {
	// Create echo Prometheus server and Middleware
	// This creates a new Echo server and a Middleware to handle Prometheus metrics.
	echo := echo.New()
	echo.GET("/metrics", echoprometheus.NewHandler())

	// Register the Middleware with the Echo server
	// This Middleware will handle Prometheus metrics.
	w.Srv.Use(echoprometheus.NewMiddleware("gosolve_api"))

	// Start metrics server
	// This starts the metrics server in a goroutine.
	go func() {
		// Log to the console the message "Metrics server started on port 8380"
		log.Info("Metrics server started on port 8380")
		// Start the metrics server on port 8380
		echo.Logger.Fatal(echo.Start(*host + ":8380"))
	}()

	// Routes
	// This registers the routes for the server.
	w.Srv.GET("/", w.hello)
	w.Srv.GET("/healthcheck", w.healthcheck)
	w.Srv.GET("/search/:value", w.getSearch) //OLD URL

	// Start app server
	// This logs to the console the message "API Server started on port 1323"
	log.Info("API Server started on port 1323")
	// This starts the server on the provided host and port
	if err := w.Srv.Start(*host + ":" + strconv.Itoa(*port)); err != nil {
		// Log the error to the console
		w.Srv.Logger.Fatal(err)
	}
}
