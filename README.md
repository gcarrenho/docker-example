# Docker Example

This project is intended to provide a little example of use docker and docker compose building and run the app with a internal mysql database.
Follow the clean architecture package by components.

## Requirements

- Docker and Docker compose (Docker compose comes with docker).
- VS code or you prefer ui.

## Setup

All development and compilation is performed in Docker. An environment Docker is setup inside the project.
In order to lunch it, just:
- Open repository in your prefered ui.
- Install the plugings needed to use docker in you ui if you prefer. In other case run docker by command line.

### OS Supported

Windows and Linuxs OS are supported.

### Environment variables

By default, all the env variables will be loaded from `dev.env`
For local development you should edit dev.env and fill it with your variables.

### Docker + Docker compose

The project is composed of some containers:
- docker-example (`docker-example-app`) exposed to host in port `45475`
- Mysql 8 Database (`mysqldb-example`) exposed to host in port `3308`

You shuld run everything using docker compose.

### Connect to db by command line

```bash
docker exec -it mysqldb-example mysql -uroot -p
```

### Conect to db from host

In order to connect to the db from your host, you should use a proper tool (DBeaver, MySQL Workbench, CLI...) with these parameters:
- Host: `localhost`
- POrt: `3308`
- User: `root`
- Password: `developer`
- DB Name: `exampledb`

### Run

In order to start the project you can start it by command line:
```bash
docker compose up
```

### Utils Docker commands

Some utils command apart of docker start and docker stop

In order to 
- Build a docker file, you should run:
```bash
docker build -t dockerexample:1.0 .
```
It take the Dockerfile and run it, check that all arg and env are setting.

- Build a specific docker file, you should run
```bash
docker build -t dockerexample:1.0 -f sonarq.Dockerfile .
```
It will take the sonarq.Dockerfile to exec and build it

- Know whats docker are running
```bash
docker ps
```

- See the volume
```bash
docker volume ls
```

- Removed the volume
```bash
docker volume rm docker-example_mysqldbdata
```

- Clean the unused resources of docker you can run
```bash
docker system prune
```

- Removed a container:
```bash
docker rm docker-example-app
```

- Removed a image:
```bash
docker rmi docker-example-docker-example-app:latest
```

### Testing

Testing can be performed throght go test ./...