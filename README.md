# Tor Scraper

Bu proje, Tor ağı üzerinden .onion uzantılı web sitelerinden HTML içerik çekmek ve ekran görüntüsü  amacıyla Go dili kullanılarak geliştirilmiştir.

---
## Projenin Amacı

- Tor ağı üzerinden dark web sitelerine erişim sağlamak  
- .onion sitelerinin HTML içeriklerini yerel olarak kaydetmek  
- Erişilebilen siteler için ekran görüntüsü almak   
---
## Özellikler

- targets.yaml dosyasından hedef URL okuma  
- Tor proxy üzerinden HTTP istekleri  
- Hata durumlarını loglama  
---
## Gereksinimler

- Go 1.20 veya üzeri  
- Tor servisi
---

## Kullanım

1. Tor servisinin çalıştığından emin olun  
2. targets.yaml dosyasına hedef .onion adreslerini ekleyin  
3. Proje dizininde aşağıdaki komutu çalıştırın:

```bash
go run .
