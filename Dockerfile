ARG GO_VERSION=1
ARG JAVA_VERSION=21

FROM golang:${GO_VERSION} as deps
RUN GOBIN=/ go install github.com/DarthSim/overmind/v2@latest


FROM golang:${GO_VERSION} as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /axiom .


FROM eclipse-temurin:${JAVA_VERSION}-jre
WORKDIR /app

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    ca-certificates \
    tmux \
    nano \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /axiom /usr/local/bin/
COPY --from=deps /overmind /usr/local/bin/
COPY lavalink /usr/local/bin/
RUN chmod +x /usr/local/bin/lavalink

COPY Lavalink.jar Procfile ./
CMD ["overmind", "start"]
