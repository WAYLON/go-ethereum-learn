pragma solidity ^0.4.24;

import "./ERC20Basic.sol";
/**
 * @title ERC20 interface
 * @dev Enhanced interface with allowance functions.
 * See https://github.com/ethereum/EIPs/issues/20
 */
contract ERC20 is ERC20Basic {
    // Check the allowed value that the _owner allows the _spender to take from his balance.
    function allowance(address _owner, address _spender) public view returns (uint256);

    // Transfer _value from the balance of holder _from to the receiver _to.
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool);

    // Approve _spender to take some _value from the balance of msg.sender.
    function approve(address _spender, uint256 _value) public returns (bool);

    // Fired when an approval is made.
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );
}