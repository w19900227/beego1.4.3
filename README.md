### 環境部署
    $ mkdir beego
    $ cd beego
    $ export GOPATH=$PWD
    
    $ go get github.com/astaxie/beego
    $ go get github.com/beego/bee
    
    $ cd $GOPATH/src
    $ export PATH=$PATH:$GOPATH/bin
    
### 下載檔案
    $ git clone https://github.com/w19900227/beego1.4.3.git
    $ cd beego1.4.3

### 執行
    $ go run main.go
或是

    $ go build .
    $ ./beego1.4.3

### 網頁瀏覽
http://localhost:8080