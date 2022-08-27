// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

// payable案例
// 实现充值和提现的功能 
// address.transfer(uint256 amount) 不会超过2300gas（gas上限）
contract money_demo {

    address public admin;
    address payable public user;
    uint256 totalAmount;

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
        if(_amount != msg.value) return;

        user = payable(msg.sender); // 记录当前充值的人
        totalAmount = _amount;
        /*
            _amount 只是一个表面的账目
            真正的钱是在msg.value里面
            涉及到金额转移的时候会自动执行增加操作
        */
        // address(this).balance += _amount;
    }

    /*
        账户名       余额
        Addr1       100
        Addr2       200
        ... ...
    */
    function getBalance() public view returns (uint256, uint256) {
        // this对象代表了合约本身
        return (address(this).balance, totalAmount);
    }

    function withdraw(uint256 _amount) public payable {
        // 对一个地址类型进行转账 需要这个地址类型是payable的
        // address payable public user;
        user.transfer(_amount);
    }
}