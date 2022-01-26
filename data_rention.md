# [Centos8 安装GO环境](https://www.cnblogs.com/fke2018/p/13739175.html)

**下载安装包**

wget https://golang.google.cn/dl/go1.15.15.linux-amd64.tar.gz

可以手动下载然后上传，也可以直接执行上面的命令

附：下载地址：

- Go官网下载地址：<https://studygolang.com/dl>

- Go官方镜像站（推荐）: <https://golang.google.cn/dl/>

  ​

**1. 解压到/usr/local目录**

执行以下命令：

tar -C /usr/local -xzf go1.15.15.linux-amd64.tar.gz

**2. home目录下, 建立一个名为gopath的目录，gopath下再建立3个目录pkg、bin、src**

执行以下命令：

mkdir -p /home/gopath/src /home/gopath/pkg /home/gopath/bin

**3. 编辑环境变量 **

执行以下命令：

vi /etc/profile

按Insert键切换到编辑模式，在最后添加以下内容

```
export GOROOT=/usr/local/go
export GOPATH=/home/gopath
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

**4.按Esc退出编辑模式，输入:wq保存并退出**

执行以下命令，使环境变量立即生效:

source /etc/profile

**5.校验环境是否正常**

执行以下命令

go env



运行GO项目

[root@docker bin]# export name=张三
[root@docker bin]# $GOPATH/bin/go_project_docker