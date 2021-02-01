## Sendsei Golang Test

Here is a simple test of your understanding of an API using an RPC client. The API will accept a json body and validate it then call an RPC server which then in turn sends an SMS using the Sendsei REST API.

You are required to complete the missing parts of this repository such that the following command will yield a successful sms sent.

```
curl -H 'x-api-key: xxx' -X POST localhost:11701/sms \
--data-raw '{
    "message": "Report to the ready room!",
    "sender": "Enterprise",
    "recipient": "61xxxxxxxxx"
}'

```

The test code will be run from the root of the repository with

```
API_KEY=xxx go run main.go
```

It is intended for this test to take roughly an hour to complete. You are free to take your time with it though be mindful of keeping it simple.

### Criteria

- Add a http handler to the api for the above request

- The api should validate that `message` does not have a character length longer than 3 SMS.

- Use your own mobile number for testing, an api key will be given to you

- The api should return data in the structure defined for you

- The http handler should initialise the already defined rpc client and use the method `SendSMS(sender, recipient, message)`

- You are required to create the SMS service method this rpc client is attempting to call

- Use the [Sendsei Beta API docs](https://documenter.getpostman.com/view/10746805/Szzei1PZ#intro) to submit an SMS from the service method

- Use only the Go standard library, no external packages

### Bonus

- Tests

### How to submit

Clone this and push to a private github repository and then add @<INSERT YOU> as a collaborator when completed.

