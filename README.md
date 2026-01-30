# arc-blog

Blog and article operations for the Arc toolkit.

## Features

- Fetch and save articles from URLs
- Import RSS/Atom feeds
- Analyze article content with AI

## Installation

```bash
go install github.com/mtreilly/arc-blog@latest
```

## Usage

```bash
# Fetch a single article
arc-blog fetch --url https://example.com/post

# Pull a feed into your research directory
arc-blog fetch --playlist feed.xml --out-dir docs/research-external/blog

# Fetch and analyze an article
arc-blog fetch --url https://example.com/post --analyze --output json
```

## License

MIT
