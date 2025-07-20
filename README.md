# Go Event Booking REST API

This project is a Go-powered REST API for event booking, built as a practice project. It provides endpoints for managing events, users, and event registrations.

## Features

* **Event Management:**
    * Get a list of all available events.
    * Get details of a specific event by ID.
    * Create new bookable events (authentication required).
    * Update existing events (authentication required).
    * Delete events (authentication required).
* **User Management:**
    * Create new user accounts.
    * Authenticate users.
* **Event Registration:**
    * Register for an event (authentication required).
    * Cancel event registration (authentication required).

## API Endpoints

The API exposes the following endpoints:

| Method   | Endpoint               | Description                     | Authentication |
| :------- | :--------------------- | :------------------------------ | :------------- |
| `GET`    | `/events`              | Get a list of available events  | None           |
| `GET`    | `/events/:id`          | Get a list of available events  | None           |
| `POST`   | `/events`              | Create a new bookable event     | Required       |
| `PUT`    | `/events/:id`          | Update an event                 | Required       |
| `DELETE` | `/events/:id`          | Delete an event                 | Required       |
| `POST`   | `/signup`              | Create new user                 | None           |
| `POST`   | `/login`               | Authenticate user               | None           |
| `POST`   | `/events/:id/register` | Register user for event         | Required       |
| `DELETE` | `/events/:id/register` | Cancel registration             | Required       |

## Technologies Used

* **Go (Golang):** The primary programming language.
* **RESTful API Principles:** Adherence to REST architectural style.

## Setup and Running Locally

**(Further instructions would go here, e.g., how to clone the repo, install dependencies, set up a database if applicable, and run the application.)**

Example (placeholders):

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/PrinceM13/go-event-booking.git](https://github.com/PrinceM13/go-event-booking.git)
    cd your-repo-name
    ```
2.  **Install dependencies:**
    ```bash
    go mod download
    ```
3.  **Set up environment variables** (e.g., database connection string, JWT secret).
4.  **Run the application:**
    ```bash
    go run main.go
    ```

## Authentication

Endpoints marked "Required" in the API Endpoints table necessitate a valid authentication token (e.g., JWT) to be sent in the request header. Users can obtain this token by successfully logging in via the `/login` endpoint.

## Project Status

This is a practice project for learning and improving Go REST API development skills.

---
**Inspired by:** Maximilian Schwarzmüller - "Go - The Complete Guide" (Udemy)