# fabric-create-ca-cert-tools
解决开始时配置的证书生成数量不够  工具可以追加证书

# 依赖
依赖地址：https://github.com/xinfei0902/fabric-sdk-go-basicPackage

# 编译
windows: ./build_base_windows.ps1 ./GeneralAddCife

linux ./build_base_linux.sh ./GeneralAddCife

## 创建组织证书 

### 工具说明

1. 命令行格式

   ```sh
   可执行文件 {子命令} {参数1} {参数2} {参数3}
   
   ```

2. 命令行说明

   | 序号 | 命令/参数  | 说明               | 备注                             |
   | ---- | ---------- | ------------------ | -------------------------------- |
   | 1.   | 可执行文件 | 生成证书可执行程序 | 是GeneralAddCife所执行的工作路径 |
   | 2.   | 子命令     | org                | 固定命令                         |
   | 3.   | 参数1      | 生成的证书存放位置 | 绝对路径                         |
   | 4.   | 参数2      | 组织名称           | 要新增组织的名称                 |
   | 5.   | 参数3      | 组织域名           | 要新增组织的名称+域名            |

### 示例

1. 依赖环境

   1. windows: .exe 可执行程序
   2. or linux: .bin 可执行程序  
   3. GeneralAddCife.exe/.bin

2. 示例

   ```sh 
   # run command
   ./GeneralAddCife.bin org ${workSpace}/crypto-config/peerOrganizations/org3msp.example.com Org3Msp org3msp.example.com
   
   # check result
   ls 
   
   # result
   # 存放证书位置 会生成以下文件夹：
   # >> ca  msp  tlsca  users
   
   ```



## 创建节点证书

### 工具说明

1. 命令行格式

   ```sh
   可执行文件 {子命令} {参数1} {参数2} {参数3} {参数4}
   
   ```

2. 命令行说明

   | 序号 | 命令/参数  | 说明                 | 备注                             |
   | ---- | ---------- | -------------------- | -------------------------------- |
   | 1.   | 可执行文件 | 生成证书的可执行程序 | 是GeneralAddCife所执行的工作路径 |
   | 2.   | 子命令     | peer                 | 固定命令                         |
   | 3.   | 参数1      | 生成的证书存放位置   | 绝对路径                         |
   | 4.   | 参数2      | 组织域名             | 需要新增peer的组织 + 域名        |
   | 5.   | 参数3      | 节点名称(peer)       | 需要新增peer的名称               |
   | 6.   | 参数4      | 用户名称(Users)      | 需要新增peer的用户               |

### 示例

1. 依赖环境

   1. windows: .exe 可执行程序
   2. or linux: .bin 可执行程序  
   3. GeneralAddCife.exe/.bin
   4. 已有组织文件夹 e.g.: 
      * ${workSpace}/crypto-config/peerOrganizations/org3msp.example.com
      * 此文件夹内证书已生成完成

2. 示例

   ```sh
   # run command
   ./GeneralAddCife.bin peer ${workSpace}/crypto-config/peerOrganizations/org3msp.example.com  org3msp.example.com  peer0  user1
   
   # check result
   ls ${workSpace}/crypto-config/peerOrganizations/org3msp.example.com/peers
   
   # result
   # >> peer0.example.com
   
   ls ${workSpace}/crypto-config/peerOrganizations/org3msp.example.com/peers/peer0.org3msp.example.com
   
   # result
   # >> msp tls
   
   ```
