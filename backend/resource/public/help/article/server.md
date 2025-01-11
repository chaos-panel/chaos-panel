# 服务端


## 部署

#### 使用 docker

```bash
WORKDIR=/opt/aliyun/
CONTAINER_NAME=api
chmod -R 777 ./*
docker rm -f ${CONTAINER_NAME} 2>/dev/null
docker run -dit \
--privileged \
--restart=always \
-p 8000:8000 \
--workdir=${WORKDIR} \
-v ${WORKDIR}:${WORKDIR}:ro \
--entrypoint ${WORKDIR}/license-server \
--name=${CONTAINER_NAME} \
loads/alpine:3.8


WORKDIR=/opt/aliyun/
CONTAINER_NAME=ui
chmod -R 777 ./*
docker rm -f ${CONTAINER_NAME} 2>/dev/null
docker run -dit \
--privileged \
--restart=always \
-p 9000:80 \
--workdir=${WORKDIR} \
-v ${WORKDIR}/spa:/usr/share/nginx/html:ro \
--name=${CONTAINER_NAME} \
nginx

```

#### 使用 docker-compose

```bash
// TODO
```