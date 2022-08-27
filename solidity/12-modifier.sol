// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

contract modifier_demo {
    address public admin;
    uint256 public amount;

    constructor() {
        admin = msg.sender;
        amount = 101;
    }

    // 本质上就是条件检测
    modifier onlyadmin() {
        require(msg.sender==admin, "only admin can do");
        require(amount > 100, "amout must > 100");
        _;
    }

    function setCount(uint256 _amount) public onlyadmin {
        amount = _amount;
    }
}