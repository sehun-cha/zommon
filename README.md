# Zommon

A Go logging library wrapper built on top of [log15](https://github.com/inconshreveable/log15) that provides structured logging with both terminal and file output support.

## Features

- **Dual Output Support**: Log to both terminal (stdout) and file simultaneously
- **Structured Logging**: Built on log15's key-value pair logging approach
- **Configurable Verbosity**: Different log levels for terminal and file output
- **File Rotation**: Automatic log file rotation using lumberjack
- **JSON & Terminal Formatting**: JSON format for files, colored terminal format for console
- **Easy Configuration**: Simple setup with a single function call

## Installation

```bash
go get github.com/sehun-cha/zommon
```

## Quick Start

```go
package main

import (
    "github.com/sehun-cha/zommon/logger"
    "github.com/sehun-cha/zommon/lib/log"
)

func main() {
    // Configure logging
    logger.SetLogRoot(
        true,           // useTerminal: enable terminal output
        true,           // useFile: enable file output
        4,              // verbosityTerminal: log level for terminal (0=Crit, 1=Error, 2=Warn, 3=Info, 4=Debug)
        2,              // verbosityFile: log level for file
        "./app.log",    // filePath: path to log file
    )

    // Start logging
    log.Info("Application started", "version", "1.0.0")
    log.Debug("Debug message", "user_id", 12345)
    log.Warn("Warning message", "retry_count", 3)
    log.Error("Error occurred", "error", "connection timeout")
}
```

## Configuration

### SetLogRoot Parameters

```go
func SetLogRoot(useTerminal bool, useFile bool, verbosityTerminal int, verbosityFile int, filePath string)
```

- **useTerminal**: Enable/disable terminal output
- **useFile**: Enable/disable file output  
- **verbosityTerminal**: Log level for terminal output (0-4)
- **verbosityFile**: Log level for file output (0-4)
- **filePath**: Path to the log file

### Log Levels

| Level | Value | Description |
|-------|-------|-------------|
| Critical | 0 | System is unusable |
| Error | 1 | Error conditions |
| Warning | 2 | Warning conditions |
| Info | 3 | Informational messages |
| Debug | 4 | Debug-level messages |

## File Rotation

The library automatically handles log file rotation with the following settings:

- **MaxSize**: 1024 MB per file
- **MaxBackups**: 5 backup files
- **MaxAge**: 28 days
- **Compression**: Enabled

## Output Formats

### Terminal Output
Colored, human-readable format:
```
INFO[06-17|21:58:10] Application started    version=1.0.0
WARN[06-17|21:58:11] Warning message        retry_count=3
```

### File Output
JSON format for machine processing:
```json
{"lvl":"info","t":"2024-01-01T21:58:10.123Z","msg":"Application started","version":"1.0.0"}
{"lvl":"warn","t":"2024-01-01T21:58:11.456Z","msg":"Warning message","retry_count":3}
```

## Advanced Usage

### Context Loggers

Create loggers with persistent context:

```go
// Create a logger with context
userLogger := log.New("user_id", 12345, "session", "abc123")

// All messages from this logger will include the context
userLogger.Info("User action", "action", "login")
userLogger.Debug("Processing request", "endpoint", "/api/users")
```

### Conditional Logging

Use different configurations for different environments:

```go
func setupLogging(env string) {
    switch env {
    case "production":
        logger.SetLogRoot(false, true, 0, 1, "/var/log/app.log") // File only, errors
    case "development":
        logger.SetLogRoot(true, true, 4, 4, "./dev.log")        // Both, debug level
    case "testing":
        logger.SetLogRoot(true, false, 2, 0, "")                // Terminal only, warnings
    }
}
```

## Dependencies

- [go-stack/stack](https://github.com/go-stack/stack) v1.8.1 - Stack trace utilities
- [natefinch/lumberjack](https://gopkg.in/natefinch/lumberjack.v2) v2.2.1 - Log rotation

## License

This project is based on [log15](https://github.com/inconshreveable/log15) which is licensed under the Apache License.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Examples

Check out the `examples/` directory for more usage examples and best practices.
