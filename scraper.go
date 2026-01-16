package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

func runScraper() {
	log.Println("[SCRAPER] HTML kazıma işlemi başladı.")

	// Tor Proxy Bağlantısı (Port: 9150)
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9150", nil, proxy.Direct)
	if err != nil {
		log.Printf("[HATA] Tor Proxy hatası: %v\n", err)
		return
	}

	client := &http.Client{
		Transport: &http.Transport{Dial: dialer.Dial},
		Timeout:   90 * time.Second, // Tor yavaşlığı için hocalarının kullandığı süre
	}

	file, _ := os.Open("targets.yaml")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue
		}

		log.Printf("[SCRAPER] taranıyor: %s\n", url)
		resp, err := client.Get(url)
		if err != nil {
			log.Printf("[HATA] %s ulaşılamadı: %v\n", url, err)
			continue
		}

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		filename := strings.ReplaceAll(strings.TrimPrefix(url, "http://"), "/", "_") + ".html"
		os.WriteFile("output/html/"+filename, body, 0644)
	}
}
