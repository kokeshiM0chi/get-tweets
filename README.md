# Go でツイートをいろいろな方法で取得

## 検索できること

1. キーワード検索(キーワードで取得したツイートのリプライも取得)
2. 特定のツイートに対して行われたリプライ
3. 特定のツイートに対して行われた引用リツイート(引用リツイートに引用リツイートしているものも取得)

<br>

## 詳解

```
.
├── README.md
├── allReply
│   └── main.go
├── dfs.go
├── keys.go
├── lib.go
├── quoteTweets
│   └── main.go
├── search
│   └── main.go
└── search.go
```

### キーワード検索

対象ディレクトリ: /search  
Twitter の Search API を使って検索する。  
lib.AllSearch の第 3 引数を True にすると、取得したツイートにリプライされているものも検索する。  
しかし、True にすると Search API を叩きまくるのでかなりの時間がかかるので注意  
やる気があればリプライも取得するか確認するための CLI 化をしたい  
<br>

### リプライ

対象ディレクトリ: /allReply  
特定ツイートに行われたリプライをすべて取得。  
こちらは True にせずとも、取得したツイートにリプライされているものも検索する。  
Search API を叩きまくるのでかなりの時間がかかるので注意  
やる気があればリプライも取得するか確認するための CLI 化をしたい  
<br>

### 引用リツイート

対象ディレクトリ: /quoteTweets
指定ツイートに行われた引用リツイートのみを取得。また、引用リツイートに引用リツイートされたものも取得する  
Search API を叩きまくるのでかなりの時間がかかるので注意  
やる気があれば引用リツイートに引用リツイートされたものも取得するか確認するための CLI 化をしたい  
<br><br>

## 前準備

1. ツイッターのデータと、twitter の json を格納するディレクトリ名を決めます。  
   keys.go を作成します。

```go
package lib

const consumerKey string = ""
const consumerSecret string = ""
const accessToken string = ""
const accessTokenSecret string = ""

// path = $HOME+path形式。ホームディレクトリ配下にPATHが作成される
const path = "/twitter-data"
```

<br><br>

### キーワード検索のバイナリ生成

```
cd search
go build
./search <検索したいワード>
```

で、ユーザーのホームディレクトリに twitter-data というディレクトリが作成され、json ファイルが配置されます
