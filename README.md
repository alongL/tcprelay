# tcprelay
tcprelay help a tcp client to access the other host unreachable


# aim
Client A can not direct access  target machine, for some reason.
The other machine can access target.
Tcp client can access target with tcp by tcprelay running in middle.


```


                                                192.168.1.101
                                             +----------------+
                       X                     |                |
        +------------------------------->    |    target      |
        |        can not access              |                |
        |                                    +----------------+
        |                                             ^ ip2:port2
        |                                             |
        |                                             |
        |                                             |
+--------------+                             +----------------+
|              |                    tcprelay |                |
|  tcp client  | +-------------------------> |     middle     |
|              |                    ip1:port1|                |
+--------------+                             +----------------+
                                                192.168.1.100



```
# how to run
```
./tcprelay -l ":2022"  -t "192.168.1.101:80"
```
it will listen on 2022, and all the client connected to :2022  will be relay to 192.168.1.101:80 

for the client just  type  http://192.168.1.100:2022   you can access the web service on target machine.


# how to build 
download this code.

cd into the dir
(on windows platform, compile linux version)
```
SET GOOS=linux
SET GOARCH=amd64
GOPROXY=https://goproxy.cn
go build .
```
compile windows version 
```
go build .
```


# download 
if you don't want to compile it. just download in release page. 
It's very easy to use.



# reference
This program is get from this two repository. they are the hero.
I do only the tcptun. 

https://github.com/xtaci/kcptun 
https://github.com/hikaricai/p2p_tun 



