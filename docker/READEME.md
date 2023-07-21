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