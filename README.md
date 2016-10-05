# Opsgenie Cardiogram

[![Build Status](https://travis-ci.org/lovoo/opsgenie-cardiogram.svg?branch=master)](https://travis-ci.org/lovoo/opsgenie-cardiogram)

**Simple Heartbeat Reporter for Opsgenie**

## Run

```
# Plain
./opsgenie-cardiogram -config /path/to/config.yml

# Docker (default config.yml expected in /data)
docker run -d --name opsgenie-cardiogram -v /dir/to/config:/data lovoo/opsgenie-cardiogram:latest
```

## Configuration

```
# default heartbeat api
url: https://api.opsgenie.com/v1/json/heartbeat/send
# generated api key from opsgenie integration
api_key: "oases-nairy-uncini-jawed-guglet-areca-azured"
# timeout for http client
timeout: 10s
# interval for sending the heartbeat
interval: 1m
targets:
  -
    # name of the configured heartbeat
    name: Google
    # url to test for the heartbeat
    url: https://www.google.com/
    # expected response status code
    status_code: 200
```

## Build

```
make build
```

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request
