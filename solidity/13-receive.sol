// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

contract receive_demo {

    uint256 totalAmount;
    address[] public addrs;

    // 用于接受外部钱包向合约地址的转账
    receive() external payable {
        totalAmount += msg.value;
        addrs.push(msg.sender);
    }

    function getBalance() public view returns (uint256, uint256) {
        return (totalAmount, address(this).balance);
    }

    // 只有在receive不存在的时候才能触发fallback函数
    // 除了金额转移 也可以检测函数调用
    // fallback多数用于解决错误调用 资金问题多用receive解决
    fallback() external payable {
        totalAmount += msg.value;
        addrs.push(msg.sender);
    }
}