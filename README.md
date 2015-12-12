# Persefoneia

Persefoneia is tool for Rest API integration testing


## Mock Server

Listens on /admin path for management commands.

Admin API: 
* add endpoint
* delete endpoint
* add endpoint scenario
* add endpoint single response 
* get list of all endpoint calls (spy mode)

### Endpoint
Listening socket for simulation of single http API. Endpoint can work in two modes:
All requests are stored and can be inspected via /admin calls. 

**spy mode** - accepts any request and returns default response (200 OK). It is possible to predefine responses for specific
requests - response is selected according to first match algorithm. It would be nice if it will be possible to generate
Request-response pairs from API Blueprint.

**scenario mode** - stricly follows provided scenario, returns error status (406) for each call that
does not match scenario.



## Request definition 

```
{
    "name": "Get users",
    "request": {
        "url": "/users"
        "method": "GET"
    }
    "response": {
        "type": "application/json",
        "payload":[
            { "name": "Bela Fleck", "age": 33 },
            { "name": "Tony Rice", "age": 28 },
        ]
    }
```

## Scenario definition


```
{
    "name": "Scenario name",
    "steps": [ list of request definitions ]
}
```


## Rest tester

Tool for sending REST request and verification of responses.

Nice to have: generator for API Blueprint 
