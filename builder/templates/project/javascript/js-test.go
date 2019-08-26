package javascript

const JAVASCRIPT_TEST_SOURCE = `
const expect = require("expect.js");
const { Client, createAccount, NetworkType } = require("orbs-client-sdk");
const { {{.AppName}} } = require("../src/{{.AppNameLowercase}}");
const { deploy, getClient } = require("../src/deploy");

describe("{{.AppName}}", () => {
    it("updates contract state", async () => {
		const contractOwner = createAccount();
		const contractName = "{{ .AppName }}" + new Date().getTime();

		await deploy(getClient(), contractOwner, contractName);
		const {{.AppNameLowercase}} = new {{.AppName}}(getClient(), contractName, contractOwner.publicKey, contractOwner.privateKey);

		const defaultValue = await {{.AppNameLowercase}}.value();
		expect(defaultValue).to.be.eql(0);

		const updatedValue = await {{.AppNameLowercase}}.add(7);
		expect(updatedValue).to.be.eql(7);
	});
});
`
