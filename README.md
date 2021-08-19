# NPM License Checker to OFF Converter

Convert NPM License Checker output (JSON) to OFF format. (see [github.com/owasp/off](https://github.com/owasp/off))

## Running

1. Get an npm license checker report in json (eg. `license-checker --exclude 'MIT, BSD, BSD-2-Clause, Apache-2.0, CC0-1.0, ISC' --json > npm-licenses.json`)
1. `go get github.com/jemurai/npmlc2off`
1. `npmlc2off npm-licenses.json`

## Releasing

Npmlc2off works to follow golang best practices.  Therefore, when updating, we need to do the following:

- `go get -u` 
- `go mod tidy`
- `git commit -m "change with version"`
- `git tag v1.0.6`
- `git push origin v1.0.6`

Run the build.sh and get the different types of artifacts and include them in the release.
