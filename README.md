# log4Go

`log4Go` is a lightweight logging library written in Go, designed to provide simple and effective logging capabilities for your Go applications.

## Features

- **Easy Integration:** Quickly integrate `log4Go` into your existing projects with minimal setup.
- **Customizable Logging Levels:** Supports various logging levels such as INFO, WARN, ERROR, etc.
- **Formatted Output:** Customize the log output format to suit your needs.
- **File and Console Logging:** Supports logging to both files and the console.

## Installation

To install `log4Go`, run:

```bash
go get github.com/KanakSasak/log4Go
```

```go
package main

import (
    "github.com/KanakSasak/log4Go"
)

func main() {
    logger := log4Go.NewLogger()
    logger.Info("This is an info message.")
    logger.Warn("This is a warning message.")
    logger.Error("This is an error message.")
}

```

![screenshot](log4Go.png)
