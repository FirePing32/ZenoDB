# ZenoDB

A lightweight, disk-based NoSQL key-value store implemented in Go.

## Overview

This project implements a basic NoSQL key-value database from scratch in Go. It's designed to be simple and minimal, focusing on the core concepts of database implementation. The database uses a B-tree as its underlying data structure and implements disk-based storage with a page size of 4KB.

## Features

- Disk-based storage
- 4KB page size
- Basic CRUD operations
- Persistence between restarts
- Freelist management for efficient space utilization

## Components

1. **Database**: Manages the program and is responsible for orchestrating transactions.
2. **Data Access Layer (DAL)**: Handles all disk operations and data organization on the disk.
3. **Freelist**: Manages free and occupied pages to avoid fragmentation.
4. **Meta Page**: Stores metadata about the database, including the freelist page number.

## Getting Started

### Prerequisites

- `Go 1.19`

### Usage

```bash
go run main.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).

