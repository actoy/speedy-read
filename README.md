# 长文总结，快速阅读

# golang环境安装
* wget https://dl.google.com/go/go1.21.5.linux-amd64.tar.gz
* rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
* 配置gopath: export GOPATH=~/go
* 设置bin: export PATH=$PATH:/usr/local/go/bin

# GO ENV 配置
* go env -w GOPROXY=https://goproxy.cn,direct
* go env -w GO111MODULE=on

# 安装服务插件
* go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
* go install github.com/cloudwego/thriftgo@latest

# git ssh 配置
* ssh-keygen -t rsa -C "ycsgldy@gmail.com"

# Mysql安装
## 下载mysql repo源
* wget https://dev.mysql.com/get/mysql80-community-release-el7-1.noarch.rpm
* rpm -ivh mysql80-community-release-el7-1.noarch.rpm 

## 安装mysql
* yum install mysql-server
* 服务启动: service mysqld restart
* 查看临时密码：grep "password" /var/log/mysqld.log
* 修改mysql密码: ALTER USER 'root'@'localhost' IDENTIFIED BY '@Free4me';

### 错误处理
如遇到mysql-community-client-plugins-8.0.35-1.el7.x86_64.rpm 的公钥尚未安装

则：rpm --checksig  /var/cache/yum/x86_64/7/mysql80-community/packages/mysql-community-client-plugins-8.0.35-1.el7.x86_64.rpm进行检查

* 解决方案
  * gpg --export -a 3a79bd29 > 3a79bd29.asc
  * rpm --import 3a79bd29.asc
  * rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2022

# 服务启动
* nohup sh build.sh 1>/dev/null 2>&1 &
* nohup sh build.sh > output.log 2>&1 &

# Thrift 生成server
* kitex -module speedy/read -service speedy-read thrift/speedy_read.thrift 