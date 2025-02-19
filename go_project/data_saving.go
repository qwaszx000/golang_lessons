package main

import (
	"database/sql"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-sql-driver/mysql"
)

var memcached_connection *memcache.Client
var mysql_connection *sql.DB
var err error

func init() {
	memcached_connection = memcache.New("memcached:11211") //memcache.New("127.0.0.1:11211")

	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "mariadb:3306", //"127.0.0.1:3307",
		DBName:               "test_db",
		AllowNativePasswords: true,
	}

	fmt.Println(cfg.FormatDSN())
	mysql_connection, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	_, err = mysql_connection.Exec(`CREATE TABLE IF NOT EXISTS test_table (
		id int primary key auto_increment,
		name varchar(64),
		value varchar(64)
	)`)
	if err != nil {
		fmt.Printf("SQL error: %s\n", err)
	}
}

// DB
// test_table: id, name, value
func save_to_db(name string, value string) bool {

	result, err := mysql_connection.Exec("INSERT INTO test_table (name, value) VALUES (?, ?)", name, value)
	if err != nil {
		fmt.Printf("SQL error: %s\n", err)
		return false
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("SQL error: %s\n", err)
		return false
	}
	fmt.Printf("SQL inserted id: %d", id)

	return true
}

func get_from_db(name string) string {
	var return_val string

	result, err := mysql_connection.Query("SELECT value FROM test_table WHERE name=? LIMIT 1", name)
	if err != nil {
		fmt.Printf("SQL error: %s\n", err)
		return return_val
	}

	result.Next()
	result.Scan(&return_val)

	return return_val
}

// Cache
func save_to_cache(key string, value []byte) bool {
	err := memcached_connection.Set(
		&memcache.Item{Key: key, Value: value},
	)

	if err != nil {
		fmt.Printf("Memcached error: %s\n", err)
		return false
	}

	return true
}

func get_from_cache(key string) []byte {
	item, err := memcached_connection.Get(key)

	if err != nil {
		return nil
	}

	return item.Value
}
