## 接口文档 

## 接口签名算法
示例：gva/server/payment_test.go

假如你请求的接口是：http://localhost:8888/payment/brcode

拼接方法+路径+参数:
method + path + query+body=
GETlocalhost:8888/payment/brcode?ts=3212&nonce=23121&apiKey=32132&body

query参数，要按照key值从小到大排列，如果没有body，则不必拼接body

然后将拼接的字符串使用hamc(sha256)算法加密，密钥是apiSecret(在个人信息主页查看)，得到的结果转为base64编码，作为query参数sig=xxx添加到请求上。

##  获取收款二维码

需要通过接口签名校验

GET /payment/brcode?ts=3213&nonce=23121&apiKey=3123

| 字段| 类型| 说明|
|---|---|---|
|ts | int64| 时间戳|
|nonce| int| 随机数|
|apiKey| string| 与apiSecret配套，在个人信息页查看|


```json
response:
{

    "code": 0,
    "msg": "xxx",
    "data": {
        "url": "",       // 二维码地址
        "creatd_at": "", //创建时间(UTC)
        "expiration": 3600,  // 有效时长，单位秒(s)
    }
}
```

## starkbank回调

测试请求

```sh
curl --location --request POST 'localhost:8888/webhook/starkbank' \
--header 'Content-Type: application/json' \
--data-raw '{
    "event": {
        "id": "5344245984526336",
        "isDelivered": false,
        "subscription": "deposit",
        "created": "2020-03-11T00:14:22.104676+00:00",
        "log": {
            "id": "5066704988143616",
            "created": "2020-03-11T00:14:22.104676+00:00",
            "type": "credited",
            "errors": [],
            "deposit": {
                "status": "created",
                "accountNumber": "1010101010101010",
                "accountType": "payment",
                "taxId": "10.203.000/0001-80",
                "transactionIds": [
                    "5862355036536832"
                ],
                "bankCode": "20018183",
                "id": "5738709764800512",
                "fee": 50,
                "name": "Iron Bank S.A",
                "created": "2023-06-11T00:14:22.104702+00:00",
                "updated": "2023-06-11T00:14:21.548876+00:00",
                "tags": [
                "deposit", "uuid=2bcf71c6-80e4-4178-bc3d-fc8fcfa7c92b",
                "payment=starkbank", "createdAt=2023-06-12T15:04:05+08:00"
                ],
                "branchCode": "0001",
                "amount": 9000,
                "type": "pix"
            }
        }
    }
}'
```