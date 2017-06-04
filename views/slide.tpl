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
        <span class="counter"></span></div>
        <div class="slide far-future" id="slide-3">
          <header><i class="fa fa-caret-right"> 组织架构</i></header>
          <section>
            <p></p>
              <img src="/static/img/组织架构meritecms.png" style="width:100%;">
            <p>
            <br>
            <br>
          </p></section>
        <span class="counter"></span></div>

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
        <span class="counter"></span></div> <!-- end slide template -->
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
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 4 --><!-- start of slide 5 -->
        <div class="slide distant-slide" id="slide-6">
          <header>并发 vs. 并行</header>
          <section>
            <p>
            并发是指同时处理很多事情。<br>
            
            <br>
            
            而并行是指同时能完成很多事情。<br>
            
            <br>
            
            两者不同，但相关。<br>
            
            <br>
            
            一个重点是组合，一个重点是执行。<br>
            
            <br>
            
            并发提供了一种方式让我们能够设计一种方案将问题(非必须的)并行的解决。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 5 --><!-- start of slide 6 -->
        <div class="slide distant-slide" id="slide-7">
          <header>一个比方</header>

          <section>
            <p>
            并发：鼠标，键盘，显示器，硬盘——同时工作。<br>
            
            <br>
            
            并行：向量数量积<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 6 --><!-- start of slide 7 -->
        <div class="slide distant-slide" id="slide-7">
          <header>并发+通信</header>

          <section>
            <p>
            并发是一种将一个程序分解成小片段独立执行的程序设计方法。<br>
            
            <br>
            
            通信是指各个独立的执行任务间的合作。<br>
            
            <br>
            
            这是Go语言采用的模式，包括Erlang等其它语言都是基于这种SCP模式：<br>
            
            <br>
            
            C. A. R. Hoare: Communicating Sequential Processes (CACM 1978)<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 7 --><!-- start of slide 8 -->
        <div class="slide distant-slide" id="slide-8">
          <header>繁忙的地鼠</header>

          <section>
            <p>
            概念太抽象。我们来点具体的。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 8 --><!-- start of slide 9 -->
        <div class="slide distant-slide" id="slide-9">
          <header>我们的问题</header>

          <section>
            <p>
            运一堆没用的手册到焚烧炉里。<br>
            </p>
            



<img src="images/gophersimple1.jpg">

            <p>
            <br>
            
            如果只有一只地鼠，这需要很长时间。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 9 --><!-- start of slide 10 -->
        <div class="slide distant-slide" id="slide-10">
          <header>更多的地鼠！</header>

          <section>
            <p></p>
            



<img src="images/gophersimple3.jpg">

            <p>
            <br>
            
            更多的地鼠还不行；他们需要更多的小推车。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 10 --><!-- start of slide 11 -->
        <div class="slide distant-slide" id="slide-11">
          <header>更多的地鼠和更多的小推车</header>

          <section>
            <p></p>
            



<img src="images/gophersimple2.jpg">

            <p>
            <br>
            
            这样快多了，但在装运处和焚烧炉处出现了瓶颈。<br>
            
            还有，这些地鼠需要能同时工作。<br>
            
            它们需要相互通知。(这就是地鼠之间的通信)<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 11 --><!-- start of slide 12 -->
        <div class="slide distant-slide" id="slide-12">
          <header>所有东西都增加一倍</header>

          <section>
            <p>
            消除瓶颈；让他们能真正的相互独立不干扰。<br>
            </p>
            



<img src="images/gophersimple4.jpg">

            <p>
            这样吞吐速度会快一倍。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 12 --><!-- start of slide 13 -->
        <div class="slide distant-slide" id="slide-13">
          <header>并发组合</header>

          <section>
            <p></p>
            



