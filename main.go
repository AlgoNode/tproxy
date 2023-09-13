package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type OKResponse struct {
	Result string `json:"result"`
}

type TLog struct {
	Host      string `json:"Host"`
	Timestamp string `json:"@timestamp"`
	Message   string `json:"Message"`
	Data      Data   `json:"Data"`
	Level     string `json:"Level"`
}

type Data struct {
	Context string  `json:"Context"`
	Details Details `json:"details"`
	V       string  `json:"v"`
}

type Details struct {
	Address string `json:"Address"`
	Round   int64  `json:"Round"`
	Period  int64  `json:"Period"`
	Step    int64  `json:"Step"`
	Weight  int64  `json:"Weight"`
}

type TEntry struct {
	Event   string `json:"event"`
	Address string `json:"address"`
	Round   int64  `json:"round"`
	Period  int64  `json:"period"`
	Step    int64  `json:"step"`
	Weight  int64  `json:"weight"`
	V       string `json:"ver"`
	UUID    string `json:"host"`
	Name    string `json:"name"`
	TS      string `json:"@timestamp"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Decompress())
	e.POST("/*", func(c echo.Context) error {
		//		b, _ := io.ReadAll(c.Request().Body)
		tlog := new(TLog)
		err := c.Bind(tlog)
		if err == nil {
			evt := strings.Split(tlog.Message, "/")
			host := strings.SplitN(tlog.Host, ":", 2)
			if len(evt) > 2 && evt[1] == "Agreement" && len(host) > 1 {
				tent := &TEntry{
					Event:   evt[2],
					Address: tlog.Data.Details.Address,
					Round:   tlog.Data.Details.Round,
					Period:  tlog.Data.Details.Period,
					Step:    tlog.Data.Details.Step,
					Weight:  tlog.Data.Details.Weight,
					V:       tlog.Data.V,
					UUID:    host[0],
					Name:    host[1],
					TS:      tlog.Timestamp,
				}
				bs, err := json.Marshal(tent)
				if err == nil {
					fmt.Println(string(bs))
				}
			}
		} else {
			fmt.Errorf("error binding %v", err)
		}
		//fmt.Println(string(b))
		return c.JSON(http.StatusOK, &OKResponse{Result: "OK"})
	})
	e.HEAD("/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
