package javascript

const JAVASCRIPT_INTERFACE_SOURCE = `
const { argString, argBytes, argUint64, argUint32 } = require("orbs-client-sdk");

function getErrorFromReceipt(receipt) {
    const value = receipt.requestStatus === "BAD_REQUEST" ? receipt.executionResult : receipt.outputArguments[0].value;
    return new Error(value);
}

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
			throw getErrorFromReceipt(receipt);
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
			throw getErrorFromReceipt(receipt);
		}

		return receipt.outputArguments[0].value;
	}
}

module.exports = {
	{{.AppName}}
};
`