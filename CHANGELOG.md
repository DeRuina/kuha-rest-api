# Changelog

## 1.0.0 (2025-05-28)


### Features

* relase please script ([a4cbb0b](https://github.com/DeRuina/KUHA-REST-API/commit/a4cbb0b9044f3d9bc51c1b4bffdcb03296de0c17))

## 2025-05-28

### Changed
- CI updated due to Ubuntu 20.04 retirement
- Audit pipeline adjustments for CI

---

## 2025-05-26

### Added
- Redis and DB health check integration in `/health` endpoint
- Uptime and goroutine count now returned in health checks
- Server metrics endpoint added

### Changed
- Improved health check documentation

---

## 2025-05-23

### Fixed
- Resolved multiple CORS issues including credential support and origin validation

### Changed
- Finalized CORS configuration for production

---

## 2025-05-19

### Added
- Sliding window rate limiting using Redis
- Fallback to fixed window rate limiting as default

---

## 2025-05-15

### Added
- Graceful shutdown support for server on SIGINT/SIGTERM

### Changed
- Replaced `context.Background()` with `r.Context()` for request-scoped context

---

## 2025-05-14

### Added
- Redis caching for endpoint logic
- Helper function for setting cached JSON
- Manual caching setup for endpoints

---

## 2025-05-12

### Added
- Base Redis client and caching infrastructure
- Redis-backed storage and limiter initialized

---

## 2025-05-09

### Added
- Logging system with rotation using Timberjack
- Custom rotation interval setup

### Changed
- Temporary log rotation frequency for backup testing

---

## 2025-05-05

### Changed
- Log file naming conventions updated

---

## 2025-05-02

### Added
- `LOG_DIR` environment variable support
- Time-based log rotation logic

### Changed
- Module cleanup with `go mod tidy`

---

## 2025-04-25

### Added
- Structured logging output to stdout and files
- Logger middleware capturing detailed request metadata

---

## 2025-04-23

### Added
- Full RBAC-based authorization and seeding
- Centralized environment variable management
- Swagger documentation for auth and token endpoints
- Health and data endpoints for UTV (Oura, Polar, Suunto)
- SQL query handling with sqlc
- Error handling, validation, and standardized response structures
