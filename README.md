# picow-rgbw-web

<!--toc:start-->

- [picow-rgbw-web](#picow-rgbw-web)
  - [Api v1 Routing Table](#api-v1-routing-table)
    - [Go Types for Request/Response](#go-types-for-requestresponse)
  - [Api V1](#api-v1)
    - [**GET** _"/api/v1/devices"_](#get-apiv1devices)
    - [**PUT** _"/api/v1/devices"_](#put-apiv1devices)
    - [**GET** _"/api/v1/devices/:id"_](#get-apiv1devicesid)
    - [**GET** _"/api/v1/events/device-update"_](#get-apiv1eventsdevice-update)
  - [TODOs](#todos)
  <!--toc:end-->

Web server for controlling all [picow-rgbw](https://github.com/knackwurstking/picow-rgbw.git) driven devices.

## Api v1 Routing Table

TODO: find a better way for this

| Method | Endpoint                                                        |
| ------ | --------------------------------------------------------------- |
| GET    | [`/api/v1/devices`](#get-apiv1devices)                          |
| PUT    | [`/api/v1/devices`](#put-apiv1devices)                          |
| GET    | [`/api/v1/devices/:id`](#get-apiv1devicesid)                    |
| SSE    | [`/api/v1/events/device-update`](#get-apiv1eventsdevice-update) |

## Api V1

> package: [v1](internal/api/v1)  
> package: [pico](internal/api/v1/pico)  

### **GET** _"/api/v1/devices"_

Get all pico devices and data about rgbw pins.

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

### **PUT** _"/api/v1/devices"_

TODO: example request (and response) with curl

### **GET** _"/api/v1/devices/:id"_

TODO: example request (and response) with curl

### **GET** _"/api/v1/events/device-update"_

> work in progress

## TODOs

- [ ] Makefile
  - [ ] script: "build"
  - [ ] script: "install" (no root) (copy empty a basic config.json file
        if not exists)
