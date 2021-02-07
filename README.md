# SRE Challenge
## Simple Checkout Basket in golang

A simple checkout HTTP API that contains below operations:
- Create a new checkout basket
- Add a product to a basket
- Get the total amount in a basket
- Remove the basket

The store only sells following 3 products:

```bash
Code         | Name              |  Price
-----------------------------------------------
PEN          | Lana Pen          |   5.00€
TSHIRT       | Lana T-Shirt      |  20.00€
MUG          | Lana Coffee Mug   |   7.50€
```

Various departments have insisted on the following discounts:

- The sales department thinks a buy 2 get 1 free promotion will work best (for each two of the same product, get the second free), and would like this to only apply to PEN items.
- The CFO insists that the best way to increase sales is with discounts on bulk purchases (buying x or more of a product, the price of that product is reduced), and requests that if you buy 3 or more TSHIRT items, the price per unit should be reduced by 25%.

Examples:

```bash
Items: PEN, TSHIRT, MUG
Total: 32.50€

Items: PEN, TSHIRT, PEN
Total: 25.00€

Items: TSHIRT, TSHIRT, TSHIRT, PEN, TSHIRT
Total: 65.00€

Items: PEN, TSHIRT, PEN, PEN, MUG, TSHIRT, TSHIRT
Total: 62.50€
```

## Requirements
- Golang v1.10.0+
- Docker

## Installation

### Manual Installation (local)

#### (Optional environment variable)
The api using a simple json db (scribble golang module). And it has an optional environment variable "DB_NAME".
Default value for DB_NAME=basket_db, if you want to change it into something else you can set the variable

```bash
$ export DB_NAME=<your-db-name>
```

```bash
$ git clone https://github.com/stvnorg/golang-checkout-basket
$ cd golang-checkout-basket
$ export DB_NAME=<LOCAL_JSON_DB> (optional)
$ go run api.go basket.go discounts.go
```
If you run it the first time, it will take some times for the go dependency manager to install the libraries needed. So, be patient! wait a little bit :)
The API has Prometheus enabled. It will listen/expose 2 TCP ports, 8080 (the API itself) and 9900 (prometheus exporter)

### Docker run
The docker image for the api is hosted in docker hub https://hub.docker.com/repository/docker/s7even/golang-checkout-basket
to run it simply by executing below command (without prometheus exposed):

```bash
$ docker run -p 8080:8080 -d s7even/golang-checkout-basket:0.2.0
```

and if you want to expose both with prometheus port, just run below command:

```bash
$ docker run -p 8080:8080 -p 9900:9900 -d s7even/golang-checkout-basket:0.2.0
```
Browse the http://localhost:8080/basket/total, to see if the API is running well

## Unit Test

To run the unit test (verbose):
```bash
$ go test -v
```

Unit test with coverage:
```bash
$ go test -cover
```

## Usage

Before you able to add products, you need to create a basket first

#### Create basket
```bash
$ curl -X POST http://<API_URL>/basket/create
```
Example:
```bash
$ curl -X POST http://localhost:8080/basket/create
```

#### Add Product to basket
```bash
$ curl -X POST http://<API_URL>/basket/add_product/<pen|mug|tshirt>
```
Example:
```bash
$ curl -X POST http://localhost:8080/basket/add_product/pen
```

#### Get the total amount in basket
```bash
$ curl -X GET http://<API_URL>/basket/total
```
Example:
```bash
$ curl -X GET http://localhost:8080/basket/total
```

#### Remove the basket
```bash
$ curl -X DELETE http://<API_URL>/basket/delete
```
Example:
```bash
$ curl -X DELETE http://localhost:8080/basket/delete
```

## CI Pipeline process
There are 2 (two) githubActions CI pipeline:

`.github/workflows/go.yml`
- Developers create a Pull Request and/or Merged the PR
- The CI will run the build and unit test

`.github/workflows/release.yml`
- Developers push new tags with semVer versioning X.Y.Z (ex: 1.0.0)
- The CI pipeline will login to docker hub
- Build the image with tagging `s7even/golang-checkout-basket:<X.Y.Z>`
- Push the image to docker hub

## (Optional) Prometheus Monitoring

The prometheus exporter is running on port 9900
If you choose to enabling it in docker run, make sure you add the correct targets in your `prometheus.yml` 

Example:
```bash
global:
  scrape_interval:     15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: prometheus
    static_configs:
      - targets: ['192.168.0.10:9900']
```
