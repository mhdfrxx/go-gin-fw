package service

import (
	"fmt"
	"go-gin-fw/database"
	"go-gin-fw/middleware"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	username := strings.TrimSpace(c.PostForm("username"))
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.HTML(http.StatusOK, "login.html", gin.H{"Error": "Semua field wajib diisi!"})
		return
	}

	var storedHash string
	err := database.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedHash)
	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{"Error": "Username tidak ditemukan"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)) != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{"Error": "Password Salah!"})
	}

	sessionID := fmt.Sprintf("%s_%d", username, time.Now().Unix())
	middleware.Sessions[sessionID] = username

	http.SetCookie(c.Writer, &http.Cookie{
		Name:  "session_id",
		Value: sessionID,
		Path:  "/",
	})

	c.Redirect(http.StatusSeeOther, "/")
}

func ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Register(c *gin.Context) {
	username := strings.TrimSpace(c.PostForm("username"))
	password := c.PostForm("password")
	confirm := c.PostForm("confirm")

	if username == "" || password == "" || confirm == "" {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Semua field harus diisi!"})
		return
	}

	if len(username) < 4 {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Username Minimal 4 Karakter!"})
		return
	}

	if len(password) < 6 {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Password Minimal 6 Karakter!"})
		return
	}

	validUser := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !validUser.MatchString(username) {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Username hanya boleh huruf, angka, atau underscore"})
		return
	}

	if password != confirm {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Konfirmasi password tidak sesuai!"})
		return
	}

	var exists int
	database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&exists)
	if exists > 0 {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Username sudah digunakan"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := database.DB.Exec("INSERT INTO users (username, password) VALUES(?,?)", username, hashedPassword)
	if err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{"Error": "Gagal menyimpan user"})
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{"Success": "Registrasi Berhasil! Silahkan Login."})
}

func Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_id")
	if err == nil {
		delete(middleware.Sessions, cookie.Value)
		http.SetCookie(c.Writer, &http.Cookie{
			Name:   "session_id",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
	}
	c.Redirect(http.StatusSeeOther, "/login")
}
