const fs = require("fs");

const IERC20 = artifacts.require("IERC20");
const contractAddresses = [
    '0xd26114cd6EE289AccF82350c8d8487fedB8A0C07',
    '0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2',
    // '0x0d8775f648430679a709e98d2b0cb6250d2887ef',
    '0xd850942ef8811f2a866692a623011bde52a462c1',
    '0x514910771af9ca656af840dff83e8264ecf986ca',
    '0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b',
    '0x6c6ee5e31d828de241282b9606c8e98ea48526e2'
];

module.exports = async function() {
    const step = 1000;
    let addresses = {};

    async function getAddresses(tokenAddress, fromBlock, toBlock) {
        const token = new web3.eth.Contract(IERC20.abi, tokenAddress);

        return new Promise((resolve) => {
            console.log("From block number: " + fromBlock);
            console.log("To block number: " + toBlock);
            console.log("Token address: " + tokenAddress);
            console.log("\n");
            token.getPastEvents("Transfer", {
                fromBlock: fromBlock,
                toBlock: toBlock
            }, function(error, events){ if (error) console.log(error) }).then(function(events){
                resolve(events.map(x=>x.returnValues.from));
            });
        });
    }

    for (let i = 0; i < contractAddresses.length; i ++) {
        let toBlockNumber = (await web3.eth.getBlock("latest")).number;
        let fromBlockNumber = toBlockNumber - step;

        try {
            let tmp = [];
            while (tmp.length < 1000) {
                // await timeout(500);
                console.log("Count: " + tmp.length);
                const res = await getAddresses(contractAddresses[i], fromBlockNumber, toBlockNumber);
                fromBlockNumber = fromBlockNumber - step;
                toBlockNumber = toBlockNumber - step;
                console.log(res)
                tmp.push(...res);
            }
            addresses[contractAddresses[i]] = tmp;

            if (contractAddresses.length - 1 === i) {
                fs.unlink(__dirname + '/addresses.json', (err) => {
                    fs.appendFileSync(__dirname + '/addresses.json', JSON.stringify(addresses));
                    console.log("Done!");
                    process.exit(1);
                });
            }
        } catch (e) {
            console.log(e)
        }
    }
};

const timeout = (time) => {
    return new Promise(resolve => {
        setTimeout(()=>{resolve(true);},time);
    });
};