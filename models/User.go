package models

import "database/sql"

// User merepresentasikan struktur data untuk seorang pengguna.
type User struct {
	ID    int64
	Name  string
	Email string
	Age   int
}

// GetUser mengambil satu data user (baris pertama) dari tabel 'user'.
//
// Catatan: Fungsi ini menggunakan `QueryRow` dengan `SELECT *` tanpa klausa `WHERE`,
// sehingga hanya akan mengembalikan baris pertama yang ditemukan oleh database.
// Untuk mengambil pengguna tertentu, sebaiknya tambahkan parameter (misalnya ID) dan klausa `WHERE`.
func GetUser(db *sql.DB) ([]User, error) {
	// Menjalankan query untuk mengambil satu baris data dari tabel user.
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err // Mengembalikan struct User kosong dan error jika query gagal.
	}

	defer rows.Close() // Menutup hasil query setelah fungsi selesai dieksekusi.

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
