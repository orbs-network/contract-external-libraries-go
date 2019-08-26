package javascript

const JAVASCRIPT_DEPLOYMENT_SOURCE = `
const { readFileSync } = require("fs");
const {
	Client, createAccount,
	PROCESSOR_TYPE_NATIVE, NetworkType
} = require("orbs-client-sdk");

function getClient() {
    const endpoint = process.env.ORBS_NODE_ADDRESS || "http://localhost:8080";
    const chain = Number(process.env.ORBS_VCHAIN) || 42;
    return new Client(endpoint, chain, NetworkType.NETWORK_TYPE_TEST_NET);
}

function getContractCode() {
	return [
		readFileSync(__dirname + "/contract/contract.go")
	];
}

async function deploy(client, contractOwner, contractName) {
    const [tx, txid] = client.createDeployTransaction(contractOwner.publicKey, contractOwner.privateKey,
		contractName, PROCESSOR_TYPE_NATIVE, ...getContractCode());
    const receipt = await client.sendTransaction(tx);
	if (receipt.executionResult !== 'SUCCESS') {
		throw new Error(receipt.outputArguments[0].value);
	}
}

module.exports = {
	getContractCode,
	getClient,
	deploy
}

if (!module.parent) {
	(async () => {
		try {
			await deploy(getClient(), createAccount(), "{{.AppName}}")
			console.log("Deployed {{.AppName}} smart contract successfully");
		} catch (e) {
			console.error(e);
		}
	})();
}
`
