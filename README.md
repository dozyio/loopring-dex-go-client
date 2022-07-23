# Loopring DEX-API-V3 Go Client

Uses openapi-generator-cli to generate a Go API client for Loopring

## Inital Setup & Regenerating Go Client
Clone https://github.com/Loopring/DEX-API-V3

run `python3 xdoc.py refresh` to download latest swagger spec

Copy swagger spec to this repo's directory

run `cp ../DEX-API-V3/meta/swagger_en.json .`

run `builder.sh`

## Fixes from DEX-API-V3 swagger
* Replaces numbers "0L" with 0
* Replaces 0 status codes with 200
* Replaces "default": "None" with 0 (integer), "" (string), false (boolean) 
* Appends V2, V3 and HebaoV3 to operationId to ensure they are unique

## Todo
* A number of response models are broken
* getToken - should return array of tokens, not token
* Most Amm functions don't work at present
* Error codes
