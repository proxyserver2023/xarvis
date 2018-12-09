# go-bootstrap

Microservices based REST API Implementation. where services talks with each other with protocol buffers.

## Concepts & Tools

1. Clean Architecture
2. Dependency Injection with Mocking Exampless
3. Following SOLID Principles.

## Objectives

1. Create decoupled system
2. Implementation of lower level domain is not concerned with the implementor, and can be replaced w/o having concern of breaking implementor function

## Aims

1. Independent of Frameworks
2. Highly Testable
3. Independent of Database
4. Independent of 3rd party libs

It has simple dependencies:

- [Chi (Router)](https://github.com/go-chi/chi)
- [Testify (Test & Mock framework)](https://github.com/stretchr/testify)
- [Mockery (Mock generator)](https://github.com/vektra/mockery)
- [Hystrix-Go (Circuit Breaker)](https://github.com/afex/hystrix-go)

Get Started:

- [Install](https://github.com/alamin-mahamud/go-bootstrap/#install)
- [Introduction](https://github.com/alamin-mahamud/go-bootstrap/#introduction)
- [Folder Structure](https://github.com/alamin-mahamud/go-bootstrap/#folder-structure)
- [Depency Injection](https://github.com/alamin-mahamud/go-bootstrap/#dependency-injection)
- [Mocking](https://github.com/alamin-mahamud/go-bootstrap/#mocking)
- [Testing](https://github.com/alamin-mahamud/go-bootstrap/#testing)
- [Circuit Breaker](https://github.com/alamin-mahamud/go-bootstrap/#circuit-breaker)

## [Install](https://github.com/alamin-mahamud/go-bootstrap/#install)

Clone the source

    git clone https://github.com/irahardianto/service-pattern-go

Setup dependencies

    go mod download
    # or
    go get -u github.com/go-chi/chi
    go get -u github.com/jinzhu/gorm
    go get github.com/stretchr/testify
    go get github.com/vektra/mockery/.../
    go get github.com/afex/hystrix-go/hystrix
    go get -u github.com/mattn/go-sqlite3

Setup sqlite data structure

    sqlite3 /var/tmp/tennis.db < setup.sql

Test first for your liking

    go test ./... -v

Run the app

    go build && ./service-pattern-go

And visit

    http://localhost:8080/getScore/Rafael/vs/Serena
