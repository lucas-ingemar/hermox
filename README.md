# Hermox

> Fast, structured logging for any command — turning raw stdout into machine-friendly JSON with context attached.

## About

Hermox is a lightweight Go utility that wraps any command and transforms its stdout into structured JSON log lines. Inspired by Hermes, the swift messenger of the gods, Hermox delivers logs with speed, consistency, and metadata you can rely on.

Whether you're piping application output into observability pipelines, shipping logs to centralized systems, or just want cleaner structured logs locally, Hermox keeps it simple.

## Features

* 🚀 Zero-config structured log transformation
* ⚡ High-performance streaming with minimal overhead
* 🏷 Attach arbitrary tags for context enrichment
* 🔄 Wrap any command without modifying application code
* 🎯 Focused on doing one thing well
* 🐹 Built in pure Go

## Usage

Run any command through Hermox and attach metadata tags:

```bash
hermox --tag1=foo --tag2 bar -- your_command
```

Everything after `--` is treated as the command to execute.

## Example

Command:

```bash
hermox --service api --env production -- ./server
```

Raw output from `./server`:

```text
Connecting to datastore...
Configuration loaded successfully
Worker pool initialized
Listening for incoming requests
```

Hermox output:

```json
{"time":"2026-07-03T11:38:48.118680073+02:00","level":"INFO","msg":"Initializing subsystem alpha","service":"api","env":"production"}
{"time":"2026-07-03T11:38:48.451227981+02:00","level":"INFO","msg":"Loading configuration bundle","service":"api","env":"production"}
{"time":"2026-07-03T11:38:49.002113420+02:00","level":"INFO","msg":"Background worker started","service":"api","env":"production"}
{"time":"2026-07-03T11:38:49.332904187+02:00","level":"INFO","msg":"Ready to accept connections","service":"api","env":"production"}
```

## Installation

Install directly with Go:

```bash
go install github.com/lucas-ingemar/hermox@latest
```

Or build from source:

```bash
git clone https://github.com/lucas-ingemar/hermox.git
cd hermox
go build -o hermox
```

## Why Hermox?

Most tools either require invasive integration or complex configuration. Hermox sits between your command and your logging backend, adding structure and context without touching your code.

It’s especially useful for:

* Wrapping legacy applications
* Standardizing logs across scripts and services
* Adding deployment metadata (environment, region, instance)
* Feeding logs into systems like Loki, Elasticsearch, or Splunk
