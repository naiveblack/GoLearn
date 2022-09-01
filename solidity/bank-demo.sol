// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

/*
    1. function
    2. bank name
    3. check moeny
*/
contract bank_demo {

    string public bankName; //银行名称
    uint256 totalAmount;    //存款总额
    address public admin;
    constructor(string memory _name) {
        bankName = _name;
        admin = msg.sender;
    }

    mapping(address=>uint256) balances; //记录用户余额

    // 查看用户余额
    function balanceOf(address _who) public view returns (uint256) {
        return balances[_who];
    }

    // 核对总额
    function getBalance() public view returns(uint256, uint256) {
        return (address(this).balance, totalAmount);
    }

    // 存款
    function deposit(uint256 _amount) public payable {
        require(_amount > 0, "amount must > 0");
        require(_amount == msg.value, "msg.value must equal amount");
        balances[msg.sender] += _amount;
        totalAmount += _amount;
        // 检测交易完后，账目是否可以对应
        require(address(this).balance == totalAmount, "bank's balance must ok");
    }

    // 取款
    function withdraw(uint256 _amount) public payable {
        require(_amount > 0, "amount must > 0");
        require(_amount <= balances[msg.sender], "user's balance not enough");
        balances[msg.sender] -= _amount;
        payable(msg.sender).transfer(_amount);
        totalAmount -= _amount;
        require(address(this).balance == totalAmount, "bank's balance must ok");
    }

    // 转移
    function transfer(address _to, uint256 _amount) public {
        require(_amount > 0, "amount must > 0");
        require(_amount <= balances[msg.sender], "user's balance not enough");
        require(address(0) != _to, "to address must valid");
        balances[msg.sender] -= _amount;
        balances[_to] += _amount;
        require(address(this).balance == totalAmount, "bank's balance must ok");
    }
}