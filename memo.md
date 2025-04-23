MCP Inspectorの準備

npx @modelcontextprotocol/inspector 

サーバー構成ファイルをカスタマイズする場合

{
  "mcpServers": {
    "my-server": {
      "command": "node",
      "args": ["build/index.js", "arg1", "arg2"],
      "env": {
        "key": "value",
        "key2": "value2"
      }
    }
  }
}