# picow-rgbw-web

<!--toc:start-->
- [picow-rgbw-web](#picow-rgbw-web)
  - [Api v1 Routing Table](#api-v1-routing-table)
    - [Quick Overview for Request and Response (Body)](#quick-overview-for-request-and-response-body)
  - [TODOs](#todos)
<!--toc:end-->

Web server for controlling all [picow-rgbw](https://github.com/knackwurstking/picow-rgbw.git) driven devices.

## Api v1 Routing Table

TODO: find a better way for this

| Method | Endpoint                       | Request             | Response        |
| ------ | ------------------------------ | ------------------- | --------------- |
| GET    | `/api/v1/devices`              | -                   | `[]pico.Device` |
| PUT    | `/api/v1/devices`              | `[]v1.ReqPutDevice` | -               |
| GET    | `/api/v1/devices/:id`          | -                   | `pico.Device`   |
| SSE    | `/api/v1/events/device-update` | @TODO               | @TODO           |

### Quick Overview for Request and Response (Body)

> package: [v1](internal/api/v1)

```go
type ReqPutDevice struct {
  Addr string       `json:"addr"`
  RGBW [4]pico.Duty `json:"rgbw"`
}
```

> package: [pico](internal/api/v1/pico)

```go
type GpPin int

type Duty int

// GpPWM
type Gp struct {
  Nr   GpPin `json:"nr"`   // Nr of gpio pin in use (gp0 - gp28)
  Duty Duty  `json:"duty"` // Duty cycle (goes from 0-100)
}

// Device
type Device struct {
  Addr string `json:"addr"` // Addr contains the ip and port <ip>:<port>
  RGBW [4]*Gp `json:"rgbw"` // RGBW holds all pins in use
}
```

## TODOs

- [ ] Makefile
  - [ ] script: "build"
  - [ ] script: "install" (no root) (copy empty a basic config.json file
        if not exists)
