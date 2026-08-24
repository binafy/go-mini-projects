<div align="center">

# 🐹 Go Mini Projects

### A hands-on collection of small, focused Go programs — from tiny CLIs to real-time web servers.

_Learn Go by building. Each project is self-contained, idiomatic, and dependency-light._

<br/>

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](#-license)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen?style=for-the-badge)](#-contributing)
[![Made with ♥](https://img.shields.io/badge/Made%20with-%E2%99%A5-red?style=for-the-badge)](https://github.com/binafy)

<br/>

[**Getting Started**](#-getting-started) ·
[**Projects**](#-projects) ·
[**Roadmap**](#-roadmap) ·
[**Contributing**](#-contributing)

</div>

---

## 📖 About

**Go Mini Projects** is a growing catalog of bite-sized programs written in pure, idiomatic Go. Every folder is an independent project you can read top to bottom in a single sitting, run in one command, and learn one concept from — flags, file I/O, encoding, concurrency, HTTP servers, real-time streaming, and more.

The goal is simple: **learn Go by shipping small things that actually work.**

- ✅ **Self-contained** — each project stands on its own, no framework, no boilerplate.
- ✅ **Idiomatic** — clean code, standard library first, dependencies only when they earn their place.
- ✅ **Runnable in seconds** — clone, `cd`, `go run`. That's it.
- ✅ **Progressive** — from a 40-line CLI to a graceful-shutdown HTTP server with SSE.

---

## 🚀 Getting Started

### Prerequisites

- [Go **1.26+**](https://go.dev/dl/) installed and on your `PATH`.

Verify your setup:

```bash
go version
```

### Clone the repository

```bash
git clone https://github.com/binafy/go-mini-projects.git
cd go-mini-projects
```

### Run a project

Each project lives in its own numbered folder. `cd` into one and run it:

```bash
cd "1. Text Wrapper"
go run main.go -text="Go is fun to learn by building" -width=15
```

> 💡 Projects that pull external packages (like the QR Code Generator) ship their own `go.mod`. Just `go run` inside the folder — Go resolves the dependencies automatically.

---

## 📦 Projects

Legend:  ✅ Implemented · 🚧 Planned

| #  | Project                                               | What it does                                       | Highlights                                               | Status |
|:--:|:------------------------------------------------------|:---------------------------------------------------|:---------------------------------------------------------|:------:|
| 1  | [**Text Wrapper**](./1.%20Text%20Wrapper)             | Wraps text to a maximum line width                 | `flag` parsing, word-boundary wrapping                   |   ✅   |
| 2  | [**QR Code Generator**](./2.%20QR%20Code%20Generator) | Turns any URL into a QR image (`png`/`jpg`/`webp`) | 3rd-party lib, random filenames, format validation       |   ✅   |
| 3  | Web Scraper                                           | Extract data from web pages                        | —                                                        |   ✅   |
| 4  | Credit Validator                                      | Validate card numbers via the Luhn algorithm       | —                                                        |   🚧   |
| 5  | URL Shortener                                         | Shorten and expand URLs                            | —                                                        |   🚧   |
| 6  | Empty File Finder                                     | Find zero-byte files in a tree                     | —                                                        |   🚧   |
| 7  | Empty Directory Finder                                | Find empty directories in a tree                   | —                                                        |   🚧   |
| 8  | Password Generator                                    | Generate strong, configurable passwords            | —                                                        |   🚧   |
| 9  | Search String                                         | Grep-like search across files                      | —                                                        |   🚧   |
| 10 | Watermark Image                                       | Overlay a watermark onto images                    | —                                                        |   🚧   |
| 11 | Encrypt / Decrypt Text                                | Symmetric text encryption                          | —                                                        |   🚧   |
| 12 | CLI Todo App                                          | Manage a todo list from the terminal               | —                                                        |   🚧   |
| 13 | XML → JSON                                            | Convert XML documents to JSON                      | —                                                        |   🚧   |
| 14 | Tic Tac Toe                                           | Terminal two-player game                           | —                                                        |   🚧   |
| 15 | String Reverse                                        | Reverse strings (Unicode-aware)                    | —                                                        |   🚧   |
| 16 | [**SSE Server**](./16.%20SSE)                         | Real-time Server-Sent Events over HTTP             | `embed`, `context`, graceful shutdown, live browser demo |   ✅   |
| 17 | WebSocket                                             | Bidirectional real-time messaging                  | —                                                        |   🚧   |

### ⭐ Featured

<details>
<summary><b>1. Text Wrapper</b> — greedy word-wrap in ~40 lines</summary>

<br/>

A tiny CLI that reflows text to a maximum line width without breaking words.

```bash
cd "1. Text Wrapper"
go run main.go -text="Im Milwad Khosravi who makes a lot of tools for developers, enjoy it!" -width=10
```

**Flags**

| Flag | Default | Description |
|------|---------|-------------|
| `-text` | _(required)_ | The text to wrap |
| `-width` | `5` | Maximum characters per line |

</details>

<details>
<summary><b>2. QR Code Generator</b> — URL → scannable image</summary>

<br/>

Generates a QR code from any URL and writes it to disk. Powered by [`skip2/go-qrcode`](https://github.com/skip2/go-qrcode).

```bash
cd "2. QR Code Generator"
go run main.go -url="https://github.com/binafy/go-mini-projects" -filename="repo" -format=png -fileSize=512
```

**Flags**

| Flag | Default | Description |
|------|---------|-------------|
| `-url` | _(required)_ | The URL to encode |
| `-filename` | _(random)_ | Output file name (without extension) |
| `-format` | `png` | Output format: `png`, `jpg`, or `webp` |
| `-fileSize` | `256` | Image size in pixels |

</details>

<details>
<summary><b>16. SSE Server</b> — real-time streaming done right</summary>

<br/>

A production-shaped HTTP server that pushes a live timestamp to the browser every second using **Server-Sent Events**. Demonstrates `//go:embed`, `context`-based cancellation, `http.NewResponseController` flushing, and graceful shutdown on `SIGINT`.

```bash
cd "16. SSE"
go run main.go -addr=":8080" -interval=1s
```

Then open **http://localhost:8080** in your browser and watch events stream in live.

**Flags**

| Flag | Default | Description |
|------|---------|-------------|
| `-addr` | `:8080` | HTTP listen address |
| `-interval` | `1s` | Delay between streamed events |

</details>

---

## 🗺️ Roadmap

The long-term vision is a comprehensive Go learning path spanning **100+ projects** across every level. A curated slice of what's coming:

<table>
<tr>
<td valign="top" width="33%">

**🟢 Beginner**
- Calculator CLI
- Notes Manager
- Base64 Encoder/Decoder
- JSON Formatter
- CSV / YAML Parser
- File Organizer
- Countdown Timer

</td>
<td valign="top" width="33%">

**🔵 Intermediate / Web**
- REST & CRUD API
- JWT Authentication
- API Rate Limiter
- Reverse Proxy
- WebSocket Chat
- Web Crawler
- Pastebin Clone

</td>
<td valign="top" width="33%">

**🔴 Advanced**
- Worker Pool
- Graceful Shutdown
- Key-Value Database
- Message Queue & Pub/Sub
- LRU Cache & Bloom Filter
- Consistent Hashing
- Mini Search Engine

</td>
</tr>
</table>

> See the full 100+ project backlog in [`todo.txt`](./todo.txt).

---

## 🗂️ Repository Structure

```text
go-mini-projects/
├── 1. Text Wrapper/         # ✅ CLI word-wrapper
│   └── main.go
├── 2. QR Code Generator/    # ✅ URL → QR image
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── 16. SSE/                 # ✅ Real-time SSE server
│   ├── main.go
│   └── index.html
├── ...                      # 🚧 Planned projects
├── todo.txt                 # Full roadmap
└── README.md
```

---

## 🤝 Contributing

Contributions are warmly welcome — whether it's a brand-new project, a bug fix, or a docs polish.

1. **Fork** the repository.
2. **Create** a branch: `git checkout -b feat/my-awesome-project`.
3. **Add** your project in its own numbered folder with a self-contained `main.go`.
4. **Keep it idiomatic** — `gofmt` your code, prefer the standard library, document your flags.
5. **Open** a Pull Request describing what you built and how to run it.

> New projects should be runnable with a single `go run` and include a short usage example.

---

## 📄 License

Released under the **MIT License**. See [`LICENSE`](./LICENSE) for details.

---

<div align="center">

**Built with 🐹 and ♥ by [Milwad Khosravi](https://github.com/binafy)**

If this repo helped you learn Go, consider giving it a ⭐

</div>
