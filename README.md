# xarvis
A Microservices in go to bootstrap.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites
1. docker
2. docker-compose 

### Installing

Clone the Repo

```bash
go get https://github.com/alamin-mahamud/xarvis
cd $GOPATH/src/github.com/alamin-mahamud/xarvis
```

``` go
docker-compose up
```

## Endpoints

* List all
```
GET - http://localhost:8080/v1/authentication

``` 

``` 
# Status Code 200 
[
   {
      "id":"5c41bde20d021b7c6d7876ed",
      "username":"alamin-mahamud",
      "email":"alamin.ineedahelp@gmail.com",
      "password":"12345-Pa$$w0rd"
   },
   {
      "id":"5c41d2110d021b3a955365bd",
      "username":"jon-anton",
      "email":"jon.anton@gmail.com",
      "password":"Pa$$w0rd-D77m!"
   }
]
```

* Create a User
```
POST - http://localhost:8080/v1/authentication
------------------------------------------------
{
	"username": "alamin-mahamud",
  	"password": "simple-password",
  	"email": "alamin.ineedahelp@gmail.com"
}

```

``` go
# Status Code 201

```

* Get a user
```
GET - http://localhost:8080/v1/authentication/5c41bde20d021b7c6d7876ed

```
```
# Status Code - 200 OK
{"id":"5c41bde20d021b7c6d7876ed","username":"alamin-mahamud","email":"alamin.ineedahelp@gmail.com","password":"simple-password"}
```


```
PATCH - http://localhost:8080/v1/authentication/5c41bde20d021b7c6d7876ed
------------------------------------------------
{
  "username": "alamin-mahamud",
  "password": "12345-Pa$$w0rd",
  "email": "alamin.ineedahelp@gmail.com"
}

``` 

```
# Status Code - 200
{"result":"success"}
```


```
DELETE - http://localhost:8080/v1/authentication/5c41bde20d021b7c6d7876ed
```

```
# Status Code - 200
{"result":"success"}
```


## Test
[WIP]

## Deployment
[WIP]

## Built With
- with standard library

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/alamin-mahamud/xarvis/tags). 

## Authors

* **Alamin Mahamud** - *Initial work* - [alamin.rocks](https://alamin-rocks.herokuapp.com)

See also the list of [contributors](https://github.com/alamin-mahamud/xarvis/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration - vinta and other awesome repo guys.
