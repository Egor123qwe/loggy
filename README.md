Loggy
========

## Overview

Package for distributed logging with tracing by trace-id
Deploy microservices logs-storage and logs-viewer and start use !!!

### Create Cart

A new cart should be created and an ID generated. The new empty cart should be returned.

```sh
POST http://localhost:3000/carts -d '{}'
```

```json
{
	"id": 1,
	"items": []
}
```

### Add to cart

A new item should added to an existing cart. Should fail if the cart does not
exist, if the product name is blank, or if the quantity is non-positive. The
new item should be returned.

```sh
POST http://localhost:3000/carts/1/items -d '{
	"product": "Shoes",
	"quantity": 10
}'
```

```json
{
	"id": 1,
	"cart_id": 1,
	"product": "Shoes",
	"quantity": 10
}
```

### Remove from cart

An existing item should be removed from a cart. Should fail if the cart does not
exist or if the item does not exist.

```sh
DELETE http://localhost:3000/carts/1/items/1
```

```json
{}
```


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




