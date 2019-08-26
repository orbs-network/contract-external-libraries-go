package javascript

const JS_INTERFACE_SOURCE = `{
const { argString, argBytes, argInt64, argInt32, 
	createTransaction, sendTransaction, sendQuery } = require("orbs-network-sdk");

class {{.AppName}} {
	constructor(orbsClient, contractName, publicKey, privateKey) {
		this.client = orbsClient;
		this.contractName = contractName;
		this.publicKey = publicKey;
		this.privateKey = privateKey;
	}

	async add(n) {
		const [ tx, txId ] = this.client.createTransaction(
			this.publicKey, this.privateKey, this.contractName,
			"add",
			[
				argUint64(n)
			]
		);

		const receipt = await this.client.sendTransaction(tx);
		if (receipt.executionResult !== 'SUCCESS') {
			throw new Error(receipt.outputArguments[0].value);
		}

		return receipt.outputArguments[0].value;
	}

	async value() {
		const query = this.client.createQuery(
			this.publicKey,
			this.contractName,
			"value",
			[]
		);

		const receipt = await this.client.sendQuery(query);
		if (receipt.executionResult !== 'SUCCESS') {
			throw new Error(receipt.outputArguments[0].value);
		}

		return receipt.outputArguments[0].value;
	}
}

module.exports = App;
`