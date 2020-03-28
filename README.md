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
  
get number of daily mewssages
<img width="1440" alt="daily-message" src="https://user-images.githubusercontent.com/45779140/77818480-497e2380-7105-11ea-9408-70e3bb046a8b.png">

  - /users/accounts/10-message
  
get top 10 accounts by messages
<img width="1440" alt="10-accounts by messages" src="https://user-images.githubusercontent.com/45779140/77818487-5d298a00-7105-11ea-85ec-c2d44096fad6.png">

  - /users/messages/10-engagement
  
get top 10 messages by engagements
<img width="1440" alt="10-messages by engagements" src="https://user-images.githubusercontent.com/45779140/77818494-6b77a600-7105-11ea-8bc6-a671631d4d6c.png">

  - /users/messages/word-clouds
  
 generate word clouds you can see the file in result/word_clouds.png ( In github , I generated it if you want to see before run this request ) 
<img width="1440" alt="word clouds" src="https://user-images.githubusercontent.com/45779140/77818498-792d2b80-7105-11ea-9302-a03cca4b1fca.png">

  - /users/messages/hashtag-clouds
  
 generate hashtag clouds you can see the file in result/hashtag_clouds.png ( In github , I generated it if you want to see before run this request ) 
<img width="1440" alt="hashtag clouds" src="https://user-images.githubusercontent.com/45779140/77818502-88ac7480-7105-11ea-8c36-009a3a075c64.png"> 
