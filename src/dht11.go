package main

import (
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
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(sensorType, pin, false, 10)
	if err != nil {
		lg.Fatal(err)
	}
	// print temperature and humidity
	lg.Infof("Sensor = %v: Temperature = %v*C, Humidity = %v%% (retried %d times)",
		sensorType, temperature, humidity, retried)
}
