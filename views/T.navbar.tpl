{{define "navbar"}}
<nav class="navbar navbar-default">
    <ul class="nav navbar-nav">
      <!-- <li {{if .IsSpider}}class="active"{{end}}>
        <a href="/getspider">水利设计院</a>
      </li> -->
      <li {{if .IsProject}}class="active"{{end}}>
        <a href="/project">项目</a>
      </li>
      <form class="navbar-form navbar-left" role="search" method="get" action="/project/search">
        <div class="form-group">
        <input id="project" type="text" class="form-control"  class="search-query span2" placeholder="Search Project" name="projecttitle"></div>
        <button type="submit" class="btn btn-default">Submit</button>
      </form>
      <!-- <li class="divider">水平分割线</li> -->
      <li>
        <a href="http://192.168.9.13:8081/standard" target="_blank">规范</a>
      </li>
      <li>
        <a href="http://192.168.9.13:8081/legislation" target="_blank">对标</a>
      </li>
      <li>
        <a href="http://192.168.9.13:8080" target="_blank">Merit</a>
      </li>
    </ul>

    <div class="pull-right">
      <ul class="nav navbar-nav">
        {{if .IsLogin}}
          <li class="dropdown">
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Uname}} <b class="caret"></b></a>
            <ul class="dropdown-menu">
              <li><a href="/user/getuserbyusername?username={{.Uname}}">用户资料</a></li>
              <li><a href="/category/viewbyuname?uname={{.Uname}}">项目列表</a></li>
              <li><a href="/topic/viewbyuname?uname={{.Uname}}">成果列表</a></li>
              <li class="divider"></li>
              <li><a href="/login?exit=true">退出</a></li>
            </ul>
          </li>
        {{else}}
          <li>
            <a href="/admin" title="管理">{{.Ip}}</a>
          </li>
        {{end}}
        <!-- <li {{if .IsWiki}}class="active"{{end}}>
          <a href="/wiki">Wiki</a>
        </li>
        <li {{if .IsTask}}class="active"{{end}}>
          <a href="/todo">Todo</a>
        </li>  -->       
      </ul>
    </div>
</nav>
{{end}}

