# Proxy Pattern with Nginx

The Proxy Pattern is a structural design pattern that provides an object representing another object. It acts as an intermediary, controlling access to the original object, allowing you to add additional functionality without changing the original object's code. This pattern is useful for implementing access control, lazy initialization, logging, and more.

## Problem

Imagine you have an application that handles various HTTP requests. You want to add rate limiting to control the number of requests a client can make to your application within a certain period. Without a proxy, you would need to add rate limiting logic directly to your application, making the code complex and difficult to maintain.

## Solution

The Proxy Pattern helps solve this problem by introducing a proxy server that controls access to the application. In this example, an `NginxServer` class is created to act as a proxy, handling rate limiting and forwarding requests to the application.

## Approach

1. **Define the Subject Interface**: This interface will have methods that both the real subject and proxy will implement.
2. **Create Real Subject**: Implement the `Subject` interface for the real subject (e.g., Application).
3. **Create Proxy Class**: Implement the `Subject` interface for the proxy class, which will control access to the real subject and add additional functionalities (e.g., rate limiting).

### Structure

- **Subject Interface**: Defines the methods that both the real subject and proxy will implement.
  - [subject.go](subject.go)
- **Real Subject**: Application class implementing the `Subject` interface.
  - [application.go](application.go)
- **Proxy Class**: NginxServer class implementing the `Subject` interface and adding rate limiting functionality.
  - [nginx_server.go](nginx_server.go)

### Usage

The main function demonstrates how to use the `NginxServer` proxy to handle requests with rate limiting.

- [main.go](main.go)

### Output

```bash
Handling request for /users with method POST
Handling request for /users with method POST
Handling request for /users with method DELETE
Handling request for /users with method DELETE

Rate limit exceeded for /users with method POST
Rate limit exceeded for /users with method DELETE
```

## Benefits

- Access Control: Controls access to the real subject, allowing you to add functionalities like rate limiting.
- Simplified Code Maintenance: Keeps the real subject code clean and focused on its primary responsibilities.
- Flexibility: Easily extendable to include new functionalities without modifying the real subject.

## Special Thanks

[Refactoring Guru: Proxy Pattern in Go](https://refactoring.guru/design-patterns/proxy/go/example)

The tutorial at Refactoring Guru provides a great example of the Proxy Pattern. I recommend readers to check their tutorial for better learning!
