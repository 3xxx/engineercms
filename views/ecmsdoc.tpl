<!DOCTYPE html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">

  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/docs.min.css"/>

<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
<script type="text/javascript" src="/static/js/docs.min.js"></script>

</head>


<!-- <div class="navbar-wrapper"> -->
  <div class="container-fill">
        {{template "navbar" .}}
  </div>
<!-- </div> -->

<body>
<div class="container">

  <div class="col-md-3" role="complementary">
    <nav class="bs-docs-sidebar hidden-print hidden-xs hidden-sm affix">
      <ul class="nav bs-docs-sidenav" data-offset-top="10">
        <li class="">
          <a href="#about">简介</a>
          <ul class="nav">
            <li class="">
              <a href="#about-function">功能</a>
            </li>
            <li class="">
              <a href="#about-use">使用方式</a>
            </li>
          </ul>
        </li>

        <li class="">
          <a href="#backroud">后台设置</a>
          <ul class="nav">
            <li class="">
              <a href="#backroud-callendar">个人日历事件</a>
            </li>
            <li class="">
              <a href="#backroud-projcate">项目分级目录</a>
            </li>
            <li>
              <a href="#backroud-organize">组织结构</a>
            </li>
            <li>
              <a href="#backroud-adduser">添加用户</a>
            </li>
            <li>
              <a href="#backroud-addip">添加IP地址段</a>
            </li>
            <li>
              <a href="#backroud-editproj">编辑项目目录</a>
            </li>
            <li>
              <a href="#backroud-addprojsynip">设置项目同步IP</a>
            </li>
          </ul>
        </li>
        
        <li>
          <a href="#addproj">添加项目</a>
        </li>

        <li>
          <a href="#addprojcallendar">项目日历事件</a>
        </li>

        <li>
          <a href="#addprod">向项目中添加成果</a>
          <ul class="nav">
            <li>
              <a href="#addprod-multi">批量模式</a>
            </li>
            <li>
              <a href="#addprod-multiattachment">多附件模式</a>
            </li>
            <li>
              <a href="#addprod-img">图文</a>
            </li>
          </ul>
        </li>

        <li>
          <a href="#synch">同步资料</a>
        </li>

        <li>
          <a href="#checkexamin">校审流程</a>
        </li>

        <li>
          <a href="#standard">规范管理</a>
        </li>

        <li>
          <a href="#onlyoffice">文档协作</a>
        </li>

      </ul>
      <a class="back-to-top" href="#top">返回顶部</a>

    </nav>
  </div>

