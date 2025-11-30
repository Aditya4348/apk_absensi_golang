package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// DB adalah variabel global untuk koneksi database yang akan digunakan di seluruh aplikasi.
var DB *sql.DB

// ConnectDB menginisialisasi koneksi ke database MySQL
func ConnectDB() error {
	// Konfigurasi untuk koneksi ke database MySQL.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "db_absensi_go",
		AllowNativePasswords: true, // Diperlukan untuk beberapa versi MySQL.
	}

	var err error
	// Membuka koneksi ke database menggunakan DSN (Data Source Name) dari konfigurasi.
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return fmt.Errorf("gagal membuka koneksi database: %v", err)
	}

	// Melakukan ping ke database untuk memastikan koneksi benar-benar berhasil.
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("gagal ping database: %v", err)
	}

	// Jika koneksi berhasil, cetak pesan sukses
	fmt.Println("Database terkoneksi!")
	return nil // Mengembalikan nil jika tidak ada error.
}
