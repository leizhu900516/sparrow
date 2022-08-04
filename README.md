![Sparrow](./doc/sparrow.png "Sparrow")
# Sparrows

#### 极简介绍
企业版wiki。麻雀虽小，五脏俱全。

##### 为什么要做这个项目：
用了 [语雀](https://www.yuque.com/)的云端知识库之后，发现真的好用，但是由于公司背景，要求不能云办公，所以开发了一套可独立部署内网环境的
知识库平台，并起名叫《麻雀》，美其名曰：麻雀虽小，五脏俱全。目前实现语雀的50%功能吧。基本满足企业内部知识平台的需要，后期会不断的进行功能新增，欢迎小伙伴们提出优质的建议。
#### 功能
- 用户注册
- 知识库
- 团队管理 [邀请、管理]
- 团队知识库=协作知识库
- 文件知识库
- 图书馆
- 搜索
- 关注用户和知识库

##### 功能截图如下
![主页](./doc/img1.png)
----
![主页](./doc/img2.png)
----
![主页](./doc/img3.png)

#### 软件架构
软件架构说明 （B/S架构）
- 前端基于Element-UI
- 后端基于gin
- mysql5.7+
- 文件存储直接存储到服务端，只需配置自定义路径即可，后期为了文件高可用性，会加入纠删码

#### 创建数据库表
1. 上传sparrow.sql到/root下
2. 登录mysql服务器
```shell script
mysql -uroot -p
CREATE DATABASE `sparrow`;
use sparrow;
source /root/sparrow.sql;
```
3. 退出

#### 前端安装教程(nginx需要提前安装)
1. 将 frontend的dist目录复制到部署主机某目录下
2. 复制frontend/sparrow.conf到ngingx的conf目录下，并修改配置文件中的后端地址和该项目的地址
3. 重启nginx即可
#### 后端编译安装教程(需提前配置go环境)

1.  克隆仓库 
```shell script
git clone https://gitee.com/leizhu/sparrow.git
```
2.  进入sparrow目录，执行make
```shell script
cd sparrow
go mod vendor
make
```
3.  编译后的可执行文件在 output目录下,进入output执行下面命令即可。(配置文件可修改output目录下的config/config.ini;可修改数据库配置信息、后台服务端口)
```shell script
./bin/sparrow 
```
4. 后台启动
```shell
nohup ./bin/sparrow  &
```

#### 使用说明

1.  部署成功之后，登录http://ip:port ,默认账号和密码admin/123456 即可快速上手

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

