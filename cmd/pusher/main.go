package main

import (
	"fmt"
	"log"
	"os"

	data "ModebusAdventure/cmd"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func main() {
	plcData := data.GetData()
	tic_100 := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tic_100_temperature",
		Help: "a temperature sensor. Unit: Celcius.",
	})
	tic_100.Set(float64(plcData["tic_100"]))
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_time",
		Help: "The timestamp.",
	})
	completionTime.SetToCurrentTime()
	if err := push.New(os.Getenv("pushgw_host"), "modbus_metrics").
		Collector(tic_100).
		Collector(completionTime).
		Grouping("modicon_m340", "sensor").
		BasicAuth(os.Getenv("pushgw_user"), os.Getenv("pushgw_pass")).
		Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