<img src="images/gophersimple4.jpg">

            <p>
            并发组合两个地鼠的工作过程。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 13 --><!-- start of slide 14 -->
        <div class="slide distant-slide" id="slide-14">
          <header>并发组合</header>

          <section>
            <p>
            现在的这种工作流程不能自动的实现并行！<br>
            
            <br>
            
            如果只有一只地鼠<br>
            
            这仍然是并发(就是目前的这种工作方式)，但它不是并行。<br>
            
            <br>
            
            然而，它是可以并行的！<br>
            
            <br>
            
            需要设计出另外的工作流程来实现并发组合。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 14 --><!-- start of slide 15 -->
        <div class="slide distant-slide" id="slide-15">
          <header>新的工作流程</header>

          <section>
            <p></p>
            



<img src="images/gophercomplex0.jpg">

            <p>
            <br>
            
            三只地鼠在工作，但看起来工作有些滞后。<br>
            
            每只地鼠都在做一种独立的工序，<br>
            
            并且相互合作(通信)。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 15 --><!-- start of slide 16 -->
        <div class="slide distant-slide" id="slide-16">
          <header>更细分工的并发</header>

          <section>
            <p>
            增加一只地鼠，专门运回空的小推车。<br>
            </p>
            



<img src="images/gophercomplex1.jpg">

            <p>
            四只地鼠组成了一个优化的工作流程，每只只做自己一种简单的工序。<br>
            
            <br>
            
            如果任务布置的合理，这将会比最初一个地鼠的工作快4倍。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 16 --><!-- start of slide 17 -->
        <div class="slide distant-slide" id="slide-17">
          <header>结果</header>

          <section>
            <p>
            我们通过在现有的工作流程里加入并发过程从而改进了执行效率。<br>
            
            <br>
            
            地鼠越多能做的越多；工作效率越高。<br>
            
            <br>
            
            这是一种比仅仅并行更深刻的认识。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 17 --><!-- start of slide 18 -->
        <div class="slide distant-slide" id="slide-18">
          <header>并发过程</header>

          <section>
            <p>
            四个地鼠有不同的工作环节：<br>
            </p>
            
            <ul>
            
            <li>往小推车里装书</li>
            
            <li>移动小推车到焚烧炉</li>
            
            <li>卸载书到焚烧炉里</li>
            
            <li>送回空的小推车</li>
            
            </ul>

            <p>
            不同的并发设计能导致不同的并行方式。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 18 --><!-- start of slide 19 -->
        <div class="slide distant-slide" id="slide-19">
          <header>更多的并行！</header>

          <section>
            <p>
            现在我们可以让并行再多一倍；按照现在的并行模式很容易实现这些。八个地鼠，全部繁忙。<br>
            </p>
            



<img src="images/gophercomplex2.jpg">

            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 19 --><!-- start of slide 20 -->
        <div class="slide distant-slide" id="slide-20">
          <header>或者它们可以完全不并行</header>

          <section>
            <p>
            请记住，只有一个地鼠在工作(零并行)，这仍然是一个正确的并发的工作方案。<br>
            </p>
            



<img src="images/gophercomplex2.jpg">

            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 20 --><!-- start of slide 21 -->
        <div class="slide distant-slide" id="slide-21">
          <header>换一种设计</header>

          <section>
            <p>
            现在我们换一种设计来组织我们的地鼠的并发工作流程。<br>
            
            <br>
            
            两个地鼠，一个中转站。<br>
            </p>
            



<img src="images/gophercomplex3.jpg">

            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 21 --><!-- start of slide 22 -->
        <div class="slide distant-slide" id="slide-22">
          <header>让常规的流程并行化</header>

          <section>
            <p>
            更多的并发流程能获得更多的吞吐量。<br>
            </p>
            



<img src="images/gophercomplex4.jpg">

            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 22 --><!-- start of slide 23 -->
        <div class="slide distant-slide" id="slide-23">
          <header>或者另外一种方法</header>

          <section>
            <p>
            在每个中转站之间都引入多个地鼠并发的模式：<br>
            
            <br>
            </p>
            



