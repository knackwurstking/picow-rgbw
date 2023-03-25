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

```go
// PWMPin
type PWMPin struct {
    Gp    int `json:"gp"`   // Gp which gpio pin is in use (goes from 0-28) (unique)
    Duty  int `json:"duty"` // Duty cycle (goes from 0-100)
}

// PicoDevice
type PicoDevice struct {
    ID   int        `json:"id"`   // ID is unique
    Addr string     `json:"addr"` // Addr contains the ip and port <ip>:<port>
    RGBW [4]*PWMPin `json:"rgbw"` // RGBW holds all pins in use
}
```

## NOTEs

- Which framework to use? Or do it pure go?
  - pure go
  - listening on port **50833**
- json logging using `slog` package
- classic /api/v1 and serve static files directory
  - svelte frontend
- sse events for "device-update"
