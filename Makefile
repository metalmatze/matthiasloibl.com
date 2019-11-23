all: build

.PHONY: clean
clean:
	rm -r public/
	rm hugo

.PHONY: npm
npm:
	cd themes/material-design-lite && npm install && cd ../..

.PHONY: hugo
hugo:
	rm -rf hugo hugo_tmp/ hugo.tar.gz
	curl -L -o hugo.tar.gz https://github.com/gohugoio/hugo/releases/download/v0.48/hugo_0.48_Linux-64bit.tar.gz
	mkdir hugo_tmp && tar xvf hugo.tar.gz -C hugo_tmp/
	mv hugo_tmp/hugo hugo
	rm -rf hugo_tmp/ hugo.tar.gz

.PHONY: gulp
gulp:
	cd themes/material-design-lite && node_modules/.bin/gulp && cd ../..

.PHONY: build
build:
	hugo
