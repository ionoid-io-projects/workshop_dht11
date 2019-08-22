package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/d2r2/go-dht"

	logger "github.com/d2r2/go-logger"
)

var lg = logger.NewPackageLogger("main",
	logger.DebugLevel,
	// logger.InfoLevel,
)

func main() {
	defer logger.FinalizeLogger()
	
	sensorType := dht.DHT11
	// sensorType := dht.AM2302
	// sensorType := dht.DHT12
	// Read DHT11 sensor data from specific pin, retrying 10 times in case of failure.
	pin := 4
	temperature, humidity, _, err :=
		dht.ReadDHTxxWithRetry(sensorType, pin, false, 10)
	if err != nil {
		lg.Fatal(err)
	}

	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	for {
		// print temperature and humidity
		// lg.Infof("Sensor = %v: Temperature = %v*C, Humidity = %v%% (retried %d times)", sensorType, temperature, humidity, retried)
		fmt.Printf("Temperature = %.2f *C, Humidity = %.2f \n",temperature,  humidity)
		time.Sleep(time.Second * 2)
	}
	
}
