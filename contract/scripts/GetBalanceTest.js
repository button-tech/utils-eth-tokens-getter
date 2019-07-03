const addresses = require("./addresses.json");
const test = require("./test.json");

const Balance = artifacts.require("Balance");
const balanceContractAddress = '0x39aA7b93bbF1b5895a7d5A41Aa70B664aF8BC29f';
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
    try {
        const balance = await Balance.at(balanceContractAddress);

        async function getBalances(tokenAddress, addresses) {
            console.log(`We have ${addresses.length} addresses for ${tokenAddress}`);
            const t = timer("getBalances");
            await balance.getBalance(addresses);
            t.stop();
        }

        async function getTokenBalancesByAddress(tokenAddresses, address) {
            console.log(`We have ${tokenAddresses.length} token addresses`);
            const t = timer("getBalances");
            await balance.getTokenBalanceByAddress(address, tokenAddresses);
            t.stop();
        }

        async function getTokenBalances(tokenAddress, addresses) {
            console.log(`We have ${addresses.length} addresses for ${tokenAddress}`);
            const t = timer("getTokenBalances");
            await balance.getTokenBalance(addresses, tokenAddress);
            t.stop();
        }

        await getTokenBalancesByAddress(test.tokens, test.user);

        for (let i in addresses) {
            await getBalances(i, addresses[i]);
            if (i === Object.keys(addresses)[Object.keys(addresses).length - 1]) {
                await getTokenBalances([contractAddresses[0]], addresses[i]);
                process.exit(1);
            }
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