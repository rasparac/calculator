# calculator

Simple math calculator which will make your life easier :). 

This project will spin up simple http server which can be use to  add, subtract, multiply and divide only two numbers via HTTP request.


To add two numbers use this URL:
```
http(s)/{host}:{port}/add?x=1&y=2
```

To subtract two numbers use this URL:
```
http(s)/{host}:{port}/subtract?x=1&y=2
```

To add multiply numbers use this URL:

```
http(s)/{host}:{port}/multiply?x=1&y=2
```

To divide two numbers use this URL:

```
http(s)/{host}:{port}/divide?x=1&y=2
```

## Getting Started


### Prerequisites
 - GO
 - Docker

### ENV

| NAME | TYPE | DEFAULT |COMMENT|
|:-----|:--------:|:--------:|:--------:| 
| HOST   | string | 0.0.0.0 ||
| PORT   | string | 9999|without colon|

### run project
To run project use this command:
```
make run
```

### build project
To build executable file run this commands:
```
make build
```
if you use windows use this command:
```
make build GOOS=windows
```

### tests
To run unit tests runs this command:
```
make test
```

### docker
If you have troubles with building and running project, you can use docker to spin up calculator server. Check `docker-compose.yml` if you want to change env vars.


```
make docker-up
```

You can also build docker image with:
```
make build-docker-image
```