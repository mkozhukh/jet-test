# Export service to excel

### config.yaml

```yml
#enable disable watermarks in the export data
watermark: false
server:
  # service path prefix
  path: /excel
  # service port 
  port: "8096"

```

### Incoming data configuration

Service expects the POST.data parameter to present with serialized json configuration

```json
{
  "name":"filename",
  "columns":[
    { "width": 100 }, ...
  ],
  "header":[
    ["Some", "Header", "Here" ],
    ...
  ],
  "data":[
    [ 1, "text", "other"],
    ...
  ],
  "styles":[
    "css":{
      "major":{ "color": "#ff00ff", "background": "#0000AA", "fontSize": 14 }
    },
    "cells":[
      [1,2,"major"]
    ]
  ]
}
```

- all top level keys are optional
- recognized color formats
    - RRGGBB
    - #RRGGBB
    - AARRGGBB


### Output format

Service responds with xlsx file, name of file is defined by POST.data.name parameter