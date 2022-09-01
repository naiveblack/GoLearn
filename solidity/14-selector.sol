// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "./IUser.sol";

contract selector_demo {

    // 函数名称进行hash得到函数选择器
    function getSelector() public pure returns (bytes32, bytes4, bytes4) {

        bytes32 hash;
        // hash = keccak256(abi.encode("addUser(string,uint8)"));
        // output: bytes32: 0xb4dc18a22b4ac5569dc0ef31f9b3971417829e3bec0f4cfc3803120773f393ce

        /*
            计算函数选择器的时候
                1、 不能有空格
                2、 不需要加abi.encode
        */
        hash = keccak256("addUser(string,uint8)");
        // output: bytes32: 0x18edda4f06495ff44ae143db709ffbcc26668fb78b2c5a674fea2285cf9372b9
        return (hash, bytes4(hash), IUser.addUser.selector);
        // IUser.addUser.selector    output: bytes4: 0x18edda4f
    }

    // 计算InterfaceID
    // 接口ID是所有函数选择器进行异或
    function getInterfaceId() public pure returns (bytes4) {

        // 存在函数重载的时候就会有问题（函数名一样 参数不一样）
        // 使用hash = keccak256("addUser(string,uint8)");方法 则更加安全
        return IUser.addUser.selector ^ IUser.getUser.selector;

    }
}