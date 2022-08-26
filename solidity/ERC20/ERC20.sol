// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "./IERC20.sol";
import "./SafeMath.sol";

contract ERC20 is IERC20 {

    string ercName;
    string ercSymbol;
    uint8 ercDecimal;
    uint256 ercTotalSupply;
    // 针对uint256类型 使用safemath库
    using SafeMath for uint256;

    mapping(address=>uint256) ercBalances;  //自定义账本
    mapping(address=>mapping(address=>uint256)) ercAllowance; //授权转账的额度

    address public owner;

    constructor(string memory _name, string memory _symbol, uint8 _decimal) {
        ercName =_name;
        ercSymbol = _symbol;
        ercDecimal = _decimal;
        owner = msg.sender;
        /*
        owner = msg.sender // 设为发布合约的人为管理员
        ercTotalSupply = 21000000;
        ercBalances[owner] = ercTotalSupply; 
        */
    }

    function name() override external view returns (string memory) {
        return ercName;
    }

    function symbol() override external view returns (string memory) {
        return ercSymbol;
    }

    function decimals() override external view returns (uint8) {
        return ercDecimal;
    }

    function totalSupply() override external view returns (uint256) {
        return ercTotalSupply;
    }

    function balanceOf(address _owner) override external view returns (uint256 balance) {
        return ercBalances[_owner];
    }

    function transfer(address _to, uint256 _value) override external returns (bool success) {
        require(_value > 0, "_value must > 0");
        require(address(0) != _to, "to address is zero");
        require(ercBalances[msg.sender] >= _value, "user's balance not enough");

        // ercBalances[msg.sender] -= _value;
        ercBalances[msg.sender] = ercBalances[msg.sender].sub(_value);

        // ercBalances[_to] += _value;
        ercBalances[_to] = ercBalances[_to].add(_value);

        emit Transfer(msg.sender , _to, _value);

        return true;
    }

    function transferFrom(address _from, address _to, uint256 _value) override external returns (bool success) {
        require(ercBalances[_from] >= _value, "value is not enough");
        require(ercAllowance[_from][msg.sender] >= _value, "approve is not enough");
        require(_value > 0, "value is zero");
        require(address(0) != _to, "to address is zero");

        ercBalances[_from] -= _value;
        ercBalances[_to] += _value;
        ercAllowance[_from][msg.sender] -= _value;

        emit Transfer(_from , _to, _value);
        return true;
    }

    function approve(address _spender, uint256 _value) override external returns (bool success) {
        // 委托人 -> 被委托人 -> 金额value
        // require(_value > 0, "_value must > 0"); //放开表示收回授权
        require(address(0) != _spender, "_spender is zero");
        require(ercBalances[msg.sender] >= _value, "uesr's balance not enough");

        ercAllowance[msg.sender][_spender] =  _value;

        emit Approval(msg.sender, _spender, _value);

        return true;
    }

    function allowance(address _owner, address _spender) override external view returns (uint256 remaining) {
        remaining = ercAllowance[_owner][_spender];
    }

    function mint(address _to, uint256 _value) public {
        require(msg.sender == owner, "only owner can do");
        require(address(0) != _to, "_to is zero");
        require(_value > 0, "value is zero");

        ercBalances[_to] += _value;
        ercTotalSupply += _value;

        emit Transfer(address(0), _to, _value);

    }

    function getTime() public view returns (uint256) {
        uint256 time = block.timestamp;
        return time;
    }
}