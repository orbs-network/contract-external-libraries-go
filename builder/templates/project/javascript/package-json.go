package javascript

const PACKAGE_JSON_SOURCE = `{
  "name": "{{.AppNameLowercase}}",
  "version": "1.0.0",
  "main": "index.js",
  "scripts": {
    "gamma:start": "gamma-cli start-local -wait",
    "gamma:stop": "gamma-cli stop-local",
    "{{.AppNameLowercase}}:local": "node ./src/deploy.js",
    "test": "mocha test --timeout 20000 --exit"
  },
  "devDependencies": {
    "expect.js": "^0.3.1",
    "mocha": "^6.1.4"
  },
  "dependencies": {
    "orbs-client-sdk": "^1.0.0"
  }
}`