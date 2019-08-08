# adaptor
Abstract Blockchain Adaptor for PalletOne
## IUtility 通用操作接口
通用接口包含了最基本的钱包和区块链的操作。钱包操作就包括公私钥的生成和地址的生成，区块链的基本操作就包括对区块和交易的构建查询。
另外为了跨链时地址的映射，有一个特殊的接口：GetMappingPalletOneAddress，用于从链上取的链上地址与PTN地址的映射。

## ICryptoCurrency Token操作接口
Token操作相关接口主要包含获取Token基本信息（小数精度），获取某个账户的余额，查询Token转账历史记录，构建转账交易，构建多签锁定地址等。
## ISmartContract 智能合约操作接口
包括智能合约的安装、部署、调用和查询。