## preparation 
brew install cockroach
brew services start cockroach

内网console：
http://localhost:26256/#/overview/list

启动db：
cockroach sql --insecure
    create database nakama
    
初始化db数据：
nakama migrate up    

编译：nakama
1. 修改go.mod中nakama-common为master：go get -u github.com/heroiclabs/nakama-common@master
2. cp nakama /usr/local/bin/nakama

## build
./build.sh

## run
 [build nakema](https://github.com/heroiclabs/nakama) on master
 
nakama --database.address "root@127.0.0.1:26257/nakama" --runtime.path "path/nakama_golang/dist"
