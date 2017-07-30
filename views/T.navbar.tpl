{{define "navbar"}}
<nav class="navbar navbar-default">
    <ul class="nav navbar-nav">
      <li {{if .IsIndex}}class="active"{{end}}>
        <a href="/index">水利设计院</a>
      </li>
      <!-- <li {{if .IsProjects}}class="active"{{end}}>
        <a href="/project">项目</a>
      </li> -->
      {{if .IsProjects}}
      <li class="active">
        <a href="/project">项目</a>
      </li>
      {{else if .IsProject}}
      <li class="active">
        <a href="/project">项目</a>
      </li>
      {{else}}
      <li>
        <a href="/project">项目</a>
      </li>
      {{end}}

      {{if .IsProjects}}
        <form class="navbar-form navbar-left" role="search" method="get" action="/search">
          <div class="form-group">
          <input type="text" class="form-control"  class="search-query span2" placeholder="Search Projects" name="keyword" id="keyword"></div>
          <button type="submit" class="btn btn-default" id="search">Submit</button>
        </form>
      {{else}}
        <form class="navbar-form navbar-left" role="search" method="get" action="/search">
          <div class="form-group">
          <input type="text" class="form-control"  class="search-query span2" placeholder="Search Products" name="keyword" id="keyword"></div>
          <input type="hidden" name="productid" id="productid" value="{{.Category.Id}}">
          <button type="submit" class="btn btn-default" id="search">Submit</button>
        </form>
      {{end}}

      <li class="dropdown">
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          文档 <b class="caret"></b>
        </a>
        <ul class="dropdown-menu">
          <li>
            <a href="/doc/ecms" title="工程师知识管理系统">EngineerCMS</a>
          </li>
          <li>
            <a href="/doc/meritms" title="价值和成果管理系统">MeritMS</a>
          </li>
        </ul>
      </li>
      <li class="dropdown">
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          API <b class="caret"></b>
        </a>
        <ul class="dropdown-menu">
          <li>
            <a href="https://gowalker.org/github.com/3xxx/EngineerCMS/controllers" title="工程师知识管理系统">EngineerCMS</a>
          </li>
          <li>
            <a href="https://gowalker.org/github.com/3xxx/MeritMS/controllers" title="价值和成果管理系统">MeritMS</a>
          </li>
        </ul>
      </li>
      <li>
        <a href="http://112.74.42.44:8081/standard" target="_blank">规范</a>
      </li>
      <li>
        <a href="http://112.74.42.44:8081/legislation" target="_blank">对标</a>
      </li>
      <li>
        <a href="http://112.74.42.44:8080" target="_blank">MeritMS</a>
      </li>
      <li {{if .IsMeetingroomCalendar}}class="active"{{end}}>
        <a href="/meetingroom">会议室</a>
      </li>
      <li {{if .IsCarCalendar}}class="active"{{end}}>
        <a href="/car">车辆</a>
      </li>
    </ul>

    <div class="pull-right">
      <ul class="nav navbar-nav">
        {{if .IsLogin}}
          {{if .IsAdmin}}
            <li>
              <a href="/admin" title="管理">{{.Ip}}</a>
            </li>
          {{else}}
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
              <ul class="dropdown-menu">
                <li><a href="/calendar" title="日程">日程</a></li>
                <li><a href="/login?exit=true">退出</a></li>
              </ul>
            </li>
          {{end}}
        {{else}}
          {{if .IsAdmin}}
            <li>
              <a href="/admin" title="管理">{{.Ip}}</a>
            </li>
          {{else}}
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Ip}}<b class="caret"></b></a>
              <ul class="dropdown-menu">
                <li><a href="/calendar" title="日程">日程</a></li>
                <li><a href="/login">登录</a></li>
              </ul>
            </li>
          {{end}}
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

