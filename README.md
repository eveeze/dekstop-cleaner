# Desktop Cleaner

![Desktop Cleaner Logo](https://verdex.id/assets/images/desktop-cleaner-banner.png)

Aplikasi sederhana untuk mengorganisir file-file di folder Downloads secara otomatis berdasarkan ekstensi file. Dibuat dengan bahasa pemrograman Go untuk membantu menjaga kerapian folder Downloads Anda.

## 🚀 Fitur

- **Organisasi Otomatis**: Memindahkan file ke folder yang sesuai berdasarkan jenis file
- **Kategori Lengkap**: Mendukung berbagai jenis file (gambar, video, dokumen, audio, arsip, aplikasi)
- **Folder Otomatis**: Membuat folder kategori secara otomatis jika belum ada
- **Log Aktivitas**: Menampilkan proses pemindahan file secara real-time
- **Aman**: Hanya memindahkan file, tidak menghapus atau merusak data

## 📁 Kategori File

| Kategori | Ekstensi yang Didukung |
|----------|------------------------|
| **Gambar** | `.jpg`, `.jpeg`, `.png`, `.gif`, `.webp`, `.heic`, `.svg` |
| **Video** | `.mp4`, `.mov`, `.mkv`, `.avi` |
| **Dokumen** | `.pdf`, `.docx`, `.doc`, `.xlsx`, `.pptx`, `.txt` |
| **Audio** | `.mp3`, `.wav`, `.m4a` |
| **Arsip** | `.zip`, `.rar`, `.7z` |
| **Aplikasi** | `.exe`, `.msi` |
| **Lainnya** | File dengan ekstensi yang tidak terdaftar |

## 🛠️ Instalasi

### Prasyarat

- Go 1.24.2 atau versi yang lebih baru
- Sistem operasi Windows, macOS, atau Linux

### Langkah Instalasi

1. **Clone repository**
   ```bash
   git clone https://github.com/username/dekstop-cleaner.git
   cd dekstop-cleaner
   ```

2. **Download dependencies**
   ```bash
   go mod tidy
   ```

3. **Build aplikasi**
   ```bash
   go build -o desktop-cleaner main.go
   ```

4. **Jalankan aplikasi**
   ```bash
   ./desktop-cleaner
   ```

## 📖 Cara Penggunaan

### Penggunaan Dasar

1. Jalankan aplikasi dengan perintah:
   ```bash
   go run main.go
   ```

2. Aplikasi akan secara otomatis:
   - Membaca semua file di folder Downloads Anda
   - Membuat folder kategori yang diperlukan
   - Memindahkan file ke folder yang sesuai
   - Menampilkan progress dan hasil akhir

### Contoh Output

```
2024/07/24 10:30:00 Memulai pembersihan di folder: /Users/username/Downloads
2024/07/24 10:30:01 Membuat folder baru: Gambar
2024/07/24 10:30:01 Memindahkan photo.jpg ke folder Gambar ...
2024/07/24 10:30:01 Memindahkan document.pdf ke folder Dokumen ...
2024/07/24 10:30:01 Membuat folder baru: Video
2024/07/24 10:30:01 Memindahkan movie.mp4 ke folder Video ...
--------------------------------------------------
2024/07/24 10:30:01 Selesai!!! 15 file berhasil dirapikan
```

## ⚙️ Kustomisasi

### Mengubah Direktori Target

Secara default, aplikasi bekerja pada folder `Downloads`. Untuk mengubah direktori target:

1. **Edit file `main.go`** pada baris 46:
   ```go
   // Ganti "Downloads" dengan folder yang diinginkan
   targetPath := filepath.Join(homeDir, "Desktop") // Contoh: Desktop
   ```

2. **Contoh direktori lain:**
   ```go
   // Untuk folder Desktop
   targetPath := filepath.Join(homeDir, "Desktop")
   
   // Untuk folder Documents
   targetPath := filepath.Join(homeDir, "Documents")
   
   // Untuk custom path
   targetPath := "/path/to/your/folder"
   ```

### Menambah Kategori File Baru

1. **Edit variabel `dirMappings`** di file `main.go`:
   ```go
   var dirMappings = map[string]string{
       // Kategori yang sudah ada...
       
       // Tambahkan kategori baru
       ".psd":  "Desain",
       ".ai":   "Desain",
       ".sketch": "Desain",
       
       // Kategori programming
       ".go":   "Code",
       ".js":   "Code",
       ".py":   "Code",
   }
   ```

### Mengubah Nama Folder Kategori

```go
var dirMappings = map[string]string{
    // Ubah nama folder dari "Gambar" menjadi "Images"
    ".jpg":  "Images",
    ".jpeg": "Images",
    ".png":  "Images",
    
    // Ubah nama folder dari "Video" menjadi "Movies"
    ".mp4": "Movies",
    ".mov": "Movies",
}
```

## 🔧 Development

### Struktur Project

```
dekstop-cleaner/
├── main.go          # File utama aplikasi
├── go.mod          # Module definition
└── README.md       # Dokumentasi
```

### Build untuk Platform Berbeda

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o desktop-cleaner.exe main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o desktop-cleaner-mac main.go

# Linux
GOOS=linux GOARCH=amd64 go build -o desktop-cleaner-linux main.go
```

## 🤝 Kontribusi

Kontribusi selalu diterima! Berikut cara berkontribusi:

1. Fork repository ini
2. Buat branch fitur baru (`git checkout -b feature/AmazingFeature`)
3. Commit perubahan (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Buat Pull Request

## 📝 Changelog

### v1.0.0
- ✨ Rilis awal dengan fitur dasar organisasi file
- 📁 Dukungan untuk 7 kategori file utama
- 🔄 Auto-create folder jika belum ada
- 📊 Logging aktivitas real-time

## ⚠️ Catatan Penting

- **Backup Data**: Selalu backup file penting sebelum menjalankan aplikasi
- **Test Dulu**: Uji coba di folder test sebelum menggunakan di data asli
- **Permissions**: Pastikan aplikasi memiliki akses write ke direktori target
- **File Duplikat**: Jika ada file dengan nama sama, proses mungkin akan gagal

## 🐛 Troubleshooting

### Error: "Gagal mendapatkan direktori home"
- **Solusi**: Pastikan environment variable HOME (Linux/Mac) atau USERPROFILE (Windows) ter-set dengan benar

### Error: "Gagal membaca direktori target"
- **Solusi**: Periksa apakah folder target ada dan aplikasi memiliki permission untuk membacanya

### File tidak terpindah
- **Solusi**: Periksa apakah file sedang digunakan oleh aplikasi lain atau ada konflik nama file

## 📞 Support

Jika Anda mengalami masalah atau memiliki pertanyaan:

- 🐛 [Buat Issue](https://github.com/username/dekstop-cleaner/issues)
- 💬 [Diskusi](https://github.com/username/dekstop-cleaner/discussions)
- 📧 Email: support@verdex.id

## 📜 Lisensi

Distributed under the MIT License. See `LICENSE` for more information.

## 🙏 Acknowledgments

- Terima kasih kepada komunitas Go Indonesia
- Inspirasi dari berbagai tool file organizer
- [Verdex.id](https://verdex.id) untuk hosting assets

---

<div align="center">

**Dibuat dengan ❤️ menggunakan Go**

![Go Version](https://img.shields.io/badge/Go-1.24.2-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey.svg)

[⭐ Star this repo](https://github.com/username/dekstop-cleaner) | [🐛 Report Bug](https://github.com/username/dekstop-cleaner/issues) | [💡 Request Feature](https://github.com/username/dekstop-cleaner/issues)

</div>
