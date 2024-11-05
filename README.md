Loggy
========

## Overview

Package for distributed logging with tracing by trace-id
Deploy microservices logs-storage and logs-viewer and start use !!!

See more examples in "example" package

### Init logger

```go
func init() {
    logOpts := logger.Options{
    Level:  logger.DEBUG,
    Module: "loggy",
    
    ToStderr: true,
    
    Server: []logger.Server{
        {
            // base api url
            URL: "http://localhost:8082",
            
            // logs chanel urls (broker hosts)
            LogsChannelsURLs: []string{"localhost:19092"},
            
            Credentials: &logger.Credentials{Username: "admin", Password: "admin"},
        },
	},
        
        File: []logger.File{
            {
                Name:      "test.log",
                MaxSizeMb: 100,
                MaxFiles:  10,
            },
        },
    }
    
    if err := logger.Init(logOpts); err != nil {
        log.Fatal(err)
    }
}
```

### Add Logs

```go
func main() {
    defer logger.Close()
    
    traceID := uuid.New().String()
    
    log := logger.Log(traceID)
    
    // tests
    log.Infof("Всё хорошо")
    log.Criticalf("Серверу не удалось подключиться")
}
```




