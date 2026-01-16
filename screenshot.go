package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func runScreenshots() {
	log.Println("[SCREENSHOT] Ekran görüntüsü görevi başladı.")

	file, _ := os.Open("targets.yaml")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("proxy-server", "socks5://127.0.0.1:9150"),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("ignore-certificate-errors", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue
		}

		log.Printf("[SCREENSHOT] Çekiliyor: %s\n", url)

		var buf []byte

		taskCtx, taskCancel := context.WithTimeout(ctx, 90*time.Second)

		err := chromedp.Run(taskCtx,
			chromedp.Navigate(url),
			chromedp.Sleep(5*time.Second),
			chromedp.FullScreenshot(&buf, 90),
		)
		taskCancel()

		if err != nil {
			log.Printf("[HATA] %s çekilemedi: %v\n", url, err)
			continue
		}

		filename := strings.ReplaceAll(strings.TrimPrefix(url, "http://"), "/", "_") + ".png"
		os.WriteFile("output/screenshots/"+filename, buf, 0644)
	}
}
