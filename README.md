# SANSA Space Science
This repository contains packages required for the acquisition, processing and persistence of data from all SANSA Space Science instruments.
 
## DAQ and Ingress for Lemi011 - Proof of Concept
Some things to take note of first:
* The app will read in all parameters from environment variables, as prescribed by 12-factor microservices architecture.
* For testing purposes, environment variables are set in the cmd/daq/config.go and cmd/ingress/config.go files - this will be removed when going into production.
* To test the application: connect a lemi011 sensor or simulator (I used an Arduino Uno) to the PC, check the COM port and change the setting in the config.go file.
* A Nats Streaming Server needs to be running somewhere as well. Adjust Nats ID variables accordingly.

To run DAQ:
```bash
cd cmd/daq
go run .
```

To run Ingress:
```bash
cd cmd/ingress
go run .
```

### Architecture
The following principles were taken into account to optimise the code for maintainability, readability, documtation and robustness:
* Clean Architecture
* Hexagonal Architecture (an interpretation of clean architecture)
* Domain Driven Design
* 12-factor microservices architecture

### Instruments
Below is a task list of all SANSA Space Science instruments that needs to be implemented. A lot of work still needs to be done.
- [x] Lemi-011
- [ ] NovAtel
- [ ] Ashtech
- [ ] Lemi-025
- [ ] DVRAS
- [ ] UltraMSK
- [ ] Wide Angle Riometer
- [ ] DTU
- [ ] Overhauser

### To Do
A list of functionality that needs to be added.
- [ ] Checksums for files
- [ ] Durable subscriptions with at-least-once publishers
- [ ] Postgres/timescaledb
- [ ] More defined API - GraphQL/REST/gRPC
- [ ] Dockerise
- [ ] Dynamic configuration
- [ ] Improve logging
