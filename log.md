# PTV-tui

## swawgger.json

/definitions are missing in file from ptv swagger.

V3.Period and V3.Operator are missing in #/definitions/

referenced in the .json here:

V3.PeriodsResponse/properties/items/\$ref
V3.OperatorsResponse/properties/items/$ref

V3.RouteResponse was borked

properties/routes was originally property/route which led to nil return

properties/routes was also originally not an array.

```"V3.RouteResponse": {
"V3.RouteResponse": {
      "type": "object",
      "properties": {
        "route": 
        "$ref": "#/definitions/V3.RouteWithStatus",
        "description": "Train lines, tram routes, bus routes, regional coach routes, Night Bus routes"
          }
        },
        "status": {
          "$ref": "#/definitions/V3.Status",
          "description": "API Status / Metadata"
        }
      }
    },
```

```



```"V3.RouteResponse": {
"V3.RouteResponse": {
      "type": "object",
      "properties": {
        "routes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/V3.RouteWithStatus",
            "description": "Train lines, tram routes, bus routes, regional coach routes, Night Bus routes"
          }
        },
        "status": {
          "$ref": "#/definitions/V3.Status",
          "description": "API Status / Metadata"
        }
      }
    },
```
## TODO

- tabs in UI for different tables
  - each stop in area
    - depatures for each route through each stop
- user input in UI
  - select parameters
    - routeType
    - stop
    - direction
- tidy packages