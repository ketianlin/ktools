package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
)

// 下面3个可以在自己项目中的配置文件中配置
const (
	jwtKey             = "diao_mao"
	tokenExpire        = 60
	refreshTokenExpire = 120
)

// 生成token
func TestGenerateToken(t *testing.T) {
	var id int64 = 666
	token, err := ktools.Jwt.GenerateToken(id, "吊毛", tokenExpire, jwtKey)
	if err != nil {
		fmt.Println("GenerateToken error,", err.Error())
		return
	}
	// 生成刷新token,刷新的token有效期是原token的一倍
	refreshToken, err := ktools.Jwt.GenerateToken(id, "吊毛", refreshTokenExpire, jwtKey)
	if err != nil {
		fmt.Println("GenerateRefreshToken error,", err.Error())
		return
	}
	fmt.Println("token: ", token)
	fmt.Println("refreshToken: ", refreshToken)
}

// 解析token
func TestAnalyzeToken(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NjY2LCJVc2VybmFtZSI6IuWQiuavmyIsImV4cCI6MTcwMTA2NzQyMX0.JLKD6egiHJ0fCMbWW4PrfU6ZWEXYcgY9Bzq8DtgCGZs"
	uc, err := ktools.Jwt.AnalyzeToken(tokenStr, jwtKey)
	if err != nil {
		fmt.Println("TestAnalyzeToken error,", err.Error())
		return
	}
	fmt.Printf("%T\t%v\n", uc, uc)
}

// 刷新token
func TestRefreshAuthorization(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NjY2LCJVc2VybmFtZSI6IuWQiuavmyIsImV4cCI6MTcwMTA2NzQ4MX0.MZ8A_j8hAJrK17Nk4S4zsKfLD7eG4Y-1wvCGNAIyHvM"
	uc, err := ktools.Jwt.AnalyzeToken(tokenStr, jwtKey)
	if err != nil {
		fmt.Println("RefreshAuthorization error,", err.Error())
		return
	}
	token, err := ktools.Jwt.GenerateToken(uc.Id, uc.Username, tokenExpire, jwtKey)
	if err != nil {
		fmt.Println("GenerateToken error,", err.Error())
		return
	}
	// 生成刷新token,刷新的token有效期是原token的一倍
	refreshToken, err := ktools.Jwt.GenerateToken(uc.Id, uc.Username, refreshTokenExpire, jwtKey)
	if err != nil {
		fmt.Println("GenerateRefreshToken error,", err.Error())
		return
	}
	fmt.Println("token: ", token)
	fmt.Println("refreshToken: ", refreshToken)
}
