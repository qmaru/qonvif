# qonvif

## Build a server

```shell
go build --tags server
```

## Build a wails

```shell
wails build --tags ui -ldflags='-s -w'
```

## Run docker

```shell
docker run --rm -p 8373:8373 -v $(pwd)/configs:/configs ghcr.io/qmaru/qonvif:latest server
```

## Player [option]

```shell
# config.toml
[player]
path = "/path/to/ffplay"
```

## API reference

### Get all devices

```shell
curl --request GET \
  --url 'http://127.0.0.1:8373/api/onvif/devices?name=Living' \
  --header 'X-API-Key: api_key'
```

### Get a device info

```shell
curl --request GET \
  --url 'http://127.0.0.1:8373/api/onvif/device/info?name=Living' \
  --header 'X-API-Key: api_key'
```

### Get a device profile (token)

```shell
curl --request GET \
  --url 'http://127.0.0.1:8373/api/onvif/device/profile?name=Living' \
  --header 'X-API-Key: api_key'
```

### Get a device media stream url

```shell
curl --request GET \
  --url 'http://127.0.0.1:8373/api/onvif/device/streamurl?name=Living&token=profile_1' \
  --header 'X-API-Key: api_key'
```

```shell
curl --request GET \
  --url 'http://127.0.0.1:8373/api/onvif/device/streamurl?name=Living&token=profile_1&username=admin&password=PASSWORD' \
  --header 'X-API-Key: api_key'
```

### Send control command

```shell
curl --request POST \
  --url http://127.0.0.1:8373/api/onvif/device/ptz/move/absolute \
  --header 'X-API-Key: api_key' \
  --header 'content-type: application/json' \
  --data '{
  "name": "Living",
  "axes": {
    "x": 0.3,
    "y": 0.2
  }
}
'
```

### Get ptz status

```shell
curl --request GET \
  --url 'http://127.0.0.1:8373/api/onvif/device/ptz/status?name=Living' \
  --header 'X-API-Key: 123456'
```

### Play Stream

```shell
curl --request POST \
  --url http://127.0.0.1:8373/api/onvif/play \
  --header 'X-API-Key: api_key' \
  --header 'content-type: application/json' \
  --data '{
  "url": "stream_url",
  "width": "1280",
  "height": "720"
}
'
```
