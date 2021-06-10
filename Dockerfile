FROM golang:1.16-alpine
WORKDIR /build
COPY . .
RUN go build -o ivysrv ./src
WORKDIR /bin
RUN cp /build/ivysrv .
EXPOSE 8000
CMD ["/bin/ivysrv"]