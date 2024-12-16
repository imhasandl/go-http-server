## Project Setup

**1. Prerequisites**

* **Go:** Ensure you have Go installed and configured on your system.
* **PostgreSQL:** Install and set up a PostgreSQL database.

**2. Project Initialization**

```bash
go mod init 
```

**3. Install Dependencies**

```
go get github.com/lib/pq # PostgreSQL driver
```

**4. Install Goose Database Migration**

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

**5. Install sqlc (SQL Code Generation)**

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## API Endpoints

* **`/api/healthz`** (GET): Checks API health.
* **`/api/metrics`** (GET): Exposes application metrics (if enabled).
* **`/api/login`** (POST): User login.
* **`/api/refresh`** (POST): Refresh access token.
* **`/api/revoke`** (POST): Revoke access token.
* **`/api/users`** (POST): Create a new user.
* **`/api/users`** (PUT): Update user information.
* **`/api/chirps`** (POST): Create a new chirp.
* **`/api/chirps`** (GET): Retrieve a list of chirps.
* **`/api/chirps/{chirpID}`** (GET): Retrieve a specific chirp.
* **`/api/chirps/{chirpID}`** (DELETE): Delete a specific chirp.
* **`/api/polka/webhooks`** (POST): Gives a user a "paid" access turns the is_chirpy_red to true  

**Headers**

* **Authorization:** `Bearer <user_token>` (for user-related endpoints)
* **Authorization:** `ApiKey <your_api_key>` (for webhook endpoint)

**Chirps Path Query Parameters**

* **sort:** `asc` (ascending), `desc` (descending)
* **author_id:** ID of the user to filter chirps by

**Request Body**

* **chirps:** `{ "body": "Chirp message" }`
* **webhooks:** `{ "data": { "user_id": "${userID}" }, "event": "user.upgraded" }`
