# final-project
### Final Project of Software Engineering Bootcamp with Go Programming Language by The Go Dragons Team
```
.
├── cmd ------------------------------------> Entry point
│   └── main.go
├── config.yml
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── img ------------------------------------> Contains application logo
│   └── go-dragon.png
├── internal
│   ├── app --------------------------------> Initialization 
│   │   └── app.go
│   ├── domain -----------------------------> Encapsulate Business Rules (Entities Layer)
│   │   ├── airline.go
│   │   ├── airplane.go
│   │   ├── airport.go
│   │   ├── city.go
│   │   ├── flightClass.go
│   │   ├── flight.go
│   │   ├── gender.go
│   │   ├── order.go
│   │   ├── passenger.go
│   │   ├── payment.go
│   │   ├── role.go
│   │   ├── ticket.go
│   │   └── user.go
│   ├── interfaces -------------------------> Interface Adapters (Controller Layer)
│   │   ├── http ---------------------------> HTTP Handlers
│   │   │   ├── authorization.go
│   │   │   ├── booking_handler.go
│   │   │   ├── flights_handler.go
│   │   │   ├── general_functions.go
│   │   │   ├── login_handler.go
│   │   │   ├── logout_handler.go
│   │   │   ├── payment_handler.go
│   │   │   ├── role_handler.go
│   │   │   ├── root_handler.go
│   │   │   ├── signup_handler.go
│   │   │   └── ticket_handler.go
│   │   └── persistence --------------------> Provides Abstract Layer to Database
│   │       ├── airline_repository.go
│   │       ├── airport_repository.go
│   │       ├── city_repository.go
│   │       ├── flight_repository.go
│   │       ├── order_repository.go
│   │       ├── passenger_repository.go
│   │       ├── payment_repository.go
│   │       ├── role_repository.go
│   │       ├── ticket_repository.go
│   │       └── user_repository.go
│   └── usecase ----------------------------> The application logic (Usecase Layer)
│       ├── booking.go
│       ├── cancelation.go
│       ├── errors.go
│       ├── flights.go
│       ├── payment_gateways.go
│       ├── payment.go
│       ├── role_usecase.go
│       ├── signup_usecase.go
│       └── ticket.go
├── LICENSE
├── pdf ------------------------------------> Temporary ticket file as PDF to download
│   └── ticket.pdf
├── pkg ------------------------------------> External Frameworks and Drivers (Port|Infra)
│   ├── config -----------------------------> Configurations
│   │   ├── config.go
│   │   └── viper.go
│   ├── database ---------------------------> Databse Layer that contains GORM
│   │   ├── db.go
│   │   └── migrations ---------------------> SQL Migration files
│   │       ├── 000001_final_project_schema.down.sql
│   │       └── 000001_final_project_schema.up.sql
│   ├── encryption
│   │   └── encryption.go
│   ├── logger -----------------------------> Logger
│   │   └── logger.go
│   ├── mock_api  --------------------------> Encapsulate the mock API
│   │   └── fake_airline_info.go
│   ├── pdf --------------------------------> Encapsulate external libs that create PDF
│   │   └── ticket.go
│   └── seeder -----------------------------> Provides seeder for secure login management
│       └── seeder.go
├── README.md
└── wait-for-it.sh
```
<p style="text-align: center; width: 100%; ">Copyright&copy; 2023 <a href="https://github.com/the-go-dragons">The Go Dragons Team</a>, Licensed under MIT</p>
