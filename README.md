# Riccardo Ossola's Code Review

This Coupon Service is a program that allows users to create, apply 
and check existing coupons.

## Prerequisites
Before running the Coupon Service,
ensure you have the following software installed:
- Go (v1.16+)
- Docker

## Usage

Navigate to the project directory:

`cd go-code-project/review/cmd/coupon_service`

Run the code:

`go run main.go`


## API Endpoints
The Coupon Service provides the following API endpoints:

- **GET** `/api/coupons`: Get the collection of all existing coupons in the database.
- **POST** `/api/coupon/:code` : Get the corresponding coupon to the provided code, if it exists.
  
    Request Example: `localhost:8080/api/coupon/SAVE50`
- **POST** `/api/create` : Allows the user to create a new coupon.

    Request Body Example:
    ```json
    {
      "discount": 50,
      "code": "SAVE50",
      "min_basket_value": 100
    }
    ```
- **POST** `/api/apply` : Allows the user to apply an existing coupon to his basket.

  Request Body Example:
    ```json
    {
      "code": "SAVE50",
      "basket": {
        "value": 100
      }  
    }
    ```

Everything was user tested on http://localhost:8080 through the Postman Agent.
## Testing
To run tests, use the following command:

`go test ./...`

## Changes

The current version of the application provides basic functionality for managing coupons
exposing API routes and storing data in a MySQL database.

1. #### Code smells

   Every out of place thing that I was able to find, bad or broken
   implementation has been fixed (both in the Go codebase and the Dockerfile)

2. #### Database

   I decided to change the map used as a database to a real
   DB implementation to store the coupon data.

3. #### API

   I divided the `/api/coupons` that initially was designed to receive a list of
    codes, to `/api/coupon/:code` to search for a specific one, and `/api/coupons` that
    return the list of all coupons in the DB (I may have have misunderstood 
    scope of the api here).
    
    In the controller all of the possible HTTP errors 
    encountered when calling the APIs should be handled.

4. #### Test

   I fixed the existing tests, now relying on a in-memory DB.