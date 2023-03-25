package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/knackwurstking/picow-rgbw-web/internal/server"
	"golang.org/x/exp/slog"
)

var (
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
	{ // flags
		flag.StringVar(&host, "host", host, "Server host.")
		flag.IntVar(&port, "port", port, "Server port.")
		flag.BoolVar(&http, "http", http, "Start HTTP server.")

		if !debug {
			flag.BoolVar(&debug, "debug", debug, "Enable debug log.")
		}

		flag.Parse()
	}

	{ // logger
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
}

func main() {
	// Get server (with handler)
	server := server.New(fmt.Sprintf("%s:%d", host, port))

	// Start server (HTTP or HTTPS)
	if http {
		if err := server.ListenAndServe(); err != nil {
			slog.Error("Server error: " + err.Error())
			os.Exit(1)
		}
	} else {
		crt := os.Getenv("HTTPS_CRT")
		if crt == "" {
			slog.Error("Missing https certificate (env: HTTPS_CRT)")
			os.Exit(1)
		}

		key := os.Getenv("HTTPS_KEY")
		if key == "" {
			slog.Error("Missing https certificate key (env: HTTPS_KEY)")
			os.Exit(1)
		}

		if err := server.ListenAndServeTLS(crt, key); err != nil {
			slog.Error("Server error: " + err.Error())
			os.Exit(1)
		}
	}
}
