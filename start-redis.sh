# !/bin/bash
# redis
rm -rf $HOME/.redis
mkdir -p $HOME/.redis && chmod -R 777 $HOME/.redis
docker ps -a | grep redis | awk '{print $1}' | xargs docker rm -f
docker run -d --name redis \
    --publish 6379:6379 \
    -e ALLOW_EMPTY_PASSWORD=yes \
    --volume $HOME/.redis/:/bitnami/redis/data \
    bitnami/redis:latest