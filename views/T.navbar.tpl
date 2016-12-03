{{define "navbar"}}
<nav class="navbar navbar-default">
    <ul class="nav navbar-nav">
      <li {{if .IsSpider}}class="active"{{end}}>
        <a href="/getspider">水利设计院</a>
      </li>
      <li {{if .IsCategory}}class="active"{{end}}>
        <a href="/project">项目</a>
      </li>
      <form class="navbar-form navbar-left" role="search" method="get" action="/searchcategory">
        <div class="form-group">
        <input id="category" type="text" class="form-control"  class="search-query span2" placeholder="Search Project" name="categoryname"></div>
        <button type="submit" class="btn btn-default">Submit</button>
      </form>

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
          <a href="/login">登陆</a>
        </li>
        {{end}}
        <li {{if .IsWiki}}class="active"{{end}}>
          <a href="/wiki">Wiki</a>
        </li>
        <li {{if .IsTask}}class="active"{{end}}>
          <a href="/todo">Todo</a>
        </li>        
      </ul>
    </div>
</nav>
{{end}}

