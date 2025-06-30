# SASTOJ

## What is SASTOJ?

SASTOJ is a microservice online judge system based on Kratos, aiming to provide a more convenient and efficient online judge system for algorithm competitions.

## Development

### Environment

```plaintext
go version: 1.22
kratos version: v2.7
```

### Download Dependency

```
go mod download
```

### Run a Service

```
kratos run
```

More details about Go-Kratos: https://go-kratos.dev/docs/

### Build and Deploy

We provide a simplified deployment script that can build and deploy specific services:

```bash
# Deploy all services
./deploy.sh

# Deploy specific services
./deploy.sh admin    # Deploy admin service
./deploy.sh contest  # Deploy contest service
./deploy.sh auth     # Deploy auth service
./deploy.sh gojudge  # Deploy Go judge service
./deploy.sh freshcup # Deploy Freshcup judge service

# Show help information
./deploy.sh --help
```

Each built image includes the Git commit hash from which it was built. This hash is stored in the `/git_commit.txt` file inside the container and is also used as the image tag. You can check it using:

```bash
# Check the Git commit hash of a running container
docker exec -it <container_id> cat /git_commit.txt

# Or when using docker-compose
docker-compose exec admin cat /git_commit.txt
```

### Use Docker Compose

You can also use Docker Compose to build and run specific services:

```bash
# Build and run specific services with Git commit hash as tag
GIT_COMMIT=$(git rev-parse HEAD | cut -c1-6) docker-compose build admin
GIT_COMMIT=$(git rev-parse HEAD | cut -c1-6) docker-compose up -d admin

# Build and run multiple services
GIT_COMMIT=$(git rev-parse HEAD | cut -c1-6) docker-compose build admin contest
GIT_COMMIT=$(git rev-parse HEAD | cut -c1-6) docker-compose up -d admin contest
```

### Build with Docker

If you need more flexible build options, you can directly use Docker commands:

```bash
# Build specific services with Git commit hash as tag
GIT_COMMIT=$(git rev-parse HEAD | cut -c1-6)
docker build --build-arg GIT_COMMIT=$GIT_COMMIT -t sastoj/admin:$GIT_COMMIT -f Dockerfile --target admin .
docker build --build-arg GIT_COMMIT=$GIT_COMMIT -t sastoj/contest:$GIT_COMMIT -f Dockerfile --target contest .
docker build --build-arg GIT_COMMIT=$GIT_COMMIT -t sastoj/public-auth:$GIT_COMMIT -f Dockerfile --target public-auth .
docker build --build-arg GIT_COMMIT=$GIT_COMMIT -t sastoj/gojudge-server:$GIT_COMMIT -f Dockerfile --target gojudge-server .
docker build --build-arg GIT_COMMIT=$GIT_COMMIT -t sastoj/freshcup-server:$GIT_COMMIT -f Dockerfile --target freshcup-server .
```
