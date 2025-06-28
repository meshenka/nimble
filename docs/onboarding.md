# Technical Onboarding Guide

Welcome to the Nimble project! This guide will help you understand the project structure,
design patterns, and key components.

## Backend

The backend is a Go application that serves a RESTful API for generating random
hero characters.

### Package Organization

The backend code is organized into several packages:

- **`cmd`**: Contains the main application entry point and command-line interface
  (CLI) utilities.
- **`handler`**: Defines the HTTP handlers for the API endpoints. Each handler is 
  responsible for processing incoming requests, calling the appropriate services,
  and writing the JSON response.
- **`internal`**: Contains the core business logic of the application.
  This package is further divided into sub-packages:
  - **`hero`**:  The hero package is responsible for creating hero characters.
    It has several sub-packages:
    - **`ancestry`**, **`background`**, **`class`**, **`motivation`**, **`origin`**,
      **`quirk`**:
      These packages define the various components of a hero character. Each package
      is responsible for selecting a random element from a predefined list.
    - **`log`**: Provides a structured logging wrapper around the standard
      `log/slog` library.
    - **`seeder`**:  Initializes and manages the random number generator with
      a seed, ensuring that character generation can be deterministic.
    - **`transport`**:  Contains HTTP server primitives and middleware.
- **`nimble`**: The root package of the application, responsible for wiring 
  together all the components and starting the server.

### Design Patterns

The backend employs the following design patterns:

- **Dependency Injection**: The `nimble` package is responsible for creating
  and injecting dependencies into the various components. This makes the code
  more modular and easier to test.
- **Options Pattern**: The `nimble.Serve` function uses the options pattern
  to allow for flexible configuration of the application.

## Frontend

The frontend is a React application that provides a user interface
for generating and displaying hero characters.

### Component Structure

The frontend is built with a component-based architecture.
The main components are:

- **`HeroApp`**: The root component of the application.
  It fetches a random hero from the API and displays it.
- **`Descriptor`**: A reusable component for displaying a hero's attributes.
- **`Background`**: A component that displays a background image.
- **`Link`**: A stylized anchor tag component.

### State Management

The frontend uses the `useState` and `useEffect` hooks to manage the application
state. When the `HeroApp` component mounts, it fetches a random hero from the API
and stores it in the component's state.
