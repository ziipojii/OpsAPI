FROM --platform=linux/amd64 asia-southeast1-docker.pkg.dev/tk-dev-micro/base-image/golang:1.22.4 as base-golang

# Set workdir for base-golang
WORKDIR /app

# download lib dependencies
COPY go.mod go.sum /app/
RUN go mod download

# copy all files to ./app
COPY . /app

# build go binary
RUN CGO_ENABLED=0 go build -o main ./src/cmd/main.go


# Setup application
FROM --platform=linux/amd64 asia-southeast1-docker.pkg.dev/tk-dev-micro/base-image/distroless-go
COPY --from=base-golang /app/main /app/main
CMD ["/app/main"]
