<pre>
</pre>

# Hermox

> Swift as the messenger god himself, Hermox transforms your mortal stdout into divine JSON logs.

## About

Hermox is a lightning-fast Rust utility that transforms plain stdout into structured JSON log lines. Like its namesake Hermes, it acts as a messenger between your programs and your logging system, delivering your logs in a format that modern tools understand.

## Features

- 🚀 Zero-config JSON log transformation
- ⚡ Blazingly fast processing
- 🔄 Seamless pipe integration
- 🎯 Single responsibility: does one thing extremely well
- 🦀 Written in pure Rust

## Usage

```bash
your_command | hermox
```

## Example

Input:
```
Server started on port 3000
Request received from 192.168.1.1
```

Output:
```json
{"timestamp":"2025-05-12T10:23:45Z","level":"INFO","message":"Server started on port 3000"}
{"timestamp":"2025-05-12T10:23:46Z","level":"INFO","message":"Request received from 192.168.1.1"}
```

## Installation

```bash
cargo install hermox
```

## License

MIT License - Copyright (c) 2025 [Your Name]
