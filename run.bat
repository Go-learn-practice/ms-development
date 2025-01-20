chcp 65001
cd devUser
docker build -t dev-user:latest .
cd ..
docker-compose up -d