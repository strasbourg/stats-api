# Strasbourg Stats API

L'API dite stats est là pour logguer les requêtes faites aux autres APIs.

## Specs

### Track request

    POST /track

nom          | type     | utilisation                                     |
-------------|----------|-------------------------------------------------|
`ip`         | `string` | adresse ip du client                            |
`service`    | `string` | nom du service parmi `cts`, `velhop`, `parkings`|
`path`       | `string` | exemple `/stops/322`                            |
`params`     | `string` | exemple `time=1234&test=1`                      |
`time`       | `string` | exemple `2013-03-02T09:39:57.321493+02:00`      |

#### Response

    Status 200 OK

```json
{
  "ip": "0.0.0.0",
  "service": "cts",
  "path": "/stops/322",
  "params": "time=1234&test=1",
  "time": "2013-03-02T09:39:57.321493+02:00"
}
```
