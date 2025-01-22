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

## Running the application locally:

Initially, we need to configure a series of environment variables. To
generate the `.env` file, you can run the command:

```bash
make local.env
```

Examples of how some variables can be filled to run in a local environment:

```bash
ENV=dev
LOG_LEVEL=debug
HOST="127.0.0.1"
PORT="1323"
```

Don't worry, this `local.env` file is ignored by git and changes to it are not pushed to the repository.


To run the app locally on your machine, run the command:

```bash
make run
```

## Cobertura de código

```
make coverage
```
