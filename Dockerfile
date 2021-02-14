FROM caddy:2-alpine
WORKDIR /var/www/html
COPY ./Caddyfile /etc/caddy/Caddyfile
COPY ./public/ /var/www/html

EXPOSE 80
