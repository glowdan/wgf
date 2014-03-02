# Demo2 － echoback

此篇demo在hello world之外，综合mvc的基本使用，并介绍在使用wgf时的一些推荐方法。

> **预期功能**
>
> 1. 用户输入自己的姓名，提交后显示用户输入的姓名及此姓名的md5值。
> 2. 要求在model中计算md5值，并在model中打印日志。
> 3. 要求内容页可以通过"/echo/姓名"形式的url快速访问。
> 4. 页面的头部显示系统当前时间。

## 文件布局

* workspace/
	* src/
		* app/
			* app.go
			* action/
				* index.go
				* echo.go
			* model/
				* text.go
	* view/
		* header.tpl
		* footer.tpl
		* index.tpl
		* echo.tpl
	
	* conf/
		* wgf.ini
		* router.ini
		* app.ini

## 代码实现

**设计思路**

action(index.go)用来显示表单，供用户输入自己的用户名并提交

	
	

