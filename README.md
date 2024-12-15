
# MicGoFiber

The Microservice boilerplate for Go Fiber project

## Run Locally

Clone the project

```bash
  git clone https://github.com/waynemamahit/micgofiber
```

Go to the project directory

```bash
  cd micgofiber
```

Install dependencies

```bash
  go mod tidy
```

Migrate schema

```bash
  go run . migrate
```

Start the server

```bash
  go run .
```

Start the server with Air for live reload

```bash
  air
```

## Run Tests

Run test with verbose output

```bash
  go test ./... -v
```

Run test with coverage

```bash
  go test ./... -cover
```

Run test with coverage output

```bash
  go test ./... -coverprofile=coverage.out
```
```bash
  go tool cover -html=coverage.out
```

## Authors

- [@waynemamahit](https://www.github.com/waynemamahit)


## License

[MIT](https://github.com/waynemamahit/micgofiber/blob/main/LICENSE)

Copyright (c) 2023-Present, Waney Wanua Mamahit
