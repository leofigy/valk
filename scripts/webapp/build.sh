docker build -t go-redis-kubernetes .
docker tag go-redis-kubernetes jovafig/go-redis-app:1.0.0
docker push jovafig/go-redis-app:1.0.0
