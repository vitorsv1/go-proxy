# HTTP Proxy with Logging

## Overview
This is a **basic HTTP proxy server** written in Go that logs all incoming HTTP requests. It captures essential details like the request method, URL, and headers, making it a useful starting point for learning about **networking, proxying, and logging in Go**.

## Features
- Forwards HTTP requests to their destination.
- Logs request details (method, URL, headers).
- Simple and lightweight, perfect for beginners.

## Installation & Usage
### Prerequisites
Ensure you have **Go installed** (version 1.18 or higher recommended). You can download it from [golang.org](https://golang.org/dl/).

### Clone the Repository
```sh
 git clone https://github.com/vitorsv1/go-proxy.git
 cd http-proxy-logger
```

### Run the Proxy Server
```sh
go run main.go
```
By default, the proxy listens on **port 8080**.

### Configure Your System to Use the Proxy
- **Browser:** Set `localhost:8080` as your HTTP proxy.
- **Curl:** Use the proxy in requests:
  ```sh
  curl -x http://localhost:8080 http://example.com
  ```
## Structure
 
cmd/  
   main.go           # Entry point for the application  
internal/  
proxy/               # Proxy folder logic  
   proxy.go          # Core proxy functionality  
   logger.go         # Logging
   filter.go         # Domain blocklist  
   config/           # Configuration settings
      config.go      # Manage proxy settings  
logs/                # Directory for storing log files  
tests/               # Unit tests for proxy functionality  
go.mod               # Go module file  
README.md            # Project documentation  


## Next Steps
Here are some enhancements you can add:
- Support for **HTTPS requests**
- Implement **domain blocklisting**
- Save logs to a **file** instead of printing to the console



