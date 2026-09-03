FROM alpine:3.23

# Create directories
WORKDIR /soft-serve
# Expose data volume
VOLUME /soft-serve

# Environment variables
ENV SOFT_SERVE_DATA_PATH "/soft-serve"
ENV SOFT_SERVE_INITIAL_ADMIN_KEYS ""
# workaround to prevent slowness in docker when running with a tty
ENV CI "1"

# Expose ports
# SSH
EXPOSE 23231/tcp
# HTTP
EXPOSE 23232/tcp
# Stats
EXPOSE 23233/tcp
# Git
EXPOSE 9418/tcp

# Set the default command
ENTRYPOINT [ "/usr/local/bin/soft", "serve" ]

RUN apk add --no-cache git bash openssh curl libcurl expat=2.8.4-r0

COPY soft /usr/local/bin/soft

