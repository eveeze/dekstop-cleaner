package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var dirMappings = map[string]string{
	// Gambar
	".jpg":  "Gambar",
	".jpeg": "Gambar",
	".png":  "Gambar",
	".gif":  "Gambar",
	".webp": "Gambar",
	".heic": "Gambar",
	".svg":  "Gambar",

	// Video
	".mp4": "Video",
	".mov": "Video",
	".mkv": "Video",
	".avi": "Video",

	// Dokument
	".pdf":  "Dokumen",
	".docx": "Dokumen",
	".doc":  "Dokumen",
	".xlsx": "Dokumen",
	".pptx": "Dokumen",
	".txt":  "Dokumen",

	// Arsip
	".zip": "Arsip",
	".rar": "Arsip",
	".7z":  "Arsip",

	// audio

	".mp3": "Audio",
	".wav": "Audio",
	".m4a": "Audio",

	// installer
	".exe": "Aplikasi",
	".msi": "Aplikasi",
}

func main() {
	// tentuakn folder target (downloads)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Gagal mendapatkan direktori home: %v", err)
	}

	// folder target (bisa kamu ubah ke folder di home direktori)
	targetPath := filepath.Join(homeDir, "Downloads")

	log.Printf("Memulai pembersihan di folder :%s", targetPath)

	time.Sleep(1 * time.Second)

	// baca semua file dan fodler di direktori target

	entries, err := os.ReadDir(targetPath)
	if err != nil {
		log.Fatalf("Gagal membaca direktori target: %v", err)
	}

	filesMoved := 0

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()
		sourcePath := filepath.Join(targetPath, fileName)

		// tentuakn fodler tujuan
		ext := strings.ToLower(filepath.Ext(fileName))
		destDirName, found := dirMappings[ext]

		if !found {
			destDirName = "Lainya"
		}

		// siapkan folder tujuan

		destDirPath := filepath.Join(targetPath, destDirName)

		if _, err := os.Stat(destDirPath); os.IsNotExist(err) {
			log.Printf("Membuat fodler baru: %s", destDirName)
			err := os.Mkdir(destDirPath, 0o755)
			if err != nil {
				log.Printf("Gagal membuat folder %s : %v", destDirPath, err)
				continue
			}
		}

		// pindahkan file

		finalDestPath := filepath.Join(destDirPath, fileName)
		log.Printf("Memindahkan %s ke folder %s ...\n", fileName, destDirName)
		err := os.Rename(sourcePath, finalDestPath)
		if err != nil {
			log.Printf("Gagal memindahkan %s: %v", fileName, err)
		} else {
			filesMoved++
		}

	}

	log.Println("--------------------------------------------------")
	if filesMoved > 0 {
		log.Printf("Selesai!!! %d file berhasil dirapikan", filesMoved)
	} else {
		log.Println("Fodler download sudah rapi, tidak ada file yang perlu dirapikan")
	}
}
