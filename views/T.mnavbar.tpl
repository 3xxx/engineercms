{{define "mnavbar"}}

<div class="blog_main">

    <div class="main_list">

      <div class="blog_article">

        <div class="blog_article_c clearfix">
          <p class="date">
            <div class="col-xs-3 col-sm-3 col-md-3 col-lg-3">
              <div id="tree"></div>
            </div>
            <div class="col-xs-9 col-sm-9 col-md-9 col-lg-9">
              <!-- 面包屑导航 -->
              <div class="breadcrumbs">
    <ol class="breadcrumb" split="&gt;">
      <li>
        <a href="javascript:void(0)"> <i class="fa fa-home" aria-hidden="true"></i>
          项目编号：{{.Category.Code}}
        </a>
      </li>

    </ol>
              </div>

              <iframe src="/project/{{.Id}}/{{.Id}}" name='iframepage' id="iframepage" frameborder="0"  width="100%" scrolling="no" marginheight="0" marginwidth="0"  onload="this.height=800"></iframe> 
            </div> 
          </p>
        </div>
      </div>
    </div>

    <div class="blog_top_wrap">
      <div class="blog_top"> <i id="menu_J" class="iconfont icon_l">导航</i>
        <form method="get" action="http://m.blog.csdn.net/search/index" id="searchform"> <i id="search_J" class="iconfont icon_r">搜索</i>
          <div id="search_c_J" class="search">
            <input type="text" placeholder="请输入" id="search" name="keyword" page="1" value="">
            <i class="iconfont icon_search"></i>
            <i class="iconfont icon_close">关闭</i>
          </div>
        </form>

        <h2 class="blog_top_t">EngineerCMS</h2>
      </div>
    </div>


    <div class="leftNav" style="left: -40rem;">

      <div class="left_top">
        <a href="http://m.blog.csdn.net/Blog?username=hotqin888">
          <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/1_hotqin888.jpg" alt="hotqin888"></a>
        <a href="http://m.blog.csdn.net/Blog?username=hotqin888" class="sign">hotqin888</a>
      </div>

      <ul class="nav_list">
        <li>
          <a href="http://m.blog.csdn.net/home/index">
            <i>•</i>
            <span>首页</span>
            <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
        </li>
        <li>
          <a href="http://m.blog.csdn.net/Column/Column?Channel=mobile&amp;Type=">
            <i>•</i>
            <span>移动开发</span>
            <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
        </li>
        <li>
          <a href="http://m.blog.csdn.net/Column/Column?Channel=enterprise&amp;Type=">
            <i>•</i>
            <span>架构</span>
            <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
        </li>
        <li>
          <a href="http://m.blog.csdn.net/Column/Column?Channel=cloud&amp;Type=">
            <i>•</i>
            <span>云计算/大数据</span>
            <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
        </li>
        <li>
          <a href="http://m.blog.csdn.net/Column/Column?Channel=www&amp;Type=">
            <i>•</i>
            <span>互联网</span>
            <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
        </li>
        <li>
          <a href="http://m.blog.csdn.net/Column/Column?Channel=system&amp;Type=">
            <i>•</i>
            <span>运维</span>
            <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
        </li>
        <li>
          <a href="http://m.blog.csdn.net/Column/Column?Channel=database&amp;Type=">
            <i>•</i>
            <span>数据库</span>
            <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
        </li>
        <li>
          <a href="http://m.blog.csdn.net/Column/Column?Channel=web&amp;Type=">
            <i>•</i>
            <span>前端</span>
            <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
        </li>
        <li>
          <a href="http://m.blog.csdn.net/Column/Column?Channel=code&amp;Type=">
            <i>•</i>
            <span>编程语言</span>
            <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
        </li>
        <li>
          <a href="http://m.blog.csdn.net/Column/Column?Channel=software&amp;Type=">
            <i>•</i>
            <span>研发管理</span>
            <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
        </li>
        <li>
          <a href="http://m.blog.csdn.net/Column/Column?Channel=other&amp;Type=">
            <i>•</i>
            <span>综合</span>
            <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-youjiantou.png" alt="img" class="arrow_r"></a>
        </li>
      </ul>
    </div>
    <!-- 这个重要，msk点一下可以缩回侧栏 -->
    <div id="mask" style="display: none;"></div>
  </div>

  <div class="backToTop" style="z-index: 9999999; display: block;">
    <img src="./beego利用casbin进行权限管理——第一节 起步、测试 - hotqin888的专栏 - CSDN博客_files/iconfont-fudongxiangshang.png" alt="img"></div>

{{end}}