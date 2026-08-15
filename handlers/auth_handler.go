package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/yourname/backend/config"
	"github.com/yourname/backend/database"
	"github.com/yourname/backend/models/request"
	"github.com/yourname/backend/models/tables"
	"github.com/yourname/backend/utils"

	"github.com/gin-gonic/gin"
)

func Register(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input request.RegisterRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "สมัครสมาชิกไม่สำเร็จ", "error": err.Error()})
			return
		}

		hashedPassword, err := utils.HashPassword(input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "สมัครสมาชิกไม่สำเร็จ", "error": "เข้ารหัสรหัสผ่านไม่สำเร็จ"})
			return
		}

		_, err = database.DB.Exec(
			"INSERT INTO staff (username, password, hospital) VALUES ($1, $2, $3)",
			input.Username,
			hashedPassword,
			input.Hospital,
		)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"message": "สมัครสมาชิกไม่สำเร็จ", "error": "username นี้มีอยู่แล้ว"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "สมัครสมาชิกสำเร็จ",
			"user": gin.H{
				"username": input.Username,
				"hospital": input.Hospital,
			},
		})
	}
}

func Login(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input request.LoginRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "เข้าสู่ระบบไม่สำเร็จ", "error": err.Error()})
			return
		}
		fmt.Println("request : ", input)
		var user tables.Staff
		err := database.DB.QueryRow(
			"SELECT username, password, hospital FROM staff WHERE username = $1 AND hospital = $2 AND status = 1",
			input.Username,
			input.Hospital,
		).Scan(&user.Username, &user.Password, &user.Hospital)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "เข้าสู่ระบบไม่สำเร็จ", "error": "username หรือ password ไม่ถูกต้อง"})
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message": "เข้าสู่ระบบไม่สำเร็จ", "error": "username หรือ password ไม่ถูกต้อง"})
			return
		}

		if !utils.CheckPasswordHash(input.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "เข้าสู่ระบบไม่สำเร็จ", "error": "username หรือ password ไม่ถูกต้อง"})
			return
		}

		token, err := utils.GenerateToken(user.Username, cfg.JWTSecret, cfg.JWTExpireHours)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "เข้าสู่ระบบไม่สำเร็จ", "error": "สร้าง token ไม่สำเร็จ"})
			return
		}

		_, err = database.DB.Exec(
			"UPDATE staff SET last_login = NOW() WHERE username = $1",
			input.Username,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "เข้าสู่ระบบไม่สำเร็จ", "error": "อัปเดตเวลาการเข้าสู่ระบบล่าสุดไม่สำเร็จ"})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "login สำเร็จ",
			"token":   token,
			"user": gin.H{
				"username": user.Username,
				"hospital": user.Hospital,
			},
		})
	}
}