<div class="col-md-9" role="main">
  <div class="bs-docs-section">

    <h1 id="about" class="page-header">
      <a class="anchorjs-link " href="#about" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      简介
    </h1>

    <p class="lead">
      EngineerCMS 具有功能：文档管理，文档协作，图文发布，团队协同，结合MeritMS进行校审流程，成果统计等。</p>
    <p class="lead">
      创新：集ONLYOFFICE文档协作，档案系统，ProjectWise协同设计，Redmine项目管理，SharePoint发布，网络云盘等特色；是基于个人电脑的知识管理与共享的微服务。
    </p>
    <p class="lead">
      特色：采用golang语音编写，天生部署简单，无需服务环境，无需数据库服务，省心。
    </p>

    <div class="bs-callout bs-callout-warning" id="jquery-required">
      <h4>1.工程师为什么要知识管理</h4>
      <p>
        设计是基于资源的，工程师的资源管理得好，应付设计会得心应手，事半功倍！
      </p>
      <h4>2.知识为何要发布和共享</h4>
      <p>
        首先是团队的效率提高，才是真的高，团队效率提高就需要资源共享，避免重复；项目资料管理也要避免重复，参见<a href="http://www.0daydown.com/07/212999.html" target="_blank">ProjectWise</a>；“……工程师有 40% 的工作时间都用在了查找和验证要使用的特定信息和文件……”；优秀的资源，也体现一个工程师的素质，从而影响整个团队的进步！
      </p>
      <h4>3.标准化管理资料</h4>
      <p>
        让工程师形成一个标准化的管理模式，方便工作交接；利用EngineerCMS的同步资料功能，可看到团队中其他人的资料；标准化后还可以方便资料整合与继承。
      </p>
      <h4>4.为什么要基于个人电脑</h4>
      <p>
        个人知识是工程师自己的财富，当然要保存在自己身边，由自己支配；Golang语言开发的服务直接运行，无需服务环境的搭建！
      </p>
      <h4>5.为什么不用成熟的PW等</h4>
      <p>
        参见<a href="http://blog.csdn.net/hotqin888/article/details/75208908" target="_blank">与ProjectWise对比</a>，PW庞杂，需要运维，EngineerCMS简介、轻量，无需部署，又可以进行校审流程和图文发布，具备设代日记，项目日历事件，大事记等特色扩展！
      </p>
    </div>

    <h2 id="about-function">
      <a class="anchorjs-link " href="#about-function" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      功能
    </h2>

  <ol>
    <li>文档、资源管理；文档协作（基于ONLYOFFICE document server）；项目资源同步；</li>
    <li>后台预定义工程目录类型和分层级别，支持无限级；</li>
    <li>建立项目时选择后台预定义的项目类型和目录层级，自动建立树形目录，支持无限级目录；项目支持公开和私有；</li>
    <li>任意层级目录下可以添加任意成果；可对任意目录进行权限设置；</li>
    <li>成果包含文章、pdf附件和非pdf附件，并将它们分别列出；一个成果如果包含一个pdf文件则直接打开，如果多于一个则打开列表，非pdf附件也是一样；文章采用富文本编辑器，支持图文word直接黏贴；</li>
    <li>多人建立相同项目和目录，可以实现目录中的成果同步显示；</li>
    <li>定义组织结构；</li>
    <li> 2种权限方式：IP地址段权限，适用于局域网内相互之间的无障碍访问；基于用户——角色——权限的方式适用于项目团队协作；独创的根据文件扩展名来进行权限管理。</li>
    <li>深度检索到成果、附件和文章全文；全局检索到局域网内其他cms上的成果、附件和文章；</li>
    <li>根据资料编号的规则，批量上传附件时，自动截取编号和名称，并归入对应的目录中；</li>
    <li>公开和私有的个人日历日程。每个项目提供一个项目日程和大事记时间轴；</li>
    <li>上传成果后，自动生成提供给MeritMS的成果清单，可提交给MeritMS系统进行设计成果校审流程和工作量统计；</li>
    <li>目的是标准化管理自己的知识体系，可以把数据库挂到到服务器上，供其他人查阅；退休后可将自己个人的cms系统导入cms服务器版，实现知识继承。</li>
    <li>√硬盘中存储文件的目录与页面目录一致。</li>
    <li>√会议室和车辆的预定，用餐人数计划统计等常用功能。</li>
    <li>√wiki技术讨论。</li>
    <li>√工程大事记，工程进展时间轴等潮玩意儿。</li>
    <li>√独创的pdf连续查阅；</li>
    <li>√IPAD、手机移动端无障碍；</li>
    <li>√设代日记图文并茂记录现场进度；连续查看文章。</li>
    <li>√设计、施工进度甘特图展示；</li>
    <li>√分享到微信；</li>
    <li>√√√支持ONLYOFFICE多人在线编辑文档（实时文档协作）。</li>

  </ol>
    <h2 id="about-use">
      <a class="anchorjs-link " href="#about-use" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      使用方式
    </h2>
  <ol>
    <li>如果是首次使用，请：</li>
    <li> 解压到d:\EngineerCMS\；（其他盘根目录下也行，因为上传附件会使得这个文件夹越来越大，所以，要考虑空间大一些的盘。）</li>
    <li> 修改配置文件conf\app.conf.sample为app.conf,打开app.conf,看到里面的httpport = 80，如果要修改成8080，请修改后保存。runmode = prod表示生产模式运行系统。</li>
    <li> 修改数据库文件database\engineercms.db.sample为engineercms.db。</li>
    <li> 运行engineercms-win64/win32.exe即可在chrome浏览器中输入本地ip（127.0.0.1）和前面设置的端口号（假设是80或8080）进行访问。如果是80端口，则端口号可省略，如http://127.0.0.1。如果运行后闪退，则可能是端口号被占用了，请修改端口号再运行。运行后不要关闭窗口，它是服务。IE浏览器支持不好，推荐使用chrome，可以使用firefox、opra。加入开机启动请自行设置。</li>
  </ol>
  <img src="/static/img/局域网meritecms.png" style="width: 100%">
    <h1 id="backroud" class="page-header">
      <a class="anchorjs-link " href="#backroud" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      后台设置
    </h1>

    <p class="lead">
      在这里可以预设和定制系统。
    </p>

    <h2 id="backroud-callendar">
      <a class="anchorjs-link " href="#backroud-callendar" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      个人日历事件
    </h2>
    <ol>
      <li>填写每天的事件，记录工作生活和计划</li>
        <img src="/static/img/个人日历事件.png" style="width: 100%">
      </ol>
    <h2 id="backroud-projcate">
      <a class="anchorjs-link " href="#backroud-projcate" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      项目分级目录
    </h2>
    <ol>
    <li>进入admin后台页面：http://127.0.0.1/admin；系统已经将127.0.0.1这个本机ip权限设置为1级，进入admin后台页面。admin页面只允许ip权限，不允许登录权限。</li>
    <li>admin页——分级目录：设置项目目录结构，按系统自带的例子理解。</li>
    <img src="/static/img/项目分级目录.png" style="width: 100%">
    </ol>

    <h2 id="backroud-organize">
      <a class="anchorjs-link " href="#backroud-organize" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      组织结构
    </h2>
    <ol>
      <p>admin页——组织结构：设置部门（分院），下级是科室（专业组），也可以无下级。</p>
    </ol>
    <h2 id="backroud-adduser">
      <a class="anchorjs-link " href="#backroud-adduser" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      添加用户
    </h2>
    <ol>
      <p>admin页——用户：可以按规定格式批量导入，也可以逐一添加。用户可以没有科室属性。端口号是这个用户运行cms系统的服务端口号。这里的权限是登录权限，1级权限建议只给自己用，可以向项目里添加成果，不能建立和删除项目；不能进入后台。2级可以下载任意成果附件。3级只能下载pdf格式附件。其他权限只能看。</p>
      <img src="/static/img/权限.png" style="width: 100%">
    </ol>
    <h2 id="backroud-addip">
      <a class="anchorjs-link " href="#backroud-addip" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      添加IP地址段
    </h2>
    <ol>
      <p>admin页——IP地址段：因为本系统适合于有固定ip地址的局域网内使用，所以这里填写自己机器的真实ip，权限设为1。这个1级权限只能给自己，可以建立项目和进入后台。IP权限为2级的，可以下载任意成果附件，3级的只能下载pdf格式附件，其余权限只能看。</p>
    </ol>
    <h2 id="backroud-editproj">
      <a class="anchorjs-link " href="#backroud-editproj" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      编辑项目目录
    </h2>
    <ol>
      <p>admin页——编辑目录：对已经建立的项目目录进行编辑，包括增加选中目录的子目录、删除选中的目录以及修改选中的目录。</p>
    </ol>
    <h2 id="backroud-editproj">
      <a class="anchorjs-link " href="#backroud-editproj" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      设置项目同步IP
    </h2>
    <ol>
      <p>admin页——同步IP：多人建立相同项目及目录，在这里向这个项目中填入他们的ip和端口号，就可以在自己cms中，这个项目的任何目录下同步他们的成果列表过来。</p>
    </ol>

    <h1 id="addproj" class="page-header">
      <a class="anchorjs-link " href="#addproj" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      添加项目
    </h1>
    <p>项目——项目列表页——添加：分级目录就是后台admin页上文第6点定义的。新建项目除了建立数据表中的目录，还在cms系统文件夹attachment中建立目录文件夹，以后所有上传的成果附件都放在这些文件夹中。</p>
    <img src="/static/img/添加项目.png" style="width: 100%">
    <h1 id="addprojcallendar" class="page-header">
      <a class="anchorjs-link " href="#addprojcallendar" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      项目日历事件
    </h1>
    <p>每个项目都有自己的日历事件，勾选大事记，可以生成时间轴。</p>
    <img src="/static/img/时间轴.png" style="width: 100%">
    <h1 id="addprod" class="page-header">
      <a class="anchorjs-link " href="#addprod" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      向项目中添加成果
    </h1>
    <p class="lead">
      有多种添加模式，并且使用3种模式展示，分别是pdf，图文和其他附件。
    </p>
    <h2 id="addprod-multi">
      <a class="anchorjs-link " href="#addprod-multi" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      批量模式
    </h2>
    <ol>
      <p>项目——具体一个项目——成果列表——添加：系统设计上按照 成果—附件来保存文件，成果里可以放pdf格式的附件，可以放dwg、dgn、word、excel、jpg等格式的附件，可以放文章，这些附件的名称和成果的编号+成果的名称是否一直取决于下面的添加方式。添加按钮鼠标放上有提示，3个添加按钮分别用于批量上传、多附件模式和添加文章。批量上传就是不用填写成果编号和成果名称，系统自动根据附件文件名截取编号和名称作为成果编号和成果名称，但上传的附件文件名必须按编号+名称命名；</p>
    </ol>
    <h2 id="addprod-multiattachment">
      <a class="anchorjs-link " href="#addprod-multiattachment" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      多附件模式
    </h2>
    <ol>
      <p>项目——具体一个项目——成果列表——添加：多附件模式指一个成果中包含多个附件，比如一张图有参照图，一份报告有多个章节等；</p>
    </ol>
    <h2 id="addprod-img">
      <a class="anchorjs-link " href="#addprod-img" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      图文模式
    </h2>
    <ol>
      <p>文章模式可以将图文word直接粘贴进来，然后按提示导入word中的图片，非常快捷了。也可以批量上传设代现场照片……</p>
    </ol>
    <h1 id="synch" class="page-header">
      <a class="anchorjs-link " href="#synch" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      同步资料
    </h1>
    <p>当团队中每个人的ecms都建立了相同的项目目录，并且在后台设置了需要同步的IP和端口号，则可以在自己的ecms中同步看到他们ecms中的资料了。</p>
    <p>比如下图模拟5个用户的ecms运行中</p>
    <img src="/static/img/模拟5个用户进行同步.png" style="width: 100%">
    <p>同步后看到他们5个人这个项目下的资料</p>
    <img src="/static/img/资料同步.png" style="width: 100%">

    <h1 id="checkexamin" class="page-header">
      <a class="anchorjs-link " href="#checkexamin" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      校审流程
    </h1>
    <p class="lead">
      当我们向EngineerCMS中上传成果的时候，在后台自动生成了成果清单，只要提交给MeritMS即可。前提是要设置MeritMS的ip地址可接受的端口。
    </p>
    <img src="/static/img/MeritMS校审流程.png" style="width: 100%">

    <h1 id="standard" class="page-header">
      <a class="anchorjs-link " href="#standard" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      规范管理
    </h1>
    <p class="lead">
      规范管理功能为查阅、上传和有效版本库管理；系统设计上不用删除旧规范，旧规范仍然可以查阅，只是系统利用有效版本库进行标识它是否是最新版本的有效规范；系统设计上将规范和有效版本库分开，好处是显而易见的，降低维护繁琐工作。使用中，我们任意上传我们收集到的规范即可，不用管这个规范是否是有效，另一方面，我们导入有效规范版本库数据即可，比如每次发布新的规范版本时，导入这些数据即可。
    </p>
    <img src="/static/img/规范查阅.png" style="width: 100%">
    <img src="/static/img/规范有效库管理.png" style="width: 100%">

    <h1 id="onlyoffice" class="page-header">
      <a class="anchorjs-link " href="#onlyoffice" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      文档协作
    </h1>
    <p class="lead">
      文档协作功能即团队成员可以同时编辑一个文档，支持word,excel,powerpoint，解决多人文档协作需要汇总的问题，比如月报，任务进度控制表，会务安排进度控制，excel计算书，会议纪要等需要多人审阅的文件，发个链接给大家，大家有意见直接在上面修改，简单高效快捷，体验良好。EngineerCMS基于ONLYOFFICE document server进行的二次开发，提供文档管理和更新后的文档存储即提供回调参数给onlyoffice。
    </p>
    <p class="lead">文档管理界面</p>
    <img src="/static/img/EngineerCMS onlyoffice.png" style="width: 100%">
    <p class="lead">文档用户权限设置</p>
    <img src="/static/img/oo doc user permission.png" style="width: 100%">
    <p class="lead">文档用户组（角色）权限设置</p>
    <img src="/static/img/oo doc role permission.png" style="width: 100%">
    <p class="lead">word协作效果</p>
    <img src="/static/img/onlyoffice word.jpg" style="width: 100%">
    <p class="lead">powerpoint协作效果</p>
    <img src="/static/img/onlyoffice powerpoint.jpg" style="width: 100%">
    <p class="lead">excel协作效果</p>
    <img src="/static/img/onlyoffice excel.jpg" style="width: 100%">

  </div>
</div>
</div>
  </body>
</html>