# request-limit-server
## 環境
- Go 版本 : 1.14.3
- Redis 版本 : 6.0.3

## 目的
設計一個每個IP每分鐘僅能接受60個requests，在首頁顯示目前的request量，超過限制的話則顯示“Error”，例如在一分鐘內第30個request則顯示30，第61個request 則顯示 Error的 Server

## 設計理念
使用 `redis` 紀錄 ip 存取的次數，透過 `ttl` 的機制作為判斷多少時間範圍內計算執行次數

## 參數配置
參數放置於 `config` 資料夾中，以下說明三個參數檔代表意義
### config
`config.yaml` 內設定關於 `server` 層級的相關參數，相關說明如下
```yaml
application:
  timeout:
    read: 1000    # 讀取 timeout 時間
    write: 1000   # 寫入 timeout 時間
  port: 8888      # server 啟動的 port
  mode: "release" # server 啟動的狀態（debug, test, release）

```
### backend
`backend.yaml` 內設定紀錄 ip 存取次數限制的 `redis`，相關說明如下
```yaml
redis:
  url: '127.0.0.1'    # redis 的 server 位置
  port: 6379          # redis 的 server port
  password: 'a12345'  # redis 存取密碼
  db: 0               # 存取 redis 的哪個 db
```
### limit
`limit.yaml` 內設定計算 ip 存取次數限制參數，相關說明如下
```yaml
limit:
  request: 60     # 限制存取次數
  range:          
    unit: minute  # 限制計算時間單位（second, minute, hour）
    per: 1        # 限制計算時間
```
## 啟動方法
### 直接啟動
#### Linux or MacOS
透過 `make` 指令直接啟動即可

#### Windows
初次啟動執行以下指令
```bash
go mod download
go run main.go
```
### Docker
#### Linux or MacOS
透過 `make docker` 指令直接啟動即可

#### Windows
初次啟動執行以下指令
```bash
docker build -t limit-request-server .
docker run -p 8888:8888 -name limit-request-server limit-request-server
```

### Docker Compose
使用 docker-compose 的方式啟動就可以不用預先安裝 `redis-server`
#### Linux or MacOS
透過 `make docker_compose` 指令直接啟動即可

#### Windows
初次啟動執行以下指令
```bash
docker build -t limit-request-server .
docker-compose up -d
```
## 測試流程
### 直接運行測試
1. 需要先設定好相關參數
2. 參照 `啟動方法` 中的方式啟動程式
3. 開啟瀏覽器，輸入 `http://<ip>:8888` 後按下 `Enter`
   - 第一次存取會看到顯示 `1`，第二次會看到 `2` 以此類推
   - 在一分鐘內存取超過 `60` 次則會顯示 `Error`

## 單元測試
### Linux or MacOS
使用 `make test_nocache` 執行
### Windows
執行以下指令
```bash
go test -v ./... -cover -count=1
```
