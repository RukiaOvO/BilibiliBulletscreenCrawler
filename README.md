## BilibiliBulletscreenCrawler

### Discription:
A small Go-based tool that can crawl a bilibiliVideo's bulletscreen   
and parse the uid of the sender you want to know.

### Project dictionary:
```
.
├── cmd
|    └── main.go  //入口
├── config
|    ├── local
|    |     └── hashMap.yaml //存放爆破过的crc32的k-v
|    └──conf.go
├── database
|       ├── data //存放爬取的弹幕数据，以bvid为文件名
|       └── model //存放xml和json数据解码的结构体
├── pkg
|    └── filecontrol
|             └── filecontrol.go //数据文件处理
└── utils
      ├── crawler.go //数据爬取和预处理
      └── decoder.go //crc32爆破
```

### Usage Instructions
- 克隆到本地
```
git clone https://github.com/RukiaOvO/BilibiliBulletscreenClawer.git
```
- 进入到根目录中配置环境
```
go mod init BilibiliBulletscreenClawer
go mod tidy
```
- 修改main函数中bvid和keyword参数并运行即可
###### Something need to know
- 由于部分b站账号是最近创建的，uid大都已经10-16位数，反向爆破crc32是不可能的，因此只有当结果是9位uid以内才显示，其余皆为unknown
- 由于爆破算量太大，即使并发也耗时过久，因此只爆破符合给定的keyword的弹幕，其余弹幕的uid皆为null
- crc32爆破所得结果不一定准确，uid可能有误差，但8-9位uid误差应该不大，勉强可以认为是准确的
- 由于接口提供的数据量限制，单个视频的弹幕数据量有限，似乎有600和1200两个阈值，因此爬取结果不全
- 由于爆破算法和算力限制，解析uid的时间较长
- 若不想用关键词匹配弹幕，在keyword参数中填入"RukiaOvO"即可，此时不会爆破uid,只会爬取弹幕内容和加密过的hash码