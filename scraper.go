package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

func runScraper() {
	fmt.Println("[SCRAPER] HTML scraping started")

	_ = os.MkdirAll("output/html", 0755)

	report, _ := os.Create("scan_report.log")
	defer report.Close()

	// Tor proxy
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9150", nil, proxy.Direct)
	if err != nil {
		fmt.Println("Tor proxy error:", err)
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
		Timeout: 30 * time.Second,
	}

	file, err := os.Open("targets.yaml")
	if err != nil {
		fmt.Println("targets.yaml not found")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue
		}

		fmt.Println("[SCRAPER] Scanning:", url)

		resp, err := client.Get(url)
		if err != nil {
			report.WriteString(url + " -> DEAD\n")
			continue
		}

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		name := strings.ReplaceAll(url, "http://", "")
		name = strings.ReplaceAll(name, "/", "_")

		_ = os.WriteFile("output/html/"+name+".html", body, 0644)
		report.WriteString(url + " -> ACTIVE\n")
	}

	fmt.Println("[SCRAPER] HTML scraping finished")
}
