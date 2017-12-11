## EngineerCMS
————工程师知识管理系统

Demo(http://112.74.42.44/)
Demo2(http://112.74.42.44:8086/)
——基于engineercms的设代资料管理平台

1. 本系统运行于工程师个人电脑，仅运行exe文件即可实现网络化管理项目知识资料，无需任何服务器环境，免维护，轻量，开源，功能齐全，技术先进，是工程师不可或缺的工具。

一 特性：

1. 后台预定义工程目录类型和分层级别，支持无限级；
2. 建立项目时选择后台预定义的项目类型和目录层级，自动建立树形目录，支持无限级目录；项目支持公开和私有；
3. 任意层级目录下可以添加任意成果；可对任意目录进行权限设置；
4. 成果包含文章、pdf附件和非pdf附件，并将它们分别列出；一个成果如果包含一个pdf文件则直接打开，如果多于一个则打开列表，非pdf附件也是一样；文章采用富文本编辑器，支持图文word直接黏贴；
5. 多人建立相同项目和目录，可以实现目录中的成果同步显示；
6. 深度检索到成果、附件和文章全文；全局检索到局域网内其他cms上的成果、附件和文章；
7. 根据资料编号的规则，批量上传附件时，自动截取编号和名称，并归入对应的目录中；
8. 3种权限方式：IP地址段权限，适用于局域网内相互之间的无障碍访问；注册用户的权限适用于远程VPN访问；用户组权限适用于项目团队协作；
9. 公开和私有的个人日历日程。每个项目提供一个项目日程和大事记时间轴；
10. 上传成果后，自动生成提供给MeritMS的成果清单，可提交给MeritMS系统进行成果统计；
11. 目的是标准化管理自己的知识体系，可以把数据库挂到到服务器上，供其他人查阅；退休后可将自己个人的cms系统导入cms服务器版，实现知识继承。

我们做工程设计的，经常要做设代，现场服务，而且人员会更换比较频繁。拥有这样一个资料管理平台，是很多人的远望，那么有没有简单免费开源的web应用呢？基于engineercms核心的系统，具有：

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

√硬盘中的资料存储与页面的目录保持一致；设代日志等文章中照片按月度存储。

二 todo:

×手机端添加图片发布文档；

×分享文章；

×考勤登记，值班安排；

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
![](https://cloud.githubusercontent.com/assets/10678867/25748719/a340448e-31de-11e7-8341-6502881fa19c.png)
IPAD移动端效果
![](https://user-images.githubusercontent.com/10678867/33826847-3b8ea162-dea1-11e7-9deb-a8b757da50ab.jpg)
项目进度展示
![snap8](https://user-images.githubusercontent.com/10678867/33214091-6b737446-d165-11e7-834e-b728ea4f590e.png)
用户——角色——权限设置，树状目录权限，区分附件扩展名，基于casbin
![snap7](https://user-images.githubusercontent.com/10678867/33214095-7133dc2c-d165-11e7-8d23-aa172042f9e5.png)
![snap3](https://user-images.githubusercontent.com/10678867/33214099-74ec231a-d165-11e7-8430-ef68c1d8610d.png)
![2](https://user-images.githubusercontent.com/10678867/33826926-7ef24c42-dea1-11e7-87a7-0b40c0906578.jpg)
![snap5](https://user-images.githubusercontent.com/10678867/33214106-7853fd70-d165-11e7-88c3-a0db71a1dbab.png)
