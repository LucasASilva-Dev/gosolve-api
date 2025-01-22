### Goal of this REPO:
The aim of this task is to create working rest service with just one endpoint.
To implement API, you can use any of existing Go libraries instead of creating it from the scratch.

`Summarize`:
- Design API for http `GET` method
- Implement functionality for searching `index` for `given` value (it should be the most efficient algorithm) 
- Add logging
- Add possibility to use configuration file where you can specify service port and log level (you should be able to choose between Info, Debug, Error)
- Add `unit tests` for created components
- Add `README.md` to describe your service
- Automate running tests with `make` file
- Remember that code structure matters
- Upload solution into `GitHub` account and share the link

Sample input file is added as `input.txt` file.

## Installing the project

Assuming that you already have GoLang working in version 1.23.5 in your envoriment.
```
make install
```

## Building the project
```
make build
```

## Running code QA tools + testing

```
make test
```

## Code coverage

```
make coverage
```

## Running the application locally:

To run the app locally on your machine, run the command:

The API have two internal servers, one for the API itself and another one for Prometheus metrics.

The prometheus metrics will start at port `8380` in endpoint `/metrics`

The default port for server is `1323`. You can change that by runing in another port

The API is configured to use commands to start the server and have options within.


Initially, we need to configure a series of environment variables. To
generate the `.env` file, you can run the command:

```bash
make local.env
```

Don't worry, this `local.env` file is ignored by git and changes to it are not pushed to the repository.


To facilite the run, the command `make run` load the default parameters, but fill free to change it.

```bash
make run
```