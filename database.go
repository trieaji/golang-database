package golangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database?parseTime=true")//untuk mengetahui bahwasanya ada tipe data date,birth, atau time, maka gunakan "parseTime=true"
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10) //pengaturan berapa jumlah koneksi minimal yang dibuat
	db.SetMaxOpenConns(100) //pengaturan berapa jumlah koneksi maximal yang dibuat
	db.SetConnMaxIdleTime(5 * time.Minute) //pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute)//pengaturan berapa lama koneksi boleh digunakan

	return db
}