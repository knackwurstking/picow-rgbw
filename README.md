# picow-rgbw

<!--toc:start-->
- [picow-rgbw](#picow-rgbw)
  - [Getting Started](#getting-started)
  - [Api v1 Routing Table](#api-v1-routing-table)
  - [Api V1](#api-v1)
    - [**GET** _"/api/v1/devices"_](#get-apiv1devices)
    - [**POST** _"/api/v1/devices"_](#post-apiv1devices)
    - [**GET** _"/api/v1/devices/:id"_](#get-apiv1devicesid)
    - [**GET** _"/api/v1/events/device-update"_](#get-apiv1eventsdevice-update)
    - [**POST** _"/api/v1/picow"_](#post-apiv1picow)
  - [TODOs](#todos)
<!--toc:end-->

Web server for controlling all [picow-rgbw](https://github.com/knackwurstking/picow-rgbw.git) driven devices.

## Getting Started

Install the latest node js version.

```bash
curl -fsSL https://deb.nodesource.com/setup_current.x | sudo -E bash -
sudo apt install nodejs
```

Build and install with make (as user, no sudo).

```bash
make build
make install
make service
```

## Api v1 Routing Table

TODO: find a better way for this

| Method | Endpoint                                                        |
| ------ | --------------------------------------------------------------- |
| GET    | [`/api/v1/devices`](#get-apiv1devices)                          |
| POST   | [`/api/v1/devices`](#post-apiv1devices)                         |
| GET    | [`/api/v1/devices/:id`](#get-apiv1devicesid)                    |
| SSE    | [`/api/v1/events/device-update`](#get-apiv1eventsdevice-update) |
| POST   | [`/api/v1/picow`](#post-apiv1picow)                             |

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

<a id="devices-list"></a>

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

### **POST** _"/api/v1/devices"_

Update device(s) duty for rgbw pins

Example Request

```bash
curl -X POST http://localhost:50833/api/v1/devices \
  -H 'Content-Type: application/json'\
  -d '[{ "addr": "192.168.178.50:80", "rgbw": [100,100,100,100] }]'
```

No Response (http status 200 on success)

### **GET** _"/api/v1/devices/:id"_

Get pico device and data per id (rgbw index)
([device index from devices list](#devices-list))

Example Request

```bash
curl http://localhost:50833/api/v1/devices/0
```

Example Response

```json
{
  "addr": "192.168.178.50:80",
  "rgbw": [
    { "nr": 0, "duty": 0 },
    { "nr": 1, "duty": 0 },
    { "nr": 2, "duty": 0 },
    { "nr": 3, "duty": 0 }
  ]
}
```

### **GET** _"/api/v1/events/device-update"_

> work in progress

### **POST** _"/api/v1/picow"_

Register a new raspberry pico device on the web server to handle.
(The [picow-rgbw](https://github.com/knackwurstking/picow-rgbw.git) will do
this automatically on startup if configured)

Example Request: _(Only picow devices will send requests here)_

```bash
curl -X POST http://localhost:50833/api/v1/picow\
  -H "Content-Type: application/json"\
  -d '{"addr": "192.168.178.50:80"}'
```

## TODOs

- [ ] frontend: control: add a color storage
