<!DOCTYPE html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">

  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/docs.min.css"/>

<script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
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
          <a href="#onekey">一键成图及裁图、打印</a>
          <ul class="nav">
            <li class="">
              <a href="#onekey-ready">供水管线布置准备工作——与程序关系不大</a>
            </li>
            <li class="">
              <a href="#onekey-run">运行程序</a>
            </li>
            <li>
              <a href="#onekey-basedata">基本设计参数输入</a>
            </li>
            <li>
              <a href="#onekey-init">excel表格初始化</a>
            </li>
            <li>
              <a href="#onekey-insertstandard">插入标准断面</a>
            </li>
            <li>
              <a href="#onekey-standard">生成标准布置图及工程量表格</a>
            </li>
            <li>
              <a href="#onekey-readycut">为布局裁图做准备</a>
            </li>
            <li>
              <a href="#onekey-cut">布局裁图</a>
            </li>
            <li>
              <a href="#onekey-readyprint">为打印所有布局做准备</a>
            </li>
            <li>
              <a href="#onekey-print">打印布局到文件</a>
            </li>
            <li>
              <a href="#onekey-system">生成系统简图</a>
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
    <img src="/static/img/HydroWSindex.png" style="width: 100%">
    <p class="lead">
      HydroWS 主要用于长距离供水工程的管线设计工作。该软件采用vb语言编写，通过连接autocad和excel这2个设计中最常用的软件，将供水工程的设计过程整合在一个小小面板上。软件特点是针对长距离管道输水工程，集管线流程设计中的计算和制图于一体，既可以一键成图，又可以分步执行，目的就是在一些通常条件下，可以在几小时内完成设计生成工程量和图纸。因为快速，软件对于管线布置（纵剖面布置）的多方案研究有很大帮助。工具独立运行，具有一键成图：自动完成工程量统计，自动计算弯头镇墩，自动绘制开挖线，自动布置排气排泥阀，自动生成标准横断面，标准平面和标准纵剖面。自动布局裁图，批量打印所有布局。用户在Excel中填一些基本参数，后续交给软件快速得到结果。</p>
      <p class="lead">特色：无需安装；镇墩、弯头、凑合段计算；K/P/S节点标识；一键成图和分步执行均可；提供不少小工具。
    </p>

    <div class="bs-callout bs-callout-warning" id="jquery-required">
      <h4>1.和市面上其他类似软件区别</h4>
      <p>
        无需安装；自动计算镇墩、弯头、凑合段；在excel中填写参数，效率提高。
      </p>
      <h4>2.使用手册</h4>
      <p>
        参见<a href="http://112.74.42.44/project/product/article/57" target="_blank">HydroWS用户手册</a>
      </p>
    </div>

    <h2 id="about-function">
      <a class="anchorjs-link " href="#about-function" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      功能
    </h2>

<ol>
    <li>软件自动完成的功能为：读取平面管线——读取纵面管线、纵面地形线——读取Excel中定义的基本参数和分段数据——完成平面和纵面管线数据合并(即弯头点数据统一)，完成镇墩长度计算，继而完成承插管节数、凑合段长度、弯头长度的分配，完成平面开挖线，完成主要工程量统计，包括开挖量、回填量、管材、镇墩、防腐，完成弯头水力学局部水头损失系数的合计，完成标准平面图和标准纵剖面图，完成自动排气阀、排泥阀布置；完成自动布局裁图(平面和纵剖面置于同一张图幅上)；完成打印所有布局，生成打印文件(plt或pdf)或送至打印机；</li>
<li>按索引沿多段线插入块——本功能的作用如，平面管线上的气阀泥阀需要征地，将征地范围做成一个块，在excel中做成索引如下
桩  号  块  名
500 排气阀井征地范围块
800 排泥阀井征地范围块
1000  检修阀块
1800  泄压阀块
用此功能就可以沿平面管线布置各种定义好的块了，并自动进行角度旋转；</li>
<li>插入典型标准断面 ——按照Excel中定义好的分段管道断面数据，批量插入管道开挖回填标准断面、计算标准断面工程量（未提供），支持单管、双管、不同包角的砂弧基础和砼管座基础、尺寸标注（未提供）；</li>
<li>地形图切剖面 ——根据相交地形线或手工赋值的高程点生成地形剖面线；</li>
<li>多段线与excel 之间互导——点选任意pl线，将坐标导入excel；选择excel两列数据x坐标和y坐标，在cad中绘制pl线；</li>
<li>多段线换向 ——pl线起点和终点进行调换；</li>
<li>根据地面线自动布置管线 ——在地形剖面线基础上，根据用户给定的覆土深度范围，拟合一条管道线，原理是对地形线进行消点处理，减少管线弯头数量；</li>
<li> 根据索引插入光栅图 ——根据excel表格中的光栅图名索引(见下表索引样式)，在cad中自动批量插入光栅图，并进行预裁剪，减轻拼图工作量和避免反复执行裁剪命令；</li>
<li>圆弧变多段线——平面管线如果转弯半径足够大，可以利用多节管道进行借转，此功能在于将圆弧变成多段给定承插管长度(比如5m、6m、12m)的折线；</li>
<li> 修改开挖线——可以根据生成的excel中的数据：地形数据，管道高程数据，用户增加几列：开挖边坡，换填深度等。这样可以分段对管道的开挖进行不同边坡定义，加入换填深度，重新利用公式算出开挖宽度和断面工程量，利用工具再导入cad作为开挖线即可；</li>
<li> 批量标注cad中pl线的坐标、节点号——比如征地范围线，一次性将所有范围线的节点标注出坐标和序号；</li> 
<li> 管网生成hammer或watercad数据——批量选择管网后，生成符合后者数据格式要求的导入文件，见下面3个表格；</li>
<li> 标注平面和纵面管线桩号，方便进行分段数据填写；</li> 
  </ol>

    <h2 id="about-use">
      <a class="anchorjs-link " href="#about-use" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      软件运行环境
    </h2>
