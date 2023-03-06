# Confinio Project ("confinio" means "border" in latin)
## Components:
* **confinio-porta** : HTTP server, reverse proxy/API gateway ("porta" means "gate")
* **confinio-machina** : mocking/testing server, CMS engine ("machina" means "engine")

## Features in development:
* Multi-domain serving of static files, indexes (HTTP/S server)
* Dispatching HTTP requests to the backends (reverse proxy): wo mutations, custom headers
* CMS engine initial support: user sessions (via cookies, URL params, headers/tokens), RBAC
* https://github.com/pkg-wire

# **Warning!!!**: Codebase is in early design stage!
