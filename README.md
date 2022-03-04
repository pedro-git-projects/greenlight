# Greenlight

## A JSON based api for managing movies.

---

# Supported endpoint actions

| Method | URL Pattern               | Action                                   |
|--------|---------------------------|------------------------------------------|
| Get    | /v1/healthcheck           | Show application health and version info |
| Get    | /v1/movies                | Show all movies details                  |
| Post   | /v1/movies                | Create a new movie                       |
| Get    | /v1/movies/:id            | Show details for a particular movie      |
| Patch  | /v1/movies/:id            | Update the details for a specific movie  |
| Delete | /v1/movies/:id            | Delete a particular movie                |
| Post   | /v1/users                 | Register a new user                      |
| Put    | /v1/users/activated       | Activate an user                         |
| Put    | /v1/users/password        | Update the password for an user          |
| Post   | /v1/tokens/authentication | Generate a new authentication token      |
| Post   | /v1/tokens/password-reset | Generate a new password-reset token      |
| Get    | /debug/vars               | Display application metrics              |
