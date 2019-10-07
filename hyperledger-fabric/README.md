#   一.智联合约（链代码）
    
    q1.是什么？
        1.实际上是控制区块链网络中的不同实体后或者相关方如何交互或者交易的业务逻辑
        2.业务网络交易的代码
    q2.作用
        获取账本或者world state
#  二.链代码结构

    //引入 github.com/hyperledger/fabric/core/chaincode/shim 包
    
    package main
 
    import "fmt"
    import "github.com/hyperledger/fabric/core/chaincode/shim"
    
    type SampleChaincode struct {
    }
    //链代码首次部署到区块链网络时调用 可以做初始化任务使用
    func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        return nil, nil
    }
    //查询方法 执行任何读取/获取/查询操作
    func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        return nil, nil
    }
    //改变区块链状态 调用Invoke方法 执行创建、更新和删除操作
    //Invoke方法的所有调用都会在区块链上记录为交易，这些交易最终被写入区块中
    func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        return nil, nil
    }
    
    func main() {
        err := shim.Start(new(SampleChaincode))
        if err != nil {
            fmt.Println("Could not start SampleChaincode")
        } else {
            fmt.Println("SampleChaincode successfully started")
        }
    }

# 三.链代码的数据模型

    ##  1.超级账本包含两部分
        WorldState 存储在rocksdb(一个键值都市byte的kv数据库)
        区块链 区块连链成 每个区块包含WorldState的hash值 并且链接到前一个区块
        ?? 区块链采用仅附加模式 (append-only) ??

    ## 2.//创建自定义数据模型/模式的代码
        //custom data models
        type PersonalInfo struct {
            Firstname string `json:"firstname"`
            Lastname  string `json:"lastname"`
            DOB       string `json:"DOB"`
            Email     string `json:"email"`
            Mobile    string `json:"mobile"`
        }
        sxsds
        type FinancialInfo struct {
            MonthlySalary      int `json:"monthlySalary"`
            MonthlyRent        int `json:"monthlyRent"`
            OtherExpenditure   int `json:"otherExpenditure"`
            MonthlyLoanPayment int `json:"monthlyLoanPayment"`
        }
        
        type LoanApplication struct {
            ID                     string        `json:"id"`
            PropertyId             string        `json:"propertyId"`
            LandId                 string        `json:"landId"`
            PermitId               string        `json:"permitId"`
            BuyerId                string        `json:"buyerId"`
            SalesContractId        string        `json:"salesContractId"`
            PersonalInfo           PersonalInfo  `json:"personalInfo"`
            FinancialInfo          FinancialInfo `json:"financialInfo"`
            Status                 string        `json:"status"`
            RequestedAmount        int           `json:"requestedAmount"`
            FairMarketValue        int           `json:"fairMarketValue"`
            ApprovedAmount         int           `json:"approvedAmount"`
            ReviewerId             string        `json:"reviewerId"`
            LastModifiedDate       string        `json:"lastModifiedDate"`
        }

    ## 3.存储检索数据

        1.将数据存储到账本中(ADD)
            //ChaincodeStubInterface 拥有实用的 API 来与区块链账本、交易上下文、调用方证书

            func CreateLoanApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
                fmt.Println("Entering CreateLoanApplication")
            
                if len(args) < 2 {
                    fmt.Println("Invalid number of args")
                    return nil, errors.New("Expected at least two arguments for loan application creation")
                }
            
                var loanApplicationId = args[0]
                var loanApplicationInput = args[1]
            
                err := stub.PutState(loanApplicationId, []byte(loanApplicationInput))
                if err != nil {
                    fmt.Println("Could not save loan application to ledger", err)
                    return nil, err
                }
            
                fmt.Println("Successfully saved loan application")
                return nil, nil
            }

        2.从账本获取数据(SELECT)    
            func GetLoanApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
                fmt.Println("Entering GetLoanApplication")
            
                if len(args) < 1 {
                    fmt.Println("Invalid number of arguments")
                    return nil, errors.New("Missing loan application ID")
                }
            
                var loanApplicationId = args[0]
                bytes, err := stub.GetState(loanApplicationId)
                if err != nil {
                    fmt.Println("Could not fetch loan application with id "+loanApplicationId+" from ledger", err)
                    return nil, err
                }
                return bytes, nil
            }
    
# 四.实现访问控制和权限

    1.成员服务
        负责向用户发放登记和交易证书来响应注册和登记
    2.登记证书（注册）(成员服务中的证书颁发机构将向希望在区块链上交易的用户发放一个登记证书，作为一种身份证明)
    3.交易证书(访问token 理解为 bearer token)（交易证书是通过数学方式从父登记证书 关联登记注册的用户）
    4.属性：每个交易证书可以包含一些由用户定义的属性。

    5.从调用方的交易证书中检索属性
        func GetCertAttribute(stub shim.ChaincodeStubInterface, attributeName string) (string, error) {
            fmt.Println("Entering GetCertAttribute")
            attr, err := stub.ReadCertAttribute(attributeName)
            if err != nil {
                return "", errors.New("Couldn't get attribute " + attributeName + ". Error: " + err.Error())
            }
            attrString := string(attr)
            return attrString, nil
        }
        
        func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
            if function == "CreateLoanApplication" {
                username, _ := GetCertAttribute(stub, "username")
                role, _ := GetCertAttribute(stub, "role")
                if role == "Bank_Home_Loan_Admin" {
                    return CreateLoanApplication(stub, args)
                } else {
                    return nil, errors.New(username + " with role " + role + " does not have access to create a loan application")
                }
        
            }
            return nil, nil
        }

