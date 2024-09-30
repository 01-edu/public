# Site Reliability Engineer - Basic API Hosting Task

## Part 1 – The Web Service

Write a web service in any language that takes in a JSON payload, does
some basic validation against an expected message format and content,
and then puts that payload into a queue of your choice or a file.

Example valid payload:
```json
{
    "ts": "1530228282",
    "sender": "curler-user",
    "message": {
        "foo": "bar",
        "hash": "bash"
    },
    "sent-from-ip": "1.2.3.4"
}
```

Validation rules:
● “ts” must be present and a valid Unix timestamp
● “sender” must be present and a string
● “message” must be present, a JSON object, and have at least one
field set
● If present, “sent-from-ip” must be a valid IPv4 address
● All fields not listed in the example above are invalid, and
should result in the message being rejected.

## Part 2 – Terraform

Deploy this application to your favourite cloud provider using Terraform.

## Part 3 – NewRelic

Implement NewRelic monitoring for this application using Terraform.