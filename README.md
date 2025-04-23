# mcp

MCPサーバー実装の練習用リポジトリです。
プロジェクトルートで`modelcontextprotocol/inspector`をつかうと色々試せます。

## mcp-time-server

現在時刻を返すMCPサーバーです。
`npx @modelcontextprotocol/inspector --config ./config.json --server current-time`
でテストできます。

## mcp-concat-file-finder

指定したディレクトリのファイルを結合して返してくれるMCPサーバーです。
`npx @modelcontextprotocol/inspector --config ./config.json --server file-finder`
でテストできます。

inspectorの設定ファイルで環境変数を指定する必要があります。
- プロジェクトルートを指定
ex) "ROOT_DIR": "./test",
- エンドポイント名、LLMへの説明、走査対象ディレクトリ、マッチさせるファイルパターンを指定するファイルのパス
ex) "CONFIG_PATH": "./finder-config.json"

エンドポイント名、LLMへの説明、走査対象ディレクトリ、マッチさせるファイルパターンを指定するファイルにパラメータを指定すると、任意のtoolsエンドポイントが作れます。

- name
エンドポイント名です。  
- description
LLMに対するエンドポイントの説明です。使うかどうかの判断がしやすいように書いてください  
- dir
走査するディレクトリです。ルートからの相対パスを指定してください。  
- path_pattern
"github.com/bmatcuk/doublestar/v4"をパターンマッチに使ってます。適当にパターン検索して使ってください。  

```
[
    {
      "name": "docs_query",
      "description": ".vsc"github.com/bmatcuk/doublestar/v4"ode以下の情報を取得します。vscodeなどの設定がほしい場合に使用してください",
      "dir": "./github-mcp-server/.vscode",
      "path_pattern": "**"
    },
    {
      "name": "e2e_code_query",
      "description": "ソースコードファイルを読み込みます。E2Eテストのコードです。",
      "dir": "./github-mcp-server/e2e",
      "path_pattern": "**/*.go"
    }
]
```