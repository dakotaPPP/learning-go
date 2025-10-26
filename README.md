# My journey of learning Go

This is a mono repo showcasing projects I'll be working on to learn the language go

I'll try to keep a devlog in each individual repo (at the very least put comments in commits showcasing current status and what I was thinking)

And at the end of each project I'll write a reflection about what I learned, what I'm most proud of, what I struggled the most with, and what AI suggests I should've done different

## Limiting my resources

- **Note: I disabled copilot inline completions and will refrain from using AI to generate code or explain step by step what I must do**
- On top of this I will be including urls to share the chats I have with AI about why something isn't working or not
- However I can't share every instance of running into AI as google has their own AI that I will use whenever looking up a topic
- My editor is NeoVim preconfigured with LazyVim, so I do have access to the Lazy Extra of lang.go
  - This gives me access to some snippets, go's LSP, linting, and automatic go formatting whenever I save a file

## Projects to Complete

I utilized Gemini to come up with projects I should do in order to learn go (you can find that in [Projects To Do](projects-to-do.md))

Obviously after I complete all these projects I won't be a master in Go but I highly believe that this will make me both comfortable and confident with coding in Go

### 1. Command-Line Utility (CLI) Tool - Completed

A simple **file rename utility** or a **text-file word/line counter** (like a simplified `wc` command).

- For this I chose to go with recreating the `wc` command to the best of my abilities

### 2. Simple Key-Value Store (In-Memory)

A simple **in-memory database** with commands like `SET <key> <value>`, `GET <key>`, and `DELETE <key>`.

### 3. Concurrent Web Scraper / Poller

A program that reads a list of URLs from a file and **fetches the HTTP status code** for each URL concurrently, reporting the results.

### 4. Basic REST API Server

An API that exposes the key-value store from Project 2 over HTTP. Endpoints might include: `POST /data` (to set a key/value) and `GET /data/{key}` (to retrieve a value).

---

## Projects I've Completed

- 1. Command-Line Utility (CLI) Tool - Recreating `wc`

## Beginning Stage

- You'll notice at the top of the repo there's the folders `hello-world` and `lesson-2`
  - hello-world was created whilst I was following along [this video](https://youtu.be/XVNvXZyU4aE?si=AAaAPq0T7LN-aT5p) by John McBride
  - And lesson-2 was created whilst I was following along [this video](https://youtu.be/q-GZ71nRe0o?si=P1PMtC9B0Ud2L2un) by John McBride
