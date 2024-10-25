## Initialization
```bash
go mod init blockchain-newsfeed-server
```

### Docker
- Remove containers not in the file docker-compose.dev.yml
  ```
  docker-compose -f docker-compose.dev.yml up --build -d --remove-orphans 
  ```

### Migrations
```sql
go run ./cmd/migration/main.go -dir migrations create ${FILE_NAME} sql
go run ./cmd/migration/main.go -dir migrations up
```

### Containerization
### Deployment
#### Koyeb
```bash
# create service
koyeb app init shortener --git github.com/taintpro98/golang-url-shortener --git-branch main --git-builder docker --instance-type free 

# create database
koyeb database create shortener-db --app shortener --instance-type free --pg-version 16 --region was
```

### Reference
- [Deploying a Go and Postgres Application using Koyeb](https://wawand.co/blog/posts/deploying-a-go-app-to-koyeb)