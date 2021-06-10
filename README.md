## ivysrv demo project

Executes Ivy commands e.g.: http://localhost:8000/2+2

Ivy source & documentation: https://robpike.io/ivy

## Local Testing
```
go mod download robpike.io/ivy
go run ./src
```
## Running with Docker
```
docker build . -t ivysrv
docker run -p 8000:8000 ivysrv
``` 