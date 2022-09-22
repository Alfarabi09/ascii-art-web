# Ascii-art-web

**Ascii-art-web** consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of Ascii-art.

**Ascii-art** is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.

### **My webpage allow the use of the different banners**
+ [shadow](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
+ [standard](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
+ [thinkertoy](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)

![image](https://highload.today/wp-content/uploads/2021/12/golang.jpeg)

## Program architecture

- cmd
	- main
		- main.go
- internal
	- handler
		- handler.go
	- server
		- server.go
- pkg
	- ascii
		- ascii.go
		- files
			- standard.txt
			- shadow.txt
			- thinkertoy.txt
- template
	- index.html
---
*To run project you need use this command in terminal*
```
go run cmd/main/main.go
```

*Or you can run Docker,use this commands in terminal*
```
$ docker image build -f Dockerfile -t imagename .

$ docker images

$ docker container run -p 8080:8080 --detach --name containername imagename

$ docker ps -a

$ docker exec -it containername ls -l

$ docker inspect imagename | jq -r .[].Config.Labels

```
