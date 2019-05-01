// Automated tests for davidgossetx.com.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	selenium "github.com/bunsenapp/go-selenium"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nAutomated testing via selenoid for davidgossetx.com\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
	}
	pageBaseURL := flag.String("p", "http://localhost:1313", "page base URL")
	seleniumURL := flag.String("s", "http://localhost:4444/wd/hub", "Selenium hub URL")
	flag.Parse()

	// Create Selenium Web Driver and session.
	caps := selenium.Capabilities{}
	caps.SetBrowser(selenium.FirefoxBrowser())
	wd, err := selenium.NewSeleniumWebDriver(*seleniumURL, caps)
	if err != nil {
		log.Fatalf("Error creating new selenium webdriver: %s\n", err)
	}
	_, err = wd.CreateSession()
	if err != nil {
		log.Fatalf("Error creating new webdriver session: %s\n", err)
	}
	defer wd.DeleteSession()

	if status, err := wd.SessionStatus(); err != nil {
		log.Fatalf("Session Status error: %s %s\n", status, err)
	}

	// Access endpoints within session
	endpoints := []string{
		"",
		"publications",
		"teaching",
		"about",
	}
	for _, ep := range endpoints {
		url := strings.Join([]string{*pageBaseURL, ep}, "/")
		//if _, err := wd.Go(url); err != nil {
		if _, err := wd.Go("https://www.google.com"); err != nil {
			fmt.Printf("Error visiting URL: %s\n", err)
		}

		shot, err := wd.Screenshot()
		if err != nil {
			log.Println(shot)
		}
		img, err := shot.ImageBytes()
		if err != nil {
			log.Fatalf("Error writing screenshot image: %s\n", err)
		}
		if err = ioutil.WriteFile(ep+".png", img, 0644); err != nil {
			log.Fatalf("Error writing screenshot image to disk: %s\n", err)
		}

		log.Printf("Successfully navigated to %s !\n", url)
	}
}
