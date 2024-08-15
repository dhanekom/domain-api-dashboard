# domain-api-dashboard

## Installation

* Install [Go](https://go.dev/doc/install) - run ```$ go version``` afterwards in a terminal to verify installation
* Install [NodeJS](https://nodejs.org/en/download/package-manager/current)
* Run ```$ go mod tidy``` to pull all the project dependencies
* Install Air by running ```$ go install github.com/air-verse/air@v1.52.3```
  * Air run your application and rebuilds it if there are any changes in your code


## Running program
* Open a terminal and navigate to the project's root directory
* Make sure your styles (in ./pubic/css/styles.css) is up to date. Anytime you use previously unused Tailwind styles in a html template Tailwind will have to rebuild the styles.css file. Use one of the following methods:
  * ```$ npm run watch-css``` (recommended) - Watches for changes and automatically updates styles.css. Run this in a separate terminal session
  * ```$ npm run build-css``` - performs a once off update of styles.css
* Run ```$ air```