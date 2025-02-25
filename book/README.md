# 场景
程序员小明需要借阅一本《西游记》，在没有线上图书管理系统的时候，他每天都要去图书馆前台咨询图书馆管理员，
* 小明：你好，请问今天《西游记》的图书还有吗？
* 管理员：没有了，明天再来看看吧。

过了一天，小明又来到图书馆，问：
* 小明：你好，请问今天《西游记》的图书还有吗？
* 管理员：没有了，你过两天再来看看吧。

就这样经过多次反复，小明也是徒劳无功，浪费大量时间在来回的路上，于是终于忍受不了落后的图书管理系统，
他决定自己亲手做一个图书查阅系统。

## 预期实现目标
* 用户登录
  依靠现有学生系统数据进行登录
* 图书检索
  根据图书关键字搜索图书，查询图书借阅情况，归还时间等。

## 系统分析
### 服务拆分
* user
    * api 提供用户登录协议
    * rpc 供search服务访问用户数据
* search
    * api 提供图书查询协议

> [!TIP]
> 这个微小的图书借阅查询系统虽然小，从实际来讲不太符合业务场景，但是仅上面两个功能，已经满足我们对go-zeroapi/rpc的场景演示了，
> 后续为了满足更丰富的go-zero功能演示，会在文档中进行业务插入即相关功能描述。这里仅用一个场景进行引入。


# 启动并验证服务
# 启动etcd、redis、mysql

# 启动user rpc
$ cd service/user/rpc
$ go run user.go -f etc/user.yaml
Starting rpc server at 127.0.0.1:8080...

# 启动user api
$ cd service/user/cmd/api
$ go run user.go -f etc/user-api.yaml

# 用户注册
curl -i -X POST \
  http://127.0.0.1:8888/user/register \
  -H 'Content-Type: application/json' \
  -d '{
    "number": "123456789",
    "name": "Tom"
    "password": "Tom123456789",
    "gender": "男"
}'
# 用户登陆
curl -i -X POST \
  http://127.0.0.1:8888/user/login \
  -H 'Content-Type: application/json' \
  -d '{
    "username":"123456789",
    "password":"Tom123456789"
}'

# 启动search api
$ cd service/search/api
$ go run search.go -f etc/search-api.yaml

# 验证服务
$ curl -i -X GET \
  'http://127.0.0.1:8889/search/do?name=%E8%A5%BF%E6%B8%B8%E8%AE%B0' \
  -H 'authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTI4NjcwNzQsImlhdCI6MTYxMjc4MDY3NCwidXNlcklkIjoxfQ.JKa83g9BlEW84IiCXFGwP2aSd0xF3tMnxrOzVebbt80'
HTTP/1.1 200 OK
Content
-Type: application/json
Date: Tue, 09 Feb 2021 06:05:52 GMT
Content-Length: 32

{"name":"西游记","count":100}
