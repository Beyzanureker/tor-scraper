Bu proje, Tor ağı üzerinden .onion uzantılı web sitelerinden HTML içerik çekmek ve ekran görüntüsü almak amacıyla Go dili kullanılarak geliştirilmiştir.

##Projenin Amacı

-Tor ağı üzerinden dark web sitelerine erişim sağlamak
-.onion sitelerinin HTML içeriklerini yerel olarak kaydetmek
-Erişilebilen siteler için ekran görüntüsü almak

##Gereksinimler

-Go 1.20 veya üzeri
-Tor servisinin çalışıyor olması

##Kullanım

-Tor servisini başlatın
-targets.yaml dosyasına hedef .onion adreslerini ekleyin
-Uygulamayı çalıştırın:
    go run .
