# Alertmanager message broker


## Prerequisites

- Go 1.16+
- Sqllite driver

## About:

The alertmanager message broker is a project made to meet some of my needs to integrate alertmanager with other platforms, through webhooks.

It has the function of registering endpoints (who will receive the alerts) so the same alert is sent to multiple platforms with specific payloads, this record is kept in a sqlite database created by the application, located in /var/lib/sql/database.db.

The logic is present in controller/controller.go, which have the main function of the application.

## Build:

* You can compile directly with ````go build -ldflags '-extldflags "-static"' -o alertmanager-broker````.

To use docker:

* ````docker build -t repo@usuario/alertmanager-broker .````
* ````docker run -t repo@usuario/alertmanager-broker -p 3000:3000````

## Variables(optional):

If you set url and endpoint as environment variables, they will create an entry in database and will be used as default.

| Variable | Description |
| --- | --- |
| PORT | Set port that the application will be listening to, default port is 3000 |
| URL | URL with protocol and hostname, example: https://api.google.com  |
| ENDPOINT | Endpoint that will Example: api/v1/message/send |

## Endpoints:

| Endpoint | Method | Description |
| --- | --- | --- |
| /api/v1/endpoint | POST  | Add/Remove endpoint  (see example of usage bellow) |
| /api/v1/endpoints | GET | List all endpoints registered, with ID (required to delete endpoint) |
| /api/v1/send | POST | Endpoint that will be used by alertmanager to receive the webhook and send alerts to endpoints |
| /api/v1/ping | GET | Send a test alert to endpoints |

Example of usage of add/delete endpoints:

Add: 
````console
$ curl -X POST http://localhost:3000/api/v1/endpoint -d '{"action":"add","url":"https://contoso.com","endpoint":"/api/v1/message/send","alertname":"test"}' 
````

Delete:
````console
$ curl -X POST http://localhost:3000/api/v1/endpoint -d '{"action":"delete","id":"1"}'
````

