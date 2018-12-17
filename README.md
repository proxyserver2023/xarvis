# gapi [WIP]

Go rest api bootstrap template with SOLID and Clean Architecture using Microservices Best Practices.

## Mongodb with go api

 Production Quality Mongodb in docker
 1. data should be spearated from the docker container
 2. custom config
 3. admin tasks

```bash
docker run --name MONGO --restart=always -d -p 27017:27017 mongo mongod --auth
docker exec -it MONGO bash
```