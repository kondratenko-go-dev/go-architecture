# go-architecture

A learning Go project designed to practice basic application architecture:

- models (`Person`, `Address`)
- interfaces (`Greeter`, `PersonStorage`)
- storage layer (`MemoryPersonStorage`)
- service layer (`PersonService`)
- a simple CLI application that works through the service and interfaces

The goal of the project is to gradually move toward an architecture similar to real web services and bots (Storage / Service / Interfaces), and to gain experience writing “proper” Go code.

---

## 🔧 Requirements

- Go 1.20+ (preferably the latest Go version)
- Any terminal (Linux / WSL / macOS / etc.)

## 🚀 Run

In the project root:

```bash
go run .