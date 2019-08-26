package javascript

const PACKAGE_JSON_SOURCE = `{
  "name": "{{.AppNameLowercase}}",
  "version": "1.0.0",
  "main": "index.js",
  "scripts": {
    "gamma:start": "gamma-cli start-local",
    "gamma:stop": "gamma-cli stop-local",
    "{{.AppNameLowercase}}:local": "node ./src/deploy.js",
    "test": "mocha test"
  },
  "devDependencies": {
    "expect.js": "^0.3.1",
    "mocha": "^6.1.4"
  },
  "dependencies": {
    "orbs-client-sdk": "^1.0.0",
  }
}`