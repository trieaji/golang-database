package golangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql" //menambah database driver itu wajib kalau ingin koneksi ke database
)

func TestEmpty(t *testing.T) {

}

//cara membuat/membuka connection ke db
func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// gunakan DB

}