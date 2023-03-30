# picow-rgbw-web

<!--toc:start-->
- [picow-rgbw-web](#picow-rgbw-web)
  - [Api v1 Routing Table](#api-v1-routing-table)
  - [Response Data (`/api/v1/device/:id`)](#response-data-apiv1deviceid)
  - [NOTEs](#notes)
<!--toc:end-->

Web server for controlling all [picow-rgbw](https://github.com/knackwurstking/picow-rgbw.git) driven devices.

## Api v1 Routing Table

| Method | Endpoint                       | Description                      |
| ------ | ------------------------------ | -------------------------------- |
| GET    | `/api/v1/device`               | _get all devices_                |
| GET    | `/api/v1/device/:id`           | _get device per :id_             |
| SSE    | `/api/v1/events/device-update` | _sse event: device data changed_ |

## Response Data (`/api/v1/device/:id`)

> package: [pico](internal/api/v1/pico/pico.go)

```go
// GpPWM
type GpPWM struct {
    Nr   int `json:"nr"`   // Nr of gpio pin in use (gp0 - gp28)
    Duty int `json:"duty"` // Duty cycle (goes from 0-100)
}

// Device
type Device struct {
    ID   int       `json:"id"`   // ID is unique
    Addr string    `json:"addr"` // Addr contains the ip and port <ip>:<port>
    RGBW [4]*GpPWM `json:"rgbw"` // RGBW holds all pins in use
}
```
