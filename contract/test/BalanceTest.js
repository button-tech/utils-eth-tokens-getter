const BalanceContract = artifacts.require("./Balance.sol");
// const TestTokenContract = artifacts.require("./TestToken.sol");

const web3 = global.web3;

const tbn = v => web3.utils.toBN(v);
const fbn = v => v.toString();
const tw = v => web3.utils.toBN(v).mul(1e18);
const fw = v => web3.utils.fromWei(v).toString();


contract('Balance', (accounts) => {
    const deployer = accounts[0];
    const testAccounts = [];
    beforeEach(async function () {
        Balance = await BalanceContract.new({from: deployer});
        // Test = await TestTokenContract.new({from: deployer});
        // for (let i = 0; i < 1000; i++) {
        //     const acc = web3.eth.accounts.create();
        //     testAccounts.push(acc.address);
        // }
    });

    describe('getBalance', () => {
        it("should get single balance", async function () {
            const balance = await web3.eth.getBalance(accounts[1]);
            const balanceFromContract = await Balance.getBalance([accounts[1]]);
            assert.equal(balanceFromContract, balance, "Balances don't match");
        });

    });

    describe('getTokenBalance', () => {
        // it("should get single balance", async function () {
        //     const balance = await web3.eth.getBalance(accounts[1]);
        //     const balanceFromContract = await Balance.getBalance([accounts[1]]);
        //     assert.equal(balanceFromContract, balance, "Balances don't match");
        // });

    });
});

