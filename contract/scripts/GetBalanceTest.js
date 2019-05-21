const addresses = require("./addresses.json").addresses;
console.log("We have: " + addresses.length + " addresses");

const Balance = artifacts.require("Balance");
const contractAddress = '0x74Cb34D5a97c808b02a5F56631a21c822CEa1204';


module.exports = async function() {

    const balance = await Balance.at(contractAddress);

    async function getBalances(addresses) {
        const t = timer("getBalances");
       await balance.getBalance(addresses);
        t.stop();
    }

    try {
        await getBalances(addresses);
        process.exit(1);
    } catch (e) {
        console.log(e)
    }
};

const timer = function(name) {
    const start = new Date();
    return {
        stop: function() {
            const end  = new Date();
            const time = end.getTime() - start.getTime();
            console.log('Timer:', name, 'finished in', time, 'ms');
        }
    }
};