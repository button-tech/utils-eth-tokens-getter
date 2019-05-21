const BalanceContract = artifacts.require("./Balance.sol");

const web3 = global.web3;

const tbn = v => web3.utils.toBN(v);
const fbn = v => v.toString();
const tw = v => web3.utils.toBN(v).mul(1e18);
const fw = v => web3.utils.fromWei(v).toString();


contract('Balance', (accounts) => {
    const account = accounts[0];

    beforeEach(async function () {
        Balance = await BalanceContract.new({from: account});
    });

    describe('Common test', () => {
        it("should get single balance", async function () {
            const balance = await web3.eth.getBalance(accounts[1]);
            const balanceFromContract = await Balance.getBalance([accounts[1]]);
            assert.equal(balanceFromContract, balance, "Balances don't match");
        });

    });
});

