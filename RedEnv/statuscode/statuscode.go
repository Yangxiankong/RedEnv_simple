package statuscode

var Success int = 0         //成功
var BadRequest int = 400    //请求格式有误
var NoSuchUser int = -1     //不存在的用户
var CantGetMoreEnv int = -2 //该用户获取红包达到上限
var Thankyou int = -3       //没抢到
var NoSuchRec int = -4      //没这个红包记录
