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

| #    | Project                                                           | What it does                                         | Highlights                                                 | Status   |
|:----:|:------------------------------------------------------------------|:-----------------------------------------------------|:-----------------------------------------------------------|:--------:|
| 1    | [**Text Wrapper**](./1.%20Text%20Wrapper)                         | Wraps text to a maximum line width                   | `flag` parsing, word-boundary wrapping                     |    ✅     |
| 2    | [**QR Code Generator**](./2.%20QR%20Code%20Generator)             | Turns any URL into a QR image (`png`/`jpg`/`webp`)   | 3rd-party lib, random filenames, format validation         |    ✅     |
| 3    | [**Web Scraper**](./3.%20Web%20Scraper)                           | Scrape GitHub profiles from a browser UI             | `gocolly/colly`, `html/template`, CSS selectors            |    ✅     |
| 4    | [**Credit Validator**](./4.%20Credit%20Validator)                 | Validate card numbers via the Luhn algorithm         | Luhn checksum, network detection, stdin batching           |    ✅     |
| 5    | [**URL Shortener**](./5.%20URL%20Shorter)                         | Shorten URLs and expand them back                    | `crypto/rand` IDs, JSON store, atomic file writes          |    ✅     |
| 6    | [**Empty File Finder**](./6.%20Empty%20File%20Finder)             | Find zero-byte files in a tree                       | `filepath.WalkDir`, resilient error handling               |    ✅     |
| 7    | [**Empty Directory Finder**](./7.%20Empty%20Directory%20Finder)   | Find empty directories in a tree                     | Recursion, transitively-empty detection                    |    ✅     |
| 8    | [**Password Generator**](./8.%20Password%20Generator)             | Generate strong, configurable passwords              | `crypto/rand`, guaranteed character classes                |    ✅     |
| 9    | [**Search String**](./9.%20Search%20String)                       | Count word matches in a text and show their offsets  | `regexp`, word boundaries, case-insensitive matching       |    ✅     |
| 10   | [**Watermark Image**](./10.%20Watermark%20Image)                  | Overlay a transparent watermark onto a photo         | `image/draw`, alpha compositing, JPEG/PNG codecs           |    ✅     |
| 11   | [**Encrypt / Decrypt Text**](./11.%20Encrypt%20Decrypt%20Text)     | Encrypt text, then decrypt it back                   | AES-256-GCM, random nonce, base64 output                   |    ✅     |
| 12   | CLI Todo App                                                      | Manage a todo list from the terminal                 | —                                                          |    🚧     |
| 13   | XML → JSON                                                        | Convert XML documents to JSON                        | —                                                          |    🚧     |
| 14   | Tic Tac Toe                                                       | Terminal two-player game                             | —                                                          |    🚧     |
| 15   | [**String Reverse**](./15.%20String%20Reverse)                    | Reverse text, whole lines or word order              | Unicode-aware (combining marks), stdin piping, `bufio`     |    ✅     |
| 16   | [**SSE Server**](./16.%20SSE)                                     | Real-time Server-Sent Events over HTTP               | `embed`, `context`, graceful shutdown, live browser demo   |    ✅     |
| 17   | WebSocket                                                         | Bidirectional real-time messaging                    | —                                                          |    🚧     |

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
<summary><b>3. Web Scraper</b> — a GitHub profile, scraped and rendered</summary>

<br/>

