# Workshop DHT11 (temperature and humidity sensor)
The DHT11 sensor is a commonly used temperature and humidity sensor, in this workshop we will see how to use the DHT11 sensor and the Raspberry Pi to display temperature and humidity information.

# Material for Raspberry Pi workshop

![wiring](doc/img/dht112.jpg)

|DHT11             |  Raspberry Pi  |
|------------------|----------------|
| Ground 		   | GND    (black) |
| Vcc / + 		   | 3V     (red)   |
| Data       	   | GPIO 4 (blue)  |

![wiring](doc/img/gpio.png)

# Code


## 1- Power your raspberry

You can achive it with connecting it to your pc trought the Micro USB Port of the raspberry pi

![power](doc/img/1-min.jpg)

## 2- Connect to your raspberry pi
Using putty if you're on windows, Ssh if you're on a linux based os
Follow the following instruction if you dont know how to connect to raspberry pi
[Connect to raspberry pi using Putty](https://github.com/ionoid-io-projects/workshop/blob/master/doc/od-iot-raspbian-rpi-zero-windows.md#5-first-boot)

## 3- Download Dht11 binary file

Assuming you're connected with... copy and past this command
If you're using Raspberry zero
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_dht11/master/bin/arm6/dht11
```

If you're using Raspberry 3 b
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_dht11/master/bin/arm7/dht11
```
## make it executable
```
chmod +x dht11
```

## 4- execute binary to make led blink
```
sudo ./dht11
```

## How to stop the program
To quit or stop the program click on **Ctrl+C**

# Ressources
[Workshop Temperature and Humidty](https://github.com/opendevices/iot.apps/tree/master/workshop-temperature-humidty-dht11-dht22)

[https://tutorials-raspberrypi.com/raspberry-pi-measure-humidity-temperature-dht11-dht22/](https://tutorials-raspberrypi.com/raspberry-pi-measure-humidity-temperature-dht11-dht22/)

