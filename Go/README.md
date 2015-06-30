## Go 练习册

**第 0000 题：**创建二维 slice

实现 `Pic`。它返回一个 slice 的长度 `dy`，和 slice 中每个元素的长度的 8 位无符号整数 `dx`。当执行这个程序，它会将整数转换为灰度（好吧，蓝度）图片进行展示。

图片的实现已经完成。可能用到的函数包括 `(x+y)/2` 、 `x*y` 和 `x^y`（使用 math.Pow 计算最后的函数）。

（需要使用循环来分配 `[][]uint8` 中的每个 `[]uint8`。）

（使用 uint8(intValue) 在类型之间进行转换。）

**第 0001 题：**用 map 记录词频

实现 `WordCount`。它应当返回一个含有 s 中每个 “词” 个数的 map。函数 wc.Test 针对这个函数执行一个测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有帮助。

**第 0002 题：**斐波纳契闭包

实现一个 fibonacci 函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数。
