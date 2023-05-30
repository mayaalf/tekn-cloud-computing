package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

type Mhs struct {
    nim        int
    nama       string
    jurusan    string
}

func main() {

		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    res, err := db.Query("SELECT * FROM mahasiswa")

    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }

    for res.Next() {

        var mhs Mhs
        err := res.Scan(&mhs.nim, &mhs.nama, &mhs.jurusan)

        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%v\n", mhs)
    }
}