# FluxScribe 📝⚡

An ephemeral data relay system built with Go. FluxScribe allows you to share text or code snippets via a unique link that automatically self-destructs after a set expiration time.

Built with professional Go practices, clean architecture, and a focus on backend engineering excellence.

## 🎯 Purpose

This project serves as a deep dive into:
- **Backend Engineering**: Robust API design with `net/http`.
- **Cloud-Native Practices**: Containerization, observability, and CI/CD.
- **Professional Go**: Clean architecture, structured logging, and type-safe templates.

It is not just a web app; it is a demonstration of system design principles required for Platform Engineering and SRE roles.

## 🛠️ Tech Stack

| Component | Technology | Why? |
|-----------|------------|------|
| **Language** | Go (Golang) | Concurrency, performance, single binary deployment. |
| **Database** | PostgreSQL | Reliability, ACID compliance, JSONB support. |
| **Frontend** | Go Templ + HTMX | Type-safe templates, dynamic interactions without heavy JS. |
| **Styling** | CSS | Custom, lightweight styling (no Tailwind). |
| **Embedding** | `go:embed` | Single binary deployment (templates + static assets). |
| **Logging** | `log/slog` | Structured, machine-readable logs. |
| **Infrastructure** | Docker + Makefile | Reproducible builds and automation. |

## 🏗️ Project Structure (Clean Architecture)

```text
fluxscribe/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/                # Private application code
│   ├── config/              # Environment configuration
│   ├── database/            # DB connections & migrations
│   ├── handlers/            # HTTP request handlers
│   ├── models/              # Data structures
│   ├── router/              # URL routing
│   └── services/            # Business logic
├── pkg/                     # Public libraries (future)
├── ui/                      # Frontend assets
│   ├── static/              # CSS, images (embedded)
│   └── templates/           # Go Templ files (embedded)
├── migrations/              # SQL migration files
├── Makefile                 # Automation shortcuts
├── go.mod                   # Dependencies
└── README.md
