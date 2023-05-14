package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
	"github.com/knackwurstking/picow-rgbw-web/pkg/log"
	"github.com/knackwurstking/picow-rgbw-web/pkg/scanner"
	"github.com/knackwurstking/picow-rgbw-web/pkg/server"
)

const (
	applicationName = "picow-rgbw-web"
)

var config = Config{
	Port:    50833,
	Debug:   isDebug(),
	Handler: pico.NewHandler(),
	Version: false,
}

func isDebug() bool {
	switch os.Getenv("DEBUG") {
	case "true", "1", "yes":
		return true
	default:
		return false
	}
}

type Config struct {
	Host    string        `json:"host"`
	Port    int           `json:"port"`
	HTTP    bool          `json:"http"`
	Debug   bool          `json:"debug"`
	Handler *pico.Handler `json:"pico-handler"`
	Version bool          `json:"-"`
}

func init() {
	initConfig()
	initFlags()
	initPicoDevices()
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Error.Printf("Failed to get the users home directory: %s", err.Error())
		return
	}

	// Load "config.json" first. If missing, load user config "~/.config/picow-rgbw-web/config.json".
	f, err := os.Open("config.json")
	if err != nil {
		f, err = os.Open(filepath.Join(home, ".config", applicationName, "config.json"))
	}

	if err == nil {
		err = json.NewDecoder(f).Decode(&config)
	}

	if err != nil {
		log.Error.Printf("Load configuration failed: %s", err.Error())
	} else {
		for _, d := range config.Handler.Devices {
			d.SSE = &config.Handler.SSE
		}
	}

	log.Debug.(*log.DebugLogger).Enabled = config.Debug
}

func initFlags() {
	flag.StringVar(&config.Host, "host", config.Host, "Server host.")
	flag.IntVar(&config.Port, "port", config.Port, "Server port.")
	flag.BoolVar(&config.HTTP, "http", config.HTTP, "Start HTTP server.")
	flag.BoolVar(&config.Debug, "debug", config.Debug, "Enable debug log.")
	flag.BoolVar(&config.Version, "version", config.Version,
		"Display verison and exit.")

	// TODO: add "--version" flag
	// TODO: Add flags for... (Need to finish the scanner first)
	//	...scan - enables the pico device scan
	//	...scan-range - 192.168.178.0 or 192.168.0.0

	flag.Parse()

	if config.Version {
		echoVersion()
		os.Exit(0)
	}
}

func initPicoDevices() {
	updated := false

	for _, device := range config.Handler.Devices {
		log.Debug.Printf("Init pico device %+v", device)

		for _, pin := range device.RGBW {
			if pin != nil {
				doUpdateDevices(device)
				updated = true
				break
			}
		}

		if updated {
			break
		}

		doSetupDevices(device)
	}

	// Start the devices scanner
	if ip, err := scanner.GetLocalIP(); err != nil {
		log.Warn.Println(err.Error())
	} else {
		ip = strings.Join(strings.Split(ip, ".")[:3], ".") + ".0"
		log.Debug.Printf("Scan for pico devices (scan-range: %s)", ip)

		// NOTE: Scan method is work in progress
		if devices, err := config.Handler.Scan(ip); err != nil {
			log.Warn.Println(err.Error())
		} else {
			config.Handler.Devices = devices
		}
	}
}

func doUpdateDevices(device *pico.Device) {
	pins := [4]pico.GpPin{
		pico.GpPinDisabled,
		pico.GpPinDisabled,
		pico.GpPinDisabled,
		pico.GpPinDisabled,
	}
	duty := [4]pico.Duty{
		pico.DutyMin,
		pico.DutyMin,
		pico.DutyMin,
		pico.DutyMin,
	}

	for i, p := range device.RGBW {
		if p == nil {
			continue
		}

		pins[i] = p.Nr
		duty[i] = p.Duty
	}

	log.Debug.Printf("Set pins (%v): %+v", pins, device)
	err := device.SetGpPins(pins)
	if err != nil {
		log.Error.Printf("Set pins (%v): %s", pins, err.Error())
	} else {
		log.Debug.Printf("Set duty (%v): %+v", duty, device)
		err = device.SetColor(duty)
		if err != nil {
			log.Error.Println("Set duty (%v): %s", duty, err.Error())
		}
	}
}

func doSetupDevices(device *pico.Device) {
	// get pins, even after set pins (in case of a failure)
	err := device.GetGpPins()
	if err != nil {
		log.Error.Printf("Get pins failed: %s", err.Error())
	}

	err = device.GetColor()
	if err != nil {
		log.Error.Printf("Get duty failed: %s", err.Error())
	}
}

func main() {
	// Get server (with handler)
	server := server.New(
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		config.Handler,
	)

	// Start server (HTTP or HTTPS)
	if config.HTTP {
		log.Info.Printf("HTTP server running: %s", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Error.Printf("Server error: %s", err.Error())
			os.Exit(1)
		}
	} else {
		crt := os.Getenv("CERT_FILE")
		if crt == "" {
			log.Error.Println("Missing server certificate (env: CERT_FILE)")
			os.Exit(1)
		}

		key := os.Getenv("KEY_FILE")
		if key == "" {
			log.Error.Println("Missing server certificate key (env: KEY_FILE)")
			os.Exit(1)
		}

		log.Info.Printf("HTTPS server running: %s", server.Addr)
		if err := server.ListenAndServeTLS(crt, key); err != nil {
			log.Error.Printf("Server error: %s", err.Error())
			os.Exit(1)
		}
	}
}
