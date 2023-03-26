package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"golang.org/x/exp/slog"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
	"github.com/knackwurstking/picow-rgbw-web/pkg/server"
)

var (
	picoHandler *pico.Handler

	host  string      // Host server will listen to
	port  = 50833     // Port server will listen to
	http  bool        // HTTP server (no HTTPS)
	debug = isDebug() // Debug logger level
)

func isDebug() bool {
	switch os.Getenv("DEBUG") {
	case "true", "1", "yes":
		return true
	default:
		return false
	}
}

func init() {
	// TODO: Loading (json) configuration?
	picoHandler = pico.NewHandler() // TODO:: pass pre-configured pico devices

	initFlags()
	initLogger()

	if ip := localIP(); ip != "" {
		if devices, err := picoHandler.Scan(ip); err == nil {
			picoHandler.Devices = devices
		} else {
			slog.Error(err.Error())
			os.Exit(1)
		}
	} else {
		slog.Warn("Local ip address not found!")
	}
}

func initFlags() {
	flag.StringVar(&host, "host", host, "Server host.")
	flag.IntVar(&port, "port", port, "Server port.")
	flag.BoolVar(&http, "http", http, "Start HTTP server.")

	// TODO: Add flags for...
	//	...scan - enables the pico device scan
	//	...scan-range - 192.168.178.0 or 192.168.0.0

	if !debug {
		flag.BoolVar(&debug, "debug", debug, "Enable debug log.")
	}

	flag.Parse()
}

func initLogger() {
	o := slog.HandlerOptions{
		AddSource: true,
	}

	if debug {
		o.Level = slog.LevelDebug
	} else {
		o.Level = slog.LevelInfo
	}

	// NOTE: Need a custom Text handler for this someday (with color support)
	h := o.NewJSONHandler(os.Stderr)
	slog.SetDefault(slog.New(h))
}

func localIP() string {
	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.To4().String()
			}
		}
	}

	return ""
}

func main() {
	// Get server (with handler)
	server := server.New(fmt.Sprintf("%s:%d", host, port), picoHandler)

	// Start server (HTTP or HTTPS)
	if http {
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
