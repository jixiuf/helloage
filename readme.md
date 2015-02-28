# My Info 
my mail hack1989huang # gmail.com
password 跟邮箱password 一样
http://jixiufcow1.appspot.com/

#  上传gae app失败的解决办法
 如 果你在国内的网络环境下，部署/上传gae程序时，遇到了上传失败的问题，那么你可以试试为 appcfg.py 套上一层 SOCKS 代理。具体做法如下（以 Mac OS X 上的 Google App Engine Launcher 的安装位置为例）：

     下载 SocksiPy，提取出其中的 socks.py，放入 /usr/local/google_appengine。
     在 appcfg.py 的 import sys 一行的下面加上这 4 行：

     import socks
     import socket
     socket.socket = socks.socksocket
     socks.setdefaultproxy(socks.PROXY_TYPE_SOCKS5, "127.0.0.1", 7070)

 其中最后一行中的 127.0.0.1:7070 即为 SOCKS 代理的地址和端口号，你可以换成自己的。至于如何获得这样的 SOCKS 代理，最简单的方式就是通过 SSH Tunnel 或 Tor。
 ssh -D 7070  deployer@thstage-japan -N
 --------------------------------------------------------------------
 appcfg.py后加一个参数（--insecure）就能部署了.
 --insecure 表示使用http和服务器通讯，默认的是https (我加了反而布署不了，不加则可以)
 google_appengine/appcfg.py --insecure update xxx
 ------------------------------------------------------------------------------------------


