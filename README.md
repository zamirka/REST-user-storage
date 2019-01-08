# REST-user-storage
run command
docker-compose -f .docker/docker-compose.yml up -d

credentials and settings are stores in environment variables in corresponfing files in .docker directory
- .docker/app.env for application itself
- .docker/mongo.env for mongodb
- .docker/mongo-express.env for mongo-express

accessing

application itself:
localhost:8080

mongo-express:
localhost:8081