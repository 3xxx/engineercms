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
              <a href="#backroud-prodcate">成果类型</a>
            </li>
            <li class="">
              <a href="#backroud-merit">定义价值</a>
            </li>
            <li>
              <a href="#backroud-secofficemerit">科室价值</a>
            </li>
          </ul>
        </li>
        <li class="">
          <a href="#achievement">校审任务流程</a>
          <ul class="nav">
            <li class="">
              <a href="#achievement-add">添加成果的5种途径</a>
            </li>
            <li class="">
              <a href="#achievement-multi">批量校审</a>
            </li>
            <li>
              <a href="#achievement-checkexamin">校审意见</a>
            </li>
            <li>
              <a href="#achievement-attachment">任务携带的附件</a>
            </li>
          </ul>
        </li>
        <li class="">
          <a href="#bigdata">大数据分析</a>
          <ul class="nav">
            <li class="">
              <a href="#bigdata-person">个人</a>
            </li>
            <li class="">
              <a href="#bigdata-secoffice">科室</a>
            </li>
            <li>
              <a href="#bigdata-department">分院</a>
            </li>
          </ul>
        </li>
        <li class="">
          <a href="#merit">价值管理</a>
          <ul class="nav">
            <li class="">
              <a href="#merit-add">添加价值</a>
            </li>
            <li class="">
              <a href="#merit-check">价值审核</a>
            </li>
            <li>
              <a href="#merit-addup">价值统计</a>
            </li>
          </ul>
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
    <p class="lead">Merit是一款在线成果（工作量）登记、价值档案管理系统，理念是尽量减少占用技术人员时间去进行工作量登记，但又能符合管理者对大数据的需求；技术人员基于Merit在线动态维护自己的价值档案，充分展示自己的价值。</p>
    <p class="lead">
      MeritMS 具有功能：EngineerCMS的所有功能，可作为团队档案管理系统，进行设计校审流程，成果统计，价值评测等。</p>
      <p class="lead">创新：集EngineerCMS工程师知识管理系统，解决团队成果统计和公示难题；采用价值评测理念，提高团队竞争力；基于个人电脑进行团队管理的微服务。
    </p>

    <div class="bs-callout bs-callout-warning" id="jquery-required">
      <h4>1.工程师为什么要绩效管理</h4>
      <p>
        绩效管理不是目的，设计工作应以人为本，适当兼顾绩效；切记不能因为绩效差而颠覆工程师的一切；应结合价值管理来综合评价一个人。
      </p>
      <h4>2.设计人员绩效管理的难点是什么</h4>
      <p>
        成果类型多，没有统一的衡量标准；因此，需要管理人员经过多年的工作经验进行标准设计，进而在MeritMS中进行设置折算系数。
      </p>
    </div>
    <img src="/static/img/成果登记系统.png" style="width: 100%">
    <h2 id="about-function">
      <a class="anchorjs-link " href="#about-function" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      功能
    </h2>

<ol>
  <li>按照 编制（制图）、设计、校核、审查等角色和权限进行成果流程管理。</li>
<li>后台定制组织架构：部门——科室——人员 或 部门——人员。</li>
<li>整合了EngineerCMS，可以集中管理项目资料。</li>
<li>技术人员无论是向MeritMS中上传成果或在个人EngineerCMS中上传 成果，都可以自动生成 成果清单，提交给MeritMS进行工作量统计。</li>
<li>基于IP权限和登录权限管理，IP权限快捷，无需登录；登录权限用于远程VPN访问。分系统管理员、部门管理者、科室管理者和技术人员4个权限等级。</li>
<li>基于科室、部门不同组织级别的成果、项目、贡献、排名等统计展示。</li>
<li>分技术人员和系统管理员2种批量添加 成果列表，后者无需经过层层校审流程，直接进入统计。</li>
<li>成果校审流程中提供附件链接和校审意见填写。</li>
<li>后台定制组织结构里的价值分类和价值列表。</li>
<li>技术人员在系统中动态维护自己的价值：注册、获奖、负责人、科研……。</li>
<li>基于科室、部门不同组织级别的技术人员价值排名。</li>
<li>为尽最大努力简化技术人员添加成果的工作量，提供了5种添加方式：1、手工添加；2、自己导入excel表格中成果清单；3、交给管理员导入excel中的成果清单，省去校审流程；4、想MeritMS中上传成果，自动生成成果清单，可批量进行修改，批量进行提交；5、在EngineerCMS中上传成果，自动生成成果清单，可批量进行修改，批量进行提交。</li> 
  </ol>

    <h2 id="about-use">
      <a class="anchorjs-link " href="#about-use" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      使用方式
    </h2>
