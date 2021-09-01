package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Location struct {
	Data []Data `json:"data"`
}

type Data struct {
	DisplayAddress string  `json:"display_address"`
	DisplayRegion  string  `json:"display_region"`
	Log            float64 `json:"lat"`
	Lat            float64 `json:"lon"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	app := echo.New()
	app.Use(middleware.CORS())
	app.Use(middleware.Logger())

	app.GET("/mor-sac/:place", func(c echo.Context) error {
		place := c.Param("place")

		url := "https://api.mymappi.com/v2/geocoding/direct?apikey=" + os.Getenv("API_TOKEN") + "&q=" + place
		res, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var location Location
		if err := json.Unmarshal(body, &location); err != nil {
			log.Fatalln(err)
		}
		return c.JSON(http.StatusOK, location)
	})

	app.Start(":" + os.Getenv("PORT"))
}
