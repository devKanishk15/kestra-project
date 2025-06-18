# kestra-project
Project for Hack Week by WeMakeDevs using Kestra

# Docker build commands
-  docker buildx build -f docker_files/Dockerfile.go-debian -t kestra-go-test .
- docker buildx build -f docker_files/Dockerfile.go-alpine -t kestra-go-alpine-test .
- docker buildx build -f docker_files/Dockerfile.go-distroless -t kestra-go-distroless-test .
- docker buildx build -f docker_files/Dockerfile.go-multistage-alpine -t kestra-go-multistage-alpine-test .
- docker buildx build -f docker_files/Dockerfile.go-scratch -t kestra-go-scratch-test .