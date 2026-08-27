# Concurrent Web Crawler

A high-performance concurrent web crawler built in Go that crawls websites, extracts links, tracks visited pages, and supports configurable crawl depth.

This project was built to demonstrate Go concurrency concepts including Goroutines, Channels, Worker Pools, WaitGroups, Synchronization, and Concurrent Data Structures.

---

## Features

- HTTP Page Fetching
- HTML Parsing
- Link Extraction
- Relative URL Resolution
- Recursive Crawling
- Depth-Limited Crawling
- Visited URL Tracking
- Thread-Safe URL Storage
- Worker Pool Architecture (In Progress)
- Concurrent Crawling (In Progress)
- Rate Limiting (Upcoming)
- JSON Export (Upcoming)
- CSV Export (Upcoming)
- Graceful Shutdown (Upcoming)

---

## Project Structure

```text
Concurrent-WebCrawler/

├── cmd/
│   └── crawler/
│       └── main.go
│
├── internal/
│   ├── crawler/
│   │   ├── crawler.go
│   │   ├── job.go
│   │   └── visited.go
│   │
│   ├── fetcher/
│   │   └── fetcher.go
│   │
│   ├── parser/
│   │   └── parser.go
│   │
│   ├── orchestrator/
│   │   └── worker.go
│   │
│   ├── limiter/
│   │   └── ratelimit.go
│   │
│   └── storage/
│       ├── json.go
│       └── csv.go
│
├── output/
│
├── go.mod
├── go.sum
└── README.md
```

---

## Technologies Used

### Go Standard Library

- net/http
- context
- sync
- time
- flag
- net/url
- encoding/json
- encoding/csv

### External Package

```go
golang.org/x/net/html
```

Used for HTML parsing and link extraction.

---

## How It Works

### Step 1

Fetch HTML

```text
URL
 ↓
HTTP Request
 ↓
HTML Response
```

### Step 2

Parse HTML

```text
HTML
 ↓
Find Anchor Tags
 ↓
Extract href Values
```

### Step 3

Resolve URLs

```text
/about
 ↓
https://example.com/about
```

### Step 4

Crawl Recursively

```text
Page A
 ↓
Page B
 ↓
Page C
```

### Step 5

Track Visited URLs

Prevents duplicate crawls and infinite loops.

---

## Current Architecture

```text
                Seed URL
                    |
                    v

              +-----------+
              |  Crawler  |
              +-----------+
                    |
                    v

               Fetch Page
                    |
                    v

               Parse HTML
                    |
                    v

              Extract Links
                    |
                    v

              Crawl Again
```

---

## Running The Project

### Clone Project

```bash
git clone https://github.com/yourusername/Concurrent-WebCrawler.git
```

### Move Into Project

```bash
cd Concurrent-WebCrawler
```

### Install Dependencies

```bash
go mod tidy
```

### Run

```bash
go run cmd/crawler/main.go -url https://example.com -depth 2
```

---

## Example Output

```text
[Depth 0] Crawling: https://example.com
[Depth 0] Found 1 links

[Depth 1] Crawling: https://iana.org/domains/example
[Depth 1] Found 31 links
```

---

## Supported Flags

### URL

```bash
-url https://example.com
```

Starting URL.

### Depth

```bash
-depth 2
```

Maximum crawl depth.

### Timeout

```bash
-timeout 10s
```

HTTP request timeout.


---

## Author

Kashish Srivastava
