// Description: nil channel
// 如果不使用make初始化channel，那么channel变量就是nil
// 对nil channel发送或接收不会引起panic，但会导致永久阻塞
// 对nil channel执行close函数会引起panic
// nil channel的用处：对于包含select语句的循环，如果不希望每次循环都等待select所涉及的所有channel
// 可以先将某些channel设置为nill，等到发送值准备就绪后，再将channel设置为非nil值并执行发送操作

package main

func main() {

}