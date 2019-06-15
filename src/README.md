# DHT11 sensor
The following python code with get informations from sensor and send it to firebase

# Installation (python app)
Before running the code you need to install dependencies
```
$ sudo apt-get install build-essential python-dev python-openssl git python-pip
$ sudo pip install firebase-admin
```
After that install dht11 python library
```
$ git clone https://github.com/adafruit/Adafruit_Python_DHT.git && cd Adafruit_Python_DHT
$ sudo python setup.py install
```
Now you're good to go
# Usage
```
$ export PROJECT_ID="dht11-data-1b7d9" python dht11-firebase-app.py
```

OR if you don't wan to use firebase

```
$ python dht11.py
```


# Installation (go version of the app)
Install dependecy first
```
go get -u github.com/d2r2/go-dht
go get -u github.com/d2r2/go-logger
```
Building the binary for rpi
 
```
env GOOS=linux GOARCH=arm GOARM=6 go build dht11.go
```

# Usage
Copy the generated binary file to your raspberry pi and execute it