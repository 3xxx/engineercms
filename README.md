
## EngineerCMS
————工程师知识管理系统

[![License](https://img.shields.io/hexpm/l/plug.svg)](LICENSE) [![Build Status](https://travis-ci.org/3xxx/engineercms.svg?branch=master)](https://travis-ci.org/3xxx/engineercms) [![GoDoc](https://godoc.org/github.com/3xxx/engineercms?status.svg)](https://godoc.org/github.com/3xxx/engineercms)

[Demo](https://zsj.itdos.com/)
——基于engineercms的设代资料管理平台（真实应用环境）

1. 本系统采用go语言（基于[beego](https://github.com/astaxie/beego)框架）开发，运行文件为编译后的二进制可执行文件，所以无需像其他语言（php、nodejs、java等语言）编写的web应用那样，需要配置运行服务环境。
本系统既可以运行于工程师个人电脑，也可以放到局域网、公网服务器上运行，仅运行可执行文件即可实现网络化管理项目知识资料，免维护，轻量，开源，功能齐全，采用大量开源的先进插件，是工程师不可或缺的工具。
2. 数据库采用sqlite嵌入式数据库，所以也无需配置数据库服务环境。
3. 开箱即用，无需网络开发知识。
4. 新增加对onlyoffice document server的二次开发，实现企业的实时文档协作，非常方便，避免了文档的汇总等繁琐事物，效率大大提高，协作更加优雅有趣。除了支持office的docx，xlsx及pptx格式外，还支持国产wps，et和dps格式。
5. 在线直接预览dwg文件，避免了图纸转换成pdf的麻烦，也避免了先下载附件然后用本机电脑打开图纸的麻烦，方便设计人员查阅图纸。
6. 采用最新的froala富文本编辑器，支持word图文直接粘贴发布，word中的图片自动上传，发布文章轻松快捷，文章支持视频和文件附件。
7. 微信小程序客户端访问和添加图文，打开微信，搜索“珠三角设代”或“青少儿书画”即可看到小程序了，小程序端也[开源](https://github.com/3xxx/wechatengineercms)。
8. 微信小程序多项目切换。
8. swagger API自动化文档，方便前后端分离。
9. 可在conf里定制9个导航条菜单。
10. 后台查看日志。
11. 通用的文档流程设置。文档审批，文档校审，合同评审流程，图纸校审流程，……
12. 整合了[mindoc](https://github.com/lifei6671/mindoc)，实现了在线创作、查阅、分享、导入、导出电子书籍。
13. 文件分享提取码。

一 特性：

1. 后台预定义工程目录类型和分层级别，支持无限级；目录采用懒加载模式，支持百万级目录数据快速显示；
2. 建立项目时选择已有项目作为模板，继承已有项目权限，或选择后台预定义的项目类型和目录层级，瞬间自动建立成百上千的树形目录，支持无限级目录；目录采用懒加载显示，百万级数据快速显示；项目支持公开和私有；
3. 任意层级目录下可以添加任意成果；可对任意目录进行权限设置；
4. 成果包含文章、pdf附件和非pdf附件，并将它们分别列出；一个成果如果包含一个pdf文件则直接打开，如果多于一个则打开列表，非pdf附件也是一样；文章采用富文本编辑器froala，支持图文word直接黏贴，自动上传图片；成果数据采用后端分页，百万级成果数据快速显示；
5. 成果间的关联：比如先出的施工图，然后一段时间后再出了对这个图纸的修改通知单，那么修改单关联上这个图纸后，每次看这个图纸，就知道这个图纸有修改了，需要引起注意；
6. 多人建立相同项目和目录，可以实现目录中的成果同步显示，方便团队协作；
7. 深度检索到成果、附件和文章全文；全局检索到局域网内其他cms上的成果、附件和文章；
8. 根据资料编号的规则，批量上传附件时，自动截取编号和名称，并归入对应的目录中；
9. 3种权限方式：IP地址段权限，适用于局域网内相互之间的无障碍访问；注册用户的权限适用于远程访问；用户组（角色）权限适用于项目团队协作；还特别为pdf扩展名添加了权限，即，这种角色只能访问pdf文件。
10. 公开和私有的个人日历日程。每个项目提供一个项目日程和大事记时间轴；
11. 上传成果后，自动生成提供给MeritMS的成果清单，可提交给MeritMS系统进行成果统计；
12. 成果提交给MeritMS后进行校审流程；详见[MeritMS](https://github.com/3xxx/MeritMS)；
13. 目的是标准化管理自己的（项目团队的、公司的）知识体系，同时方便其他人根据权限查阅；退休后可将自己个人的cms系统导入cms服务器版，实现知识继承。

二 应用案例：

我们做工程设计的，经常要做设代，现场服务，而且人员会更换比较频繁。拥有这样一个资料管理平台，是很多人的愿望，那么有没有简单免费开源的web应用呢？基于engineercms核心的系统，具有：

√参建单位各自自由建立自己的目录，共享资料，避免大家重复存储项目资料，比如会议纪要，法律法规。

√关键字检索和全文检索。

√图纸易得，能更好地控制工程质量——pdf设计文件参建单位都可以阅读，提高效率，减少障碍。

√基于“用户——角色——权限”的权限设置。独创的根据文件扩展名来进行权限管理。

√会议室和车辆的预定，用餐人数计划统计等常用功能。

√wiki技术讨论。

√工程大事记，工程进展时间轴等潮玩意儿。

√独创的pdf连续查阅；

√IPAD、手机移动端无障碍；

√设代日记图文并茂记录现场进度，支持视频格式；

√项目甘特图展示工作进度；

√硬盘中的资料存储与页面的目录保持一致；设代日志等文章中照片按月度存储；

√分享文章到微信。

√√√ONLYOFFICE实时文档协作支持。除了支持office的docx，xlsx及pptx格式外，还支持wps，et和dps格式。

√ONLYOFFICE中文档的历史版本对照功能。

√ONLYOFFICE中文档的格式转换功能，如word转pdf。

√ONLYOFFICE中文档协作的权限设置，采用casbin。

√在线预览dwg文件。

√手机端添加图片、视频发布文档；——已开发小程序“珠三角设代”应用中并[开源](https://github.com/3xxx/wechatengineercms)。

√考勤登记，值班安排；

√小程序图纸查阅权限，用户注册设计；

√小程序打卡；

√小程序多人写同一天的设代日志；

√小程序公告，展示车辆、会议室的调度、占用；

√文档关联完善：编辑关联，关联打开链接；

√自定义业务流程：单线传递，并行传递；

√资料“购物车”，管理员批准后可下载资料；

二 todo:

× 规范标准库编辑，爬虫；

× 首页搜索后转入新页面，首页vue.js；

× 读取文件属性——完成时间，作为月度统计依据，或提供选择，以上传时间为统计口径，生成某个目录下月度成果报表，与上个月对比柱状图；

× 项目目录编辑完善：ztree目录拖动；

× 多关键字检索；

× 项目合并和拆分：将2个项目目录合并起来，其中一个项目合并到另一个项目的子目录中，手动合并文件夹，自动修改数据库；

× 流媒体服务；

× websocket聊天室；

## 下载和安装

在release标签中下载二进制文件和源码压缩包(不再提供最新编译文件，请到网盘下载)。[https://github.com/3xxx/EngineerCMS/releases](https://github.com/3xxx/EngineerCMS/releases)

将二进制文件放到源码文件夹下直接运行即可。

或者去百度网盘下载，直接运行。[链接：https://pan.baidu.com/s/1MDJ-QfCmv_LiychDSLn8jw 提取码：nhar ](https://pan.baidu.com/s/1MDJ-QfCmv_LiychDSLn8jw)

Linux系统下请替换掉执行文件engineercms(linux)

不清楚的，可以加我QQ504284或者微信hotqin999聊，也可参考[quickstart](https://github.com/3xxx/engineercms/blob/master/quickstart%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B.txt)。

[wiki](https://github.com/3xxx/EngineerCMS/wiki)。包括linux系统下的编译。linux系统下的部署参见网盘中的文档。

技术开发过程详见[CSDN我的博客](https://blog.csdn.net/hotqin888)
系统简单使用见[mindoc文档](https://zsj.itdos.com/docs/engineercms_008)

## Quick Start

* 参见[quickstart](https://github.com/3xxx/engineercms/blob/master/quickstart%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B.txt)快速开始。包括后台导入Excel用户，设置用户角色和权限。

* linux系统上的部署参见网盘中的文档。

## Documentation

* [中文文档]——请查阅document文件夹
* 系统简单使用见[mindoc文档](https://zsj.itdos.com/docs/engineercms_008)

## 免费开源和问题反馈

* 开源地址[https://github.com/3xxx/engineercms/](https://github.com/3xxx/engineercms/)
* 问题反馈: [https://github.com/3xxx/engineercms/issues](https://github.com/3xxx/engineercms/issues)

## 第三方包

1. [beego框架](https://github.com/astaxie/beego)
2. [casbin权限管理](https://github.com/casbin/casbin)
3. [simplejson](https://github.com/bitly/go-simplejson)
4. [excelize](https://github.com/360EntSecGroup-Skylar/excelize)
5. [xlsx](https://github.com/tealeg/xlsx)
6. [goquery](https://github.com/PuerkitoBio/goquery)
7. [go-sqlite](https://github.com/mattn/go-sqlite3)
8. [mahonia](https://github.com/axgle/mahonia)
9. [flow文档流程](https://github.com/js-ojus/flow)[我改造的flow文档流程](https://github.com/3xxx/flow)
10. [xorm](https://github.com/go-xorm/xorm)
11. [gorm](https://github.com/jinzhu/gorm)
12. [unioffice](github.com/unidoc/unioffice)
13. [pdfcpu](https://github.com/pdfcpu/pdfcpu)

## 前端
1. [onlyoffice document server文档协作](https://github.com/ONLYOFFICE/DocumentServer)
2. [pdf.js阅览pdf](https://github.com/mozilla/pdf.js)
3. [html2canvas](https://github.com/niklasvh/html2canvas)
4. [jQueryGantt甘特图](https://github.com/robicch/jQueryGantt)
5. [multyselect](https://github.com/davidstutz/bootstrap-multiselect)
6. [bootstrap treeview树状目录](https://github.com/patternfly/patternfly-bootstrap-treeview)
7. [fullcalendar日历](https://github.com/fullcalendar/fullcalendar)
8. [daterangerpicker](https://github.com/dangrossman/daterangepicker)
9. [datetimepicker](https://github.com/smalot/bootstrap-datetimepicker)
10. [datepicker](https://github.com/uxsolutions/bootstrap-datepicker)
11. [select2](https://github.com/select2/select2)
12. [x-editable表格在线编辑](https://github.com/vitalets/x-editable)
13. [ztree树状目录](https://github.com/zTree/zTree_v3)
14. [request](https://github.com/mozillazg/request)
15. [froala富文本编辑器](https://github.com/froala/wysiwyg-editor)
16. [ueditor百度富文本编辑器](https://github.com/fex-team/ueditor)
17. [webupload百度文件上传](https://github.com/fex-team/webuploader)
18. [jqfileupload](https://github.com/blueimp/jQuery-File-Upload)
19. [bootstrap-table表格](https://github.com/wenzhixin/bootstrap-table)
20. [jquery.form.js表单上传文件](http://jquery.malsup.com/form/)
21. [vue.js](https://github.com/vuejs/vue)
22. [axios](https://github.com/axios/axios)
23. [element组件](https://github.com/ElemeFE/element)
24. [vxe-table表格在线编辑](https://github.com/xuliangzhan/vxe-table)

## LICENSE

EngineerCMS source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).

ONLYOFFICE实时文档协作效果——除了支持office的docx，xlsx及pptx格式外，还支持国产wps，et和dps格式。

![engineercms onlyoffice1](https://user-images.githubusercontent.com/10678867/38768484-0a55e06e-4027-11e8-9871-fc65e1686408.png)

![onlyoffice word](https://user-images.githubusercontent.com/10678867/36413270-9fc09c00-1658-11e8-817e-3e58021a8253.jpg)

![onlyoffice powerpoint](https://user-images.githubusercontent.com/10678867/36413278-a61110ee-1658-11e8-9955-14241b8b8bd7.jpg)

![onlyoffice excel](https://user-images.githubusercontent.com/10678867/36413285-ab0dcd8a-1658-11e8-9a11-4c94709efea8.jpg)

在线直接预览dwg图纸文件：

![default](https://user-images.githubusercontent.com/10678867/41599295-4fb72070-7405-11e8-82fb-e746950ffa0d.gif)

手机端访问文章效果

![Snap4](https://user-images.githubusercontent.com/10678867/59145056-6d0a9d00-8a11-11e9-88b0-f057100f4e3a.png)

微信小程序富文本编辑器添加文章：

![wx05](https://user-images.githubusercontent.com/10678867/59144747-c1ac1900-8a0d-11e9-9735-cab114eca5e8.png)

IPAD移动端效果

![](https://user-images.githubusercontent.com/10678867/33826847-3b8ea162-dea1-11e7-9deb-a8b757da50ab.jpg)
项目进度展示
![snap8](https://user-images.githubusercontent.com/10678867/33214091-6b737446-d165-11e7-834e-b728ea4f590e.png)
用户——角色——权限设置，树状目录权限，区分附件扩展名，基于casbin
![snap7](https://user-images.githubusercontent.com/10678867/33214095-7133dc2c-d165-11e7-8d23-aa172042f9e5.png)
![snap3](https://user-images.githubusercontent.com/10678867/40973567-44c8dfc8-68f7-11e8-8d92-b67c56ed9c08.png)
![2](https://user-images.githubusercontent.com/10678867/33826926-7ef24c42-dea1-11e7-87a7-0b40c0906578.jpg)
![snap5](https://user-images.githubusercontent.com/10678867/33214106-7853fd70-d165-11e7-88c3-a0db71a1dbab.png)

froala富文本编辑器支持word图片自动上传，视频和附件
![snap13](https://user-images.githubusercontent.com/10678867/42722537-3c71c216-8780-11e8-8065-f1538bbcad18.png)
![snap6](https://user-images.githubusercontent.com/10678867/42722539-40376bc6-8780-11e8-8173-1f6e9e60ef14.png)
