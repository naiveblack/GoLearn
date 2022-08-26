pragma solidity^0.6.6;
// 0.8版本以后增了检测机制

contract addtest {
    uint8 public data = uint8(-1); //最大值  255
    uint8 public data2 = 200;

    function add(uint8 a, uint8 b) public pure returns (uint8) {
        return a + b;
    }

}