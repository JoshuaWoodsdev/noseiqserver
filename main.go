package main

import (
    
    "net/http"
    
    "github.com/labstack/echo/v4"
)


func main() {
    // Initialize Echo
    e := echo.New()

   
	// Define routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is the Index Page!")
	})
    
    // Start the server
    e.Start(":8080")
}
