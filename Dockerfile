# Must use golang:1.12, since go download was introduced then
FROM golang:1.12rc1-alpine3.9

WORKDIR /src

# Alpine images do not include git by default, we'll need to add it
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Copy only project dependency files
COPY go.mod .
COPY go.sum .

# Download our dependencies. Since dependencies change far less frequently 
# than source files do, this make subsequent builds a lot faster
RUN go mod download

# Include files.
ADD . .

# Building app
RUN mkdir /app; go build -o goapp && cp goapp /app/

# Just the runtime, no build tools. 
# This 2nd stage only includes files that we explicitly copy over,
# essentially starting from scratch.
FROM alpine:3.9
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=0 /app /app
EXPOSE 8000
ENTRYPOINT ["./goapp"]