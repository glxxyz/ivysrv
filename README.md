## ivysrv demo project

Executes Ivy commands e.g.: http://localhost:8080/ivy/2+2

Ivy source & documentation: https://robpike.io/ivy

## Local Testing
```
go mod download robpike.io/ivy
```
## Deploying
```
docker build . -t ivysrv
docker run -p 8000:8080 ivysrv
``` 