<ol>
    <li>供水管线设计工具软件 for AutoCAD 2014、Excel2003～2013</li>
  <li> 适用平台：Windows：Win7-64位</li>
  <li> AutoCAD： 2014</li>
  <li> Excel：2003～2013</li>
  <li> 语言：中文English</li>
  </ol>
  <!-- <img src="/static/img/局域网meritecms.png" style="width: 100%"> -->

    <h1 id="onekey" class="page-header">
      <a class="anchorjs-link " href="#onekey" aria-label="Anchor link for: about" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      一键成图及裁图、打印
    </h1>

    <!-- <p class="lead">
      在这里可以预设和定制系统。
    </p> -->

    <h2 id="onekey-ready">
      <a class="anchorjs-link " href="#onekey-ready" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      供水管线布置准备工作——与程序关系不大
    </h2>
    <ol>
      <li>在测量图上布置平面管线，或其他底图上；</li>
      <li>切地形剖面，用本软件：给剖切线加高程点——地形图切剖面，或其他软件；</li>
      <li>在地面剖切线下进行管道纵向布置；</li>
      <li>用本软件标注平面和纵面桩号，方便填写管线分段数据；</li>
      <li>填写excel表格基本参数表和分段数据</li>
        <!-- <img src="/static/img/个人日历事件.png" style="width: 100%"> -->
      </ol>
    <h2 id="onekey-run">
      <a class="anchorjs-link " href="#onekey-run" aria-label="Anchor link for: whats included precompiled" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      运行程序
    </h2>
    <ol>
    <li>程序运行前将cad和excel表格并排置于屏幕上，方便操作和查看提示；程序运行期间，不要点击excel，否则数据会填写错误。</li>
    <li>程序处理用户误操作功能不强，程序如果进入死循环，需要启动任务管理器强制退出程序。提示用户输入数值时，可回车接受默认值，不接受空格键。ESC键会退出程序，需要重新启动。请谨慎操作，程序力求完善错误处理机制。</li>
    <li>注意：
    多段线分为：优化多段线LightWeightPolyline和三维多段线Polyline，前者又叫二维多段线，没有z值。本程序要求使用前者，即用pl命令生成的二维多段线。
    平面管线、纵面管线、纵面地形线等这3条线是最基本的cad输入条件，要求必须无重复节点，如果有，需要用“desp多段线除重点.lsp”处理一下；
    要求管道分段处要有节点，可以在平面上或纵面上二选一设置一个节点，这里的桩号与分段数据表中的定义管线断面一致。桩号的精度也要完全相同。
    PL线上增加和删除节点用“plve更强多线段定点编辑.LSP”。</li>

    <!-- <img src="/static/img/项目分级目录.png" style="width: 100%"> -->
    </ol>

    <h2 id="onekey-standard">
      <a class="anchorjs-link " href="#onekey-standard" aria-label="Anchor link for: whats included source" data-anchorjs-icon="" style="font-family: anchorjs-icons; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: normal; line-height: inherit; position: absolute; margin-left: -1em; padding-right: 0.5em;"></a>
      生成标准布置图及工程量表格
    </h2>
    <ol>
      <p>此程序工作的前提条件是布置好了平面管线、切好了地面线、布置好了纵面管线、0桩号0高程点。</p>
      <li>点击“getdata2drawing”</li>
      <li>根据提示，先选择布置好的平面管线，数据读入excel表中。</li>
      <li>再选择布置好的纵面管线，数据读入excel表中。</li>
      <li>选择切好的地面线，程序根据以上3条线的数据，结合“基本参数表”和“分段数据”，经过计算，生成“管线中间数据”“管线最终数据”“土方工程量”“”“平纵、数据分离”“设施”。</li>
     </ol>
    <ol>
      <p>程序连续执行所实现的功能如下：</p>
      <li>将平面管线、纵面管线、地形剖切线的节点读入excel表中。进行节点的数据合并等</li>
      <li>将分段数据分配到管线中。</li>
      <li>计算土方工程量</li>
      <li>计算镇墩、弯头等</li>
      <li>汇总管材</li>
      <li>平、纵数据分离</li>
      <li>绘制开挖线</li>
      <li>绘制完整平面图</li>
      <li>绘制完整纵剖面图</li>
      <li>自动布置纵剖面排气排泥设施</li>
      <li>将纵面的排气排泥设施布置到平面上，并导入到excel的设施标签表中。</li>
    </ol>
    <img src="/static/img/HydroWS.png" style="width: 100%">
  </div>
</div>
</div>
  </body>
</html>