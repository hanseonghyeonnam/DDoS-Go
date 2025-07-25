# DDoS-Go (dGo)

`DDoS-Go` is a simple DDoS (Distributed Denial of Service) simulation tool written in Go. It is designed for network stress testing and security training. **Use only for educational purposes.**

---

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Options](#options)
- [Examples](#examples)
- [Disclaimer](#disclaimer)
- [License](#license)

---

## Features

- Sends high-rate HTTP GET requests to a single URL per second
- Lightweight and user-friendly CLI interface
- Leverages Go goroutines for high-concurrency performance

## Installation

1. Install Go 1.18 or higher
2. Clone this repository:
   ```bash
   git clone https://github.com/hanseonghyeonnam/DDoS-Go.git
   cd DDoS-Go
   ```
3. Build the binary:
   ```bash
   go build -o dgo
   ```

## Usage

```bash
./dgo -url <TARGET_URL> -DPS <REQUESTS_PER_SECOND>
```

## Options

| Option     | Description                         | Required | Default |
| ---------- | ----------------------------------- | -------- | ------- |
| `-url`     | Target URL to attack                | Yes      | http://localhost:3000       |
| `-DPS`     | Requests per second (DPS)           | Yes      | 10/S                             |


## Supporting Protocols
 | Procotol | Supporting |
 | -------- | ---------- |
 | `http`   |     Y      |
 | `https`  |     N      |
 | `ETC` |     N      |
## Examples

1. Send 50 requests per second to `http://example.com`:

   ```bash
   ./dgo -url http://example.com -DPS 50/s
   ```

2. Set a 10-second timeout per request:

   ```bash
   ./dgo -url http://example.com -DPS 10/s
   ```

## Disclaimer

- **For educational and security testing only.**
- Unauthorized use against targets is **illegal** and may result in legal consequences.
- Always obtain permission from network or system administrators before use.

## License

Distributed under the [MIT License](LICENSE). © 2025 hanseonghyeonnam



