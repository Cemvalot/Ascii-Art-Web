# ASCII Art Web Generator

This project is a web application built using Go that allows users to create and download ASCII art from custom text input. It offers various banner styles, light and dark themes, input validation, and the ability to export generated ASCII art as .txt or .doc files.

## Table of Contents

- [Getting Started](#getting-started)
- [Features](#features)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Usage](#usage)
- [Error Handling](#error-handling)
- [Styling](#styling)
- [Future Enhancements](#future-enhancements)

## Getting Started

To run the ASCII Art Web Generator:

1. Clone this repository
2. Run the server:
   ```bash
   go run main.go
   ```
3. Access the web application by opening a web browser and navigating to http://localhost:8080

## Features

- Generate ASCII art from custom text input
- Multiple banner styles to choose from
- Input validation
- Export art in multiple file formats

## Project Structure

The project has the following structure:

- `main.go`: Main entry point of the application
- `banners/`: Contains text files for different banner styles
- `internal/`: Contains the core application logic
  - `config/`: Configuration-related code
  - `handlers/`: HTTP request handlers
  - `server/`: Server logic
  - `service/`: Service implementations
- `static/`: Static assets (including CSS)
- `templates/`: HTML templates
- `dockerfile`: Creates docker image when running the appropriate command

## Configuration

Modify settings in `internal/config/config.go` or via environment variables.

## Usage

1. Input custom text
2. Select a banner style
3. Generate ASCII art
4. Export the generated art in either txt, json or xml format

## Error Handling

- Input validation ensures only supported characters are accepted
- Proper HTTP status codes and user-friendly error messages are provided

## Styling

- A gradient background with animation
- Custom styled form elements
- Hover and focus effects
- Custom scrollbars
- Loading animations

## Future Enhancements

- Add more ASCII art styles
- Implement custom fonts
- Improve UI/UX
- Add more export formats
- Implement user authentication

## Contributing

Cemvalot
Pkalliag
Gpatoula
