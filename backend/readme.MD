# Backend Service Docs

## End Point

| No  |Method|Endpoint|Payload (Example)|
|:---:|:---:|:---:|:---:|
|   1| POST |localhost:8080/api/temperature            |```json{"month": 4, "year": 2022}```|
|   2| POST |localhost:8080/api/temperature/download   |```json{"image":{base64}, "month": 4, "year": 2022}```   |