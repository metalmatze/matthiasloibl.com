.PHONY: all clean deps build

all: build

clean:
	rm -r public/
	rm hugo

deps:
	cd themes/material-design-lite && npm install -q && cd ../..
	curl -L -o hugo.tar.gz https://github.com/spf13/hugo/releases/download/v0.17/hugo_0.17_Linux-64bit.tar.gz
	tar xvf hugo.tar.gz
	mv hugo_0.17_linux_amd64/hugo_0.17_linux_amd64 hugo
	rm -rf hugo_0.17_linux_amd64 hugo.tar.gz

build:
	cd themes/material-design-lite && node_modules/.bin/gulp && cd ../..
	./hugo
