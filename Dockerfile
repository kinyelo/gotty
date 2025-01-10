FROM node:16 as js-build
WORKDIR /gotty
COPY js /gotty/js
COPY Makefile /gotty/
RUN make bindata/static/js/gotty.js.map

FROM golang:1.23.4 as go-build
WORKDIR /gotty
COPY . /gotty
COPY --from=js-build /gotty/js/node_modules /gotty/js/node_modules
COPY --from=js-build /gotty/bindata/static/js /gotty/bindata/static/js
RUN CGO_ENABLED=0 make

FROM alpine:latest
RUN apk update && \
    apk upgrade && \
    apk --no-cache add ca-certificates bash openssh

COPY --from=go-build /gotty/gotty /usr/bin/
RUN addgroup -S gotty && adduser -S gotty -G gotty
RUN chown -R root:root /home/gotty

RUN mkdir -p /home/gotty/bin
### The Bash shell can detect when it has been invoked using "rbash" instead of "bash." ###
RUN ln -s /bin/bash /home/gotty/bin/rbash
RUN ln -s /usr/bin/gotty /home/gotty/bin
RUN ln -s /usr/bin/ssh /home/gotty/bin
RUN chown -R gotty:gotty /home/gotty/bin

ENV PATH=/home/gotty/bin
WORKDIR /home/gotty
CMD ["gotty", "-w", "rbash"]
