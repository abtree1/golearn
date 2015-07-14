测试中web_server_test目录中内容来源：https://github.com/gSchool/go.git

slice 
	len slice的元素个数
 	cap slice的容量

每次slice容量被用完时
	append自动分配两倍插入值数量的容量
	如，原来len = 2 cap = 2，插入3个元素后，len = 5（2 + 3），cap = 8（2 + 3*2），
	此时再插入2个元素， len = 7， cap = 8

go func(Params) 开启轻量级协程
chan（channel）传输数据 （箭头为传输方向）常用于临界区数据传输管理
c := make(chan int, size) 创建,size为缓存大小，当缓存未满时，插入不会阻塞，满时阻塞，直到腾出空间才会接受阻塞
c <- v              	  插入
v := <-c                  取出