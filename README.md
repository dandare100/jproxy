## A simple transparent tcp proxy

### To build 

```CGO_ENABLED=0 GOOS=linux go build -a -o prox  .```

```docker build .```

### To use

#### Transparent TCP Proxy

```Usage : ./prox lport rhost rport```

#### Docker image usage (We all know why this is useful :-)

```docker run -it -p443:4444 jport:latest /prox 4444 192.168.8.1 80```
