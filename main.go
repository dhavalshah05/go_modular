package main

import (
	"encoding/json"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
	"os"
	"strings"
)

type Channel struct {
	Number   string
	Name     string
	Language string
	Genre    string
	App      string
}

func (c Channel) PrintDetails() {
	fmt.Println(c)
}

func getChannelFromHandle(entry playwright.ElementHandle) (*Channel, error) {
	pHandles, err := entry.QuerySelectorAll("p")
	if err != nil {
		return nil, err
	}

	var channel = &Channel{}
	for pIndex, pHandle := range pHandles {
		content, err := pHandle.TextContent()
		if err != nil {
			continue
		}
		if pIndex == 0 {
			channel.Number = strings.TrimSpace(content)
		} else if pIndex == 1 {
			channel.Name = strings.TrimSpace(content)
		} else if pIndex == 2 {
			channel.Language = strings.TrimSpace(content)
		} else if pIndex == 3 {
			channel.Genre = strings.TrimSpace(content)
		} else if pIndex == 4 {
			channel.App = strings.TrimSpace(content)
		}
	}
	if channel.Number != "" && channel.Name != "Channel Name" {
		return channel, nil
	}
	return nil, nil
}

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	var headless = false
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: &headless,
		Args:     []string{"--start-maximized"},
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage(playwright.BrowserNewContextOptions{})
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	if _, err = page.Goto("https://www.dthhelp.net/dth/jio-tv-channel-list-language-wise.html"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	entries, err := page.QuerySelectorAll("body > article > table > tbody > tr")
	if err != nil {
		panic(err)
	}

	var channels []Channel

	for _, entry := range entries {
		channel, err := getChannelFromHandle(entry)
		if err != nil {
			panic(err)
		}
		if channel != nil {
			channels = append(channels, *channel)
		}
	}

	err = saveChannels(channels)
	if err != nil {
		panic(err)
	}

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}

func saveChannels(channels []Channel) error {
	jsonBytes, err := json.Marshal(channels)
	if err != nil {
		return err
	}
	err = os.WriteFile("channels.json", jsonBytes, 0777)
	if err != nil {
		return err
	}
	return nil
}
