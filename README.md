# 恋爱合约

## 说明
男女双方均是单身情况下,确定恋爱关系,写入合约,哪怕海枯石烂也无法修改.  

1.男方调用 loveContract 的 Boy 方法,申请确定和女方的恋爱关系
2.女方调用 loveContract 的 Girl 方法,确定和男方的恋爱关系

注意:如果男方申请后,女方拒绝写入恋爱关系,男方的记录也无法清除.

 ## 使用示例
 ```shell
##男方
./bin/xchain-cli native invoke --method Boy -a '{"girl":"cmr"}' --fee 110000 loveContract
##女方
./bin/xchain-cli native invoke --method Girl -a '{"boy":"zys"}' --fee 110000 loveContract
 ```
