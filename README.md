# URL Shortener in Go-lang

## Tech Stack used

- Gin, MongoDB, Mongoose

### Design

- Takes in long URL, stores it in DB while creating an id,
- generates a new URL using helper functions,
- maps the new URL to original one using id
- Returns a new shorter URL.
