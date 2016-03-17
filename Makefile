.PHONY: all clean deps build

HUGO ?= 0.15_linux_amd64
THEME ?= themes/material-design-lite

all: build

clean:
	rm -r public/
	rm hugo

deps:
	cd $(THEME) && npm install -q && cd ../..
	curl -L -o /tmp/hugo_$(HUGO).tar.gz https://github.com/spf13/hugo/releases/download/v0.15/hugo_$(HUGO).tar.gz
	tar xvf /tmp/hugo_$(HUGO).tar.gz -C /tmp/
	mv /tmp/hugo_$(HUGO)/hugo_$(HUGO) hugo

build:
	cd $(THEME) && node_modules/.bin/gulp && cd ../..
	./hugo
