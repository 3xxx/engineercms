package tool

// 假设项目的目录结构如下:

// demo/main.go demo/a/1.go demo/b/c/2.go
// 强调一下,这时候/demo/b下,没有文件
// 也就是说,没有 package b
// 如果这时候,在main.go或者1.go里,import "package demo/b/c"的话,就会出现"malformed module path"错误.

// 如果项目目录是如下:

// demo/main.go demo/a/1.go demo/b/3.go // 这里增加了一个文件 demo/b/c/2.go
// 这个时候,肯定就已经有package b了.
// 你会发现"malformed module path"错误没有了.
