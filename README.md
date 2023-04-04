# picow-rgbw-web

<!--toc:start-->
- [picow-rgbw-web](#picow-rgbw-web)
  - [Api v1 Routing Table](#api-v1-routing-table)
  - [Response Data (`/api/v1/device/:id`)](#response-data-apiv1deviceid)
  - [TODOs](#todos)
<!--toc:end-->

Web server for controlling all [picow-rgbw](https://github.com/knackwurstking/picow-rgbw.git) driven devices.

## Api v1 Routing Table

TODO: find a better way for this

| Method | Endpoint                       | Description                      |
| ------ | ------------------------------ | -------------------------------- |
| GET    | `/api/v1/devices`              | _get all devices_                |
| PUT    | `/api/v1/devices`              | _set rgbw for devices_           |
| GET    | `/api/v1/devices/:id`          | _get device per :id_             |
| SSE    | `/api/v1/events/device-update` | _sse event: device data changed_ |

Response Data (`/api/v1/device/:id`)

> package: [pico](internal/api/v1/pico/pico.go)

```go
type GpPin int

type Duty int

// GpPWM
type Gp struct {
  Nr   GpPin `json:"nr"`   // Nr of gpio pin in use (gp0 - gp28)
  Duty Duty  `json:"duty"` // Duty cycle (goes from 0-100)
}

type ReqPutDevice struct {
  Addr string  `json:"addr"`
  RGBW [4]Duty `json:"rgbw"`
}
```

## TODOs

- [ ] Makefile
  - [ ] script: "build"
  - [ ] script: "install" (no root) (copy empty a basic config.json file
        if not exists)
