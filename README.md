# Social Dashboard

## Getting started
With Dockerfile
```
docker build -t social-dashboard .
docker run --rm -p 8080:8080 --name social-dashboard -it social-dashboard
```

With go cmd ( recommend )
```
go run main.go
```

Note: If you run by dockerfile , Some API path use much time for request that if ram in docker file is out, it'll be stopped

## API
  - /users/daily/message
```
![number of daily message](/images/daily-message.png?raw=true)
```)