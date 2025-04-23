package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer("時刻答える君", "0.0.1")

	currentTimeTool := mcp.NewTool("current_time",
		mcp.WithDescription("現在時刻を返します"),
	)

	s.AddTool(currentTimeTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		now := time.Now().In(jst)
		message := fmt.Sprintf("現在の時刻: %s", now.Format("2006-01-02 15:04:05"))
		return mcp.NewToolResultText(message), nil
	})

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("サーバーエラー: %v\n", err)
	}
}
