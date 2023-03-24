# picow-rgbw-web

<!--toc:start-->

- [picow-rgbw-web](#picow-rgbw-web)
  - [Api v1 Routing Table](#api-v1-routing-table)
  - [Response Data (`/api/v1/device?id=`)](#response-data-apiv1deviceid)
  - [NOTEs](#notes)
  <!--toc:end-->

Web server for controlling all [picow-rgbw](https://github.com/knackwurstking/picow-rgbw.git) driven devices.

## Api v1 Routing Table

| Method | Endpoint             | Description                  |
| ------ | -------------------- | ---------------------------- |
| GET    | `/api/v1/device`     | _get all devices_            |
| GET    | `/api/v1/device?id=` | _get device with :id number_ |
| GET    | `/api/v1/`           | \_...\_                      |

## Response Data (`/api/v1/device?id=`)

```go
type Device struct {
  ID   int    `json:"id"`   // ID is unique
  Addr string `json:"addr"` // <ip>:<port>
  // TODO: ...
}
```

## NOTEs

- Which framework to use? Or do it vanilla?
- json logging using `slog` package
- classic /api/v1 and serve static files directory
