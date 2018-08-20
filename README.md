# IZIGo [![Build Status](https://travis-ci.org/diepdt/izigo.svg?branch=master)](https://travis-ci.org/diepdt/izigo) [![GoDoc](http://godoc.org/github.com/izi-global/izigo?status.svg)](http://godoc.org/github.com/izi-global/izigo) [![Foundation](https://img.shields.io/badge/Golang-Foundation-green.svg)](http://golangfoundation.org) [![Go Report Card](https://goreportcard.com/badge/github.com/izi-global/izigo)](https://goreportcard.com/report/github.com/izi-global/izigo)


izigo is used for rapid development of RESTful APIs, web apps and backend services in Go.
It is inspired by Tornado, Sinatra and Flask. izigo has some Go-specific features such as interfaces and struct embedding.

###### More info at [go.izi.asia](http://go.izi.asia).

## Quick Start

#### Download and install

    go get github.com/izi-global/izigo

#### Create file `hello.go`
```go
package main

import "github.com/izi-global/izigo"

func main(){
    izigo.Run()
}
```
#### Build and run

    go build hello.go
    ./hello

#### Go to [http://localhost:8080](http://localhost:8080)

Congratulations! You've just built your first **izigo** app.

###### Please see [Documentation](http://go.izi.asia/docs) for more.

## Features

* RESTful support
* MVC architecture
* Modularity
* Auto API documents
* Annotation router
* Namespace
* Powerful development tools
* Full stack for Web & API

## Documentation

* [English](http://go.izi.asia/docs/intro/)
* [中文文档](http://go.izi.asia/docs/intro/)
* [Русский](http://go.izi.asia/docs/intro/)

## Community

* [http://go.izi.asia/community](http://go.izi.asia/community)
* Welcome to join us in Slack: [https://izigo.slack.com](https://izigo.slack.com), you can get invited from [here](https://github.com/izi-global/izidoc/issues/232)

## License

izigo source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).
