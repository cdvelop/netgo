# netgo

netgo is a Go library that provides utilities for basic network operations.

## Main Features

- **Get Host IP**: `GetHostIP() (string, error)`
  - Returns the current host's IP address
  - Handles errors if IP cannot be obtained

- **Assign IP based on machine**: `AssignIPBasedOnCurrentMachine(lasNumber string) (string, string, error)`
  - Generates an IP address based on the current machine
  - Useful for dynamic network configurations

- **Ping IP addresses**: `Ping(ip string) error`
  - Performs a ping to the specified IP address
  - Returns error if IP cannot be reached

## Installation


go get github.com/cdvelop/netgo


## Basic Usage


package main

import (
    "fmt"
    "github.com/cesar/netgo"
)

func main() {
    // Get host IP
    ip, err := netgo.GetHostIP()
    if err != nil {
        panic(err)
    }
    fmt.Println("Host IP:", ip)

    // Ping an IP
    err = netgo.Ping("8.8.8.8")
    if err != nil {
        fmt.Println("Ping failed:", err)
    } else {
        fmt.Println("Ping successful")
    }
}


## Contributing

Contributions are welcome. Please open an issue or pull request on GitHub.

## License

MIT License - See [LICENSE](LICENSE) for more details.