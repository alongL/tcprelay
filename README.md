# tcptun
a tcp tunnel to serve two host


# usage
client can not direct access  target machine,for some reason.
the other machine can access target.
tcp client can access target with tcp by tcptun running in middle.


```
                                               +----------------+
                         X                     |                |
          +------------------------------->    |    target      | (192.168.1.101)
          |        can not access              |                |
          |                                    +----------------+
          |                                             ^ ip2:port2
          |                                             |
          |                                             |
          |                                             |
  +--------------+                             +----------------+
  |              |                             |                |
  |  tcp client  | +-------------------------> |    tcptun      | (192.168.1.100)
  |              |                    ip1:port1|                |
  +--------------+                             +----------------+

```
# how to use
./tcptun -l ":2022"  -t "192.168.1.101:80"

for the client just  type  http://192.168.1.100:2022   you can access the web service on target machine.


# reference
This program is get from this two repository. they are the hero.
I do only tcptun.
https://github.com/xtaci/kcptun 
https://github.com/hikaricai/p2p_tun



