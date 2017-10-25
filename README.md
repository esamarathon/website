# ESA Marathon
The greatest speedrunning event Europe has ever seen

## Requirements
* Go
* Docker / RethinkDB
* Gulp
* Npm

## Local development
1. Set up your local .env file base don the .example-env
2. Download dependencies with `go get` and `npm install` (or `yarn`)
3. Run `docker-compose up -d` to initialize the Database
4. Migrate and seed the DB by running `go run dbinit.go`
5. Compile styles with `npm run gulp`
6. Run [`fresh`](https://github.com/pilu/fresh) or `go run main.go`
7. Add some cool features