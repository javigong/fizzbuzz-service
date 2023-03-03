# General
- Create a simple application that consists of frontend(FE) and backend(BE)
- Submit a video showing how it works and the GitHub repository
- Let us know how long it took you to complete the task
- You can use any library

# Requirements
- It's a simple application showing FizzBuzz
- [Demo](https://www.figma.com/proto/E1pHkTtyT0q6dGugCyTQMD/Tech-Interview?node-id=1%3A2&scaling=min-zoom&page-id=0%3A1&starting-point-node-id=1%3A2)
- If a user pushes the buttons, the count number is incremented.
- If the number is a multiple of 3, FE shows the message "Fizz"
- If the number is a multiple of 5, FE shows the message "Buzz"
- If the number is a multiple of 3 and 5, FE shows the message "FizzBuzz". It doesn't mean showing total 3 messages "Fizz", "Buzz", and "FizzBunzz", but only "FizzBuzz".
- On pushing the button, FE sends BE the current count number
- BE returns a message like "Fizz", "Buzz", or "FizzBuzz". If the number isn't in the condition, it returns an empty string
- FE shows the returned message
- The count number is stored in FE, not BE. And it doesn't need to be persisted.
- Since FE and BE are supposed to be hosted on different servers, care about cross-origin resource sharing.

## Frontend
- SPA on React and TypeScript
- Follow [the Figma design] (https://www.figma.com/file/E1pHkTtyT0q6dGugCyTQMD/Tech-Interview?node-id=1%3A2&t=xtxraTKOhv2Z4UrV-1)
- The URL of BE server must be given as an environment variable.

## Backend
- API server on Golang
- Handle errors properly
- The messages like "Fizz", "Buzz" and "FizzBuzz" must be given as environment variables. Don't do hard coding.

## API
- Based on the RESTful way


```
openapi: 3.1.0
info:
  title: FizzBuzz Service
  version: '1.0'
  summary: FizzBuzz returning server
servers:
  - url: 'http://localhost:3000'
paths:
  /fizzbuzz:
    post:
      summary: Get a fizzbuzz message
      description: Get a fizzbuzz message
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                count:
                  type: number
                  description: A count number to get FizzBuzz message
      responses:
        '200':
          $ref: '#/components/schemas/DefaultResponse'
        '400':
          $ref: '#/components/schemas/DefaultResponse'
        '500':
          $ref: '#/components/schemas/DefaultResponse'
components:
  schemas:
    DefaultResponse:
      type: object
      properties:
        status:
          type: number
          example: 200
        message:
          type: string
      required:
        - status
        - message
```