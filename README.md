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
│   ├── database----------------------------> Databse Layer that contains GORM
│   │   └── db.go
│   └── logger -----------------------------> Logger
│       └── logger.go
└── README.md ------------------------------> Documentations
```
<p style="text-align: center; width: 100%; ">Copyright&copy; 2023 <a href="https://github.com/the-go-dragons">The Go Dragons Team</a>, Licensed under MIT</p>
