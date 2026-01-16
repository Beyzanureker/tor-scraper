package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Log dosyasını oluştur ve yapılandır
	logFile, _ := os.OpenFile("scan_report.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer logFile.Close()
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	log.Println("[SİSTEM] Tor Scraper başlatılıyor...")

	// Gerekli klasörleri oluştur [cite: 42]
	os.MkdirAll("output/html", 0755)
	os.MkdirAll("output/screenshots", 0755)

	// Görevleri sırasıyla çalıştır
	runScraper()
	runScreenshots()

	log.Println("[SİSTEM] Tüm tarama görevleri tamamlandı.")
}
