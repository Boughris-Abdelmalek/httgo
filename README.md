1. Start by Understanding the RFC
   Read RFC 2616 (HTTP/1.1) or RFC 1945 (HTTP/1.0).

Break it down into key components:

Request parsing

Response formatting

Headers & status codes

Connection handling (persistent vs. non-persistent)

2. Set Up a Basic TCP Server
   Use Go’s net package to accept raw TCP connections.

Read incoming data and inspect it manually to understand raw HTTP messages.

3. Implement HTTP Request Parsing
   Parse the request line (method, path, version).

Parse headers and handle different types of line endings (\r\n).

Handle different request methods (GET, POST, etc.).

Handle request bodies (especially for POST).

4. Construct and Send HTTP Responses
   Implement response status codes (e.g., 200 OK, 404 Not Found).

Format headers properly (Content-Type, Content-Length, etc.).

Implement chunked transfer encoding for streamed responses.

5. Handle Persistent Connections (Keep-Alive)
   Support Connection: keep-alive and Connection: close.

Implement request pipelining (multiple requests over the same connection).

6. Handle Edge Cases
   Malformed requests

Large requests and timeouts

Handling different content types and encodings

7. Write a Simple HTTP Client
   Implement a basic HTTP client to send requests to your server.

8. Compare with Go’s net/http Package
   Once you have a working implementation, compare it with Go’s built-in net/http to see how it handles things differently