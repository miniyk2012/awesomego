tcp粘包

参考: https://www.cnblogs.com/li-peng/p/13153569.html

解决方案: 
example0: 会产生粘包
example1: 规定所有的数据包长度为1024byte，如果不够则补充至1024长度
example2: 通过数据头部来解析数据包长度，比如用4个字节来当数据头，保存每个实数据包的长度。




