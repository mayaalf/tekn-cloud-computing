package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Mahasiswa struct {
		nim     int
		nama    string
		jurusan string
	}
	router := gin.Default()

	router.GET("/mahasiswa/:id", func(c *gin.Context) {
		var (
			mahasiswa Mahasiswa
			result    gin.H
		)
		id := c.Param("nim")
		row := db.QueryRow("select * from mahasiswa where id = ?;", id)
		err = row.Scan(&mahasiswa.nim, &mahasiswa.nama, &mahasiswa.jurusan)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": mahasiswa,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// GET all mahasiswas
	router.GET("/mahasiswas", func(c *gin.Context) {
		var (
			mahasiswa  Mahasiswa
			mahasiswas []Mahasiswa
		)
		rows, err := db.Query("select * from mahasiswa;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&mahasiswa.nim, &mahasiswa.nama, &mahasiswa.jurusan)
			mahasiswas = append(mahasiswas, mahasiswa)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": mahasiswas,
			"count":  len(mahasiswas),
		})
	})

	// POST new mahasiswa details
	router.POST("/mahasiswa", func(c *gin.Context) {
		var buffer bytes.Buffer
		nama := c.PostForm("nama")
		jurusan := c.PostForm("jurusan")
		stmt, err := db.Prepare("insert into mahasiswa (nama, jurusan) values(?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nama, jurusan)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(nama)
		buffer.WriteString(" ")
		buffer.WriteString(jurusan)
		defer stmt.Close()
		identitas := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", identitas),
		})
	})

	// PUT - update a mahasiswa details
	router.PUT("/mahasiswa", func(c *gin.Context) {
		var buffer bytes.Buffer
		nim := c.Query("nim")
		nama := c.PostForm("nama")
		jurusan := c.PostForm("jurusan")
		stmt, err := db.Prepare("update mahasiswa set nama= ?, jurusan= ? where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(nama, jurusan, nim)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(nama)
		buffer.WriteString(" ")
		buffer.WriteString(jurusan)
		defer stmt.Close()
		identitas := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", identitas),
		})
	})
}