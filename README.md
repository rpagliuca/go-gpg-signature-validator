# GPG signature validation

## About

Validates a signed GPG message using Golang.

## How to use

Send an HTTP request to the microservice containing a JSON with the properties ArmoredMessage and ExpectedSignatureKeyId.
The service will respond with a JSON containing the original message contents and whether the real signature matches the expected siganture or not.

    $ curl localhost:8080 -d '{"ArmoredMessage":"-----BEGIN PGP MESSAGE-----<MESSAGE_CONTENTS...>", "ExpectedSignatureKeyId":"C94AAC02A4D11A51FCC92DCD77F41A0A51E58C75"}'

    {"messageContents":"Who else but myself?\n","signatureIsValid":true}

## Presentation slides

Link to the presentation: https://docs.google.com/presentation/d/1kIDikOH5Rr0FjyQdUGpmszWIdL_Bqe2pq2RYFcw6T38/edit?usp=sharing
