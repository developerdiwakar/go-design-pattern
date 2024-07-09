# Design Patterns in GO Programming

There are three main categories of Design Patterns.
- Creational Design Patterns
- Structural Design Patterns
- Behavioral Design Patterns

## Creational 
Creational design patterns provide various object creation mechanisms, which increase flexibility and reuse of existing code.

### Some commonly used Creational design patterns:
 - **Factory Method** 
    - **Concept:** Defines an interface for creating objects but lets subclasses decide which class to intantiate.
    - **Use Case:** When you want to create objects of different related classes based on specific criteria.
    - **Benefits:** Promotes loose coupling, allows for subclass control over object creation, hides concrete implementation.
    - **Drawbacks:** Introduces an extra layer of abstraction, might require subclass proliferation.
 - **Abstract Factory**
    - **Concept:** Create families of related objects without specifying their concrete classes.
    - **Use Case:** When you need to create sets of related objects that belong to different families.
    - **Benefits:** Promotes loose coupling, simplifies client code, enforces consistent family creation.
    - **Drawbacks:** More complex than Factory Method, can lead to a large number of concrete factory classes. 
 - **Builder**
    - **Concept:**
    - **Use Case:**
    - **Benefits:**
    - **Drawbacks:**
 - **Prototype**:
    - **Concept:**
    - **Use Case:**
    - **Benefits:**
    - **Drawbacks:**
 - **Singleton**:
    - **Concept:**
    - **Use Case:**
    - **Benefits:**
    - **Drawbacks:**


## How to run the program
- **1st Method:**
Ensure All Files Have package main at the Top: Make sure every .go file has package main at the top.
Correct Import Statements: Ensure main.go correctly imports the necessary packages.

Build the Project: Use the build command from the directory(for eg. factory-method, singleton, facade):

```sh
go build -o myapp
./myapp
```
With the correct package declarations and file structure, your program should compile without the "undefined" errors.

- **2nd Method**
When you use go run main.go, it only compiles and runs the main.go file. It does not automatically include other files in the same package. To run the program using go run, you need to specify all the .go files in the directory(for eg. factory-method, singleton, facade).

You can use the following command to run the program:
```sh
go run main.go ak47.go gun.go gunFactory.go maverick.go
```

Alternatively, you can use a wildcard to include all .go files in the directory:
```sh
go run ./*.go
```