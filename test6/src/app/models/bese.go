package models
// DBテーブルのベース
import (
	// "crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"test6/config"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	// tableNameTodo    = "todos"
	// tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	// cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	uuid STRING NOT NULL UNIQUE,
	// 	name STRING,
	// 	email STRING,
	// 	password STRING,
	// 	created_at DATETIME)`, tableNameUser)
	// ↓追加
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		birth_day STRING,
		age int,
		height float,
		weight float)`, tableNameUser)
		// created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// func Encrypt(plaintext string) (cryptext string) {
// 	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
// 	return cryptext
// }
// uuid=重複しないID
