# picow-rgbw

Web server for controlling all [picow-rgbw](https://github.com/knackwurstking/picow-rgbw.git) driven devices.

# Index

<!-- vscode-markdown-toc -->
* 1. [Getting Started](#GettingStarted)
* 2. [Api v1 Routing Table](#Apiv1RoutingTable)
* 3. [Api V1](#ApiV1)
	* 3.1. [**GET** _"/api/v1/devices"_](#GET_apiv1devices_)
	* 3.2. [**POST** _"/api/v1/devices"_](#POST_apiv1devices_)
	* 3.3. [**GET** _"/api/v1/devices/:id"_](#GET_apiv1devices:id_)
	* 3.4. [**GET** _"/api/v1/events/device-update"_](#GET_apiv1eventsdevice-update_)
	* 3.5. [**POST** _"/api/v1/picow"_](#POST_apiv1picow_)
* 4. [TODOs](#TODOs)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc --># picow-rgbw

##  1. <a name='GettingStarted'></a>Getting Started

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

##  2. <a name='Apiv1RoutingTable'></a>Api v1 Routing Table

TODO: find a better way for this

| Method | Endpoint                                                        |
| ------ | --------------------------------------------------------------- |
| GET    | [`/api/v1/devices`](#get-apiv1devices)                          |
| POST   | [`/api/v1/devices`](#post-apiv1devices)                         |
| GET    | [`/api/v1/devices/:id`](#get-apiv1devicesid)                    |
| SSE    | [`/api/v1/events/device-update`](#get-apiv1eventsdevice-update) |
| POST   | [`/api/v1/picow`](#post-apiv1picow)                             |

##  3. <a name='ApiV1'></a>Api V1

> package: [v1](internal/api/v1)  
> package: [pico](internal/api/v1/pico)

###  3.1. <a name='GET_apiv1devices_'></a>**GET** _"/api/v1/devices"_

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

###  3.2. <a name='POST_apiv1devices_'></a>**POST** _"/api/v1/devices"_

Update device(s) duty for rgbw pins

Example Request

```bash
curl -X POST http://localhost:50833/api/v1/devices \
  -H 'Content-Type: application/json'\
  -d '[{ "addr": "192.168.178.50:80", "rgbw": [100,100,100,100] }]'
```

No Response (http status 200 on success)

###  3.3. <a name='GET_apiv1devices:id_'></a>**GET** _"/api/v1/devices/:id"_

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

###  3.4. <a name='GET_apiv1eventsdevice-update_'></a>**GET** _"/api/v1/events/device-update"_

> work in progress

###  3.5. <a name='POST_apiv1picow_'></a>**POST** _"/api/v1/picow"_

Register a new raspberry pico device on the web server to handle.
(The [picow-rgbw](https://github.com/knackwurstking/picow-rgbw.git) will do
this automatically on startup if configured)

Example Request: _(Only picow devices will send requests here)_

```bash
curl -X POST http://localhost:50833/api/v1/picow\
  -H "Content-Type: application/json"\
  -d '{"addr": "192.168.178.50:80"}'
```

##  4. <a name='TODOs'></a>TODOs

- [x] frontend: control: add a color storage
- [ ] frontend: devices: add some device setup popup for changing the name ('host:port')
- [ ] backend: Fix crash if device goes offline?
- [ ] backend: if device is offline on a set duty cycle command set configured pins first
