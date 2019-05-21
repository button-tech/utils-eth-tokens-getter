const addresses = require("./addresses.json");

const Balance = artifacts.require("Balance");
const balanceContractAddress = '0x74Cb34D5a97c808b02a5F56631a21c822CEa1204';

module.exports = async function() {
    try {

        const balance = await Balance.at(balanceContractAddress);

        async function getBalances(tokenAddress, addresses) {
            console.log(`We have ${addresses.length} addresses for ${tokenAddress}`);
            const t = timer("getBalances");
            await balance.getBalance(addresses);
            t.stop();
        }

        for (let i in addresses) {
            await getBalances(i, addresses[i]);
            if (i === Object.keys(addresses)[Object.keys(addresses).length - 1])
                process.exit(1);
        }
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