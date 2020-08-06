# todo
[![Build Status](https://cloud.drone.io/api/badges/prologic/todo/status.svg)](https://cloud.drone.io/prologic/todo)
[![GoDoc](https://godoc.org/github.com/prologic/todo?status.svg)](https://godoc.org/github.com/prologic/todo)
[![Go Report Card](https://goreportcard.com/badge/github.com/prologic/todo)](https://goreportcard.com/report/github.com/prologic/todo)
[![CodeCov](https://codecov.io/gh/prologic/todo/branch/master/graph/badge.svg)](https://codecov.io/gh/prologic/todo)
[![Sourcegraph](https://sourcegraph.com/github.com/prologic/msgbus/-/badge.svg)](https://sourcegraph.com/github.com/prologic/msgbus?badge)
[![Docker Version](https://images.microbadger.com/badges/version/prologic/todo.svg)](https://microbadger.com/images/prologic/todo)
[![Image Info](https://images.microbadger.com/badges/image/prologic/todo.svg)](https://microbadger.com/images/prologic/todo)

todo is a self-hosted todo web app that lets you keep track of your todos in a easy and minimal way. 📝

## Screenshots
_Nord Theme_

<img src="screenshots/mobile-nord.png" alt="Mobile Nord Theme" height="500"/>
<img src="screenshots/desktop-nord.png" alt="Desktop Nord Theme" height="500"/>
<br />

_Dracula Theme_

<img src="screenshots/mobile-dracula.png" alt="Mobile Dracula Theme" height="500"/>
<img src="screenshots/desktop-dracula.png" alt="Desktop Dracula Theme" height="500"/>

See all themes in the "Preset Color Themes" section below

## Demo
There is also a public demo instance avilable at: [https://todo.mills.io](https://todo.mills.io)

## Deployment

### Docker Compose
`docker-compose.yml`
```
version: '3'

services:
  todo:
    image: prologic/todo
    container_name: todo
    restart: always
    ports:
      - 8000:8000
    volumes:
      - todo_db:/usr/local/go/src/todo/todo.db

volumes:
  todo_db:
```
This file:
* Creates the `todo` container using the latest image from `prologic/todo` in [Docker Hub](https://hub.docker.com/r/prologic/todo).
* Binds port 8000 on your host machine to port 8000 in the container (you may change the host port to whatever you wish).
* Volume mounts the database path, saving your todo items so that your todo list will be saved in between container restarts.

Bring the container up with:
```
$ docker-compose up
```

### Docker
Alternatively, you can run the container without docker-compose:
```
$ docker run -p 8000:8000 -v todo_db:/usr/local/go/src/todo/todo.db prologic/todo
```

## Configuration

### Preset Color Themes
todo comes with 12 different color themes based on some of the most popular programming themes:
```
ayu, dracula, gruvbox-dark, gruvbox-light, lucario, monokai, nord, solarized-dark, solarized-light, tomorrow, tomorrow-night, zenburn
```

You can set the theme by passing the `COLOR_THEME` environment variable to the docker container, for example:

`docker-compose.yml`
```
version: '3'

services:
  todo:
    image: prologic/todo
    container_name: todo
    environment:
      COLOR_THEME: ayu
    restart: always
    ports:
      - 8000:8000
    volumes:
      - todo_db:/usr/local/go/src/todo/todo.db

volumes:
  todo_db:
```

### Custom Color Themes
You can set your own color theme by passing in the appropriate environment variables.

Set the `COLOR_THEME` environment variable to `custom`, and the five following environment variables to the colors of your choice (in hex format, omitting the `#`):

| Environment Variable           | Description                       |
|--------------------------------|-----------------------------------|
| COLOR_PAGEBACKGROUND           | Web page background               |
| COLOR_INPUTBACKGROUND          | Text boxes and buttons background |
| COLOR_FOREGROUND               | Input and item text               |
| COLOR_CHECK                    | Check mark on button              |
| COLOR_X                        | X mark on button                  |
| COLOR_LABEL                    | Heading text and button hover     |

An example configuration:

`docker-compose.yml`
```
version: '3'

services:
  todo:
    image: prologic/todo
    container_name: todo
    environment:
      COLOR_THEME: custom
      COLOR_PAGEBACKGROUND: 282a36
      COLOR_INPUTBACKGROUND: 44475a
      COLOR_FOREGROUND: f8f8f2
      COLOR_CHECK: 50fa7b
      COLOR_X: ff5555
      COLOR_LABEL: ffffff
    restart: always
    ports:
      - 8000:8000
    volumes:
      - todo_db:/usr/local/go/src/todo/todo.db

volumes:
  todo_db:
```

### Additional Configuration
| Environment Variable           | Description                                      | Default Value |
|--------------------------------|--------------------------------------------------|---------------|
| MAXITEMS                       | Maximum number of items allowed in the todo list | 100           |
| MAXTITLELENGTH                 | Maximum length of a todo list item               | 100           |

## Development / Non-Dockerized Deploy
You can quickly run a todo instance from source using the Makefile:
```
$ git clone https://github.com/prologic/todo.git
$ cd todo
$ make
```
Then todo will be running at: http://localhost:8000

By default todo stores todos in `todo.db` in the local directory.

This can be configured with the `-dbpath /path/to/todo.db` option.

All the environment variables that you set will be read upon running the Makefile.
For example:
```
$ export COLOR_THEME=nord
$ make
```
^ Will run the application with the nord color theme.

## License
MIT

Icon made by [Smashicons](https://smashicons.com/) from [flaticon.com](https://flaticon.com)
