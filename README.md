# Commurz
Commurz is a simple e-commerce system to demonstrate a DDD & Functional Core Imparative Shell approach to build a system.

This system will have 2 distinct modules:
- frontoffice
    - This module will be used by the customers.
- backoffice
    - This module will be used by the internal teams to manage the system.

## Table of Contents
- [Usecases](docs/Usecase.md)
- [Tech Stack]()

## Tech Stack
- Golang
- PostgreSQL
    - SQLite for now
- SQLC used in repository layer
- [connect.build](https://connect.build/) for service layer
    - this serve as imparative shell in Functional Core, Imperative Shell concept
    or as a application layer in Hexagonal Architecture
    - it is the public APIs of the system, it can serve as gRPC or JSON-RPC server without too much configs.
    - the protobuf can be used as SDK to build a user interface SPA, CLI, Desktop, Mobile
    - the serivce methods can be used directly to build an SSR App

### TODO
- [ ] SSR App
- [ ] add monitoring & tracing