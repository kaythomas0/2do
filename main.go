package main

import (
	"io"
	"log"
	"os"

	"github.com/namsral/flag"
	"github.com/prologic/bitcask"
)

var (
	db *bitcask.Bitcask
)

func main() {
	var (
		bind string
	)

	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.Parse()

	var err error
	db, err = bitcask.Open("./db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	selectColorTheme()

	newServer(bind).listenAndServe()
}

func selectColorTheme() {
	colorTheme := os.Getenv("COLOR_THEME")

	if colorTheme == "custom" {
		pageBackgroundColor := os.Getenv("COLOR_PAGE_BACKGROUND")
		inputBackgroundColor := os.Getenv("COLOR_INPUT_BACKGROUND")
		foregroundColor := os.Getenv("COLOR_FOREGROUND")
		checkMarkColor := os.Getenv("COLOR_CHECK_MARK")
		xMarkColor := os.Getenv("COLOR_X_MARK")
		labelColor := os.Getenv("COLOR_LABEL")

		customThemeFile, err := os.OpenFile("./static/css/color-theme.css", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer customThemeFile.Close()

		_, err = customThemeFile.WriteString(":root {" +
			"\n\t--page-background: #" + pageBackgroundColor + ";" +
			"\n\t--input-background: #" + inputBackgroundColor + ";" +
			"\n\t--foreground: #" + foregroundColor + ";" +
			"\n\t--check: #" + checkMarkColor + ";" +
			"\n\t--x: #" + xMarkColor + ";" +
			"\n\t--label: #" + labelColor + ";" +
			"\n}")
	} else {
		if colorTheme == "" {
			colorTheme = "dracula"
		}

		from, err := os.Open("./static/color-themes/" + colorTheme + ".css")
		if err != nil {
			log.Fatal(err)
		}
		defer from.Close()

		to, err := os.OpenFile("./static/css/color-theme.css", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer to.Close()

		_, err = io.Copy(to, from)
		if err != nil {
			log.Fatal(err)
		}
	}
}
