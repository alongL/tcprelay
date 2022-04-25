# tcprelay
tcprelay helps client accessing the host unreachable.
For example:  
- Virtual machine vm1 runing on server A, with NAT ip 192.168.122.101.  
- Server A has public IP 1.2.3.4.  
- Client can access server A, but can not access tcp port 80 of vm1 directly.  
- With tcprelay running on server A.  
- Client can access tcp port 80 of vm1 directly by 1.2.3.4:8080.  


# Figure
Client A can not direct access  target machine, for some reason firewall or NAT etc.  
The other middle machine can access target.
Tcp client can access target with tcp by tcprelay running on middle machine.

```
                                                
                                             +----------------+
                       X                     |                |
        +------------------------------->    |     target     |(vm1)
        |        can not access              |                |
        |                                    +----------------+
        |                                             ^  192.168.122.101:80
        |                                             |
        |                                             | tcprelay
        |                                             |
+--------------+                             +----------------+
|              |                             |                |
|  tcp client  | +------------------------>+ |     middle     |(server A)
|              |                1.2.3.4:8080 |                |
+--------------+                             +----------------+
                                                192.168.122.1



```
# how to run 
on server A, run tcprelay.  
```
./tcprelay -l ":8080"  -t "192.168.122.101:80"
```
It will listen on 8080, and all the client connected to :8080  will be relay to 192.168.1.101:80 

For the client, just  type  http://1.2.3.4:8080   you can access the web service on target machine.


# how to build  
download this code, cd into the dir.
type command as follows:  
```
SET GOOS=linux
SET GOARCH=amd64
SET GOPROXY=https://goproxy.cn
go build .
```


# download 
if you don't want to compile it. just download in release page.  
It's very easy to use.



# reference
This program is taken from this two repository. they are the hero.

https://github.com/xtaci/kcptun 

https://github.com/hikaricai/p2p_tun 



