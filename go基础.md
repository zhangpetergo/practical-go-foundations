# go基础

事项表

## 第一天

### 议程

- 字符串和格式化输出
  - 什么是字符串？
  - Unicode 基础
  - 使用 fmt 包进行格式化输出
- 调用 REST API
  - 使用 net/http 进行 HTTP 调用
  - 定义结构
  - 序列化 JSON
- 处理文件
  - 处理错误
  - 使用 defer 管理资源
  - 使用 io.Reader 和 io.Writer 接口





1.配置vscode环境



![2](./image/1.png)







### 1.字符串

#### Unicode

Unicode是一种编码，原来使用ASCII码，只能对英文和一些字符进行编码。

由于越来越多的国家加入互联网。需要对其他语言进行编码。

采用Unicode



一个Unicode字符通常表示为

> U+4AE0，由U+开始，后面是16进制数，总共两个字节，16位





一个Unicode编码常用的数据类型是int32，这也是Go里面的rune类型



#### UTF8

> **UTF8是一个将Unicode码点编码为字节序列的变长编码**。UTF8编码由Go语言之父Ken Thompson和Rob Pike共同发明的，现在已经是Unicode的标准。



### 2.REST API和JSON



我们输出的body和我们浏览器看到的body输出格式不一样的原因是

User-Agent



浏览器的body是面向人的，我们在程序中输出的body，是面向go的





#### JSON

如果我们想自定义结构体字段的名称，使用tag

```
type Reply struct{
	Name string
	NumRepos `json:"public_repos" `
}
```





在vscode中，我们ctrl + shift + p，使用go tag，可以给struct加tag

![image-20221024193437446](./image/2.png)



> 





### 3.处理文件

> 为什么打开文件后要关闭文件，
>
> 因为一台服务器上的文件描述符打开个数是有限制的
>
> ulimit -a，
>
> 为了避免资源的耗尽，出现错误，所以我们要关闭文件。



指向文件的指针是nil

```go
package main

import "os"

func main() {
	var f *os.File
	f.Close()
}
```

文件依然正常关闭



gzip

> gz结尾的文件表示是用gzip压缩的文件
>
> 但通常gzip仅用来压缩单个文件。多个文件的压缩归档通常是首先将这些文件合并成一个[tar](https://zh.m.wikipedia.org/wiki/Tar_(计算机科学))文件，然后再使用gzip进行压缩，最后生成的`.tar.gz`或者`.tgz`文件就是所谓的“tar压缩包”或者“tarball”。







## 第二天

- Sorting
  - 使用切片
  - 写methods
  - 理解接口
- 捕获恐慌
  - 内置的recover功能
  - 命名的返回值
- 处理文本
  - 使用bufio.Scanner逐行读取
  - 使用正则表达式
  - 使用maps



捕获panic

> 当我们不想程序随着恐慌而崩溃

```go
defer func() {
		if e := recover(); e != nil {
			log.Println("Error", e)
			err = fmt.Errorf("%v", e)
		}
}()
```

使用内置的recover，同时日志输出error
