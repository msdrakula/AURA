# syntax=docker/dockerfile:1

# --- Сборка -----------------------------------------------------------------
# CGO не нужен: SQLite — чистый Go (modernc.org/sqlite),
# GTK/WebKit-окно отключаем флагом --window=false (в контейнере есть только web-UI).
FROM golang:1.25-bookworm AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/aura ./cmd/server

# --- Рантайм ----------------------------------------------------------------
FROM debian:bookworm-slim
RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates \
 && rm -rf /var/lib/apt/lists/* \
 && useradd --system --uid 10001 --create-home aura

WORKDIR /app
COPY --from=build /out/aura /app/aura
COPY web /app/web
RUN mkdir -p /app/data && chown -R aura:aura /app

USER aura
ENV AURA_ROOT=/app
VOLUME ["/app/data"]

# 1337 — UI/API, 8080 — HTTP-прокси (8082 — callback, слушает только
# контур контейнера, см. internal/callback/server.go)
EXPOSE 1337 8080

ENTRYPOINT ["/app/aura", \
    "--window=false", \
    "--api-host=0.0.0.0", \
    "--proxy-host=0.0.0.0", \
    "--db-path=data/aura.db", \
    "--ca-dir=data"]
