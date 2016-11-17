FROM abiosoft/caddy

COPY Caddyfile /etc/Caddyfile
COPY public/ /srv

EXPOSE 2015