<img src="images/gophercomplex5.jpg">

            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 23 --><!-- start of slide 24 -->
        <div class="slide distant-slide" id="slide-24">
          <header>全程优化</header>

          <section>
            <p>
            使用这种技术策略，16个地鼠都很繁忙！<br>
            
            <br>
            </p>
            



<img src="images/gophercomplex6.jpg">

            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 24 --><!-- start of slide 25 -->
        <div class="slide distant-slide" id="slide-25">
          <header>习得</header>

          <section>
            <p>
            有很多分解流程的方式。<br>
            
            <br>
            
            这都是并发设计。<br>
            
            <br>
            
            一旦完成了分解，并行可能会丧失，但很容易纠正。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 25 --><!-- start of slide 26 -->
        <div class="slide distant-slide" id="slide-26">
          <header>回到计算机世界</header>

          <section>
            <p>
            将我们的运书工作替换成如下：<br>
            </p>
            
            <ul>
            
            <li>书堆 =&gt; Web内容</li>
            
            <li>地鼠 =&gt; CPU</li>
            
            <li>小推车 =&gt; 调度，渲染或网络传输</li>
            
            <li>焚烧炉 =&gt; 代理，浏览器或其他消费源</li>
            
            </ul>
            



            <p>
            我们现在的这种设计就是一种可扩展的Web服务的并发设计。<br>
            
            地鼠提供Web内容服务。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 26 --><!-- start of slide 27 -->
        <div class="slide distant-slide" id="slide-27">
          <header>关于Go语言的一点背景知识</header>

          <section>
            <p>
            这里不是一个详细的教材，只是快速做一些重点介绍。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 27 --><!-- start of slide 28 -->
        <div class="slide distant-slide" id="slide-28">
          <header>Go例程(Goroutines)</header>

          <section>
            <p>
            一个Go例程就是一个和其它Go例程在同一地址空间里但却独立运行的函数。<br>
            </p>
            




            <p>
            <br>
            
            就像是在shell里使用 &amp; 标记启动一个命令。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 28 --><!-- start of slide 29 -->
        <div class="slide distant-slide" id="slide-29">
          <header>Go例程不是线程</header>

          <section>
            <p>
            (很像线程，但比线程更轻量。)<br>
            
            <br>
            
            多个例程可以在系统线程上做多路通信。<br>
            
            <br>
            
            当一个Go例程阻塞时，所在的线程会阻塞，但其它Go例程不受影响。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 29 --><!-- start of slide 30 -->
        <div class="slide distant-slide" id="slide-30">
          <header>通道 Channels</header>

          <section>
            <p>
            通道是类型化的值，能够被Go例程用来做同步或交互信息。<br>
            
            <br>
            </p>
            




            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 30 --><!-- start of slide 31 -->
        <div class="slide distant-slide" id="slide-31">
          <header>Select</header>

          <section>
            <p>
            这select语句很像switch，但它的判断条件是基于通信，而不是基于值的等量匹配。<br>
            
            <br>
            </p>
            




            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 31 --><!-- start of slide 32 -->
        <div class="slide distant-slide" id="slide-32">
          <header>Go语言非常的支持并发</header>

          <section>
            <p>
            非常。<br>
            
            <br>
            
            一个程序里产生成千上万个Go例程很正常。<br>
            
            (有一次调试一个程序发现有130万个例程。)<br>
            
            <br>
            
            堆栈初始很小，但随着需求会增长或收缩。<br>
            
            <br>
            
            Go例程不是不耗资源，但它们很轻量级的。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 32 --><!-- start of slide 33 -->
        <div class="slide distant-slide" id="slide-33">
          <header>闭包在这里也是重要角色</header>

          <section>
            <p>
            它让一些并发运算更容易表达。<br>
            
            <br>
            
            它们是局部函数。<br>
            
            下面是一个非并发例子。<br>
            </p>
            




            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 33 --><!-- start of slide 34 -->
        <div class="slide distant-slide" id="slide-34">
          <header>一些例子</header>

          <section>
            <p>
            通过实例学习Go语言并发<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 34 --><!-- start of slide 35 -->
        <div class="slide distant-slide" id="slide-35">
          <header>启动后台程序</header>

          <section>
            <p>
            使用闭包封装一个后台操作。<br>
            
            <br>
            
            下面是从输入通道拷贝数据到输出通道。<br>
            </p>
            





            <p>
            这个for range操作会一直执行到处理掉通道内最后一个值。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 35 --><!-- start of slide 36 -->
        <div class="slide distant-slide" id="slide-36">
          <header>一个简单的负载均衡的例子(1)</header>

          <section>
            <p>
            数据类型：<br>
            
            <br>
            </p>
            




            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 36 --><!-- start of slide 37 -->
        <div class="slide distant-slide" id="slide-37">
          <header>一个简单的负载均衡的例子(2)</header>

          <section>
            <p>
            一个worker的任务：<br>
            </p>
            




            <p>
            必须保证当一个worker阻塞时其他worker仍能运行。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 37 --><!-- start of slide 38 -->
        <div class="slide distant-slide" id="slide-38">
          <header>一个简单的负载均衡的例子(3)</header>

          <section>
            <p>
            runner:<br>
            </p>
            





            <p>
            <br>
            
            很简单的任务，但如果没有并发机制，你仍然很难这么简单的解决。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 38 --><!-- start of slide 39 -->
        <div class="slide distant-slide" id="slide-39">
          <header>并发是并行成为可能</header>

          <section>
            <p>
            这个负载均衡的例子具有很明显的并行和可扩展性。<br>
            
            <br>
            
            Worker数可以非常巨大。<br>
            
            <br>
            
            Go语言的这种并发特征能的开发一个安全的、好用的、可扩展的、并行的软件变得很容易。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 39 --><!-- start of slide 40 -->
        <div class="slide distant-slide" id="slide-40">
          <header>并发简化了同步</header>

          <section>
            <p>
            没有明显的需要同步的操作。<br>
            
            <br>
            
            程序的这种设计隐含的实现了同步。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 40 --><!-- start of slide 41 -->
        <div class="slide distant-slide" id="slide-41">
          <header>真是太简单了</header>

          <section>
            <p>
            让我们实现一个更有意义的负载均衡的例子。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 41 --><!-- start of slide 42 -->
        <div class="slide distant-slide" id="slide-42">
          <header>负载均衡</header>

          <section>
            <p></p>
            



