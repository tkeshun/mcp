package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type QueryConfig struct {
	Name        string `json:"name"`         // クエリ名
	Description string `json:"description"`  // MCP説明
	Dir         string `json:"dir"`          // 環境変数ROOT_ENVからの相対パス
	PathPattern string `json:"path_pattern"` // パスパターン（glob）
}

func loadConfig(filename string) ([]QueryConfig, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cfg []QueryConfig
	err = json.Unmarshal(file, &cfg)
	return cfg, err
}

func concatFilesWithGlob(root, pattern string) (string, error) {
	var builder strings.Builder
	matches, err := doublestar.Glob(os.DirFS(root), pattern)
	if err != nil {
		return "", err
	}

	for _, match := range matches {
		fullPath := filepath.Join(root, match)
		data, err := os.ReadFile(fullPath)
		if err != nil {
			continue
		}
		builder.WriteString(fmt.Sprintf("==== %s ====\n", match))
		builder.Write(data)
		builder.WriteString("\n\n")
	}

	return builder.String(), nil
}

func main() {
	// 環境変数取得
	rootDir := os.Getenv("ROOT_DIR")
	if rootDir == "" {
		panic("環境変数 ROOT_DIR が未設定です")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		slog.Error("CONFIG_PATHが未設定")
	}

	// MCPサーバー初期化
	s := server.NewMCPServer("Dynamic MCP Server", "1.0.0")

	// 設定ファイル読み込み
	configs, err := loadConfig(configPath)
	if err != nil {
		slog.Error(fmt.Errorf("設定ファイル読み込み失敗: %w", err).Error())
	}

	// 各ツールを登録
	for _, cfg := range configs {
		tool := mcp.NewTool(cfg.Name,
			mcp.WithDescription(cfg.Description),
		)

		localCfg := cfg // クロージャで固定

		// rootDir + cfg.Dir に解決（再帰探索のベース）
		fullSearchRoot := filepath.Join(rootDir, localCfg.Dir)

		s.AddTool(tool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := concatFilesWithGlob(fullSearchRoot, localCfg.PathPattern)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("ファイル探索エラー: %v", err)), nil
			}
			return mcp.NewToolResultText(result), nil
		})
	}

	// 起動
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("サーバーエラー: %v\n", err)
	}
}
