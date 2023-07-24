# Docker Basic Operation

挂载文件卷
```shell 
docker run -d -P \
    --name container_name \
    # -v /src/webapp:/usr/share/nginx/html \
    --mount type=bind,source=path_local,target=path_in_container \
    image_name
```


打包镜像
```shell
docker stop container_id
docker commit -a "name" -m "info" container_id  image:tags
```



## 常用路径
docker run -d -P \
    --name go_dev \
     --mount type=bind,source=/Users/yixinda/file/util/go_util ,target=/home/go_util \
    --mount type=bind,source=/Users/yixinda/file/school/overlaySR ,target=/home/overlaySR \ 
    --mount type=bind,source=/Users/yixinda/file/cni/knetwork ,target=/home/knetwork \
    --mount type=bind,source=/Users/yixinda/file/cni/knetwork ,target=/home/knetwork
    bendyx/dev_go