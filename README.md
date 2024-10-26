## Initialization
```bash
go mod init blockchain-newsfeed-server
```

### Docker
### Migrations
```sql
go run ./cmd/migration/main.go -dir migrations create ${FILE_NAME} sql
go run ./cmd/migration/main.go -dir migrations up
```

### Containerization
```bash
# Remove containers not in the file docker-compose.dev.yml
docker-compose -f docker-compose.dev.yml up --build -d --remove-orphans
# Build an image on local
docker build --build-arg TELEGRAM_TOKEN=$(grep TELEGRAM_TOKEN .env | cut -d '=' -f2) \
             --build-arg TELEGRAM_CHAT_ID=$(grep TELEGRAM_CHAT_ID .env | cut -d '=' -f2) \
            -t golang-url-shortener .
# Run container
docker run -d -p 8080:8080 --name golang-url-shortener-container golang-url-shortener

# Start container
docker start golang-url-shortener-container
```
### Deployment
#### Koyeb
```bash
# create service
koyeb app init shortener --git github.com/taintpro98/golang-url-shortener --git-branch main --git-builder docker --instance-type free --env "POSTGRES_HOST=ep-quiet-night-a4ehz4z4.us-east-1.pg.koyeb.app" --env "POSTGRES_PORT=5432" --env "POSTGRES_DB=koyebdb" --env "POSTGRES_USER=koyeb-adm" --env "POSTGRES_PASSWORD=[password]"

# deploy with yaml file
koyeb services deploy -f koyeb.yaml

# create database
koyeb database create shortener-db --app shortener --instance-type free --pg-version 16 --region was
```

### Reference
- [Deploying a Go and Postgres Application using Koyeb](https://wawand.co/blog/posts/deploying-a-go-app-to-koyeb)