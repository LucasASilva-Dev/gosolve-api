package server

import (
	"fmt"
	"gosolve/internal/version"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type Response struct {
	Index int `json:"index"`
}

// Handler is a function that handles the hello route.
// It returns a 200 OK response with the message "GoSolve API <version>".
func (w *Webserver) hello(c echo.Context) error {
	// Return a 200 OK response with the message "GoSolve API <version>"
	// Where <version> is the version of the API.
	return c.String(http.StatusOK, fmt.Sprintf("GoSolve API %s", version.Version()))
}

// healthcheck is a route handler to check the health of the server.
// It logs the response to the console with the specified log level,
// and returns a 200 OK response with the message "WORKING".
func (w *Webserver) healthcheck(c echo.Context) error {
	w.Srv.Logger.Errorf("This is an error level log")
	w.Srv.Logger.Debugf("This is a debug level log")
	w.Srv.Logger.Warnf("This is a warning level log")
	w.Srv.Logger.Infof("This is an info level log")

	return c.String(http.StatusOK, string("WORKING"))
}

// getSearch is a route handler to search for a value in the index.
// It takes the value as a parameter and returns a JSON response with
// the index of the value. If the value is not found, it returns a
// 404 StatusNotFound response.
func (w *Webserver) getSearch(c echo.Context) error {
	// Get the value from the URL parameter.
	value, err := strconv.Atoi(c.Param("value"))

	if err != nil {
		// If the value is not a valid integer, return a bad request response.
		w.Srv.Logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid value parameter"})
	}

	// Look up the value in the index.
	positionFound, found := w.Im.Lookup(value)
	if !found {
		// If the position is not found, return a 404 StatusNotFound response.
		w.Srv.Logger.Error(fmt.Sprintf("Position for value %d not found", value))
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Position not found"})
	}

	// Create a response object to return the index of the position.
	routesResponse := positionFound
	response := Response{}
	response.Index = routesResponse

	// Return the response as JSON.
	return c.JSON(http.StatusOK, response)
}
