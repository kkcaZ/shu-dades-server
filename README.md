# shu-dades-server

## Overview
This project has been created for the designing & developing enterprise systems module for Sheffield Hallam University. It is a basic stock tracking system that accepts TCP requests to carry out functions on the server.

## How to run
### Prerequisites
[go v1.21+](https://go.dev/doc/install)  
[docker](https://docs.docker.com/engine/install/)  

### Option 1  

> go run ./cmd/main.go  

### Option 2 

> docker build -t server .  
> docker run --network host server  

_Please note that some functionality may not work if run on windows with docker due to the lack of support for --network_