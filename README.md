# calculator

Simple math calculator which will make your life easier :). 

This project will spin up simple http server which can be use to  add, subtract, multiply and divide only two numbers via HTTP request.


## Getting Started


### Prerequisites
 - GO
 - Docker

### ENV

| NAME | TYPE | DEFAULT |COMMENT|
|:-----|:--------:|:--------:|:--------:| 
| HOST   | string | 0.0.0.0 |           |
| PORT   | string | 9999    |without colon|


## Operations
Here is the list of endpoints you can use. To use this server you must start project.
Calculator will store results with the same problem in cache and return the answer from cache for the same problem. Cache will expire after 1 min. 

Add two numbers use this URL:
```
http/{host}:{port}/add?x=1&y=2
```

Subtract two numbers use this URL:
```
http/{host}:{port}/subtract?x=1&y=2
```

Multiply numbers use this URL:

```
http/{host}:{port}/multiply?x=1&y=2
```

Divide two numbers use this URL:

```
http/{host}:{port}/divide?x=1&y=2
```

Successfull response:
```
{
    "action": "add|divide|multiply|subtract",
    "answer": 0,
    "x": 0,
    "y": 0,
    "cached": true
}
```

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