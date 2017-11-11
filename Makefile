.PHONY: all clean deps build

all: build

clean:
	rm -r public/
	rm hugo

deps:
	cd themes/material-design-lite && yarn install && cd ../..
	rm -rf hugo hugo_tmp/ hugo.tar.gz
	curl -L -o hugo.tar.gz https://github.com/gohugoio/hugo/releases/download/v0.30.2/hugo_0.30.2_Linux-64bit.tar.gz
	mkdir hugo_tmp && tar xvf hugo.tar.gz -C hugo_tmp/
	mv hugo_tmp/hugo hugo
	rm -rf hugo_tmp/ hugo.tar.gz

build:
	cd themes/material-design-lite && node_modules/.bin/gulp && cd ../..
	./hugo
