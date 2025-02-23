# Composite Pattern

The Composite Pattern is a structural design pattern that allows you to compose objects into tree structures to represent part-whole hierarchies. This pattern lets clients treat individual objects and compositions of objects uniformly.

## Problem

Imagine you have a file system where files and folders can contain other files and folders. You want to be able to perform operations (e.g., search) on both individual files and folders containing multiple files and subfolders. Without a unified interface, the client code has to handle files and folders differently, making the code complex and difficult to maintain.

## Solution

The Composite Pattern helps solve this problem by defining a unified interface for both individual objects (files) and compositions of objects (folders). In this example, `File` and `Folder` classes are created to represent files and folders, respectively, and both implement a common interface.

## Approach

1. **Define the Component Interface**: This interface will have methods that both files and folders will implement.
2. **Create Leaf Class**: Implement the `Component` interface for individual objects (e.g., File).
3. **Create Composite Class**: Implement the `Component` interface for compositions of objects (e.g., Folder).
4. **Implement Operations**: Implement operations (e.g., search) that can be performed on both individual objects and compositions.

### Structure

- **Component Interface**: Defines the methods that both files and folders will implement.
  - [component.go](component.go)
- **Leaf Class**: File class implementing the `Component` interface.
  - [file.go](file.go)
- **Composite Class**: Folder class implementing the `Component` interface.
  - [folder.go](folder.go)

### Usage

The main function demonstrates how to use the `File` and `Folder` classes to create a file system and perform a search operation.

- [main.go](main.go)

### Output

```bash
Searching for rose in file2
Searching for rose in file3
Searching for rose in file1
Searching for rose in file2
Searching for rose in file3
```

## Benefits

- Unified Interface: Provides a unified interface for both individual objects and compositions of objects.
- Simplified Client Code: Simplifies client code by allowing it to treat individual objects and compositions uniformly.
- Scalability: Easily extendable to include new types of components without changing the client code.

## Special Thanks

Tutorial Link: [Refactoring Guru: Composite Pattern in Go](https://refactoring.guru/design-patterns/composite/go/example)

The tutorial at Refactoring Guru provides a great example of the Composite Pattern. I recommend readers to check their tutorial for better learning!
