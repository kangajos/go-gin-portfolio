Hi.. 

Applikasi Prtfolio bersifat static yg di build dengan Go Gin Gonic

Anda boleh mengembangkan aplikasi ini lebih sempurna dan menyesuaiakan sesuai kebutuhan.

cara jalankan aplikasi

1. download atau clone github.com/kangajos/go-gin-portfolio.git
2. extract aplikasi yg Anda download
3. pastikan Anda sudah menginstall Golang Latest Version
4. lalu bukan command line interface (CMD/terminal) dan arahkan ke root project
5. copy file .env.example menjadi .env dan edit bagian APP_PORT jika Anda butuh jalan di port yg diingingkan
6. lalu jalankan perintah "go mod tidy" untuk download package yg di perlukan dalam aplikasi ini
7. jalankan server aplikasi dengan perintah "go run main.go"
8. selamat aplikasi sudah jalan di localhsot:APP_PORT port bisa lihat di file .env bagian APP_PORT