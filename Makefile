PREFIX = /usr
BINDIR = $(PREFIX)/bin
STATICDIR = $(PREFIX)/share/beehive

beehive-server: server.go
	go build -o $@ server.go

convert: convert.go
	go build -o $@ convert.go

STATICFILES = COPYING README.md robots.txt beehive.html beehive.css beehive.js
BACKENDPORT = 29092
NGINXCONF = /etc/nginx/nginx.conf
SYSTEMDCONF = /lib/systemd/system/beehive.service

install: beehive-server
	install -m 0755 -t $(BINDIR) beehive-server
	install -m 0644 -D -t $(STATICDIR) $(STATICFILES)

uninstall:
	rm -rf $(STATICDIR)
	rm -f $(BINDIR)/beehive-server

configure-nginx:
	install -m 0644 --backup nginx.conf $(NGINXCONF)
	sed -i "s#HOSTNAME#$(HOSTNAME)#g" $(NGINXCONF)
	sed -i "s#STATICDIR#$(STATICDIR)#g" $(NGINXCONF)
	sed -i "s#BACKENDPORT#$(BACKENDPORT)#g" $(NGINXCONF)

configure-systemd:
	install -m 0644 beehive.service $(SYSTEMDCONF)
	sed -i "s#BINDIR#$(BINDIR)#g" $(SYSTEMDCONF)
	sed -i "s#BACKENDPORT#$(BACKENDPORT)#g" $(SYSTEMDCONF)

clean:
	rm -f beehive-server convert

.PHONY: install uninstall configure-nginx configure-systemd clean
