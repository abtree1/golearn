[go基础]
hello-world  :go第一个程序
values :值
variables :变量
short-declarations  :变量定义
functions  :普通方法基础
methods  :实例方法基础,绑定在struct实例上
variadic-functions  :不定参数个数的函数
multiple-return-values  :多返回值函数
closures  :闭包，一种计数器的实现
pointers  :指针基础
random-numbers  :random基础
constants  :常量定义
environment-variables  :环境变量设置
exit  :程序强行退出
reading-files  :文件读取

[cmd相关]
spawning-processes  :通过命令行启动其它程序
execing-processes  :命令行命令调用，根据系统不同会有一些区别
command-line-arguments  :带参数带main函数
command-line-flags  :flag相关操作，暂时不知道有什么用
line-filters  :命令行内容输入,scanner相关

[数据库］
postgres  :数据库sql操作

[error 处理]
defer  :defer基础
errors  :error基础
panic  :panic基础

[time 相关]
time  :时间基本函数
epoch  :时间戳
time-formatting-parsing  :时间类型的格式化
tickers  :定时器,循环返回
timers  :定时器，一次返回
timeouts  :超时

[interface]
interfaces  :interface基础
interface_reflect  :interface反射

[控制相关]
if-else  :基础
for  :基础for循环
range  :range循环
range-over-channels  :通道的range循环
recursion  :递归函数
switch  :switch基础

[数组相关]
arrays  :基本数组
slices  :可变长数组
collection-functions  :字符串数组相关的遍历查询操作
maps  :map基础，map非线程安全，大量异步操作存在性能瓶颈且不安全
redis  :redis基础
sorting  :数组排序
sorting-by-functions  :数组自定义规则排序
structs  :相当于类的结构

[异步相关]
atomic-counters  :原子计数器
				  可以实现共享内存的异步算法执行，不推荐使用
goroutines  :协程基础
channels  :无buffer的channel最基本的使用
channel-synchronization :简单的channel多线控制
channel-buffering  :带buffer的channel，先写后读，也就是写不依赖读
channel-directions  :channel作为函数参数，且只能单向使用，或读或写
select  :通道select基础
non-blocking-channel-operations  :通道select的非阻赛操作
closing-channels  :channel关闭，不可再写入
mutexes  :一种轻量级的锁
rate-limiting  :channel的消息控制机制
signals  :unix sign的notify操作
stateful-goroutines  :带状态带协程，struct指针通道
worker-pools  :工作池，相当于临界去，一个通道放入数据，一个读取

[字符处理相关]
base64-encoding  :简单的字符串编码转换，也可以实现其它基本类型的编码转换
			      encoding库包含大量与字符串相关的东西 
json  :json简单操作
number-parsing  :strconv的数字转换操作
string-formatting  :字符串格式化
string-functions  :操作string的一些函数
regular-expressions  :字符串模式匹配
sha1-hashes  :sha1加密

[http相关］
hello-web  :web hello程序
static-content  :http基础
streaming-http-servers  :http基础
basic-authentication  :基本数据验证,用委托的方式，实现数据验证的监听
canonical-hosts	 :拼接url并监听，wrap方法
http-client  :http get 模拟
middleware  :中间件，相当于绑定的事件处理器
request-logging  :中间件中加入log处理机制
url-parsing  :url解析
request-routing  :url参数获取
responses  :http response基础

[tcp相关]
tcp-base :tcp的基本使用
graceful-shutdown  :协程控制退出,优雅的程序退出机制