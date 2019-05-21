const fs = require("fs");

const IERC20 = artifacts.require("IERC20");
const contractAddress = '0xd26114cd6EE289AccF82350c8d8487fedB8A0C07';

module.exports = async function() {
    const step = 1000;
    let toBlockNumber = (await web3.eth.getBlock("latest")).number;
    let fromBlockNumber = toBlockNumber - step;

    let addresses = [];

    const token = await IERC20.at(contractAddress);
    function getAddresses() {
        return new Promise((resolve) => {
            console.log(fromBlockNumber)
            token.getPastEvents("Transfer", {
                fromBlock: fromBlockNumber,
                toBlock: toBlockNumber
            }, function(error, events){ console.log(error); }).then(function(events){
                fromBlockNumber = fromBlockNumber - step;
                toBlockNumber = toBlockNumber - step;

                resolve(events.map(x=>x.returnValues.from));
            });
        });
    }

    try {
        while (addresses.length < 2500) {
            const res = await getAddresses();
            addresses.push(...res);
        }
        fs.unlink(__dirname + '/addresses.json', (err) => {
            fs.appendFileSync(__dirname + '/addresses.json', JSON.stringify({addresses: addresses}));
            console.log("Done!");
            process.exit(1);
        });
    } catch (e) {
        console.log(e)
    }
};