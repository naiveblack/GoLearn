// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

contract event_demo {

    address public admin;
    address payable public user;
    uint256 totalAmount;


    // 增加event定义
    event Deposit(address _who, uint256 _amount);
    event Withdraw(address _who, address _operator, uint256 _amount);

    constructor(address _owner) {
        admin = _owner;
    }

    // 充值方法
    function deposit(uint256 _amount) public payable {
        /*
            存在问题：
            当执行该函数时，实际上已经完成了余额的增加
            但如果在if判断时不满足，则会导致后续的表面上的账户余额没有增加

            修改使用断言处理
        */
        // if(_amount != msg.value) return;
        require(_amount == msg.value, "amount must == msg.value");
        assert(_amount > 0);
        
        user = payable(msg.sender); // 记录当前充值的人
        totalAmount = _amount;
        emit Deposit(msg.sender, _amount);
    }

    function getBalance() public view returns (uint256, uint256) {
        // this对象代表了合约本身
        return (address(this).balance, totalAmount);
    }

    function withdraw(uint256 _amount) public payable {
        // 对一个地址类型进行转账 需要这个地址类型是payable的
        // address payable public user;
        user.transfer(_amount);
        emit Withdraw(user, msg.sender, _amount);
    }
}