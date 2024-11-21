# Airline Ticket Management System
### Final Project of Software Engineering Bootcamp with Go Programming Language by The Go Dragons Team

## Overview
Welcome to the **Airline Ticket Management System**! This project is designed to manage airline ticket bookings efficiently. Built using Go (Golang), it leverages modern software engineering practices to ensure robust and scalable management of airline tickets.

## Features

- **Booking Management**: Create, view, update, and delete airline ticket bookings.
- **User Authentication**: Secure user login and registration system.
- **Flight Information**: Manage flight schedules and availability.
- **Payment Processing**: Integration with payment gateways for ticket purchases.
- **Reporting**: Generate detailed reports on bookings and flights.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/dl/) (1.16 or later)
- [MySQL](https://www.mysql.com/) or any other supported database
- [Docker](https://www.docker.com/) (optional, for containerization)

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/ArdeshirV/Airline-Ticket-Management.git
    cd Airline-Ticket-Management
    ```

2. **Set up the database:**

    Ensure you have a running instance of MySQL (or your preferred database).

    ```sh
    mysql -u root -p
    CREATE DATABASE airline_ticket_management;
    ```

3. **Configure the environment:**

    Create a `.env` file in the root directory and configure your database connection and other environment variables:

    ```env
    DB_USER=root
    DB_PASSWORD=yourpassword
    DB_NAME=airline_ticket_management
    DB_HOST=localhost
    DB_PORT=3306
    ```

4. **Install dependencies:**

    ```sh
    go mod tidy
    ```

5. **Run the application:**

    ```sh
    go run main.go
    ```

## Usage

Once the application is running, you can use the following endpoints to interact with the system:

- **Booking Management:**
  - `GET /bookings` - View all bookings
  - `POST /bookings` - Create a new booking
  - `GET /bookings/{id}` - View a specific booking
  - `PUT /bookings/{id}` - Update a booking
  - `DELETE /bookings/{id}` - Delete a booking

- **User Authentication:**
  - `POST /register` - Register a new user
  - `POST /login` - Login a user

- **Flight Information:**
  - `GET /flights` - View all flights
  - `POST /flights` - Create a new flight
  - `GET /flights/{id}` - View a specific flight
  - `PUT /flights/{id}` - Update a flight
  - `DELETE /flights/{id}` - Delete a flight

- **Payment Processing:**
  - `POST /payments` - Process a payment

## Testing

To run tests, use the following command:

```sh
go test ./...
```

## Contributing
Contributions are welcome! Please fork this repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact
For any questions or inquiries, please contact ArdeshirV at [e-job(at)protonmail.com](e-job@protonmail.com).

<p style="text-align: center; width: 100%; ">Copyright&copy; 2023 <a href="https://github.com/the-go-dragons">The Go Dragons Team</a>, Licensed under MIT</p>
