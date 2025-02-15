# Proxy Design Pattern

## What is Proxy Pattern?

Proxy is a structural design pattern that provides an object that acts as a substitute for a real service object used by a client. A proxy receives client requests, does some work (access control, caching, etc.) and then passes the request to a service object.

The proxy object has the same interface as a service, which makes it interchangeable with a real object when passed to a client.

## Conceptual Example

A web server such as Nginx can act as a proxy for your application server:

- It provides controlled access to your application server.
- It can do rate limiting.
- It can do request caching.

## Reference

- [https://refactoring.guru/design-patterns/proxy/go/example](https://refactoring.guru/design-patterns/proxy/go/example)
