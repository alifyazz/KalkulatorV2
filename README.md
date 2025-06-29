# Kalkulator Scientific dengan Go

Aplikasi kalkulator scientific sederhana yang dibangun menggunakan Go, HTML, dan CSS. Kalkulator ini memiliki interface yang modern dan responsif dengan fungsi-fungsi matematika dasar dan scientific, serta fitur history perhitungan.

## Fitur

- **Operasi Dasar**: Penjumlahan, pengurangan, perkalian, pembagian
- **Fungsi Scientific**: 
  - Sinus (sin)
  - Cosinus (cos) 
  - Tangen (tan)
  - Logaritma basis 10 (log)
  - Logaritma natural (ln)
  - Akar kuadrat (sqrt)
  - Pangkat (pow)
- **History Perhitungan**: Menyimpan riwayat perhitungan terakhir (maksimal 10 item)
- **Interface Modern**: Desain glassmorphism dengan gradient background
- **Responsif**: Bekerja dengan baik di desktop dan mobile
- **Real-time Display**: Menampilkan expression dan hasil secara real-time

## Struktur Proyek

```
kalkulator2/
├── main.go              # Entry point aplikasi Go
├── templates/
│   └── index.html       # Template HTML untuk interface
├── static/
│   └── style.css        # File CSS untuk styling
├── go.mod               # Go module file
├── .gitignore           # Git ignore file
└── README.md           # Dokumentasi proyek
```

## Cara Menjalankan

1. **Pastikan Go sudah terinstall** di sistem Anda
2. **Clone atau download** proyek ini
3. **Buka terminal/command prompt** di direktori proyek
4. **Jalankan aplikasi** dengan perintah:
   ```bash
   go run main.go
   ```
5. **Buka browser** dan akses `http://localhost:8080`

## Penggunaan

### Operasi Dasar
- Klik tombol angka (0-9) untuk input angka
- Klik operator (+, -, ×, ÷) untuk operasi matematika
- Klik tombol `=` untuk melihat hasil

### Fungsi Scientific
- **sin**: Menghitung sinus (dalam derajat)
- **cos**: Menghitung cosinus (dalam derajat)  
- **tan**: Menghitung tangen (dalam derajat)
- **log**: Logaritma basis 10
- **ln**: Logaritma natural
- **√**: Akar kuadrat
- **x^y**: Pangkat (format: pow(base,exponent)

### Tombol Khusus
- **C**: Clear - menghapus input dan hasil, reset ke 0

### Fitur History
- **Otomatis**: Setiap perhitungan berhasil otomatis disimpan ke history
- **Tampilan**: History ditampilkan di panel sebelah kanan kalkulator
- **Informasi**: Setiap item history menampilkan:
  - Expression (perhitungan)
  - Result (hasil)
  - Timestamp (waktu perhitungan)
- **Limit**: Maksimal 10 item history (yang terlama otomatis terhapus)

## Teknologi yang Digunakan

- **Backend**: Go (Golang)
- **Frontend**: HTML5, CSS3, JavaScript
- **Server**: Go HTTP Server
- **Template Engine**: Go HTML Templates
- **Storage**: In-memory storage untuk history (akan hilang saat server restart)

## API Endpoints

- `GET /` - Halaman utama kalkulator
- `POST /calculate` - Menghitung expression dan menyimpan ke history
- `POST /clear` - Clear display dan reset ke 0
- `GET /history` - Endpoint JSON untuk mendapatkan data history

## Catatan Pengembangan

Aplikasi ini menggunakan evaluator expression yang sederhana. Untuk penggunaan production, disarankan menggunakan library parser expression yang lebih robust seperti:

- `github.com/Knetic/govaluate`
- `github.com/antonmedv/expr`

### Fitur History
- History disimpan dalam memory (akan hilang saat server restart)
- Untuk persistence, bisa ditambahkan database atau file storage
- Maksimal 10 item history untuk performa optimal
- History hanya tersimpan saat perhitungan berhasil (tombol =), tidak saat clear

## Lisensi

Proyek ini dibuat untuk tujuan pembelajaran dan dapat digunakan secara bebas.

## Kontribusi

Silakan berkontribusi dengan membuat pull request atau melaporkan bug melalui issues. 