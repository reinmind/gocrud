# !/bin/bash
# minio
rm -rf $HOME/.minio
mkdir -p $HOME/.minio && chmod -R 777 $HOME/.minio
docker ps -a | grep minio | awk '{print $1}' | xargs docker rm -f
docker run -d --name minio \
    --publish 9000:9000 \
    --publish 9001:9001 \
    --volume $HOME/.minio/:/data \
    --env MINIO_ACCESS_KEY="minio" \
    --env MINIO_SECRET_KEY="minio123" \
    bitnami/minio:latest

