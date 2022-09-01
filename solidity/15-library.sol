// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "./libstring.sol";

// 调用libstring库
contract library_demo {
    // 先声明 库给string类型的数据使用
    using libstring for string;
    // isEqual(string memory _a, string memory _b)
    // 等于变成了string类型的内部函数 isEqual是string的成员函数
    // string本身就是成员中的一个元素
    // 则此时只需要传递一个参数就行 因为第一个参数默认就变成了string对象本身

    function isMyEqual(string memory a, string memory b) public pure returns (bool) {
        return a.isEqual(b); // 等价于 isEqual(a, b);
    }
}