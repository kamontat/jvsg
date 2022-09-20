# Mock server for http request on large response

Default host: `localhost:3333`

Endpoints:
1. `/mirror` - [POST] Response with request body
2. `/json` - [POST] Response with large json body from fs (26MB)
3. `/json/cached` - [POST] Response with large json body from memory (26MB)
