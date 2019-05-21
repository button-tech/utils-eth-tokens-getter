const Balance = artifacts.require("Balance");

module.exports = function(deployer) {
  deployer.deploy(Balance, {gas: 6720000, overwrite: false});
};
