# redirect

A simple webserver that redirects all requests to a specified target URL.
The path and query string from the original request are included in the redirect.

This is mainly intended to provide the very simple kind of redirect service domain registrars sometimes provide.

Environment variables:

- TARGET - Absolute URL of redirect target. Required.
- PORT - Listening port. Default: 8080
- STATUS - HTTP status code in redirect response. Default: 302
