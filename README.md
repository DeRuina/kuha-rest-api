# KUHA REST API [![Go Report Card](https://goreportcard.com/badge/github.com/DeRuina/kuha-rest-api)](https://goreportcard.com/report/github.com/DeRuina/kuha-rest-api) ![Audit](https://github.com/DeRuina/kuha-rest-api/actions/workflows/audit.yaml/badge.svg) ![Version](https://img.shields.io/github/v/tag/DeRuina/kuha-rest-api?sort=semver)

Centralized API for the [KUHA program](https://cemis.fi/kainuun-urheilun-ja-hyvinvoinnin-data-analytiikan-ohjelma-kuha-3) (sports and wellbeing data ecosystem). The service unifies data from multiple provider databases into a single access point.

## Context

- Program: KUHA (Kainuu Program for Sports and Well-being Data Analytics) in collaboration with CSC, University of Jyväskylä, KAMK, Finnish Olympic Committee, and others.
- Mission: Collect, integrate, and expose sports and wellbeing data through one API to enable data-driven coaching, research, and visualizations.
- Hosting: Provider databases are hosted in CSC infrastructure; this API brokers access and unifies schemas (“one API to rule them all”).

## Data domains supported

- Competition data (FIS results, rankings, events)
- Wearables and athlete monitoring (Garmin, Oura, Suunto, Polar via UTV)
- Training diaries and questionnaires (e.g., Tietoevry 360, KAMK forms)
- Performance testing (K-Lab, Coachtech/Vuokatti Sport test center)
- Motion analytics (Archinisis)
- Auth service for issuing access and refresh tokens

## Project layout

- `cmd/api`: HTTP server, routing, middleware, and handlers for each data domain/provider.
- `cmd/migrate`: SQL migrations and seeding entrypoint.
- `internal`: shared packages (DB connections, auth, caching, logging, rate limiting, stores).
- `docs`: Swagger definitions and generated artifacts.
