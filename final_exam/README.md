### 接口

1.“/user/register”:实现注册

传入格式:
```html
data: {"name": name, "password": password,"token":token}
```
返回格式:
```go
{     "Code":  "002",      "Message": string,   }//（0开头为错误，1开头为正确）
```
2."/user/login":实现登录 （返回格式和注册格式与注册类似）

3."/user/creatroom/:id":实现创建房间，加入房间和随机创建房间

传入格式:
```html
url: u+"user/creatroom/0"
```
返回格式:

随机创建房间时：

```go
{  "Code":    "000",   "Message": string,   "Room":    k,} //k为房间号 
```

加入和创建房间都与注册与登录返回格式类似

4."/ws/:id":每个房间对应一个id (这是一个webstock)

### 补充

上云后部署的前端页面无法储存cookie，导致后续程序无法正常运行，而用本地页面访问却能储存

登录注册页面[101.201.140.26]()

postman: [https://www.getpostman.com/collections/e60395736e01efc441dc]()