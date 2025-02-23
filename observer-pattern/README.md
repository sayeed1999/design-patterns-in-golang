# Observer Pattern

The Observer Pattern is a behavioral design pattern that defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically. This pattern is useful for implementing distributed event-handling systems.

## Problem

Imagine you have a notification system where you need to notify users through different channels (e.g., email and SMS) whenever an event occurs. Without a unified interface, the client code has to interact with multiple notification channels directly, making the code complex and difficult to maintain.

## Solution

The Observer Pattern helps solve this problem by defining a subscription mechanism to allow multiple objects (observers) to listen to and react to events (notifications) from a subject. In this example, a `NotificationService` class is created to manage observers and notify them of events.

## Approach

1. **Define the Observer Interface**: This interface will have a method `Update` that all concrete observers will implement.
2. **Create Concrete Observers**: Implement the `Observer` interface for different notification channels (e.g., Email and SMS).
3. **Define the Subject Interface**: This interface will have methods to add, remove, and notify observers.
4. **Create Concrete Subject**: Implement the `Subject` interface to manage observers and notify them of events.

### Structure

- **Observer Interface**: Defines the `Update` method.
  - [observer/observer.go](observer/observer.go)
- **Concrete Observers**: Email and SMS observers implementing the `Observer` interface.
  - [observer/email.go](observer/email.go)
  - [observer/sms.go](observer/sms.go)
- **Subject Interface**: Defines methods to add, remove, and notify observers.
  - [publisher/subject.go](publisher/subject.go)
- **Concrete Subject**: NotificationService implementing the `Subject` interface.
  - [publisher/notification_service.go](publisher/notification_service.go)

### Usage

The main function demonstrates how to use the `NotificationService` to add observers and notify them of events.

- [main.go](main.go)

### Output

```bash
Sending email notification...
Sending SMS notification...
```

## Benefits

- Decoupling: Decouples the subject from its observers, allowing them to vary independently.
- Flexibility: You can add or remove observers at runtime without modifying the subject.
- Scalability: Easily extendable to include new types of observers without changing the subject.

## Special Thanks

Tutorial Link: <https://refactoring.guru/design-patterns/observer>

The tutorial at Refactoring Guru provides a great example of the Observer Pattern. I recommend readers to check their tutorial for better learning!
