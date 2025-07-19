package main

import (
	"dgo/library"
	"flag"
	"log/slog"
	"fmt"
	"strings"
)

type Settings struct {
	URL string
	DPS string // DDoS Per Second, DPS
}

var (
	urlFlag *string
	dpsFlag *string
)

func ShowBanner() {
	banner := `
   ▐▌ ▗▄▄▖ ▄▄▄
   ▐▌▐▌   █   █
▗▞▀▜▌▐▌▝▜▌▀▄▄▄▀
▝▚▄▟▌▝▚▄▞▘      v1.0.0
`
	fmt.Printf("\033[36m%s\033[0m\n", banner)
}



func init() {
	// flag.String returns *string
	urlFlag = flag.String("url", "http://localhost:3000/", "Set the URL to be DDoSed")
	dpsFlag = flag.String("DPS", "10/S", "Set the DDoS-Per-Second, DPS.")
	flag.Parse()
}

func main() {
	// Log startup
	ShowBanner()
	fmt.Printf("\n")
	fmt.Printf("\033[31mTHIS CLI TOOL IS DESIGNED & CREATED FOR EDUCATION. WE WILL NOT BE LIABLE FOR ABUSE OF THIS TOOL.\033[0m\n\n")

	settings := Settings{
		URL: *urlFlag,
		DPS: *dpsFlag,
	}

	if ! strings.HasSuffix(settings.URL, "/") {
		settings.URL = settings.URL + "/"
	}

	if err := ddos.Run(settings.URL, settings.DPS); err != nil {
		slog.Error("DDoS attack failed", "error", err)
	}
}