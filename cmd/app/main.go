package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware" // 正しいパス
)

func main() {
	e := echo.New()

	// CORSミドルウェアを追加
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// ルートハンドラーを登録
	e.GET("/todos", func(c echo.Context) error {
		todos := []map[string]string{
			{"id": "1", "title": "Learn Go", "status": "pending"},
			{"id": "2", "title": "Build an API", "status": "completed"},
		}
		return c.JSON(200, todos)
	})

	// サーバーを開始
	e.Logger.Fatal(e.Start(":8080"))
}
