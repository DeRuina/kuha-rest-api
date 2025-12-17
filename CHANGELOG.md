# Changelog

## [1.3.0](https://github.com/DeRuina/kuha-rest-api/compare/v1.2.1...v1.3.0) (2025-12-17)


### Features

* Archinisis GET endpoints ([e05e5b9](https://github.com/DeRuina/kuha-rest-api/commit/e05e5b9af75dd8bd40682631898684db88e0b39f))
* archinisis GET sportti_ids endpoint ([683f519](https://github.com/DeRuina/kuha-rest-api/commit/683f519f24f9d649859ffee1b5da409d2208d674))
* Archinisis POST & GET data endpoints ([9b09972](https://github.com/DeRuina/kuha-rest-api/commit/9b09972eaa148c05b984f3452d3382ac514c79d8))
* archinisis sqlc ([754a3e6](https://github.com/DeRuina/kuha-rest-api/commit/754a3e6c54360a9176e8d11dd9c4f1db8975d44a))
* deleted-users tietoevry GET endpoint ([98c590e](https://github.com/DeRuina/kuha-rest-api/commit/98c590eb80a103aabbfb6a29821f57b5d23b9be0))
* garmin endpoints ([60bdcc5](https://github.com/DeRuina/kuha-rest-api/commit/60bdcc542e4edf4381bcb3c9f697e4ec47d0d528))
* GET activity_zones tietoevry endpoint ([ded46f4](https://github.com/DeRuina/kuha-rest-api/commit/ded46f48d2eed2e4591a0a58e1e31319325d85e3))
* GET customers klab endpoint ([befbb8d](https://github.com/DeRuina/kuha-rest-api/commit/befbb8d67422481068d9137267d81da626896062))
* GET exercises endpoint ([2f24d46](https://github.com/DeRuina/kuha-rest-api/commit/2f24d46d00f2fb6ee15531e3bcda7a7ae442749c))
* GET linked devices utv endpoint ([cbab932](https://github.com/DeRuina/kuha-rest-api/commit/cbab9321b4feacb0ab17a12c9e9732518d2a85f9))
* GET questionnaires tietoevry endpoint ([05db2c5](https://github.com/DeRuina/kuha-rest-api/commit/05db2c506f40814c7861363130b07684d84224b0))
* GET sport_id from utv klab and archinisis ([095d5bf](https://github.com/DeRuina/kuha-rest-api/commit/095d5bfbeceb356559c9e92f8047bf13bdbe8284))
* GET symptoms tietoevry endpoint ([e44f2f3](https://github.com/DeRuina/kuha-rest-api/commit/e44f2f31035808050c56ebbd4b3e5b9e2193817c))
* GET test-results tietoevry endpoint ([e23ac6c](https://github.com/DeRuina/kuha-rest-api/commit/e23ac6c4c9487cf2fae54bbfc9a2ed923976e231))
* GET user tietoevry endpoint ([8ff9fca](https://github.com/DeRuina/kuha-rest-api/commit/8ff9fcaa48edf1e0c1f4ab7220edcf39a5775cf0))
* Graceful database connection handling ([4ff5d46](https://github.com/DeRuina/kuha-rest-api/commit/4ff5d46535b6e7faf18d04d090b3c0e59958aab5))
* gzip decompression POST endpoints middleware ([00098dc](https://github.com/DeRuina/kuha-rest-api/commit/00098dceb1c30790b20c386eeb177acaf735084a))
* klab sqlc ([e443699](https://github.com/DeRuina/kuha-rest-api/commit/e4436996e3feabf2a1f2bf300302c9d19a63d7d3))
* POST activity zones tietoevry endpoint ([11f809d](https://github.com/DeRuina/kuha-rest-api/commit/11f809d18bc036dca3735bfc50909b65c394e6ee))
* POST measurements tietoevry endpoint ([bcaee13](https://github.com/DeRuina/kuha-rest-api/commit/bcaee1308be57bac0f9bc17f67396376aadf8f4b))
* POST questionnaires tietoevry endpoint ([5176ac2](https://github.com/DeRuina/kuha-rest-api/commit/5176ac288b8fa097d63d9175eac476eb931a14ea))
* POST test_results endpoint tietoevry ([9d66f3b](https://github.com/DeRuina/kuha-rest-api/commit/9d66f3bfb1fb4f1a0a1846be1fd2e6699eb98003))
* relase please script ([a4cbb0b](https://github.com/DeRuina/kuha-rest-api/commit/a4cbb0b9044f3d9bc51c1b4bffdcb03296de0c17))
* Tietoevry POST Symptopms endpoint ([f3b1866](https://github.com/DeRuina/kuha-rest-api/commit/f3b18667add65296640926e6a5dd94b6ad98965a))
* update api version automatically ([0955bc5](https://github.com/DeRuina/kuha-rest-api/commit/0955bc540cfbac14addf13ad6c3c7f252c0c9d16))
* utv archinisis_tokens table endpoints ([5608bea](https://github.com/DeRuina/kuha-rest-api/commit/5608bead0166426cdc761e66d9aaf032fa086ed9))
* utv source_cache endpoints ([b722d4e](https://github.com/DeRuina/kuha-rest-api/commit/b722d4eb0046100037c7a1c3788af27e060a2d03))


### Bug Fixes

* 401, 503 error documentations ([6f3bba9](https://github.com/DeRuina/kuha-rest-api/commit/6f3bba9bda633804ec142f3b6684957e2164dc43))
* allow_anonymous_data: true--&gt; allow_anonymous_data: 0 ([667b4a8](https://github.com/DeRuina/kuha-rest-api/commit/667b4a8a5208fb8369a2f95b7d3405f6bd90df6e))
* bulk inserts instead of single inserts in POST exercise & symptoms ([97a94b9](https://github.com/DeRuina/kuha-rest-api/commit/97a94b9a3f456256246ef1df204d406cfc2f629c))
* cache adjustments ([6b41b6d](https://github.com/DeRuina/kuha-rest-api/commit/6b41b6dba05507cb03bed1f2c530f196cae5a7e5))
* data4update token4update response fix ([49d5593](https://github.com/DeRuina/kuha-rest-api/commit/49d559332ffdd506520d31b89f74a9a9cc489a58))
* DataTimeout modification ([076a30f](https://github.com/DeRuina/kuha-rest-api/commit/076a30fdc3ff7bfa1dad8a918a62b0e72f57e422))
* decompress EOF error ([a8226eb](https://github.com/DeRuina/kuha-rest-api/commit/a8226eb130253304efc9bc042f314c8b890b74d7))
* DELETE user klab. remove sportti_id_list table ([2930dfa](https://github.com/DeRuina/kuha-rest-api/commit/2930dfab21446ab4a5bab40ea2e5f62f1f71a6fe))
* duration ISO8601 parsing ([101c8c3](https://github.com/DeRuina/kuha-rest-api/commit/101c8c3110e766d6bffee7d5495906271310224d))
* fis queries ([5d68ab3](https://github.com/DeRuina/kuha-rest-api/commit/5d68ab3c79fe392e68cfda2d7128d782c2f1cd91))
* fis schema ([5ea88dd](https://github.com/DeRuina/kuha-rest-api/commit/5ea88dd3465a17d3d8193ef3c1841873b9656213))
* GET all utv endpoint ([7ee9688](https://github.com/DeRuina/kuha-rest-api/commit/7ee9688cc4b992aa223004c2f94ec513529ba32d))
* GET exercises tietoevry endpoint ([0917dac](https://github.com/DeRuina/kuha-rest-api/commit/0917dac39dccf00596406ed41ae77a599518508c))
* GET sport-ids klab ([931ede4](https://github.com/DeRuina/kuha-rest-api/commit/931ede497def4a24b6faf289123d3c141d4b8af4))
* gzip middleware for Archinisis & kamk Injury endpoints ([ba4dc3f](https://github.com/DeRuina/kuha-rest-api/commit/ba4dc3f2dd236fdadcfad2769373eaa0eeee0fb9))
* health endpoint ([c509ed0](https://github.com/DeRuina/kuha-rest-api/commit/c509ed00632b7162ec4b1651c1235f456c599eb4))
* id numeric error ([773b8e8](https://github.com/DeRuina/kuha-rest-api/commit/773b8e824cb881c81a09dd2ae0e0f9edcfbfff1c))
* insert and not upsert in POST exercises ([e2cc960](https://github.com/DeRuina/kuha-rest-api/commit/e2cc9601347c397b8ab0c5acf3baf9b81e8b9cf0))
* klab GET not found 404 ([b685d9d](https://github.com/DeRuina/kuha-rest-api/commit/b685d9da1801c77ec1e90e656aae7314ebb32422))
* latest default to 1 ([ad2e775](https://github.com/DeRuina/kuha-rest-api/commit/ad2e775797a1811d6f98c7c35ba2c74d78290157))
* log deleted users on DELETE endpoint ([f3107c0](https://github.com/DeRuina/kuha-rest-api/commit/f3107c0a2b8e2323b97c21beea74fbc6be179251))
* prevalidation errors ([9989843](https://github.com/DeRuina/kuha-rest-api/commit/998984319ae6c60d5b5953184ecc87daa8262b5a))
* prevalidation to avoid reverted queries if one fails ([8dbe21a](https://github.com/DeRuina/kuha-rest-api/commit/8dbe21aca128522a80edb40838a4b2a39d458c24))
* raise json request limit ([44ed37f](https://github.com/DeRuina/kuha-rest-api/commit/44ed37f46c245b4f38a6dc8d9a4cbad20cdcc182))
* raise ratelimit for utv ([5448dce](https://github.com/DeRuina/kuha-rest-api/commit/5448dce0c63cbcf1b52159fd2a20c67ed2b8a2f4))
* refresh + token error ([fb70e40](https://github.com/DeRuina/kuha-rest-api/commit/fb70e40f1fc934e7cec233732794933e07341f7b))
* remove sportti_id_list and create DELETE archinisis endpoint ([686173f](https://github.com/DeRuina/kuha-rest-api/commit/686173fb268968dc8dd899c150d447f533331c8f))
* remove symptom references symptoms ([8af1027](https://github.com/DeRuina/kuha-rest-api/commit/8af102794de00a0b03a0d9b42972d729cbd1f881))
* RFC3339 format modifications ([54ce102](https://github.com/DeRuina/kuha-rest-api/commit/54ce102fdb3c2b3f5a33fa9f347c596e987e3f18))
* roles update ([9dfc91c](https://github.com/DeRuina/kuha-rest-api/commit/9dfc91cdab851b5db0af9212b91c5fadf8381a8f))
* swagger docs utv token endpoint ([79044bc](https://github.com/DeRuina/kuha-rest-api/commit/79044bc0d375b0d4decf40f6969dd95231b901ef))
* test_results swagger docs ([c598e08](https://github.com/DeRuina/kuha-rest-api/commit/c598e0859b4e29569818bc8567b13ac96d028fe8))
* tietoevry cache fix ([17c54bc](https://github.com/DeRuina/kuha-rest-api/commit/17c54bc605da0fd13cfa68a0db9b41dc2b463db9))
* underscore parameters unity ([754afb1](https://github.com/DeRuina/kuha-rest-api/commit/754afb1538aa4b09381e12e2912f4a8d9824c639))
* update api version automatically fix ([addd09d](https://github.com/DeRuina/kuha-rest-api/commit/addd09d58ab90d2a4a4306f532c8bdde36feb6f7))
* UPSERT tietoevry ([13b3f51](https://github.com/DeRuina/kuha-rest-api/commit/13b3f5163f2467fb8a2c205c360d261d76dff1b4))
* utv klab archinisis sport_id integration ([c461470](https://github.com/DeRuina/kuha-rest-api/commit/c46147000bccb800ca0d222ba6b7da46fd616253))


### Performance Improvements

* bump timberjack version ([8c8fe2e](https://github.com/DeRuina/kuha-rest-api/commit/8c8fe2e2f29a226550a525ac857b3ee4a5519e80))

## 2025-12-17

### Changed
- Make repo public

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
