# Cartographer 🗺️

> Auto-generated living documentation and dependency graphs for your codebase

## Overview
Cartographer is a lightweight internal developer portal that parses your codebase to auto-generate API contracts, interactive dependency graphs, and real-time CI/CD health dashboards. Think Backstage — but instant.

## Features
- 📄 AST-based code parser that extracts API contracts automatically
- 🕸️ Interactive dependency graph with D3.js force simulation
- 🚦 CI/CD health aggregator across GitHub Actions, deploys & alerts
- 🌡️ Real-time service map with latency heatmap overlay
- 🔍 Full-text search across all service documentation

## Architecture
```
GitHub Repo → Go Parser (AST) → GraphQL API → React Dashboard
                              ↓
                         D3.js Graph Engine
```

## Tech Stack
| Layer | Technology |
|-------|-----------|
| Frontend | React, D3.js, TypeScript |
| Backend | Go, GraphQL (gqlgen) |
| Data | PostgreSQL, Redis |
| Integration | GitHub API, Kubernetes API |

## Getting Started
```bash
git clone https://github.com/yankah-julian/cartographer.git
cd cartographer

# Start backend
cd server
go mod tidy
go run ./cmd/server

# Start frontend
cd ../web
npm install
npm run dev
```

## Project Structure
```
cartographer/
├── server/
│   ├── cmd/server/      # Go entry point
│   ├── parser/          # AST code analysis
│   ├── graph/           # GraphQL schema + resolvers
│   └── integrations/    # GitHub, K8s clients
├── web/
│   ├── src/
│   │   ├── components/
│   │   │   ├── DependencyGraph.tsx
│   │   │   └── ServiceMap.tsx
│   │   └── pages/
└── docker-compose.yml
```

## License
MIT © Julian Yankah
