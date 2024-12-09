# qonvif

## Run api

```shell
go run main.go server
```

## Run docker

```shell
docker run --rm -p 8373:8373 -v $(pwd)/configs:/configs ghcr.io/qmaru/qonvif:latest server
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

### Send control command

```shell
curl --request POST \
  --url http://127.0.0.1:8373/api/onvif/device/ptz/control \
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
