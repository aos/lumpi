## Enable monitoring of Nginx logs

1. We use GoAccess for this (https://goaccess.io)
2. Use this command to run the server:

```
sudo goaccess /var/log/nginx/access.log -o report.html --log-format=COMBINED --real-time-html --ws-url=wss://lumpi.host/stats &
```

This will generate a "report.html" file, we connect it to the server
websocket @ "lumpi.host/stats". The server actually runs on:
"localhost:7890"

3. This will be proxied via nginx to localhost:7890, see nginx conf file @
/etc/nginx/conf.d/default.conf

```nginx
	# Proxy the real time monitor report
	location /monitor {
		alias /home/aos/servers;
		index report.html;
	}

	# Proxy the websocket server for monitoring
	location /stats {
		proxy_pass http://localhost:7890;
		proxy_http_version 1.1;
		proxy_set_header Upgrade $http_upgrade;
		proxy_set_header Connection "upgrade";
		proxy_read_timeout 86400;
	}
```
