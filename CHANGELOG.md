# API v1.0.3

## Aug, 2020

+ 增加资料购物车，管理员进行批准。

## Jun, 2020

+ 小程序显示文章、文件等混排
+ 增加成果flow开关
+ jsoneditor的update增加admin权限控制
+ 初始化数据库增加admin角色，anonymous角色等
* 接口调整：maincontrollers下的pdf，wxpdf，getwxpdf等调整到attachmentcontrollers里
* attatch拼写错误，改成attach

## Nov, 2019

+ 采用bootstrap fileinput来实现上传excel到服务端，解析后和规范标准库进行比对，将规范号填入excel，供用户下载。
+ 增加了flow和flowchart，engineercms从此具备了文档状态功能，具备了通用事务流程工程。

## Aug, 2019

* 修复mysql中数据表casbin_rule；
* 修复standards规范标准库的导入错误；
* 修复批量上传文档修改命名中的#和/；
+ 增加小程序公告——车辆安排信息；
+ 增加onlyoffice历史版本；
+ 增加onlyoffice文档转换API；


# API v1.0.2

## Oct, 2019

+ vue.js结合beego，增加了文档状态和事务流程管理

## Aug, 2019

* 修复mysql中数据表casbin_rule；
* 修复standards规范标准库的导入错误；
* 修复批量上传文档修改命名中的#和/；
+ 增加小程序公告——车辆安排信息；
+ 增加onlyoffice历史版本；
+ 增加onlyoffice文档转换API；


# API v1.0.1

## may, 2019

* 修复mysql中数据表casbin_rule-未完成
* 修复后台权限配置时，权限查询显示树状目录不正确；
+ 增加小程序打卡，月度统计；
+ 增加小程序待办事件；


## jan, 2019

* 修复批量上传文件不带编号时，存储的文件名重复，导致无法下载为bug
+ 增加proj products页面为2种供选择，一个是协作页面，可以编辑dwg和office文档，一个是直接下载


## Dec, 2018

* 增加项目里office或wps的协作
* 增加项目里dwg文件的在线编辑和保存到服务器中，利用mxdraw梦想控件，必须用ie浏览器或梦想控件官方提供的企业版chrome浏览器
* 小程序端增加注册，服务端增加资源访问确权

## Nov, 2018

* 修复搜索项目里的成果，分页，搜索结果里再次搜索，采用无限条件动态条件一个字段对应多值的检索方法
* 将规范管理standardms和对标系统集成进来
* 小程序增加搜索某级目录下所有成果，分享链接，查阅pdf
* 小程序增加检索规范和分享链接

## Oct, 2018

* 修复project赋权时，casbin_rule的资源路径出现两个//的bug，原因是判断里少了||projparentidpath==“$#”
* 完成了项目模板里权限继承代码和测试
* 小程序体验版增加底部的tabnav和顶部的switch导航条

## Sep, 2018

* 后台查看日志，读取beego生成的日志文件，生成json数据，填充表格
* conf定制导航条9个菜单，并设置默认首页
* conf里定制微信小程序访问的目录id和域名
* swagger API自动化文档
* 自动建立空数据库时，添加admin用户，密码也是admin
* onlyoffice协作页面的文档地址url带sessionid，方便确权
* dwg预览页面的文件下载地址url带sessionid
* 小程序上传图文，记录用户openId，方便和注册用户比对

# TODO

* 成果状态
* 成果确认
* ztree拖曳项目目录
* 成果移动
* 每月各个目录里成果数量增量图形展示
* 只显示自己权限的目录
* 根据选择，多选框和单选框进行切换，方便批量操作
* 建立项目时，建立默认admin角色的权限；系统默认anonymous角色权限的设置
* 集成电子规范管理系统进来
* 添加文章评论，完善文章展示页面（参考WordPress）

# API v1.0 

## Aug, 2018

* 项目模板，包含权限继承
* 微信小程序访问文章和发布图文
* merit成果里根据权限查看组织结构
