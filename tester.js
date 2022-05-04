const algosdk = require('algosdk');
const fetch = require('node-fetch');

async function tester() {
    try {
 
        const algosdk = require('algosdk');
		const baseServer = 'https://testnet-algorand.api.purestake.io/ps2'
		const port = '';
		const token = {'X-API-Key': ''}
		

		const algodClient = new algosdk.Algodv2(token, baseServer, port);
		let params = await algodClient.getTransactionParams().do();

        const senderSeed = "brand mind toe laugh awful deal tell goose wall royal property bridge era region year goose cruise grant delay tray social vapor cabbage abandon engage";
        let senderAccount = algosdk.mnemonicToSecretKey(senderSeed);


        const creator = senderAccount.addr;
        const defaultFrozen = false;
        const unitName = "test";
        const assetName = "test";
        const assetURL = "example/url";
        let note = undefined;
        const manager = senderAccount.addr;
        const reserve = undefined;
        const freeze = undefined;
        const clawback = undefined;
        let assetMetadataHash = undefined;
        const total = 1;
        const decimals = 0;

        const mintNFT = algosdk.makeAssetCreateTxnWithSuggestedParams(
            creator,
            note,
            total,
            decimals,
            defaultFrozen,
            manager,
            reserve,
            freeze,
            clawback,
            unitName,
            assetName,
            assetURL,
            assetMetadataHash,
            params,
        );

        let signedTxn = algosdk.signTransaction(mintNFT, senderAccount.sk);
        
  
        console.log(signedTxn.txID)
        console.log(signedTxn)
        let transaction = []
        let i = 0;
        while (i < signedTxn.blob.length) {
            transaction[i] = signedTxn.blob[i]
            i++;
        } 
    

        return fetch("http://localhost:500/broadcast", {method: "POST", body: JSON.stringify({signedTXN: transaction, txID: signedTxn.txID})}).then(console.log).catch(console.error)

  
    }
    catch (err) {
        console.log("err", err);
    }
    process.exit();
};

tester();