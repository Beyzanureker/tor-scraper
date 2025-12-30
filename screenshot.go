package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func runScreenshots() {
	fmt.Println("[SCREENSHOT] Screenshot task started")

	_ = os.MkdirAll("output/screenshots", 0755)

	file, err := os.Open("targets.yaml")
	if err != nil {
		fmt.Println("targets.yaml not found")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("proxy-server", "socks5://127.0.0.1:9150"),
		chromedp.Flag("ignore-certificate-errors", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue
		}

		fmt.Println("[SCREENSHOT] Capturing:", url)

		var buf []byte
		err := chromedp.Run(ctx,
			chromedp.Navigate(url),
			chromedp.Sleep(5*time.Second),
			chromedp.FullScreenshot(&buf, 90),
		)

		if err != nil {
			fmt.Println("[SCREENSHOT] FAILED:", url)
			continue
		}

		name := strings.ReplaceAll(url, "http://", "")
		name = strings.ReplaceAll(name, "/", "_")

		_ = os.WriteFile("output/screenshots/"+name+".png", buf, 0644)
		fmt.Println("[SCREENSHOT] Saved:", name)
	}

	fmt.Println("[SCREENSHOT] Screenshot task finished")
}
