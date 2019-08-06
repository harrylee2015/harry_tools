pragma solidity ^0.5.5;

contract quickgame {


    // currency unit
    uint decimal = 1e18;
    struct gameInfo {
        address payable creator;
        string name;
        //游戏轮次
        uint round;
        //创建房间所需押金
//      uint deposit;
        //当前参与的次数
        uint size;
        mapping(uint => Record) data;
//        Record  []RecordList;
        uint bonusPool;
        uint createTime;
        uint closeTime;
        uint luckyNum;
    }
    //投注,规则可调整,这里相当于猜数字 1-20,猜5个
    struct Record {
        address payable addr;
        uint  value;
    }
    // 房间总数
    uint RoomCount;
    // 记录每个地址创建得房间数
    mapping (address =>uint) RoomMap;
    // 记录房间名
    mapping (uint =>string) RoomNameMap;
    // 记录房间名是否被占用
    mapping (string =>bool) RoomName;
    // 房间游戏轮次,类似于计数器
    mapping(string => uint) RoundMap;

    // 记录每一个房间每一轮游戏信息
    mapping (string => mapping(uint => gameInfo))  gameMap;

    // administrator address
    address admin;

    // guessRecord
//    address[] guessList;

    modifier isAdmin() {
        require(msg.sender == admin, "Only admin is permited");
        _;
    }

    constructor() payable public {
        admin = msg.sender;
    }


    //获取区块时间
    function getBlockTime() public view returns(uint) {
        return now;
    }
    function checkGameRoomName(string memory _name)public view returns (bool) {

        return RoomName[_name];
    }

    function createGameRoom(string memory _name) public {
        require(msg.sender != admin, "administrator is not allowed");
        address payable creator = msg.sender;
        require(RoomMap[creator]==5, "create up to 5 rooms per person!");
        require(RoomName[_name]==true, "create up to 5 rooms per person!");
        //初始化
        gameInfo memory info = gameInfo({
        creator:creator,
        name:_name,
        round:1,
            size:0,
        bonusPool:0,
        createTime:block.timestamp,
        closeTime:0,
        luckyNum:0
    });
        uint lastRound = 1;
        RoundMap[_name]=lastRound;
        gameMap[_name][lastRound]=info;
    }



    function guess(string memory _name,uint _value) payable public{
        require(msg.sender != admin, "administrator is not allowed");
        uint value = msg.value;
        require(value > decimal, "min value is 1.0 bty");
//        require(value < decimal+decimal/10, "max value is 1.09 bty");
        require(RoomName[_name]==false,"room name not exist!");
        // 这里需要对投注数字进行检查
        uint lastRound = RoundMap[_name];
        Record memory record = Record(msg.sender,_value);

        gameInfo storage info = gameMap[_name][lastRound];
        info.size=info.size+1;
        info.data[info.size]=record;
        info.bonusPool=info.bonusPool+decimal;
        if (info.size==3){
            //满足条件自动开奖
            uint luckyNum = (block.number+now)%10;
            info.closeTime=block.timestamp;
            info.luckyNum=luckyNum;
            uint index =0;
            // find the lucky people
            for(uint idx=1; idx<=info.size;idx++){
                if (info.data[idx].value == luckyNum ){
                    index++;
                }
            }
            // 一个中奖得都没有
            if (index==0){
                gameMap[_name][lastRound]=info;
                RoundMap[_name]=info.round+1;
                gameInfo memory lastInfo = info;
                lastInfo.round=lastInfo.round+1;
                lastInfo.createTime=block.timestamp;
                lastInfo.size=0;
                lastInfo.closeTime=0;
                lastInfo.luckyNum=0;
                gameMap[_name][lastRound+1]=lastInfo;
                return ;
            }
            // give bonus

            uint totalBonus = info.bonusPool/2;
            uint oneBonus = totalBonus/index;

            for(uint idx=1; idx<=info.size;idx++){
                if (info.data[idx].value == luckyNum ){
                    info.data[idx].addr.transfer(oneBonus);
                }
            }
            gameMap[_name][lastRound]=info;
            RoundMap[_name]=info.round+1;
            gameInfo memory lastInfo = info;
            lastInfo.round=lastInfo.round+1;
            lastInfo.createTime=block.timestamp;
            lastInfo.bonusPool=totalBonus;
            lastInfo.size=0;
            lastInfo.closeTime=0;
            lastInfo.luckyNum=0;
            gameMap[_name][lastRound+1]=lastInfo;
            //具体分配奖金方案待定
        }

        gameMap[_name][lastRound]=info;
    }

    function getRound(string memory _name) public view returns(uint) {
        return RoundMap[_name];
    }

    function getLastLuckyNum(string memory _name) public view returns(uint) {
        uint lastRound = RoundMap[_name];
        return gameMap[_name][lastRound].luckyNum;
    }

    // solidity中不支持返回数组形式,后续调整游戏模式
//    function getRecord(string memory _name,uint round) public view returns(Record){
//
//        return ;
//    }

}
