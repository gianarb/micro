This service expose an http server on port 8000 and serve your current ip like output. It has also a healthcheck root `/health` to be sure that it's ready.

```
MYSQL_USERNAME=hello MYSQL_PASSWORD=******** MYSQL_ADDR=mysql.cc.net ./micro
```

This application has a proper repository in [hub.docker](https://hub.docker.com/r/gianarb/micro/) with few version of this service, feel free to use them.
```
docker run \
  -e MYSQL_USERNAME=hello 
  -e MYSQL_PASSWORD=******** \
  -e MYSQL_ADDR=mysql.cc.net \
  -p 8000:8000 \
  --net your-net \
  gianarb/micro:2.0.0
```

I create this application to have a good example around few topic:
* HealthCheck
* HTTP Microservices
* Docker
* 12factor application

And also to have an application to use during my talk/blogpost/article, right now I used that in few occasions:
* http://gianarb.it/blog/docker-1-12-orchestration-built-in
* DockerMeetup in Dublin, resources will be provided soon.

There are two version of this service, they provide both same route, index and health but with few differences:
* `1.0.0` the healthcheck is really simple, in practice doesn't do any sanity check for Thrid Party service.
* `2.0.0` the healthcheck was written to simulate a dependency with MySQL for this reason it checks if there is alive a MySQL instances. The index has a nice input (nice is too much :smile:)

## Home
Request
```
GET
/
```
Response
```
<ip>
```

## Healthcheck
Request
```
GET
/
```
Response: it return 200 if the health is good, in other case 500.
```
{
  "status": false,
  "info": [
    "database": "Tcp connection doesn't work"
  ]
}
```

## Docker
There is also a docker image that you can use
```
git pull gianarb/micro
```
