# 2do
Minimal self-hosted todo list forked from [prologic/todo](https://github.com/prologic/todo) aimed at being modern and mobile friendly.

## Screenshots
_Nord Theme_

<img src="mobile-nord.png" alt="Mobile Nord Theme" height="500"/>
<img src="desktop-nord.png" alt="Desktop Nord Theme" height="500"/>
<br />

_Dracula Theme_

<img src="mobile-dracula.png" alt="Mobile Dracula Theme" height="500"/>
<img src="desktop-dracula.png" alt="Desktop Dracula Theme" height="500"/>

See all themes in the "Preset Color Themes" section below

## Demo
Try out a public demo instance here: https://2do-demo.page

## Deployment

### Docker Compose (Recommended Method)
`docker-compose.yml`
```
version: '3'

services:
  2do:
    image: kevinnthomas/2do
    container_name: 2do
    ports:
      - 8000:8000
    volumes:
      - db:/2do/db

volumes:
  db:
```
This file:
* Creates the `2do` container using the latest image from `kevinnthomas/2do` in [Docker Hub](https://hub.docker.com/r/kevinnthomas/2do).
* Binds port 8000 on your host machine to port 8000 in the container (you may change the host port to whatever you wish).
* Volume mounts the database path, saving your todo items so that your todo list will be saved in between container restarts.

Bring the container up with:
```
$ docker-compose up
```

### Docker
Alternatively, you can run the container without docker-compose:
```
$ docker run -p 8000:8000 -v db:/2do/db kevinnthomas/2do
```

## Configuration

### Preset Color Themes
2do comes with 12 different color themes based on some of the most popular programming themes:
```
ayu, dracula, gruvbox-dark, gruvbox-light, lucario, monokai, nord, solarized-dark, solarized-light, tomorrow, tomorrow-night, zenburn
```

You can set the theme by passing the `COLOR_THEME` environment variable to the docker container, for example:

`docker-compose.yml`
```
version: '3'

services:
  2do:
    image: kevinnthomas/2do
    container_name: 2do
    environment:
      COLOR_THEME: ayu
    ports:
      - 8000:8000
    volumes:
      - db:/2do/db

volumes:
  db:
```

#### Screenshots

You can find screenshots and the CSS of all the preset color themes in both mobile and desktop views on the [Wiki Page](https://gitlab.com/KevinNThomas/2do/-/wikis/Color-Themes).

### Custom Color Themes
You can set your own color theme by passing in the appropriate environment variables.

Set the `COLOR_THEME` environment variable to `custom`, and the five following environment variables to the colors of your choice (in hex format, omitting the `#`):

| Environment Variable           | Description                       |
|--------------------------------|-----------------------------------|
| COLOR_PAGE_BACKGROUND          | Web page background               |
| COLOR_INPUT_BACKGROUND         | Text boxes and buttons background |
| COLOR_FOREGROUND               | Input and item text               |
| COLOR_CHECK_MARK               | Check mark on button              |
| COLOR_X_MARK                   | X mark on button                  |
| COLOR_LABEL                    | Heading text and button hover     |

An example configuration:

`docker-compose.yml`
```
version: '3'

services:
  2do:
    image: kevinnthomas/2do
    container_name: 2do
    environment:
      COLOR_THEME: custom
      COLOR_PAGE_BACKGROUND: 282a36
      COLOR_INPUT_BACKGROUND: 44475a
      COLOR_FOREGROUND: f8f8f2
      COLOR_CHECK_MARK: 50fa7b
      COLOR_X_MARK: ff5555
      COLOR_LABEL: ffffff
    ports:
      - 8000:8000
    volumes:
      - db:/2do/db

volumes:
  db:
```

### Additional Configuration
| Environment Variable           | Description                                      | Default Value |
|--------------------------------|--------------------------------------------------|---------------|
| MAX_ITEMS                      | Maximum number of items allowed in the todo list | 100           |
| MAX_TITLE_LENGTH               | Maximum length of a todo list item               | 100           |

## Development
You can quickly run a local version of 2do from source using the Makefile:
```
$ git clone https://gitlab.com/KevinNThomas/2do
$ cd 2do
$ make
```

## License
MIT

Icon made by [Smashicons](https://smashicons.com/) from [flaticon.com](https://flaticon.com).
