
## 通过 ssh 免密登陆 github 

### 创建私钥公钥「本地机器」

```shell
ssh-keygen -t ed25519 -C "your_email@example.com"
```

然后就顺着向下走，完成创建过程

### 粘贴 .pub 公钥到 github ssh
完成登陆设置