# 五.创建和发出自定义事件

    1.已有的事件 
        区块事件
        链代码事件
        拒绝事件
        注册事件
    2.自定义事件
        type customEvent struct {
            Type        string `json:"type"`
            Description string `json:"description"`
        }
        
        func CreateLoanApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
            fmt.Println("Entering CreateLoanApplication")
        
            if len(args) < 2 {
                fmt.Println("Invalid number of args")
                return nil, errors.New("Expected at least two arguments for loan application creation")
            }
        
            var loanApplicationId = args[0]
            var loanApplicationInput = args[1]
        
            err := stub.PutState(loanApplicationId, []byte(loanApplicationInput))
            if err != nil {
                fmt.Println("Could not save loan application to ledger", err)
                return nil, err
            }
        
            var event = customEvent{"createLoanApplication", "Successfully created loan application with ID " + loanApplicationId}
            eventBytes, err := json.Marshal(&event)
            if err != nil {
                    return nil, err
            }
            err = stub.SetEvent("evtSender", eventBytes)
            if err != nil {
                fmt.Println("Could not set event for loan application creation", err)
            }
        
            fmt.Println("Successfully saved loan application")
            return nil, nil
        
        }

# 六.处理日志

    1.使用 shim 包中的 ChaincodeLogger 类型
        ChaincodeLogger 支持以下日志级别：
            CRITICAL
            ERROR
            WARNING
            NOTICE
            INFO
            DEBUG
    2.设置日志级别
        可以通过 3 种方式设置日志级别：
            shim.SetChaincodeLoggingLevel()：此方法将采用 CORE_LOGGING_CHAINCODE 集中的 core.yaml 文件内指定的日志级别。core.yaml 文件包含设置和部署区块链网络所需的所有配置信息。
            shim.SetLoggingLevel(level LoggingLevel)：此方法将在 shim 级别上设置日志级别。
            ChaincodeLogger.SetLevel(level LoggingLevel)：此方法将在单个记录器实例级别上设置日志级别。
    3.创建、配置和使用 ChaincodeLogger
        func SampleLogging() {
            //Different Logging Levels
            criticalLevel, _ := shim.LogLevel("CRITICAL")
            errorLevel, _ := shim.LogLevel("ERROR")
            warningLevel, _ := shim.LogLevel("WARNING")
            noticeLevel, _ := shim.LogLevel("NOTICE")
            infoLevel, _ := shim.LogLevel("INFO")
            debugLevel, _ := shim.LogLevel("DEBUG")
    
            //Logging level at the shim level
            shim.SetLoggingLevel(infoLevel)
    
            //Create a logger instance
            myLogger := shim.NewLogger("SampleChaincodeLogger")
    
            //Set logging level on logger instance
            myLogger.SetLevel(infoLevel)
    
            //Check logging level
            fmt.Println(myLogger.IsEnabledFor(infoLevel))
    
            //Log statements
            myLogger.Info("Info Message")
            myLogger.Critical("Critical Message")
            myLogger.Warning("Warning Message")
            myLogger.Error("Error Message")
            myLogger.Notice("Notice Message")
            myLogger.Debug("Debug Message")
    
        }

# 七.常见问题和最佳实践
    q1.如何将文件（图像、音频、视频、PDF 等）存储在区块链中？
        an1.将所有文件/对象存储为 base64 编码字符串。客户端应用程序将文件/对象转换为 base64 编码字符串，并将它作为输入参数发送给链代码函数。然后链代码将它作为字节数组存储在键/值存储中。
        an2.实际的文件/对象内容存储在区块链以外的地方；例如，存储在 IBM Bluemix Object Storage 服务 中。仅将文件/对象的链接/引用/ID 连同文件/对象的哈希值一起存储在区块链上。存储哈希值可确保在区块链外对文件/对象的任何篡改都能被相关方/实体检测出来。

    q2.如何避免将私有业务逻辑/合同细节泄漏给网络中的所有对等节点？
            如何避免将私有业务逻辑/合同细节泄漏给网络中的所有对等节点？
        此问题是在一个供应链场景中提出的，区块链解决方案的一个最终用户不满意在对所有对等节点可见的智能合约中共享私有业务逻辑/合同信息（比如与不同供应商谈判的不同价格）。在 v0.6 中，可以使用外部系统集成来处理这种情况。

        解决方案：对等节点希望保持为私有的业务逻辑/规则/合同，可以作为一组业务规则在外部应用程序（比如服务）中运行。链代码本身能够执行出站调用。所以举例而言，链代码可对业务规则/逻辑服务执行 REST API 调用并获取结果，以便隐藏逻辑，让实际链代码看不见它。

        可以从链代码内与区块链外的系统集成。例如，可以使用链代码与外部数据库、API 等通信。但重要的是确保与这些系统的交互不会让链代码变得不确定。

        局限性和问题

        修改业务规则服务时无需了解其他对等节点，因为它在区块链外运行。根据区块链网络中不同参与方之间的业务交互类型，这可能导致信任问题。
        业务规则服务必须可供区块链网络中运行智能合约/链代码的所有对等节点使用。
        该解决方案可能导致链代码变得不确定。链代码必须是确定性的。简言之，如果多方使用相同的参数调用链代码中的同一个函数，那么结果应该是相同的。例如，如果在与响应关联的链代码函数中使用一个时间戳或计数器，那么多方对链代码的调用会导致不同的结果。这将在区块链网络中的不同对等节点之间产生不一致的账本状态。
        请记住，每次调用链代码都会导致一致性网络中所有对等节点在自己的本地账本副本上调用链代码。

        备注：在目前正在开发的 Hyperledger Fabric v1.0 中，此问题已通过更改架构本身得到系统解决。
