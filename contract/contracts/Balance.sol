pragma solidity ^0.5.2;

contract IERC20 {
       function totalSupply() public view returns (uint);
       function balanceOf(address tokenOwner) public view returns (uint balance);
       function allowance(address tokenOwner, address spender) public view returns (uint remaining);
       function transfer(address to, uint tokens) public returns (bool success);
       function approve(address spender, uint tokens) public returns (bool success);
       function transferFrom(address from, address to, uint tokens) public returns (bool success);

       event Transfer(address indexed from, address indexed to, uint tokens);
       event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}

contract IBalance {
    function getTokenBalance(address[] memory _addresses, address[] memory _tokenAddresses) public view returns (uint256[] memory balances);
    function getBalance(address[] memory _addresses) public view returns (uint256[] memory balances);
}

contract Balance {

   function getTokenBalance(address[] memory _addresses, address[] memory _tokenAddresses) public view returns (uint256[] memory balances) {
       require(_tokenAddresses.length >= 1);
       require(_addresses.length >= 1);
       uint256 counter = 0;
       balances = new uint[](_addresses.length * _tokenAddresses.length);
       for (uint i = 0; i < _tokenAddresses.length; i++) {
           for (uint j = 0; j < _addresses.length; j++) {
                IERC20 token = IERC20(_tokenAddresses[i]);
                balances[counter] = uint256(token.balanceOf(_addresses[j]));
                counter++;
           }
       }
       return balances;
   }

   function getBalance(address[] memory _addresses) public view returns (uint256[] memory balances) {
       require(_addresses.length >= 1);
       uint256 counter = 0;
       balances = new uint[](_addresses.length);
       for (uint j = 0; j < _addresses.length; j++) {
            balances[counter] = _addresses[j].balance;
            counter++;
       }
   }

}