<img src="images/gopherchart.jpg">

            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 42 --><!-- start of slide 43 -->
        <div class="slide distant-slide" id="slide-43">
          <header>定义请求</header>

          <section>
            <p>
            请求者向均衡服务发送请求。<br>
            </p>
            




            <p>
            注意这返回的通道是放在请求内部的。<br>
            
            通道是first-class值<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 43 --><!-- start of slide 44 -->
        <div class="slide distant-slide" id="slide-44">
          <header>请求函数</header>

          <section>
            <p>
            没有实际用处，但能很好的模拟一个请求者，一个负载产生者。<br>
            
            <br>
            </p>
            




            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 44 --><!-- start of slide 45 -->
        <div class="slide distant-slide" id="slide-45">
          <header>Worker定义</header>

          <section>
            <p>
            一些请求通道，加上一些负载记录数据。<br>
            
            <br>
            </p>
            





            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 45 --><!-- start of slide 46 -->
        <div class="slide distant-slide" id="slide-46">
          <header>Worker</header>

          <section>
            <p>
            均衡服务将请求发送给压力最小的worker。<br>
            </p>
            




            <p>
            请求通道(w.requests)将请求提交给各个worker。均衡服务跟踪请求待处理的数量来判断负载情况。<br>
            
            每个响应直接反馈给它的请求者。<br>
            
            <br>
            
            你可以将循环体内的代码当成Go例程从而实现并行。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 46 --><!-- start of slide 47 -->
        <div class="slide distant-slide" id="slide-47">
          <header>定义负载均衡器</header>

          <section>
            <p>
            负载均衡器需要一个装很多worker的池子和一个通道来让请求者报告任务完成情况。<br>
            
            <br>
            </p>
            




            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 47 --><!-- start of slide 48 -->
        <div class="slide distant-slide" id="slide-48">
          <header>负载均衡函数</header>

          <section>
            <p>
            简单！<br>
            
            <br>
            </p>
            




            <p>
            <br>
            
            你只需要实现dispatch和completed方法。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 48 --><!-- start of slide 49 -->
        <div class="slide distant-slide" id="slide-49">
          <header>储存通道的堆(heap)</header>

          <section>
            <p>
            将负载均衡的池子用一个Heap接口实现，外加一些方法：<br>
            
            <br>
            </p>
            




            <p>
            <br>
            
            现在我们的负载均衡使用堆来跟踪负载情况。<br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 49 --><!-- start of slide 50 -->
        <div class="slide distant-slide" id="slide-50">
          <header>Dispatch</header>

          <section>
            <p>
            需要的东西都有了。<br>
            
            <br>
            </p>
            




            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 50 --><!-- start of slide 51 -->
        <div class="slide distant-slide" id="slide-51">
          <header>Completed</header>

          <section>
            <p></p>
            



            <p>
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 51 --><!-- start of slide 52 -->
        <div class="slide distant-slide" id="slide-52">
          <header>习得</header>

          <section>
            <p>
            一个复杂的问题可以被拆分成容易理解的组件。<br>
            
            <br>
            
            它们可以被并发的处理。<br>
            
            <br>
            
            结果就是容易理解，高效，可扩展，好用。<br>
            
            <br>
            
            或许更加并行。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 52 --><!-- start of slide 53 -->
        <div class="slide distant-slide" id="slide-53">
          <header>最后一个例子</header>

          <section>
            <p>
            我们有几个相同的数据库，我们想最小化延迟，分别询问他们，挑选第一个响应的。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 53 --><!-- start of slide 54 -->
        <div class="slide distant-slide" id="slide-54">
          <header>查询数据库</header>

          <section>
            <p></p>
            





            <p>
            并发和垃圾回收机制让这成为一个很小很容易解决的问题。<br>
            
            <br>
            
            (作业练习：处理晚来的响应。)<br>
            
            <br>
            
            <br>
            
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 54 --><!-- start of slide 55 -->
        <div class="slide distant-slide" id="slide-55">
          <header>结论</header>

          <section>
            <p>
            并发很强大。<br>
            
            <br>
            
            并发不是并行。<br>
            
            <br>
            
            并发帮助实现并行。<br>
            
            <br>
            
            并发使并行(扩展等)变得容易。<br>
            
            <br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
      <!-- end of slide 55 --><!-- start of slide 56 -->
        <div class="slide distant-slide" id="slide-56">
          <header>更多信息</header>

          <section>
            <p>
            Go: golang.org<br>
            
            <br>
            
            一些历史: swtch.com/~rsc/thread/<br>
            
            <br>
            
            另一个视频: tinyurl.com/newsqueak<br>
            
            <br>
            
            并行不是并发(Harper): tinyurl.com/pincharper<br>
            
            <br>
            
            一个并发window系统(Pike): tinyurl.com/pikecws<br>
            
            <br>
            
            并发系列(McIlroy): tinyurl.com/powser<br>
            
            <br>
            
            最后，并行但不是并发：<br>
            
            research.google.com/archive/sawzall.html<br>
            </p>
            



            <p>
          </p></section>
        <span class="counter"></span></div> <!-- end slide template -->
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