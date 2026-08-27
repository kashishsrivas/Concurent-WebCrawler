# Concurrent Web Crawler

A high-performance concurrent web crawler built in Go that crawls websites, extracts links, tracks visited pages, exports crawl results, and supports configurable crawl depth.

---

## Features

- HTTP Page Fetching
- HTML Parsing
- Link Extraction
- Relative URL Resolution
- Duplicate URL Removal
- Link Filtering (`mailto`, `tel`, `javascript`, fragments)
- Depth-Limited Crawling
- Concurrent Crawling
- Worker Pool Architecture
- Thread-Safe URL Tracking
- Rate Limiting
- Crawl Statistics
- JSON Export
- CSV Export
- Graceful Shutdown (Ctrl + C)

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
│       ├── model.go
│       ├── json.go
│       └── csv.go
│
├── output/
│   ├── crawl.json
│   └── crawl.csv
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
- os
- os/signal
- syscall
- encoding/json
- encoding/csv

### External Package

```go
golang.org/x/net/html
```

Used for HTML parsing and link extraction.

---

## How It Works

### Step 1 - Fetch HTML

```text
URL
 ↓
HTTP Request
 ↓
HTML Response
```

### Step 2 - Parse HTML

```text
HTML
 ↓
Find Anchor Tags
 ↓
Extract href Values
```

### Step 3 - Normalize URLs

```text
/about
 ↓
https://example.com/about
```

### Step 4 - Filter Unwanted Links

```text
javascript:void(0)
mailto:test@example.com
#section
tel:+911234567890

        ↓

Ignored
```

### Step 5 - Create Crawl Jobs

```text
Extracted Links
        ↓
Generate Jobs
        ↓
Push To Queue
```

### Step 6 - Concurrent Processing

```text
Multiple Workers
       ↓
Fetch Pages
       ↓
Parse Content
       ↓
Extract Links
```

### Step 7 - Store Results

```text
Processed URLs
      ↓
Store Results
      ↓
Export JSON / CSV
```

---

## System Architecture

```text
                         Seed URL
                             │
                             ▼

                    ┌─────────────────┐
                    │     main.go     │
                    └─────────────────┘
                             │
                             ▼

                    ┌─────────────────┐
                    │   Worker Pool   │
                    └─────────────────┘
                             │
                             ▼

                    ┌─────────────────┐
                    │   Jobs Channel  │
                    └─────────────────┘
                             │
      ┌──────────────────────┼──────────────────────┐
      │                      │                      │
      ▼                      ▼                      ▼

┌─────────────┐      ┌─────────────┐      ┌─────────────┐
│  Worker 1   │      │  Worker 2   │      │  Worker N   │
└─────────────┘      └─────────────┘      └─────────────┘
      │                      │                      │
      └──────────────┬───────┴──────────────┬───────┘
                     ▼                      ▼

                 Fetch Page          Rate Limiter
                     │
                     ▼

                 Parse HTML
                     │
                     ▼

                Extract Links
                     │
                     ▼

              Generate New Jobs
                     │
                     ▼

                Store Results
                     │
                     ▼

          JSON Export / CSV Export
```

---

## Concurrency Workflow

```text
Seed URL
    │
    ▼

Jobs Channel
    │
    ▼

┌──────────┐
│ Worker 1 │
└──────────┘

┌──────────┐
│ Worker 2 │
└──────────┘

┌──────────┐
│ Worker 3 │
└──────────┘

┌──────────┐
│ Worker 4 │
└──────────┘

┌──────────┐
│ Worker 5 │
└──────────┘

    │
    ▼

Fetch → Parse → Extract → Queue New Jobs
```

---

## Running The Project

### Clone Repository

```bash
git clone https://github.com/<your-username>/Concurrent-WebCrawler.git
```

### Move Into Project Directory

```bash
cd Concurrent-WebCrawler
```

### Install Dependencies

```bash
go mod tidy
```

### Run the Crawler

```bash
go run cmd/crawler/main.go \
-url https://example.com \
-depth 2 \
-workers 5
```

---

## Example Output

```text
Worker 1 started
Worker 2 started
Worker 3 started
Worker 4 started
Worker 5 started

[Worker] Crawling: https://example.com (Depth 0)
[Depth 0] Found 1 links

[Worker] Crawling: https://iana.org/domains/example (Depth 1)
[Depth 1] Found 26 links

[Worker] Crawling: https://iana.org/about (Depth 2)
[Depth 2] Found 53 links
```

---

## Crawl Summary

```text
================================
Pages Crawled : 28
Workers Used  : 5
Max Depth     : 2
================================
```

---

## Generated Output

### JSON Export

```json
[
  {
    "url": "https://example.com",
    "depth": 0
  },
  {
    "url": "https://iana.org/domains/example",
    "depth": 1
  }
]
```

### CSV Export

```csv
url,depth
https://example.com,0
https://iana.org/domains/example,1
```

---

## Supported Flags

### URL

```bash
-url https://example.com
```

Starting URL for crawling.

### Depth

```bash
-depth 2
```

Maximum crawl depth.

### Workers

```bash
-workers 5
```

Number of concurrent worker goroutines.

### Timeout

```bash
-timeout 10s
```

HTTP request timeout.

---

## Author

**Kashish Srivastava**  