A small web app rather than a CLI. It serves a page on port `8080`, and when you submit a username it fetches that profile from github.com with [`gocolly/colly`](https://github.com/gocolly/colly), pulls the fields out with CSS selectors, and renders them back through `html/template`.

```bash
cd "3. Web Scraper"
go run .
```

Then open **http://localhost:8080** and enter a username. The form accepts a bare handle, an `@handle`, or a full profile URL — the page normalises it to `https://github.com/<user>` before submitting.

**Scraped fields**

| Field | Selector it comes from |
|-------|------------------------|
| Title | `<title>` |
| Name | `span.p-name` |
| Bio | `div.p-note` |
| Location | `span.p-label`, `span[itemprop='location']` |
| Followers / Following | the profile's `tab=followers` and `tab=following` links |
| Repositories | `tab=repositories` counter, falling back to the `meta[name='description']` text |
| Avatar | `img[itemprop='image']`, resolved to an absolute URL |

Counts are parsed loosely, so GitHub's `1.2k` shorthand and comma-grouped numbers both come back as plain integers. A missing field is labelled rather than left blank, so it is obvious when a selector has gone stale — GitHub changes its markup often, and keeping up with it is really the point of the exercise.

> Run it from inside the project directory: `index.html` is parsed at startup by relative path.

</details>

<details>
<summary><b>4. Credit Validator</b> — Luhn, in a single pass over the digits</summary>

<br/>

Checks card numbers with the **Luhn algorithm** — the checksum every card in your wallet satisfies — and names the network the number was issued by. Spaces and dashes are stripped first, so a number can be pasted in exactly the way it is printed on the card.

```bash
cd "4. Credit Validator"
go run main.go -number="4539 1488 0343 6467"
go run main.go 4111111111111111 378282246310005
cat cards.txt | go run main.go
```

```text
'4539 1488 0343 6467' is a valid Visa card number.
'378282246310005' is a valid American Express card number.
'4111111111111112' is not a valid card number.

Checked 3 card number(s): 2 valid, 1 invalid.
```

Numbers can arrive three ways — the `-number` flag, arguments, or one per line on stdin — so a whole file of them can be checked in one run. Visa, Mastercard, American Express, Discover, JCB, Diners Club and UnionPay are recognised by their leading digits; anything else still gets validated, just without a network name.

**Flags**

| Flag | Default | Description |
|------|---------|-------------|
| `-number` | _(arguments, then stdin)_ | The card number to validate |

> A number passing the Luhn check only means it is well-formed — it says nothing about whether the account exists or has funds.

</details>

<details>
<summary><b>6. Empty File Finder</b> — zero-byte files across a whole tree</summary>

<br/>

Walks a directory tree and reports every zero-byte file it finds. Sizes come from the directory entry rather than reading each file, so a huge tree never gets pulled through memory, and a directory it cannot open is reported and skipped instead of ending the scan.

```bash
cd "6. Empty File Finder"
go run main.go -d testdata
```

```text
The 'testdata/empty.txt' file is empty!
The 'testdata/milwad/empty2.txt' file is empty!

Found 2 empty file(s) in 'testdata'.
```

**Flags**

| Flag | Default | Description |
|------|---------|-------------|
| `-d` | `.` | The directory to search (scanned recursively) |

</details>

<details>
<summary><b>7. Empty Directory Finder</b> — empty folders, including the ones hiding empties</summary>

<br/>

Walks a tree and reports directories with nothing in them. With `-nested` it also reports directories that hold nothing but other empty directories — the ones worth pruning even though they are not literally empty. Unreadable directories are reported and skipped, and never counted as empty.

```bash
cd "7. Empty Directory Finder"
go run main.go -d testdata
go run main.go -d testdata -nested
```

Given `outer/middle/inner` where only `inner` is literally empty, the default run reports just `inner`, while `-nested` reports `inner`, `middle` and `outer`.

**Flags**

| Flag | Default | Description |
|------|---------|-------------|
| `-d` | `.` | The directory to search (scanned recursively) |
| `-nested` | `false` | Also report directories holding nothing but empty directories |

> ⚠️ Git cannot track empty directories, so `testdata/emptyFolder` has to be created locally: `mkdir -p "7. Empty Directory Finder/testdata/emptyFolder"`.

</details>

<details>
<summary><b>8. Password Generator</b> — random you can actually rely on</summary>

<br/>

Generates passwords from `crypto/rand` rather than `math/rand`, so the output is not predictable from previous passwords. Each enabled character class is guaranteed to appear at least once, and the result is shuffled so the class order never leaks into the layout.

```bash
cd "8. Password Generator"
go run main.go
go run main.go -count=5 -length=24
go run main.go -length=20 -symbols=false
```

**Flags**

| Flag | Default | Description |
|------|---------|-------------|
| `-length` | `16` | Length of each password |
| `-count` | `1` | How many passwords to generate |
| `-lower` | `true` | Include lowercase letters |
| `-upper` | `true` | Include uppercase letters |
| `-digits` | `true` | Include digits |
| `-symbols` | `true` | Include symbols |

Disable a class with `=false`, e.g. `-symbols=false`. Asking for a length smaller than the number of enabled classes is rejected, since such a password could not contain them all.

</details>

<details>
<summary><b>9. Search String</b> — every match, with its exact offset</summary>

<br/>

Searches a built-in paragraph about Go for a word and reports where each occurrence sits. The search term is escaped with `regexp.QuoteMeta` and wrapped in `\b` word boundaries, so it is matched as a whole word and never as a fragment of a longer one — searching for `go` finds `Go` and `Go`, but not the `go` inside `Google`. Matching is case-insensitive.

```bash
cd "9. Search String"
go run main.go -text="go"
```

```text
0 Index: [22:24]
1 Index: [136:138]
Text: go
Repeat count: 2
```

Each line is one match, printed as the half-open byte range `[start:end)` it occupies in the text — the same slice bounds you would use to cut it back out with `description[start:end]`.

**Flags**

| Flag | Default | Description |
|------|---------|-------------|
| `-text` | _(required)_ | The word to search for |

> The text being searched is the `description` constant in `main.go`. Edit it to search something else.

</details>

<details>
<summary><b>10. Watermark Image</b> — compositing two images with the standard library</summary>

<br/>

Stamps a transparent PNG watermark onto a JPEG photo using nothing but `image`, `image/draw` and the standard codecs — no third-party imaging library. The photo is decoded and copied into a writable `*image.RGBA` with `draw.Src`, then the watermark is composited on top with `draw.Over`, so the PNG's alpha channel is respected and whatever is transparent in the watermark stays see-through.

```bash
cd "10. Watermark Image"
go run main.go
```

The filenames are fixed: it reads `original.jpg` and `watermark.png` from the working directory and writes `watermarked.jpg` next to them, re-encoded as JPEG at quality 90. All three ship with the project, so a plain `go run main.go` produces a result immediately.

**Files**

| File | Role |
|------|------|
| `original.jpg` | The base photo to stamp |
| `watermark.png` | The overlay — transparency comes from its alpha channel |
| `watermarked.jpg` | The output, JPEG quality 90 (overwritten on each run) |

The watermark is placed by subtracting its size from the photo's, leaving a 20px margin — so with a watermark smaller than the photo it lands neatly in the **bottom-right corner**.

> ⚠️ There is no scaling step. The bundled `watermark.png` is 1536×1024 while `original.jpg` is 1200×675, so the computed corner position is negative and `draw.Draw` clips the overlay against the photo's bounds — the watermark ends up spread across the whole frame with its edges cut off, rather than tucked into the corner. Swap in a watermark smaller than your photo to see the intended corner placement.

</details>

<details>
<summary><b>11. Encrypt / Decrypt Text</b> — text in, ciphertext and back again</summary>

<br/>

Encrypts text with **AES-256-GCM**, prints it base64 encoded, and decrypts it straight back — both directions in one run. Every run uses a fresh random nonce, so the same text never encrypts to the same output twice.

```bash
cd "11. Encrypt Decrypt Text"
go run main.go -text="meet me at noon"
```

```text
Encrypted: 02VpL3MKu5ub7POYh40y6xyLRErDAzKPnkQZwbDXW5frqqWX0v8Z0LGXUA==
Decrypted: meet me at noon
```

**Flags**

| Flag | Default | Description |
|------|---------|-------------|
| `-text` | _(required)_ | The text to encrypt |

The key is a constant in `main.go`, which keeps the demo to a single flag — swap it for a real secret before encrypting anything you care about.

</details>

<details>
<summary><b>15. String Reverse</b> — reversing text without mangling Unicode</summary>

<br/>

Reverses text from a flag, from arguments, or straight from a pipe. Naive `[]rune` reversal detaches accents from their letters, so characters are reversed together with the combining marks that follow them — `café` comes back as `éfac`, not with a floating accent.

```bash
cd "15. String Reverse"
go run main.go -text="Milwad Khosravi"
go run main.go -words -text="one two three"
cat notes.txt | go run main.go
```

Input is read line by line, so piped files keep their line structure — each line is reversed on its own.

**Flags**

| Flag | Default | Description |
|------|---------|-------------|
| `-text` | _(stdin)_ | The text to reverse; falls back to arguments, then stdin |
| `-words` | `false` | Reverse the order of the words instead of the characters |

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
├── 6. Empty File Finder/    # ✅ Recursive zero-byte file finder
│   ├── main.go
│   └── testdata/
├── 7. Empty Directory Finder/   # ✅ Recursive empty-directory finder
│   ├── main.go
│   └── testdata/
├── 8. Password Generator/   # ✅ crypto/rand password generator
│   └── main.go
├── 9. Search String/        # ✅ Word search with match offsets
│   └── main.go
├── 10. Watermark Image/     # ✅ PNG watermark composited onto a JPEG
│   ├── main.go
│   ├── original.jpg
│   └── watermark.png
├── 11. Encrypt Decrypt Text/    # ✅ AES-256-GCM text encryption
│   └── main.go
├── 15. String Reverse/      # ✅ Unicode-aware string reverser
│   └── main.go
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
