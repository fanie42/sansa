This repository contains packages required for the acquisition, processing and persistence of data from all SANSA Space Science instruments.
 
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
