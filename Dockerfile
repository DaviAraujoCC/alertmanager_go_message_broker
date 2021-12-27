FROM golang:1.16 AS build

WORKDIR /src/

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags '-extldflags "-static"' -o /bin/alertmanager-broker

RUN mkdir -p /var/lib/sql

FROM centurylink/ca-certs

COPY --from=build /bin/alertmanager-broker  /bin/alertmanager-broker
COPY --from=build /var/lib/sql  /var/lib/sql

ENTRYPOINT [ "/bin/alertmanager-broker" ]