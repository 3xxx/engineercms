## EngineerCMS
————工程师知识管理系统

Demo(http://112.74.42.44/)

Demo2(http://112.74.42.44:8086/)
——基于engineercms的设代资料管理平台（真实应用环境）

1. 本系统采用go语言开发，运行文件为编译后的二进制exe文件，所以无需像其他语言（php、nodejs、java等语言）编写的web应用那样，需要配置运行服务环境。
本系统既可以运行于工程师个人电脑，也可以放到服务器上运行，仅运行exe文件即可实现网络化管理项目知识资料，免维护，轻量，开源，功能齐全，采用大量开源的先进插件，是工程师不可或缺的工具。

2. 数据库采用sqlite嵌入式数据库，所以也无需配置数据库服务环境。

3. 开箱即用，无需网络开发知识。

4. 新增加对onlyoffice document server的二次开发，实现企业的实时文档协作，非常方便，避免了文档的汇总等繁琐事物，效率大大提高，协作更加优雅有趣。

一 特性：

1. 后台预定义工程目录类型和分层级别，支持无限级；
2. 建立项目时选择后台预定义的项目类型和目录层级，瞬间自动建立树形目录，支持无限级目录；项目支持公开和私有；
3. 任意层级目录下可以添加任意成果；可对任意目录进行权限设置；
4. 成果包含文章、pdf附件和非pdf附件，并将它们分别列出；一个成果如果包含一个pdf文件则直接打开，如果多于一个则打开列表，非pdf附件也是一样；文章采用富文本编辑器，支持图文word直接黏贴；
5. 多人建立相同项目和目录，可以实现目录中的成果同步显示；
6. 深度检索到成果、附件和文章全文；全局检索到局域网内其他cms上的成果、附件和文章；
7. 根据资料编号的规则，批量上传附件时，自动截取编号和名称，并归入对应的目录中；
8. 3种权限方式：IP地址段权限，适用于局域网内相互之间的无障碍访问；注册用户的权限适用于远程访问；用户组（角色）权限适用于项目团队协作；
9. 公开和私有的个人日历日程。每个项目提供一个项目日程和大事记时间轴；
10. 上传成果后，自动生成提供给MeritMS的成果清单，可提交给MeritMS系统进行成果统计；
11. 成果提交给MeritMS后进行校审流程；详见MeritMS；
11. 目的是标准化管理自己的（项目团队的、公司的）知识体系，同时方便其他人根据权限查阅；退休后可将自己个人的cms系统导入cms服务器版，实现知识继承。

二 应用案例：

我们做工程设计的，经常要做设代，现场服务，而且人员会更换比较频繁。拥有这样一个资料管理平台，是很多人的愿望，那么有没有简单免费开源的web应用呢？基于engineercms核心的系统，具有：

√参建单位各自自由建立自己的目录，共享资料，避免大家重复存储资料。

√关键字检索和全文检索。

√图纸易得，能更好地控制工程质量——pdf文件参建单位都可以阅读，提高效率，减少障碍。 

√基于“用户——角色——权限”的权限设置。独创的根据文件扩展名来进行权限管理。

√会议室和车辆的预定，用餐人数计划统计等常用功能。

√wiki技术讨论。

√工程大事记，工程进展时间轴等潮玩意儿。

√独创的pdf连续查阅；

√IPAD移动端无障碍；

√设代日记图文并茂记录现场进度；

√项目甘特图进度；

√硬盘中的资料存储与页面的目录保持一致；设代日志等文章中照片按月度存储；

√分享文章到微信。

√√√ONLYOFFICE实时文档协作支持。

二 todo:

×手机端添加图片发布文档；

×考勤登记，值班安排；

## 下载和安装

在release标签中下载二进制文件和源码压缩包。[https://github.com/3xxx/EngineerCMS/releases](https://github.com/3xxx/EngineerCMS/releases)

将二进制文件放到源码文件夹下直接运行即可。

## Quick Start

* 参加quickstart快速开始。

## Documentation

* [中文文档]——请查阅document文件夹

## 免费开源和问题反馈

* 开源地址[https://github.com/3xxx/engineercms/](https://github.com/3xxx/engineercms/)
* 问题反馈: [https://github.com/3xxx/engineercms/issues](https://github.com/3xxx/engineercms/issues)

## LICENSE

EngineerCMS source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).

ONLYOFFICE实时文档协作效果

![engineercms onlyoffice](https://user-images.githubusercontent.com/10678867/36413347-e5d75f26-1658-11e8-8818-daaeec4746e5.jpg)

![onlyoffice word](https://user-images.githubusercontent.com/10678867/36413270-9fc09c00-1658-11e8-817e-3e58021a8253.jpg)

![onlyoffice powerpoint](https://user-images.githubusercontent.com/10678867/36413278-a61110ee-1658-11e8-9955-14241b8b8bd7.jpg)

![onlyoffice excel](https://user-images.githubusercontent.com/10678867/36413285-ab0dcd8a-1658-11e8-9a11-4c94709efea8.jpg)

手机端访问文章效果

![](https://user-images.githubusercontent.com/10678867/34637355-7391a26e-f2ef-11e7-8c9d-9f3edcce9004.png)

IPAD移动端效果

![](https://user-images.githubusercontent.com/10678867/33826847-3b8ea162-dea1-11e7-9deb-a8b757da50ab.jpg)
项目进度展示
![snap8](https://user-images.githubusercontent.com/10678867/33214091-6b737446-d165-11e7-834e-b728ea4f590e.png)
用户——角色——权限设置，树状目录权限，区分附件扩展名，基于casbin
![snap7](https://user-images.githubusercontent.com/10678867/33214095-7133dc2c-d165-11e7-8d23-aa172042f9e5.png)
![snap3](https://user-images.githubusercontent.com/10678867/33214099-74ec231a-d165-11e7-8430-ef68c1d8610d.png)
![2](https://user-images.githubusercontent.com/10678867/33826926-7ef24c42-dea1-11e7-87a7-0b40c0906578.jpg)
![snap5](https://user-images.githubusercontent.com/10678867/33214106-7853fd70-d165-11e7-88c3-a0db71a1dbab.png)
