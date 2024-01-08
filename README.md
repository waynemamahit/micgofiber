
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

Build CLI generator

```bash
  go build ./cmd/; ./micgo
```

Migrate schema

```bash
  ./micgo migrate
```

Start the server

```bash
  go run main.go
```

Start the server with Air

```bash
  air
```


## Authors

- [@waynemamahit](https://www.github.com/waynemamahit)


## License

[MIT](https://github.com/waynemamahit/micgofiber/blob/main/LICENSE)

Copyright (c) 2023-Present, Waney Wanua Mamahit
