package models

import "database/sql"

// Hobi merepresentasikan struktur data untuk sebuah hobi.
type Hobi struct {
	ID          int64
	Name        string
	Description string
}

// GetAllHobi mengambil semua data hobi dari database.
// Fungsi ini menggunakan db.Query() untuk mendapatkan semua baris,
// lalu melakukan iterasi untuk membaca setiap baris ke dalam slice []Hobi.
func GetAllHobi(db *sql.DB) ([]Hobi, error) {
	// Menjalankan query untuk mengambil semua data dari tabel hobi.
	rows, err := db.Query("SELECT * FROM hobi") // Perbaikan query untuk mengambil semua data hobi.d
	if err != nil {
		return nil, err // Mengembalikan slice nil dan error jika query gagal.
	}
	// defer rows.Close() sangat penting untuk memastikan koneksi ke database
	// dilepaskan setelah fungsi selesai dieksekusi.
	defer rows.Close()

	var hobies []Hobi // Slice untuk menampung semua data hobi.

	// Looping untuk setiap baris yang ditemukan.
	for rows.Next() {
		var hobi Hobi // Variabel untuk menampung data per baris.
		// Membaca data dari baris saat ini ke dalam variabel hobi.
		if err := rows.Scan(&hobi.ID, &hobi.Name, &hobi.Description); err != nil {
			return nil, err // Jika ada error saat scanning, hentikan dan kembalikan error.
		}
		hobies = append(hobies, hobi) // Tambahkan hobi yang sudah dibaca ke dalam slice.
	}

	// Mengembalikan slice yang berisi semua data hobi dan nil untuk error.
	return hobies, nil
}

func GetHobiByID(db *sql.DB, id int64) (Hobi, error) {
	var hobi Hobi

	err := db.QueryRow("SELECT * FROM hobi WHERE id = ?", id).Scan(&hobi.ID, &hobi.Name, &hobi.Description)
	if err != nil {
		return Hobi{}, err
	}

	return hobi, nil
}
