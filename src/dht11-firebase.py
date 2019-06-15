#!/usr/bin/python
import os
import sys
import Adafruit_DHT
import time
import firebase_admin
from firebase_admin import credentials
from firebase_admin import db

project_id = os.getenv('PROJECT_ID', 'dht11-data-1b7d9')
cred = credentials.Certificate("/home/pi/serviceAccountKey.json")
firebase_admin.initialize_app(cred, {'databaseURL': 'https://{}.firebaseio.com/'.format(project_id)})
ref = db.reference('/')

dht_sensor = Adafruit_DHT.DHT11
dht_pin = 4

while True:

    humidity, temperature = Adafruit_DHT.read_retry(dht_sensor, dht_pin)
    print ('Temp: {0:0.1f} C  Humidity: {1:0.1f} %'.format(temperature, humidity))

    data = {
        "Date": time.strftime('%H:%M:%S'),
        "Time": time.strftime('%d/%m/%Y'), 
        "Temperature": '{0:0.1f}*'.format(temperature),
        "Humidity": '{0:0.1f}*'.format(humidity)
    }

    ref.push(data)
    time.sleep(2)
