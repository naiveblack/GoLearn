pragma solidity ^0.8.7;

// 商品合约
contract Goods {

    struct TraceData {

        address operator;   // 信息的操作者
        uint8 status;        // 商品的状态 {0:create, 1:transfer, 2:upload, 3:consume}
        uint256 timestamp;
        string remark;

    }

    uint256 goodsID;
    uint8 currentStatus;
    uint8 constant STATUS_CREATE = 0;
    

    TraceData[] traceDatas; // 溯源的信息

    event NewStatus(address _operator, uint8 _status, uint256 _time, string _remark);

    constructor(uint256 _goodsID) {
        goodsID = _goodsID;
        traceDatas.push(TraceData(tx.origin, STATUS_CREATE, block.timestamp, "create"));
        currentStatus = STATUS_CREATE;
        emit NewStatus(tx.origin,  STATUS_CREATE, block.timestamp, "create");
    }

    // 商品状态变更
    function changeStatus(uint8 _status, string memory _remark) public returns (bool) {
        currentStatus = _status;
        traceDatas.push(TraceData(tx.origin, _status, block.timestamp, _remark));
        emit NewStatus(tx.origin,  _status, block.timestamp, _remark);
        return true;
    }

    function getStatus() public view returns (uint8) {
        return currentStatus;
    }

    function getTraceInfo() public view returns (TraceData[] memory) {
        return traceDatas;
    } 
}