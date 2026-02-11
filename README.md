# Public holidays

![Go Tests](https://github.com/softwarespot/public-holidays/actions/workflows/go.yml/badge.svg)

This project showcases my proficiency in Go by creating a clear and readable
public holidays API. The primary focus is on writing clean, maintainable code
that effectively demonstrates the logic behind retrieving public holidays for a
country's ISO 3166-1 alpha-2 code.

## GET /holidays/v1/{country}/{year}

This endpoint retrieves a list of public holidays for a specified country and
year. By replacing `{country}` with the two-letter ISO 3166-1 alpha-2 code of
the desired country and `{year}` with the target year, users can access detailed
information about each holiday, including dates, local names and English names.
This API is useful for applications that need to display or utilize holiday data
for various purposes, such as scheduling or planning.

## Hosted by [render.com](https://render.com/)

This API is available at https://public-holidays.onrender.com.

**IMPORTANT**
The instance might be down due to inactivity, therefore, wait about 50s for the instance to be started again.

### Example request

```bash
curl -s "https://public-holidays.onrender.com/holidays/v1/FI/2024" | jq

# or locally

curl -s "http://localhost:10000/holidays/v1/FI/2024" | jq
```

### Response format

The response will be in JSON format and includes an array of holiday objects
with the following details:

- **date**: The date of the holiday.
- **name**: The name of the holiday in the local language.
- **englishName**: The name of the holiday in English.

```json
[
    {
        "date": "2024-01-01",
        "name": "Uudenvuodenpäivä",
        "englishName": "New Year's Day"
    },
    {
        "date": "2024-01-06",
        "name": "Loppiainen",
        "englishName": "Epiphany"
    },

    ...
]
```

## Countries

Here is a list of the supported countries.

| Country    | ISO 3166-1 alpha-2 | Wikipedia Link                                                                         |
| ---------- | ------------------ | -------------------------------------------------------------------------------------- |
| 🇩🇰 Denmark | DK                 | [Public Holidays in Denmark](https://en.wikipedia.org/wiki/Public_holidays_in_Denmark) |
| 🇫🇮 Finland | FI                 | [Public Holidays in Finland](https://en.wikipedia.org/wiki/Public_holidays_in_Finland) |
| 🇮🇸 Iceland | IS                 | [Public Holidays in Iceland](https://en.wikipedia.org/wiki/Public_holidays_in_Iceland) |
| 🇳🇴 Norway  | NO                 | [Public Holidays in Norway](https://en.wikipedia.org/wiki/Public_holidays_in_Norway)   |
| 🇸🇪 Sweden  | SE                 | [Public Holidays in Sweden](https://en.wikipedia.org/wiki/Public_holidays_in_Sweden)   |

## Prerequisites

- Go 1.26.0 or above
- make (if you want to use the `Makefile` provided)
- Docker

## Dependencies

**IMPORTANT:** No 3rd party dependencies are used.

I could easily use [Cobra](https://github.com/spf13/cobra) (and usually I do,
because it allows me to write powerful CLIs), but I felt it was too much for
such a tiny project. I only ever use dependencies when it's say an adapter for
an external service e.g. Redis, MySQL or Prometheus.

## Setup

1. Create and edit the `.env` (used when developing locally) and `.env.production` (used when deploying to production)

```bash
cp .env.example .env
cp .env.example .env.production
```

## Run not using Docker

```bash
go run .
```

or when using `make`

```bash
make

./bin/public-holidays
```

### Version

Display the version of the application and exit.

```bash
# As text
./bin/public-holidays --version

# As JSON
./bin/public-holidays --json --version
```

### Help

Display the help text and exit.

```bash
./bin/public-holidays --help
```

## Run using Docker

1. Build the Docker image with the tag `public-holidays`.

```bash
docker build -t public-holidays .
```

2. Run the Docker image.

```bash
# Port number should the same as defined in ".env.production"
docker run -p "10000:10000" --rm public-holidays
```

### Version

Display the version of the application and exit.

```bash
# As text
docker run --rm public-holidays --version

# As JSON
docker run --rm public-holidays --json --version
```

### Help

Display the help text and exit.

```bash
docker run --rm public-holidays --help
```

## Tests

Tests are written as [Table-Driven Tests](https://go.dev/wiki/TableDrivenTests).

```bash
go test -cover -v ./...
```

or when using `make`

```bash
make test
```

### Linting

Docker

```bash
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run --tests=false --default=none -E durationcheck,errorlint,exhaustive,gocritic,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```

Local

```bash
golangci-lint run --tests=false --default=none -E durationcheck,errorlint,exhaustive,gocritic,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```

## Additional information

This section documents any additional information which might be deemed important for the reviewer.

### Decisions made

- Despite using 1.26.0+ and the `slices` pkg being available, I have opted not
  to use it, and instead went for how I've been writing Go code before the
  `slices` pkg existed. Although for production code, I have started to use it
  where applicable.
- Loosely used https://jsonapi.org/.

### License

The code has been licensed under the [MIT](https://opensource.org/license/mit) license.
