package service

import (
	"go-gin-fw/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
}

func ListUsers(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, username From users")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error membaca database")
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Username)
		users = append(users, u)
	}

	c.HTML(http.StatusOK, "users.html", gin.H{"Users": users})
}

func EditUserForm(c *gin.Context) {
	id := c.Param("id")
	var u User
	err := database.DB.QueryRow("SELECT id, username FROM users WHERE id = ?", id).Scan(&u.ID, &u.Username)
	if err != nil {
		c.String(http.StatusNotFound, "User tidak ditemukan")
		return
	}

	c.HTML(http.StatusOK, "edit.html", gin.H{"User": u})
}

func EditUser(c *gin.Context) {
	id := c.Param("id")
	username := strings.TrimSpace(c.PostForm("username"))
	password := strings.TrimSpace(c.PostForm("password"))

	if username == "" {
		c.HTML(http.StatusOK, "edit.html", gin.H{"Error": "Username tidak boleh kosong"})
		return
	}

	if password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		_, err := database.DB.Exec("UPDATE users SET username=?,password=? WHERE id=?", username, hashed, id)
		if err != nil {
			c.String(http.StatusInternalServerError, "Gagal Update User")
			return
		}
	} else {
		_, err := database.DB.Exec("UPDATE users SET username=?, password=?", username, id)
		if err != nil {
			c.String(http.StatusInternalServerError, "Gagal Update user")
			return
		}
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func DeleteUserConfirm(c *gin.Context) {
	id := c.Param("id")
	var u User
	err := database.DB.QueryRow("SELECT id, username FROM users WHERE id = ?", id).Scan(&u.ID, &u.Username)
	if err != nil {
		c.String(http.StatusNotFound, "User tidak ditemukan")
		return
	}
	c.HTML(http.StatusOK, "confirm_delete.html", gin.H{"User": u})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Gagal menghapus User")
		return
	}
	c.Redirect(http.StatusSeeOther, "/users")
}
