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
<img width="1440" alt="daily-message" src="https://user-images.githubusercontent.com/45779140/77818055-11c1ac80-7102-11ea-9600-cee69d57b79e.png">

  - /users/accounts/10-message
  
get top 10 accounts by messages
<img width="1440" alt="10-account by message" src="https://user-images.githubusercontent.com/45779140/77818084-51889400-7102-11ea-9447-535097e07095.png">

  - /users/messages/10-engagement
  
get top 10 messages by engagements
<img width="1440" alt="10-message by engagement" src="https://user-images.githubusercontent.com/45779140/77818126-99a7b680-7102-11ea-8f2b-f3d69ea1b0d6.png">

  - /users/messages/word-clouds
  
 generate word clouds you can see the file in result/word_clouds.png ( In github , I generated it if you want to see before run this request ) 
<img width="1440" alt="wordclouds" src="https://user-images.githubusercontent.com/45779140/77818142-ba700c00-7102-11ea-839d-b0fd6756d775.png">

  - /users/messages/hashtag-clouds
  
 generate hashtag clouds you can see the file in result/hashtag_clouds.png ( In github , I generated it if you want to see before run this request ) 
 <img width="1440" alt="hashtagcloulds" src="https://user-images.githubusercontent.com/45779140/77818203-379b8100-7103-11ea-9596-f64744a3f218.png">



 
