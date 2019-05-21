const env = process.env;

const HDWalletProvider = require('truffle-hdwallet-provider');

module.exports = {
    compilers: {
        solc: {
            version: "0.5.2",
            settings: {
                optimizer: {
                    enabled: true,
                    runs: 200
                }
            }
        }
    },
    networks: {
        testrpc: {
            host: 'localhost',
            port: 8545,
            network_id: '*', // eslint-disable-line camelcase
        },
        ganache: {
            host: 'localhost',
            port: 8545,
            skipDryRun: true,
            network_id: '*', // eslint-disable-line camelcase
        },
        rinkeby: {
            provider: new HDWalletProvider(env.MNEMONIC, "https://rinkeby.infura.io/", 0, 10),
            skipDryRun: true,
            network_id: 4
        },
        mainnet: {
            provider: new HDWalletProvider(env.MNEMONIC, "https://mainnet.infura.io/1u84gV2YFYHHTTnh8uVl", 0, 10),
            skipDryRun: true,
            network_id: 1
        }
    },
};