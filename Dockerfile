ARG GO_VERSION=1
ARG JAVA_VERSION=21

FROM eclipse-temurin:${JAVA_VERSION} as jre
RUN $JAVA_HOME/bin/jlink \
    --add-modules java.base,java.logging,java.desktop,java.scripting,java.security.sasl,jdk.management,jdk.unsupported \
    --strip-debug \
    --no-man-pages \
    --no-header-files \
    --compress=zip-6 \
    --output /jre


FROM golang:${GO_VERSION} as deps
RUN GOBIN=/ go install github.com/DarthSim/overmind/v2@latest


FROM golang:${GO_VERSION} as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /axiom .


FROM debian:bookworm-slim
WORKDIR /app

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    ca-certificates \
    tmux \
    && rm -rf /var/lib/apt/lists/*

ENV JAVA_HOME=/opt/java/openjdk
ENV PATH "${JAVA_HOME}/bin:${PATH}"
COPY --from=jre /jre $JAVA_HOME

COPY --from=deps /overmind /usr/local/bin/
COPY --from=build /axiom /usr/local/bin/

COPY Lavalink.jar Procfile ./
CMD ["overmind", "start"]
