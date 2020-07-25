# Build
FROM golang:1.13.11-buster AS build
ARG DEBIAN_FRONTEND=noninteractive
WORKDIR /
COPY . /dispatcher

# Build dispatcher
ENV GO111MODULE=on CGO_ENABLED=0
RUN cd dispatcher && go build .

# Run
FROM scratch
COPY --from=build \
	/dispatcher/dispatcher /dispatcher
COPY --from=build \
	/dispatcher/static /static

ENV IFCONFIGIS_DISPATCHER_PORT=5080
ENV IFCONFIGIS_API="http://api:5000"
ENV IFCONFIGIS_STATIC=/static
ENTRYPOINT ["/dispatcher"]
