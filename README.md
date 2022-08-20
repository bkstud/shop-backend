# Backend application for shop
## Main features:
 - Third party OAuth with Google, Facebook, GitHub with custom bearer users / tokens kept in database
 - Integration with stripe checkout that pulls proper items from database using just items IDs provided
 - CRUD models for basket, feedback, transaction, item

## Starting app:
For the app to work correctly, proper secrets need to be provided either in `secrets/` directory or via environment variables.
### List of files from `secrets/` directory and corresponding variables starting with `VAR_ID_` and `VAR_SECRET_`
- facebook-creds.json or `FACEBOOK`
- github-creds.json or `GITHUB`
- google-creds.json or `GOOGLE`
- stripe-creds.json or `STRIPE` (for stripe "cid" or VAR_ID_STRIPE should be created, but left empty)

Example: `VAR_ID_FACEBOOK=abcdefg`

.json template:
```
{
    "cid":"yourapplicationid",
    "csecret":"your application secret"
}
```
### Running:
Run / build app in production or local mode, or use docker image:

Local run scenario:
```
go get ./...
go run -tags local server.go
```
Local build runs on HTTPS and will require custom certificate / key to work properly. <br>
See `GetCertAndKey()` from `server.go`

Production build scenario:
```
go get ./...
go build -tags production
./shop
```
The production build is desired to be run on azure cloud that provides HTTPS proxy containers, so the application will be started on port 80. <br>
See `config/config_production.go`.

## Tests:
To run endpoint tests, start or build server with test tag:
`go run -tags test server.go` <br>
Run tests: <br>
`go test shop/api/v1/test  -tags test -v -failfast`


