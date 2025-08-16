# httpfromtcp

A learning project where I implement a minimal HTTP server in **Go**, starting from raw **TCP sockets** instead of relying on Go’s built-in `net/http` package.
The goal is to deeply understand how HTTP works at the protocol level and how data flows between a client and a server over the network.

---

## 🚀 Motivation

Most Go developers rely on `net/http` for building servers, but that abstracts away a lot of the lower-level details. Special thanks to boot.dev as this project is a follow along to their course Learn HTTP Protocol in Go.
This project is my attempt to peel back the layers and learn:

- How TCP connections are established and handled
- How HTTP requests are structured and parsed
- How to construct and send valid HTTP responses
- How concurrency plays a role in handling multiple client requests

---

## 📚 Learning Objectives

- Understand the **TCP protocol** in Go
- Parse **HTTP request lines, headers, and bodies**
- Write minimal but correct **HTTP response messages**
- Experiment with **concurrency (goroutines)** for handling multiple clients
- Build a deeper intuition for how modern web servers work under the hood

---
