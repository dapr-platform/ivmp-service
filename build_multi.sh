./swagger.sh
docker buildx build --platform linux/amd64,linux/arm64 -t ivmp-service  .