<ol>
    <li>如果是首次使用，请：</li>
  <li> 解压到d:\merit\；（其他盘根目录下也行，因为上传附件会使得这个文件夹越来越大，所以，要考虑空间大一些的盘。）</li>
  <li> 修改配置文件conf\app.conf.sample为app.conf,打开app.conf,看到里面的httpport = 80，如果要修改成8080，请修改后保存。runmode = prod表示生产模式运行系统。</li>
  <li> 修改数据库文件database\merit.db.sample为merit.db。</li>
  <li> merit-win64/win32.exe即可在chrome浏览器中输入本地ip（127.0.0.1）和前面设置的端口号（假设是80或8080）进行访问。如果是80端口，则端口号可省略，如http://127.0.0.1。如果运行后闪退，则可能是端口号被占用了，请修改端口号再运行。运行后不要关闭窗口，它是服务。IE浏览器支持不好，推荐使用chrome，可以使用firefox、opra。加入开机启动请自行设置。</li>
  </ol>
  <img src="/static/img/组织架构meritecms.png" style="width: 100%">

    <h1 id="backroud" class="page-header">
      <a class="anchorjs-link " href="#backroud" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      后台设置
    </h1>

    <p class="lead">
      在这里可以预设和定制系统。
    </p>

    <h2 id="backroud-prodcate">
      <a class="anchorjs-link " href="#backroud-prodcate" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      成果类型
    </h2>
    <ol>
      <li>登记成果类型</li>
      <img src="/static/img/成果类型.png" style="width: 100%">
    </ol>
    <h2 id="backroud-merit">
      <a class="anchorjs-link " href="#backroud-merit" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      定义价值
    </h2>
    <ol>
    <li>需要考核的价值分类和价值分值。</li>
    <img src="/static/img/定义价值.png" style="width: 100%">
    </ol>

    <h2 id="backroud-secofficemerit">
      <a class="anchorjs-link " href="#backroud-secofficemerit" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      科室价值
    </h2>
    <ol>
      <p>给不同的科室设置上面定义的价值分类</p>
      <img src="/static/img/科室价值.png" style="width: 100%">
    </ol>
    <h1 id="achievement" class="page-header">
      <a class="anchorjs-link " href="#achievement" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      校审任务流程
    </h1>
    <p class="lead">
      采用web方式走完一个校审流程。
    </p>

    <h2 id="achievement-add">
      <a class="anchorjs-link " href="#achievement-add" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      添加成果的5种途径
    </h2>
    <ol>
      <p>为了尽最大努力将技术人员的成果登记时间缩短到最少，目前有5种方式。一种比一种快捷。</p>
      <li>手工添加；</li>
      <li>自己导入excel表格中成果清单；</li>
      <li>交给管理员导入excel中的成果清单，省去校审流程；</li>
      <li>向MeritMS中上传成果，自动生成成果清单，可批量进行修改，批量进行提交；</li>
      <li>在EngineerCMS中上传成果，自动生成成果清单，可批量进行修改，批量进行提交。</li>
      <img src="/static/img/5种登记成果方式.png" style="width: 100%">
      <p>Ⅰ 成果在线登记的重要原则：</p>
      <li>自己可以从 制图、设计和校核任意位置发起，不能直接发起审查；</li>
      <li>发起后面必须至少有一位存在，进行下一级处理，不能发起一个没有下级的流程；</li>
      <li>下一级可以 跳级，比如我发起一个 编制，下一级就是 校核 了，没有 设计；</li>
      <li>重复数据判断原则：成果编号、成果名称和成果类型3者都相同则为重复数据；</li>
      <li>自己对应的系数要填上，否则统计结果为0。</li>
      <li>系统不提供计量单位选择，数量必须按系统默认标准单位折算后填写。</li>

      <li>对于多人出差同一项目，成果编号里填日期和自己的名字，比如：20160926秦晓川，成果名称：出差珠三角水资源配置工程。填写制图系数为1，下一级设计可以给项目负责人或主任进行确认即可，难度系统没填，系统默认为1。</li>
      <li>我直接发起一个校核，没有校核系数咋办？由下一级（审查）进行填写。</li>
      <li>我直接发起一个综合说明的编制，下一级就是审查。</li>
      <li>制图和设计是同一人，制图和设计不是同一人。</li>

      <p>Ⅱ 处理提交的原则</p>
      <li>下一级可以修改前一级系数和难度系数。</li>
      <li>必须填写自己的系数。</li>
      <li>最后一级确认难度系数。</li>
      <p>Ⅲ 关于人名选择提示</p>
      <li>添加里的人名选择，采用输入用户拼音形式。当输入2个字母时，请停顿一下，系统会查出含这2个字母的人名。可继续输入，系统进行筛选，也可以用上下箭头或鼠标进行选择。</li>
 
      <li>在修改表格里，人名选择是采用汉字形式，当输入汉字后，系统开始筛选，也可以用鼠标或键盘选择。</li>

    </ol>
    <h2 id="achievement-multi">
      <a class="anchorjs-link " href="#achievement-multi" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      批量校审
    </h2>
    <ol>
      <p>批量修改成果内容、校审人员、批量提交等</p>
      <img src="/static/img/批量处理校审.png" style="width: 100%">
    </ol>
    <h2 id="achievement-checkexamin">
      <a class="anchorjs-link " href="#achievement-checkexamin" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      校审意见
    </h2>
    <ol>
      <p>任务发起人可添加说明，校审人员可添加校审意见</p>
      <img src="/static/img/校审意见.png" style="width: 100%">
    </ol>
    <h2 id="achievement-attachment">
      <a class="anchorjs-link " href="#achievement-attachment" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      任务携带的附件
    </h2>
    <ol>
      <p>任务发起人可添加修改附件，校审人员可查阅附件</p>
      <img src="/static/img/任务携带的附件.png" style="width: 100%">
    </ol>
    <h1 id="bigdata" class="page-header">
      <a class="anchorjs-link " href="#bigdata" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      大数据分析
    </h1>
    <p class="lead">
      大数据分析师最终目的，对绩效进行排序，对项目的人力成本进行评估。
    </p>

    <h2 id="bigdata-person">
      <a class="anchorjs-link " href="#bigdata-person" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      个人
    </h2>
    <ol>
      <li>用户查看年度内情况：一年以来每个月的成果分值和排名；一个月以来成果类型情况；一年以来成果类型情况。</li>
      <li>用户查看自己参与的项目列表和自己的贡献率</li>
      <li>查看某个项目里本专业的全部成果列表</li>
      <li>查看某个项目的每个人贡献率</li>
      <li>查看某个项目的成果类型分布</li>
      <img src="/static/img/个人绩效.png" style="width: 100%">
    </ol>
    <h2 id="bigdata-secoffice">
      <a class="anchorjs-link " href="#bigdata-secoffice" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      科室
    </h2>
    <ol>
      <li>如果是主任，可以查看本科室情况。</li>
      <li>点击可查看成员详细。</li>
      <li>点击表头可以进行排序。</li>
      <li>主任查看本科室年度所有项目。</li>
      <li>科室主任查看本月所有成果。</li>
      <img src="/static/img/科室绩效.png" style="width: 100%">
    </ol>
    <h2 id="bigdata-department">
      <a class="anchorjs-link " href="#bigdata-department" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      部门
    </h2>
    <ol>
      <li>如果是分院领导，可以看分院所有科室。</li>
      <img src="/static/img/部门绩效.png" style="width: 100%">
    </ol>
    <h1 id="merit" class="page-header">
      <a class="anchorjs-link " href="#merit" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      价值管理
    </h1>
    <p class="lead">
      基于用户自我维护价值的理念，管理者确定价值目标，对价值进行排序。
    </p>

    <h2 id="merit-add">
      <a class="anchorjs-link " href="#merit-add" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      添加价值
    </h2>
    <ol>
      <li>用户查看年度内情况：一年以来每个月的成果分值和排名；一个月以来成果类型情况；一年以来成果类型情况。</li>
      <img src="/static/img/添加价值.png" style="width: 100%">
    </ol>
    <h2 id="merit-check">
      <a class="anchorjs-link " href="#merit-check" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      价值审核
    </h2>
    <ol>
      <li>对可是人员添加的价值进行审核。</li>
      <img src="/static/img/价值审核.png" style="width: 100%">
    </ol>
    <h2 id="merit-addup">
      <a class="anchorjs-link " href="#merit-addup" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      价值统计
    </h2>
    <ol>
      <li>统计科室人员的价值高低</li>
      <img src="/static/img/价值统计.png" style="width: 100%">
    </ol>


  </div>
</div>
</div>
  </body>
</html>