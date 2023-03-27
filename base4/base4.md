### Base4 的关键： **动态路由**

其实这一章节的关键点就是要理解前缀树

什么是前缀树？

参考leetcode的题目：
https://leetcode.cn/problems/implement-trie-prefix-tree/

这个题目非常推荐去做一下

简述一下思路：

> 思考一下树的结构
> 
> 这里有三个功能函数： insert search startsWith
> 
> 然后我插入一串字符串，目的要达到
> 
> 1  能够插入字符串
> 
> 2  如果我们再search之前已经插入了，就返回true，否则返回false
> 
> 3  如果我们匹配startWith，比如说插入了`connect`，`t.startWith("conn")`返回true
> 
> 那我们其实就是说，这里只插入英文，所以我们一个节点存储26个英文字母就可以