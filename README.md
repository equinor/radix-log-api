# radix-log-api

The Radix Log API gives access to container logs from Azure Log Analytics Workspaces for applications hosted in Radix.


## Configuration

**Command line arguments**

| Name | Type | Required | Description | Default |
| ---- | ---- | -------- | ----------- | ------- |
| host | string | No | Host name or IP for the server | "8000" |
| port | string | No | Port that the server listens to | "" |
| auth-issuer | string | Yes | The expected issuer (`iss` claim) for JWT used as bearer token in authorization header for incoming requests | "" |
| auth-audience | string | Yes | The expected audience (aud) for JWT used as bearer token in authorization header for incoming requests | "" |
| log-analytics-workspace-id | string | Yes | The workspace ID for the Azure Log Analytic Workspace to collect container logs from | "" |
| radix-api-host | string | Yes | FQDN to the Radix API server  | "" |
| radix-api-path | string | No | Base path for Radix API | "/api/v1" |
| radix-api-scheme | string | No | The Radix API HTTP scheme (https or http) | "https" |

Every command line argument can be specified as an environment variable by prefixing it with `LOG_API_`, capitalizing it, and replicaing hyphens (`-`) with underscores (`_`), e.g. `auth-issuer` becomes `LOG_API_AUTH_ISSUER`.

## Developing

You need Go installed. Make sure `GOPATH` and `GOROOT` are properly set up.

Also needed:

- [`gomock`](https://github.com/golang/mock) (go install github.com/golang/mock/mockgen@v1.6.0)

Clone the repo into your `GOPATH` and run `go mod download`.

### Swagger

Swagger docs are generated using [https://github.com/swaggo/swag](https://github.com/swaggo/swag). The [Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format) describes how to configure the swagger spec.

### Generating mocks
We use gomock to generate mocks used in unit test.
You need to regenerate mocks if you make changes to any of the interface types used by the application

`make mocks`

### Running locally

The following envirnoment variables are needed.  
Copy file `.env.template` to `.env`

* `LOG_API_AUTH_ISSUER`
* `LOG_API_AUTH_AUDIENCE`
* `LOG_API_LOG_ANALYTICS_WORKSPACE_ID`
* `LOG_API_RADIX_API_HOST`

## Deployment

Radix Log API follows the [standard procedure](https://github.com/equinor/radix-private/blob/master/docs/how-we-work/development-practices.md#standard-radix-applications) defined in _how we work_. 

Radix Log API is installed as a Radix application in [script](https://github.com/equinor/radix-platform/blob/master/scripts/install_base_components.sh) when setting up a cluster.
