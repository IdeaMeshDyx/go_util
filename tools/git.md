

## [同步远程仓库最新修改] 
拉取远程分支并创建本地分支,之后 merge 到 main 或者 master 完成更新

### 查看远程分支
使用如下git命令查看所有远程分支：
``` shell
git branch -r
```

### 拉取远程分支并创建本地分支
方法一
使用如下命令：
```shell
git checkout -b 本地分支名x origin/远程分支名x
```
使用该方式会在本地新建分支x，并自动切换到该本地分支x, **采用此种方法建立的本地分支会和远程分支建立映射关系**

方式二
使用如下命令：
``` shell
git fetch origin 远程分支名x:本地分支名x
```

使用该方式会在本地新建分支x，但是不会自动切换到该本地分支x，需要手动checkout。**采用此种方法建立的本地分支不会和远程分支建立映射关系**

### 修改远程仓库
方法有三：

1.直接命令修改

git remote set-url origin [url]
例如：git remote set-url origin https://ddns.togit.cc:7777/rffanlab/tensorflow_get_started_doc_cn.git
复制
2.先删后添加

git remote rm origin
git remote add origin [url]
复制
3.修改config文件
 打开.git目录，修改config文件里面的url替换就OK了。