pow-protector
====
Server and client services for testing PoW as ddos protector.

Using
----
* create network between containers
```dockerfile
docker network create pow
```
* build server and client containers
```dockerfile
docker build -t pow-client client/.
docker build -t pow-server server/.
```
* run containers
```dockerfile
docker run --rm --network pow --name pow-client -d pow-client
docker run --rm --network pow --name pow-server -d pow-server
```
* exec pow client
```dockerfile
docker exec -it  pow-client bash
```
* and run pow-client
```dockerfile
pow-client -host pow-server
```

Pow-client knows only commands: 
* **get** - get text from server, after make work and send proof
* **exit** - close connect and exit
