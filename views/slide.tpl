<!DOCTYPE html>
<!-- {{template "header"}} -->
<!-- <html manifest="cache.appcache" class=" transition"><head> -->
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=Edge;chrome=1">
    <title>EngineerCMS：工程师知识管理系统</title>
    <link href="/static/css/slidedefault.css" class="theme" rel="stylesheet" media="screen">
    <link href="/static/css/slidecommon.css" rel="stylesheet" media="screen">
    <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  </head>

  <body>
<!--   <div class="navbar navba-default navbar-fixed-top">
    <div class="container-fill">{{template "navbar" .}}</div>
  </div> -->
    <nav id="helpers">
      <button id="nav-first" type="checkbox" class="nav-prev" title="去首页">|&lt;</button>
      <button title="Previous slide" id="nav-prev" class="nav-prev">&lt;</button> 
      <button title="Jump to a random slide" id="slide-no">4</button> 
      <button title="Next slide" id="nav-next" class="nav-next">&gt;</button>
      <button id="nav-last" type="checkbox" class="nav-next" title="去尾页">&gt;|</button>
      <menu>
        <button type="checkbox" data-command="share" class="share" title="分享">分享</button>
        <button type="checkbox" data-command="help" title="View Help">?</button>
      </menu>
    </nav>
    <div class="presentation">
      <div id="presentation-counter">4</div>
      <div class="slides" style="display: block;">
        <div class="slide distant-slide" id="landing-slide">
          <style>
            #landing-slide p {
              font-size: 35px;
            }
          </style>
          <section class="middle">
            <p>这是一个使用HTML5制作的幻灯片</p>
            <p>使用 <span id="left-init-key" class="key">→</span> 键开始播放。</p>
          </section>
          <aside class="note">
            <section>
              Welcome! (This field is for presenter notes and commentary.)
            </section>
          </aside> 
        <span class="counter"></span></div>

        <!-- <div class="slide far-past" id="controls-slide">
          <header>控制键</header>
          <style>
            #controls-slide li, #controls-slide p {
              font-size: 32px;
            }
            #controls-slide .key {
              bottom: 2px;
            }
          </style>
          <section>
            <ul>
              <li>使用 <span class="key">←</span> 和 <span class="key">→</span> 进行前后翻页。</li>
              <li>使用 <span class="key">Ctrl/Command</span> 和 <span class="key">+</span> 或 <span class="key">-</span> 进行页面缩放。</li>
              <li>按 <span class="key">3</span> 切换3D效果。</li>
              <li>按 <span class="key">0</span> 显示帮助。</li>
            </ul>
          </section>
        <span class="counter"></span></div> -->

        <!-- <div class="slide past" id="hello">
          <section class="middle">
            <hgroup><img src="images/bumper640x360.png"></hgroup>
          </section>
        <span class="counter"></span></div> -->
        <div class="slide current" id="title-slide">
          <style>
            #title-slide h1, #title-slide h2 {
              color: black;
            }
            #title-slide h1 {
              font-size: 80px;
            }
            #title-slide h2 {
              font-size: 36px;
            }
          </style>
          <section class="middle">
            <hgroup>
              <h1>
                  工程师知识管理系统(EngineerCMS)<br>
                  <!-- 简介<br> -->
                  简介<br>
              </h1>
              <h2>
                <br>504284@qq.com<br><span class="trans"><a href="https://github.com/3xxx">3xxx</a>制作</span><br>
              </h2>
            </hgroup>
          </section>
        <span class="counter"></span></div>

	<!-- start of slide 0 fa-slack fa-yelp fa-houzz fa-ioxhost fa-stop fa-long-arrow-right fa-hand-o-right fa-arrow-right fa-arrow-circle-right  fa-chevron-circle-right fa-caret-square-o-right fa-gear fa-fighter-jet  fa-square fa-square-o fa-star  fa-star-o fa-map-marker  fa-dot-circle-o fa-circle  fa-circle-o fa-cubes fa-bell fa-asterisk-->
        <div class="slide future" id="slide-0">
          <header><i class="fa fa-square"> 视频总揽</i></header>
          <section>
            <p>
            <i class="fa fa-slack"> 第一节 工程师需求</i>
            <br>
            <br>
            <i class="fa fa-slack"> 第二节 新建项目</i>
            <br>
            <br>
            <i class="fa fa-slack"> 第三节 往项目里添加成果</i>
            <br>
            <br>
            <i class="fa fa-slack"> 第四节 添加大事记</i>
            <br>
            <br>
            <i class="fa fa-slack"> 第五节 同步资料</i>
            <br>
            <br>
            <i class="fa fa-slack"> 第六节 用户和权限</i>
            <br>
            <br>
            <i class="fa fa-slack"> 第七节 定制目录和部门</i>
            <br>
            <br>
            <i class="fa fa-slack"> 第八节 后记</i>
            </p>
            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 0 --><!-- start of slide 1 -->
        <div class="slide far-future" id="slide-1">
          <header><i class="fa fa-slack"> 第一节 工程师需求</i></header>

          <section>
            <p>
            <i class="fa fa-arrow-right"> 工程师个人资料标准化管理</i>
            <br>
            </p>
            <ul>
            <li><i class="fa fa-caret-right">个人增加时间成本——但降低总时间成本</i></li>
            <li><i class="fa fa-caret-right">团队协作方便——不用为别人找资料，发资料</i></li>
            <li><i class="fa fa-caret-right">资料交接快捷——工作调换方便</i></li>
            </ul>

            <p>
            <i class="fa fa-arrow-right"> 知识共享——知识继承</i>
            <br>
            </p>
            <ul>
            <li><i class="fa fa-caret-right">持久化展示自己的知识总结——拓宽团队视野，避免重复造轮子</i></li>
            
            <li><i class="fa fa-caret-right">展示工程现场图片——展示其他业绩</i></li>
            <li><i class="fa fa-caret-right">参建各方远程查询工程资料——设计资料易得是保证工程质量的必备条件</i></li>
            </ul>

            <p>
            <i class="fa fa-arrow-right"> 网络化的必然趋势——基于个人电脑的微服务</i>
            <br>
            </p>
            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->

        <div class="slide far-future" id="slide-2">
          <header><i class="fa fa-caret-right"> 局域网</i></header>
          <section>
            <p></p>
              <img src="/static/img/局域网meritecms.png" style="width:100%;">
            <p>
            <br>
            <br>
          </p></section>
          <span class="counter"></span>
        </div>
        <div class="slide far-future" id="slide-3">
          <header><i class="fa fa-caret-right"> 组织架构</i></header>
          <section>
            <p></p>
              <img src="/static/img/组织架构meritecms.png" style="width:100%;">
            <p>
            <br>
            <br>
          </p></section>
          <span class="counter"></span>
        </div>

      <!-- end of slide 2 --><!-- start of slide 3 -->
        <div class="slide distant-slide" id="slide-4">
          <header><i class="fa fa-slack"> 第二节 新建项目</i></header>
          <section>
            <p>
            <i class="fa fa-arrow-right"> 添加</i>
            <br>
            </p>
            <ul>
            <li><i class="fa fa-caret-right">目录——树状目录在后台预设</i></li>
            <li><i class="fa fa-caret-right">标签——项目类型和特点</i></li>
            </ul>
            <br>

            <p>
            <i class="fa fa-arrow-right"> 编辑</i>
            <br>
            </p>
            <ul>
            <li><i class="fa fa-caret-right">名称</i></li>
            <li><i class="fa fa-caret-right">目录——在后台编辑树状目录</i></li>
            </ul>
            <br>

            <p>
            <i class="fa fa-arrow-right"> 删除</i>
            <br>
            </p>
            <ul>
            <li><i class="fa fa-caret-right">成果——所带附件一并删除</i></li>
            <li><i class="fa fa-caret-right">附件——文件夹和数据库</i></li>
            <li><i class="fa fa-caret-right">文章</i></li>
            </ul>
            <br>

            <p>
          </p></section>
          <span class="counter"></span>
        </div> <!-- end slide template -->
          <!-- end of slide 3 --><!-- start of slide 4 -->
        <div class="slide distant-slide" id="slide-5">
          <header><i class="fa fa-slack"> 第三节 往项目里添加成果</i></header>
          <section>
            <p>
            <i class="fa fa-arrow-right"> 添加</i>
            <br>
            </p>
            <ul>
            <li><i class="fa fa-caret-right">批量上传——电子文件命名规则</i></li>
            <li><i class="fa fa-caret-right">多附件模式</i></li>
            <li><i class="fa fa-caret-right">文章模式——发布图文并茂的文章</i></li>
            </ul>
            <br>

            <p>
            <i class="fa fa-arrow-right"> 编辑</i>
            <br>
            </p>
            <ul>
            <li><i class="fa fa-caret-right">编辑成果信息</i></li>
            <li><i class="fa fa-caret-right">编辑成果附件</i></li>
            </ul>
            <br>

            <p>
            <i class="fa fa-arrow-right"> 删除</i>
            <br>
            </p>
            <br>

            <p>
            <i class="fa fa-arrow-right"> 同步——后台设置本项目同步的ip列表</i>
            <br>
            </p>
            <br>

            <p>
          </p></section>
          <span class="counter"></span>
        </div> <!-- end slide template -->
        <!-- end of slide 4 --><!-- start of slide 5 -->
        <!-- end slide template -->
      <!-- end of slide 56 -->

        <div class="slide distant-slide" id="goodbye">
          <section class="middle">
            <hgroup><img src="images/bumper640x360.png"></hgroup>
          </section>
        <span class="counter"></span></div>

      </div>

      <div id="speaker-note" class="invisible" style="display: none;"></div> <!-- speaker note -->
      <aside id="help" class="sidebar invisible" style="display: hidden;">
        <table>
          <caption>帮助</caption>
          <tbody>
            <tr>
              <th>翻页</th>
              <td>←&nbsp;→</td>
            </tr>
            <tr>
              <th>源文件</th>
              <td>s</td>
            </tr>
            <tr>
              <th>3D切换</th>
              <td>3</td>
            </tr>
            <tr>
              <th>帮助</th>
              <td>0</td>
            </tr>
          </tbody>
        </table>
      </aside>

    </div> <!-- presentation -->

    <!--[if lt IE 9]>
    <script 
      src="http://ajax.googleapis.com/ajax/libs/chrome-frame/1/CFInstall.min.js">
    </script>
    <script>CFInstall.check({ mode: "overlay" });</script>
    <![endif]-->

    <script src="/static/js/slideutils.js"></script>
  

</body></html>