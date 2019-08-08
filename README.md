# adaptor
Abstract Blockchain Adaptor for PalletOne
## IUtility 通用操作接口
通用接口包含了最基本的钱包和区块链的操作。钱包操作就包括公私钥的生成和地址的生成，区块链的基本操作就包括对区块和交易的构建查询。
另外为了跨链时地址的映射，有一个特殊的接口：GetPalletOneMappingAddress，用于从链上取的链上地址与PTN地址的映射。

## ICryptoCurrency Token操作接口
Token操作相关接口主要包含获取Token基本信息（小数精度），获取某个账户的余额，查询Token转账历史记录，构建转账交易，构建多签锁定地址等。
## ISmartContract 智能合约操作接口
包括智能合约的安装、部署、调用和查询。

## Token跨链的实现
### BTC跨链

### ETH跨链
#### 充币（ETH换PETH）
1. 官方发布一个多签合约到以太坊，该合约提供了函数：ABI
2. 用户使用PalletOne钱包，获得一个PalletOne地址。
3. 用户用自己的ETH钱包账户调用多签合约的PTNAddress映射函数，将ETH地址和PTN地址映射，ABI
4. 用户用自己的ETH钱包，转对应数量的ETH到多签地址。
5. 用户到PalletOne充币提币ETH合约中，调用提取PETH的方法，该方法逻辑如下：
   1. 根据当前调用者的PTN地址，调用Adaptor的GetPalletOneMappingAddress，获得对应的ETH地址
   2. 调用GetAddrTxHistory获得该地址往多签地址的转账记录
   3. 查询StateDB，如果是新的转账，则转对应PETH给用户
   4. 更新StateDB，标记该TXID为已处理
#### 提币（PETH换ETH）
1. 复用充币对应的那个以太坊多签合约
2. 用户将PETH转账到PalletOneETH提币合约，并调用了申请提币函数。
3. 用户使用ETH钱包，获得一个ETH钱包地址，然后调用PalletOne ETH提币合约的提币函数，传入ETH地址
4. 陪审团验证并处理提币请求，
   1. 调用Adaptor的CreateTransferTokenTx，生成一个交易，
   2. 调用SignTransaction，生成签名
   3. 收集陪审员的签名，并调用BindTxAndSignature，生成一个签名的交易
   4. 更新StateDB，将提币金额更新
   5. 由某个广播节点，将签名后的Tx广播到以太坊网络

### ERC20跨链