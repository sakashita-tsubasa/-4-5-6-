# Golang練習コンテナ
## 操作方法
### 開始
1. Dockerコンテナを起動する
```
cd [golang_practiceのディレクトリパス]
docker-compose up -d
```
2. VSCodeで「Attach to Running Container...」を選択し「/go-practice-container」を選択
3. コンテナに接続後、「フォルダを開く」→「/go/src」を選択
### コード修正
1. コンテナに接続しているVSCodeで編集 又は golang_practice/src 内のコードを編集
1. コンテナのターミナルで「go run main.go」を実行
### 終了
```
cd [golang_practiceのディレクトリパス]
docker-compose down
```