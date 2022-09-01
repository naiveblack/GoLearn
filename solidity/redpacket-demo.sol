// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

/*
    1. 发送红包 ： 随机发 || 指定金额
    2. 抢红包
    3. 退回
*/

contract redpacket {

    bool public rType;
    uint8 public rCount;
    uint256 public rTotalAmount;
    address payable public tuhao;
    mapping(address=>bool) isStake;

    // true：固定数值 false：随机
    constructor(bool isAvg, uint8 _count, uint256 _amount) payable {
        rType = isAvg;
        rCount = _count;
        rTotalAmount = _amount;
        require(_amount == msg.value, "redpacket's balance is ok");
        tuhao = payable(msg.sender);
    }

    function getBalance() public view returns (uint256) {
        return address(this).balance;
    }

    // 抢红包
    function stakePacket() public payable {
        require(rCount > 0, "redpacket must left");
        require(getBalance() > 0, "contract's balance must enough");
        require(!isStake[msg.sender], "user has been stake");
        isStake[msg.sender] = true;
        if(rType == true) { // 等值
            uint256 amount = getBalance() / rCount;
            payable(msg.sender).transfer(amount);
        } else {            // 随机
            if (rCount == 1) {
                uint256 amount = getBalance();
                payable(msg.sender).transfer(amount);
            } else {
                uint256 randNum = uint256(keccak256(abi.encode(tuhao, rTotalAmount, rCount, block.timestamp, msg.sender)))%10; // 10以内的随机数
                // uint256 amount = randNum / 10 * getBalance(); // 如此计算会得到 0 * getBalance;
                uint256 amount = getBalance() * randNum / 10;
                payable(msg.sender).transfer(amount);
            }
        }
        
        rCount --;
    }

    function kill() public payable {
        selfdestruct(tuhao);    // 合约销毁
    }
}