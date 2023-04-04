# picow-rgbw-web

<!--toc:start-->
- [picow-rgbw-web](#picow-rgbw-web)
  - [Api v1 Routing Table](#api-v1-routing-table)
    - [Go Types for Request/Response](#go-types-for-requestresponse)
  - [Api V1](#api-v1)
    - [*GET* **"/api/v1/devices"**](#get-apiv1devices)
    - [*PUT* **"/api/v1/devices"**](#put-apiv1devices)
    - [*GET* **"/api/v1/devices/:id"**](#get-apiv1devicesid)
    - [*GET* **"/api/v1/events/device-update"**](#get-apiv1eventsdevice-update)
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

### Go Types for Request/Response

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

## Api V1

### *GET* **"/api/v1/devices"**

Example Request

```bash
curl http://localhost:50833/api/v1/devices
```

Example Response

```json
[
  {
    "addr": "192.168.178.50:80",
    "rgbw": [
      { "nr": 0, "duty": 0 },
      { "nr": 1, "duty": 0 },
      { "nr": 2, "duty": 0 },
      { "nr": 3, "duty": 0 }
    ]
  }
]
```

### *PUT* **"/api/v1/devices"**

TODO: example request (and response) with curl

### *GET* **"/api/v1/devices/:id"**

TODO: example request (and response) with curl

### *GET* **"/api/v1/events/device-update"**

> work in progress

## TODOs

- [ ] Makefile
  - [ ] script: "build"
  - [ ] script: "install" (no root) (copy empty a basic config.json file
        if not exists)
