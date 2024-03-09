FROM golang:1.22.1-bullseye

# Set destination for COPY
WORKDIR ./


# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./
