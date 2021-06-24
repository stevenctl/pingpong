# pingpong

A simple TCP server and client.
This can be used to test for long lived connections being interrupted.

## Components

### Deployments

The `deployments/` directory contains two YAML files that contain nearly identical deployments/services
to run on Kubernetes that run with the `server` but also contain the `client` which can be used
by `exec`ing to the Pods.

### Server

The server will always respond with "pong" after receiving a line of input.

Usage:

``` 
server <opts>
-h|--host: the hostname to listen on (default "0.0.0.0")
-p|--port: the port to listen on (default "5000")
```

### Client 

The client  will send "ping", wait to receive a line of input, then sleep for 5 seconds before repeating.

Usage:

``` 
client <opts>
-h|--host: the hostname to connect to (default "localhost")
-p|--port: the port to connect to (default "5000")
-i|--interval: the polling interval to send "ping"s (default "5s")
```
