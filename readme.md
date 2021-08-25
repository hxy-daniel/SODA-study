# SODA study
## block collector
```SODA_code/go-ethereum/core/state_processor.go Process()```
## trans collector
external trans ```SODA_code/go-ethereum/core/state_processor.go ApplyTransaction()```
internal trans ```SODA_code/go-ethereum/core/vm/interpreter.go Run()```and```SODA_code/go-ethereum/core/vm/instructions.go opCreate() opCreate2() opCall() opCallCode() opDelegateCall() opStaticCall()```

## ins collector

## update location of go-ethereum
1. ```SODA_code/go-ethereum/cmd/pluginManage/* new file```
2. ```SODA_code/go-ethereum/core/state_processor.go Process()```
3. ```SODA_code/go-ethereum/core/vm/evm.go```
4. ```SODA_code/go-ethereum/core/vm/instructions.go```
5. ```SODA_code/go-ethereum/core/vm/interpreter.go Run()```
6. ```SODA_code/go-ethereum/core/vm/stack.go```
7. ```SODA_code/go-ethereum/eth/api.go```
8. ```SODA_code/go-ethereum/log/pluginlog.go new file```
9. ```SODA_code/go-ethereum/miner/worker.go```
10. ```SODA_code/go-ethereum/params/config.go```
11. ```SODA_code/go-ethereum/tingrong/transOpcodeStructure.go new file```


# SODA_code
SODA is a novel generic online detection framework for smart contracts on blockchains that support Ethereum virtual machine (EVM). We released the 8 detection apps and the source code of the framework here.  

We developed the framework SODA based on go-ethereum v1.9.0 (https://github.com/ethereum/go-ethereum/tree/v1.9.0) in Ubuntu 16.06.

The source code of 8 detection apps is under the path ```SODA_code/plugin/plugin```.

## How to use this framework and the 8 detection apps
1. Use ```go env``` to check your paths of ```GOPATH``` and ```GOROOT``` in your Ubuntu.
2. Copy the file ```collector.go``` in the path ```SODA_code/collector``` to the path ```GOROOT/src/github.com/ethereum/collector``` (if a directory does not exist, create it).
3. Copy the folder ```json-iterator``` and ```modern-go``` in the path ```SODA_code/go-ethereum/vendor/github.com``` to the path ```GOPATH/src/github.com``` (if a directory does not exist, create it).
4. Enter the folder ```SODA_code/go-ethereum```, set```GO111MODULE=auto```, use ```make geth``` to compile the framework, and then you can get ```geth``` from the path ```SODA_code/go-ethereum/build/bin```.
5. Enter the path ```SODA_code/plugin/plugin/P1```, and then use ```go build –buildmode=plugin P1.go``` to get ```P1.so```.
6. Make two new directories ```plugin``` where to put the ```P1.so``` and ```public``` where to store sync data in the same directory as ```geth```.
7. In the directory where ```geth``` is, use ```./geth -syncmode full -datadir public``` or ```nohup ./geth -syncmode full -datadir public > gethlog.txt 2>&1 &``` to start syncing.
8. Finally, you will find the result of each app in the folder ```plugin_log```.

# Result
P1 is an app for detecting a malicious re-entrancy aiming at stealing ETH. The result of P1 is listed in the table ```P1_result.xlsx```.   
We have listed all 8 apps' results at https://drive.google.com/drive/folders/1gHAlmivO1zntSaAoZjoSymG0sQS8lv32?usp=sharing.

# Paper
You can find our paper about the design, implementation, and experimental results of SODA at https://www.ndss-symposium.org/wp-content/uploads/2020/02/24449.pdf.

# Citing in Academic Work
Welcome to cite our paper:
Ting Chen, Rong Cao, Ting Li, Xiapu Luo, Guofei Gu, Yufei Zhang, Zhou Liao, Hang Zhu, Gang Chen, Zheyuan He, Yuxing Tang, Xiaodong Lin, Xiaosong Zhang. SODA: A Generic Online Detection Framework for Smart Contracts. In NDSS 2020.

# Contact us
If you have any problems in using our tool, please send emails to chenting19870201@163.com and 1797258848@qq.com.
