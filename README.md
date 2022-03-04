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

---

# Usage

## Flags

To change the port the application will run on you can invoke the flag -port=\<port number\>

To change the environment you can pass -env=\<environment\>

---

# Dependencies

## httprouter

When building API endpoints without thrid party libraries we face the limitation that the http.ServeMux does not allow routing to different handlers based on the request method. It doesn't provide support for clean URLs with interpolated paramters either.

Thus the choice was made to use  julienschmidt's httprouter for it solves both this problems and is extremely performant because it uses a radix sort algorithm for URL matching. 

---
# Miscellaneous

## URL prefixing

As one can see in the endpoints, all of them are prefixed with v1. As in real business users often need to change endpoints functionality overtime, sometimes breaking backwards compatibility, versioning needs to be implemented. There are two common approaches for doing so:

- Prefixing URLs, like done here

- Using custom *Accept* and *Content-Type* headers on requests and responses, ie: Accept: application/vnd.greenlight-v1.

Even though custom headers are arguably "purer", I think URL prefixes come out on top regarding ease of use, thus the API was written this way.
