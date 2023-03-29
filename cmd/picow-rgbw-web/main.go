package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/exp/slog"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
	"github.com/knackwurstking/picow-rgbw-web/pkg/scanner"
	"github.com/knackwurstking/picow-rgbw-web/pkg/server"
)

const (
	applicationName = "picow-rgbw-web"
)

var (
	config = Config{
		Port:    50833,
		Debug:   isDebug(),
		Handler: pico.NewHandler(),
	}
)

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
}

func init() {
	initConfig()
	initFlags()
	initLogger()
	initPicoDevices()
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("failed to get the users home directory: " + err.Error())
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
		slog.Error("load config failed: " + err.Error())
	}
}

func initFlags() {
	flag.StringVar(&config.Host, "host", config.Host, "Server host.")
	flag.IntVar(&config.Port, "port", config.Port, "Server port.")
	flag.BoolVar(&config.HTTP, "http", config.HTTP, "Start HTTP server.")
	flag.BoolVar(&config.Debug, "debug", config.Debug, "Enable debug log.")

	// TODO: Add flags for... (Need to finish the scanner first)
	//	...scan - enables the pico device scan
	//	...scan-range - 192.168.178.0 or 192.168.0.0

	flag.Parse()
}

func initLogger() {
	o := slog.HandlerOptions{
		AddSource: true,
	}

	if config.Debug {
		o.Level = slog.LevelDebug
	} else {
		o.Level = slog.LevelInfo
	}

	// NOTE: Need a custom Text handler for this someday (with color support)
	h := o.NewJSONHandler(os.Stderr)
	slog.SetDefault(slog.New(h))
}

func initPicoDevices() {
	for _, device := range config.Handler.Devices {
		update := false

		slog.Debug(fmt.Sprintf("init pico device %+v", device))
		for _, pin := range device.RGBW {
			if pin != nil {
				update = true // A pin was set for this device
				break
			}
		}

		if update {
			pins := [4]int{-1, -1, -1, -1}
			for i, p := range device.RGBW {
				if p == nil {
					continue
				}
				pins[i] = p.Nr
			}
			if err := device.SetPins(pins); err != nil {
				slog.Error(err.Error())
			}
		}

		// get pins, even after set pins (in case of a failure)
		err := device.GetPins()
		if err != nil {
			slog.Error("get pins failed: " + err.Error())
		}

		err = device.GetDuty()
		if err != nil {
			slog.Error("get duty failed: " + err.Error())
		}
	}

	// Start the devices scanner
	if ip, err := scanner.GetLocalIP(); err != nil {
		slog.Warn(err.Error())
	} else {
		ip = strings.Join(strings.Split(ip, ".")[:3], ".") + ".0"
		slog.Debug(fmt.Sprintf("Scan for pico devices (scan-range: %s)", ip))

		// NOTE: Scan method is work in progress
		if devices, err := config.Handler.Scan(ip); err != nil {
			slog.Warn(err.Error())
		} else {
			config.Handler.Devices = devices
		}
	}
}

func main() {
	// Get server (with handler)
	server := server.New(fmt.Sprintf("%s:%d", config.Host, config.Port), config.Handler)

	// Start server (HTTP or HTTPS)
	if config.HTTP {
		slog.Info("HTTP server running " + server.Addr)
		if err := server.ListenAndServe(); err != nil {
			slog.Error("Server error: " + err.Error())
			os.Exit(1)
		}
	} else {
		crt := os.Getenv("CERT_FILE")
		if crt == "" {
			slog.Error("Missing server certificate (env: CERT_FILE)")
			os.Exit(1)
		}

		key := os.Getenv("KEY_FILE")
		if key == "" {
			slog.Error("Missing server certificate key (env: KEY_FILE)")
			os.Exit(1)
		}

		slog.Info("HTTPS server running " + server.Addr)
		if err := server.ListenAndServeTLS(crt, key); err != nil {
			slog.Error("Server error: " + err.Error())
			os.Exit(1)
		}
	}
}
