FROM golang:1.22.1-bullseye

# Set destination for COPY
WORKDIR ./
COPY . ./
CMD go run main.go ${HOST_IP}