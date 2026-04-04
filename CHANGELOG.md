# Changelog

## [1.5.0](https://github.com/BaoDuong254/gopher-social/compare/v1.4.0...v1.5.0) (2026-04-04)


### Features

* add check for open API version pull request before syncing main into develop ([a1c91db](https://github.com/BaoDuong254/gopher-social/commit/a1c91db5b0bb377295c4fb4f3c5913622eaa8704))
* add check for open API version pull request before syncing main into develop ([97cb21f](https://github.com/BaoDuong254/gopher-social/commit/97cb21fb1b70b47f3ddd67d88a92c00e2a003372))

## [1.4.0](https://github.com/BaoDuong254/gopher-social/compare/v1.3.0...v1.4.0) (2026-04-04)


### Features

* use dynamic feed and post keys in SWR hooks ([f70296e](https://github.com/BaoDuong254/gopher-social/commit/f70296e6207ef59a3e2ef26694e36e086fac486a))
* use dynamic feed and post keys in SWR hooks ([8635d3f](https://github.com/BaoDuong254/gopher-social/commit/8635d3f194fbe3cf73ad211ae2f9a4961a6f997c))

## [1.3.0](https://github.com/BaoDuong254/gopher-social/compare/v1.2.0...v1.3.0) (2026-04-04)


### Features

* add checkout step for syncing repository after release ([22f3867](https://github.com/BaoDuong254/gopher-social/commit/22f3867e5b702aaa7ab3a790e4a031b5dbe3cc43))

## [1.2.0](https://github.com/BaoDuong254/gopher-social/compare/v1.1.0...v1.2.0) (2026-04-04)


### Features

* enhance user feed handler with authentication check and add vercel configuration ([2dd8401](https://github.com/BaoDuong254/gopher-social/commit/2dd8401fed03cf8bd9b5cf4448ab52225c7be09b))

## [1.1.0](https://github.com/BaoDuong254/gopher-social/compare/v1.0.0...v1.1.0) (2026-04-03)


### Features

* add comments functionality with database integration and update post structure ([6e99909](https://github.com/BaoDuong254/gopher-social/commit/6e9990960cdb65aad1da508f261bdb5046e7f487))
* add context timeout for database queries in store package ([11caf29](https://github.com/BaoDuong254/gopher-social/commit/11caf29c1245420ceab6b150cb24bb7fe69aa4c0))
* add Dockerfile for building and running the API ([0696e79](https://github.com/BaoDuong254/gopher-social/commit/0696e79363db7ac2c3483df4fe9ebf9fe469726b))
* add expvar metrics for version and runtime statistics ([99fe477](https://github.com/BaoDuong254/gopher-social/commit/99fe477f9a87e18d3f56ed129bc6b0a16180aa62))
* add get post handler and update routing for posts ([731f698](https://github.com/BaoDuong254/gopher-social/commit/731f6982046c4c948ead628dfa7c6bc2c2ac7735))
* add JSON handling functions and update healthcheck response ([8a781ed](https://github.com/BaoDuong254/gopher-social/commit/8a781ed9701cbb5fc7b926e2b13fd9b92d07617f))
* add migration scripts for creating and dropping database indexes ([107d5e9](https://github.com/BaoDuong254/gopher-social/commit/107d5e9b1436d165795c5d56e47bcd4c78a43453))
* add React and Vite setup with routing and account confirmation page ([9b5d59a](https://github.com/BaoDuong254/gopher-social/commit/9b5d59a265a8a3a158bafdf9135be79a376988be))
* add roles table and user role association in database migrations ([4dd9b91](https://github.com/BaoDuong254/gopher-social/commit/4dd9b9163951126aae460a5dbbabce3ffa3bfd7a))
* add Swagger documentation and update API routes for user retrieval ([4265112](https://github.com/BaoDuong254/gopher-social/commit/426511272ecab17028925816fc4db757fd884867))
* add Swagger documentation files for API endpoints ([c7ac518](https://github.com/BaoDuong254/gopher-social/commit/c7ac518af162bb07097c4863e29f280bed4d1a37))
* add test utilities and mock implementations for user authentication and storage ([52195df](https://github.com/BaoDuong254/gopher-social/commit/52195dfdba318e4603577108586983d2fce2adc1))
* add token for creating pull request in update API version workflow ([191d1bc](https://github.com/BaoDuong254/gopher-social/commit/191d1bc7e746477ec8ddff3687d2fa4b634bd66b))
* add user confirmation, login, post creation, and single post view functionality ([2d0be51](https://github.com/BaoDuong254/gopher-social/commit/2d0be51e15c6d3fdce910b49d4950ef7edc2f9c0))
* add user feed retrieval functionality with handler and store method ([30ad663](https://github.com/BaoDuong254/gopher-social/commit/30ad663352d8fc1f7b8d0e3263eed16be056bb35))
* add user retrieval functionality with GetByID method and handler ([7655fff](https://github.com/BaoDuong254/gopher-social/commit/7655fff7bc179f5ab87dbd498fa18d87f514d169))
* add validation for post creation payload and initialize validator ([83ca657](https://github.com/BaoDuong254/gopher-social/commit/83ca657e8d6a7510dc2778f10a83396532299924))
* add versioning to posts for optimistic concurrency control ([eba8dcd](https://github.com/BaoDuong254/gopher-social/commit/eba8dcd890128c6c7a296966b85e1ef11c32147a))
* create base ([5fd1436](https://github.com/BaoDuong254/gopher-social/commit/5fd14361cdeb51babd1d69adcb5d072cf5f101d9))
* enhance post retrieval with search and tag filtering in user feed ([72d8f69](https://github.com/BaoDuong254/gopher-social/commit/72d8f69210084c7d4cf1571d613ba15b66bf20ce))
* enhance user experience and error handling across the application ([56b80c0](https://github.com/BaoDuong254/gopher-social/commit/56b80c01a9c49b5d0cd0609b395342f022283d53))
* enhance version resolution logic in GitHub Actions workflow ([11e488b](https://github.com/BaoDuong254/gopher-social/commit/11e488bbc34636aa295d8ac01415a35849704a15))
* implement authentication middleware and enhance user context handling ([a4fddfb](https://github.com/BaoDuong254/gopher-social/commit/a4fddfb380ffaf1a97ab5918dcbe1e8b26c9a00f))
* implement basic authentication middleware and error handling ([c0d8e57](https://github.com/BaoDuong254/gopher-social/commit/c0d8e5785880d7708a323e54c8ff59b6991b93ad))
* implement comments creation and seeding functionality ([7faf50e](https://github.com/BaoDuong254/gopher-social/commit/7faf50ed2f081a36218ca2abea809560a7f0616c))
* implement create post handler and update database migrations for tags ([bc6fbfd](https://github.com/BaoDuong254/gopher-social/commit/bc6fbfdbe52dd0c7eb31b6ba25e220fd3d27de16))
* implement CRUD operations for posts with context middleware ([05ed9ac](https://github.com/BaoDuong254/gopher-social/commit/05ed9ac65ff638b85d7253ff7a658a11f84ea109))
* implement email authentication with SendGrid and Mailtrap integration ([739ea54](https://github.com/BaoDuong254/gopher-social/commit/739ea5407ede20bf33d9de7f66154dd5b71c2765))
* implement follow and unfollow functionality with context middleware ([f65f2cc](https://github.com/BaoDuong254/gopher-social/commit/f65f2cc461b528077578b1d01ed9cd630402482c))
* implement graceful shutdown for HTTP server ([a6e7c1b](https://github.com/BaoDuong254/gopher-social/commit/a6e7c1b5e23c226118243735601fb725f1a0202a))
* implement jsonResponse method and update handlers to use it for consistent JSON responses ([9e0b0ee](https://github.com/BaoDuong254/gopher-social/commit/9e0b0eef9f8f060490426c795996db846517ccc7))
* implement JWT authentication with token generation and validation ([df4af0a](https://github.com/BaoDuong254/gopher-social/commit/df4af0a6dc5f4da8328a9bb7a504e6bf61178561))
* implement pagination for user feed retrieval with query parameters ([ab47c57](https://github.com/BaoDuong254/gopher-social/commit/ab47c57d01e24eb6e759e9d63a7b809f48cf0fb9))
* implement rate limiting middleware and fixed window rate limiter ([bfbc0c9](https://github.com/BaoDuong254/gopher-social/commit/bfbc0c99f9624357f365093d27ab587fa8b6bb25))
* implement repository pattern with database migrations and environment configuration ([31fe6ab](https://github.com/BaoDuong254/gopher-social/commit/31fe6abbafc91edfb4ce0e084bf54515a3460b48))
* implement role-based access control and user role management ([3ae05ef](https://github.com/BaoDuong254/gopher-social/commit/3ae05ef4c2573a96c53723e3424525e7264a9bf9))
* implement user registration and activation with email invitations ([827fc13](https://github.com/BaoDuong254/gopher-social/commit/827fc1338ac5b9f44384572d3f234a83b045c906))
* improve code formatting and structure in App and SinglePost components ([d78f47d](https://github.com/BaoDuong254/gopher-social/commit/d78f47dc8668d993b5f5ad0b7a32233697720008))
* improve code formatting and structure in App and SinglePost components ([bb5d4f8](https://github.com/BaoDuong254/gopher-social/commit/bb5d4f8f93748d250a7802df778b3154d072983f))
* improve environment variable loading and server URL logging in main function ([c6b87c0](https://github.com/BaoDuong254/gopher-social/commit/c6b87c04fd94b774a8a6fbdbba8dcbe4d358cfaf))
* integrate Redis caching for user data and add Redis configuration support ([65a7791](https://github.com/BaoDuong254/gopher-social/commit/65a77918208c19f07561c47ae537cf751a9baa4e))
* integrate zap logger for improved logging in application ([7f972d4](https://github.com/BaoDuong254/gopher-social/commit/7f972d4d9fb342c69f4e13348700149401bef6b1))
* refactor error handling in healthcheck and post handlers; add centralized error responses ([05c243b](https://github.com/BaoDuong254/gopher-social/commit/05c243b0b4f621471ff6f91612048d5398bafc91))
* update GitHub Actions workflow to extract and update API version from CHANGELOG.md ([88d44db](https://github.com/BaoDuong254/gopher-social/commit/88d44db3288ef0480060e1d9512b06ac2141b29d))


### Bug Fixes

* ci ([401853f](https://github.com/BaoDuong254/gopher-social/commit/401853f9d9248945e17db41ca0639c0754c4f117))
* update API endpoint for user feed fetching ([f0fe53d](https://github.com/BaoDuong254/gopher-social/commit/f0fe53d2c0264f812f819929d13e821e9e1624fe))
* update description in package.json for clarity ([81196d4](https://github.com/BaoDuong254/gopher-social/commit/81196d42fc657fa0e52d5ac82279fee8300cca03))
* update homepage URL in package.json ([b0cb032](https://github.com/BaoDuong254/gopher-social/commit/b0cb032e66feab020fcf8d05302aae11eba3db77))
