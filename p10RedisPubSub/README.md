[Redis]
-  docker run -it -p 6379:6379 -v $PWD/redis/data:/data -v $PWD/redis/config:/usr/local/etc/redis redis


[Executable]
- go run publisher.go
- go run subscriber.go