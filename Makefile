# -*- coding:utf-8 -*-
.PHONY: deploy
deploy:
# goapp deploy
# or
# appcfg.py后加一个参数（--insecure）就能部署了.
# --insecure 表示使用http和服务器通讯，默认的是https
# google_appengine/appcfg.py --insecure update xxx
	appcfg.py -v  --email=hack1989huang@gmail.com update .
