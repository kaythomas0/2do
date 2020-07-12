# Fork of [prologic/todo](https://github.com/prologic/todo) aimed at being more mobile friendly.

# todo
[![Build Status](https://cloud.drone.io/api/badges/prologic/todo/status.svg)](https://cloud.drone.io/prologic/todo)
[![GoDoc](https://godoc.org/github.com/prologic/todo?status.svg)](https://godoc.org/github.com/prologic/todo)
[![Go Report Card](https://goreportcard.com/badge/github.com/prologic/todo)](https://goreportcard.com/report/github.com/prologic/todo)
[![CodeCov](https://codecov.io/gh/prologic/todo/branch/master/graph/badge.svg)](https://codecov.io/gh/prologic/todo)
[![Sourcegraph](https://sourcegraph.com/github.com/prologic/msgbus/-/badge.svg)](https://sourcegraph.com/github.com/prologic/msgbus?badge)
[![Docker Version](https://images.microbadger.com/badges/version/prologic/todo.svg)](https://microbadger.com/images/prologic/todo)
[![Image Info](https://images.microbadger.com/badges/image/prologic/todo.svg)](https://microbadger.com/images/prologic/todo)

todo is a self-hosted todo web app that lets you keep track of your todos in a easy and minimal way. 📝

![animated screenshot](screenshot.gif)

## Quickstart
### Docker

You can also use the [Todo Docker Image](https://hub.docker.com/r/prologic/todo):

```#!bash
$ docker pull prologic/todo
$ docker run -d -p 8000:8000 prologic/todo
```

## Demo
There is also a public demo instance avilable at: [https://todo.mills.io](https://todo.mills.io)


## Installation
### Source

```#!bash
$ go get github.com/prologic/todo
```

## Usage
Run todo:

```#!bash
$ todo
```
Then visit: http://localhost:8000/

## Configuration
By default todo stores todos in `todo.db` in the local directory.

This can be configured with the `-dbpath /path/to/todo.db` option.

## License
MIT
Icon made by [Smashicons](https://smashicons.com/) from [flaticon.com](https://flaticon.com)
