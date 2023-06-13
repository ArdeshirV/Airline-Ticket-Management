# final-project
### Final Project of Software Engineering Bootcamp with Go Programming Language by The Go Dragons Team
```
.
├── cmd ------------------------------------> Entry point
│   └── main.go
├── internal
│   ├── app --------------------------------> Initialization 
│   │   └── app.go
│   ├── domain -----------------------------> Encapsulate Business Rules (Entities Layer)
│   │   ├── booking.go
│   │   ├── airline.go
│   │   ├── airport.go
│   │   ├── city.go
│   │   ├── flightClass.go
│   │   ├── flight.go
│   │   ├── gender.go
│   │   ├── passenger.go
│   │   ├── payment.go
│   │   ├── role.go
│   │   ├── ticket.go
│   │   └── user.go
│   ├── interfaces -------------------------> Interface Adapters (Controller Layer)
│   │   ├── http ---------------------------> HTTP Handlers
│   │   │   ├── booking_handler.go
│   │   │   └── flight_search_handler.go
│   │   └── persistence --------------------> Provides Abstract Layer to Database
│   │       ├── booking_repository.go
│   │       └── flight_repository.go
│   └── usecase ----------------------------> The application logic (Usecase Layer)
│       ├── booking.go
│       ├── cancelation.go
│       └── flight_search.go
├── pkg ------------------------------------> External Frameworks and Drivers (Port|Infra)
│   ├── config -----------------------------> Configurations
│   │   └── config.go
│   └── logger -----------------------------> Logger
│       └── logger.go
└── README.md ------------------------------> Documentations
```
