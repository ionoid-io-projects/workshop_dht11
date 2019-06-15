#!/usr/bin/python
import os
import sys
import Adafruit_DHT
import time

dht_sensor = Adafruit_DHT.DHT11
dht_pin = 4

while True:

    humidity, temperature = Adafruit_DHT.read_retry(dht_sensor, dht_pin)
    print ('Temp: {0:0.1f} C  Humidity: {1:0.1f} %'.format(temperature, humidity))
    time.sleep(2)
