package pluginManage

//add new file
// 注册所有.so插件
import (
	"fmt"
	"github.com/ethereum/collector"
	"os"
	"path/filepath"
	"plugin"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RegisterInfo struct {
	PluginName string   `json:"pluginname"`
	OpCode     map[string]string `json:"option"`
}

// 注册所有.so插件
func SetUpPlugin(manage *PluginManages){
	pluginFiles,_ := filepath.Glob("./plugin/*.so")
	log_path := "./plugin_log"
	_,err := os.Stat(log_path)
	if err == nil || os.IsNotExist(err){
		os.Mkdir(log_path,os.ModePerm)
	}
	for _, value := range pluginFiles {
		fmt.Println("plugin:", value)
		RegisterPlugin(manage, value)
	}
	
}

// 注册某插件
func RegisterPlugin(manage *PluginManages, path string) bool {
	// 使用plugin包调用.so动态库
	plugin, err := plugin.Open(path)
	if err != nil {
		fmt.Println("error open plugin: ", err, "from path :", path)
		os.Exit(-1)
	}
	// fmt.Println("ex",plugin)
	// 获取.so中的Register函数
	register_method, err := plugin.Lookup("Register")
	if err != nil {
		fmt.Println("Can not find register function:Register() in plugin", err, "from path :", path)
		panic(err)
	}
	// 调用Register函数，返回注册信息和错误信息字节数组
	register_res, b_err := register_method.(func() []byte)
	if !b_err{
		panic(b_err)
	}
	var register_info RegisterInfo
	// 将注册信息存入register_info
	err = json.Unmarshal(register_res(), &register_info)
	if err != nil {
		fmt.Println("Can not parse the struct RegisterInfo from the function:Register() in plugin", err, "from path :", path)
		panic(err)
	}
	// 日志文件动态路径
	fmt.Println("Data log path:./plugin_log/" , register_info.PluginName , "datalog")
	register_map := register_info.OpCode
	// 遍历该插件OpCode字典
	for opcode,sendfunc := range(register_map){
		// 每个plugin的每个opcode构造一个对应的monitor
		var monitor MonitorType	// ./monitor.go
		monitor.SetPluginName(register_info.PluginName)
		monitor.SetLogger(register_info.PluginName)
		// fmt.Println("opcode:",opcode,"sendfunc:",sendfunc)
		symGreeter, err := plugin.Lookup(sendfunc)
		if err != nil {
			fmt.Println("Can not find function",sendfunc," in plugin", err, "from path :", path)
			panic(err)
		}
		rcvefunc, ok := symGreeter.(func(*collector.AllCollector) (byte,string))
		if !ok {
			fmt.Println("unexpected type from module symbol")
			os.Exit(0)
		}
		// fmt.Println("rcve",rcvefunc)
		monitor.SetSendFunc(rcvefunc)
		monitor.SetOpcode(opcode)
		monitor.SetIAL_Optinon(opcode)
		// 调用manage注册Opcode(opcode,对应得monitor对象)
		manage.RegisterOpcode(opcode,&monitor)	
	}
	return true
}
