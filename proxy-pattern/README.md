# Proxy Design Pattern

The Proxy Design Pattern a structural design pattern is a way to use a placeholder object to control access to another object. Instead of interacting directly with the main object, the client talks to the proxy, which then manages the interaction. This is useful for things like controlling access, delaying object creation until it’s needed (lazy initialization), logging, or adding security checks.

## How to implement Proxy Design Pattern?

Below are the simple steps to implement the Proxy Design Pattern:

1. Create the Real Object Interface: Define an interface or abstract class that represents the operations the real object will provide. Both the real object and proxy will implement this interface.

2. Create the Real Object: This class implements the interface and contains the actual logic or operation that the client wants to use.

3. Create the Proxy Class: The proxy class also implements the same interface as the real object. It holds a reference to the real object and controls access to it. The proxy can add extra logic like logging, caching, or security checks before calling the real object’s methods.

4. Client Uses the Proxy: Instead of creating the real object directly, the client interacts with the proxy. The proxy decides when and how to forward the client’s request to the real object.
