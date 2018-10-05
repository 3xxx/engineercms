{{define "mnavbar"}}
<nav class="navbar navbar-default navbar-static-top">
    <ul class="nav navbar-nav">
      <!-- <li {{if .IsIndex}}class="active"{{end}}>
        <a href="/index">首页</a>
      </li>
      <li {{if .IsProject}}class="active"{{end}}>
        <a href="/project/25001">项目</a>
      </li> -->
      <li {{if .IsBuild}}class="active"{{end}}>
        <a href="/project/25005">建设</a>
      </li>
      <li {{if .IsDesign}}class="active"{{end}}>
        <a href="/project/25002">设计</a>
      </li>
      <li {{if .IsSupervision}}class="active"{{end}}>
        <a href="/project/25004">监理</a>
      </li>
      <li {{if .IsConstruct}}class="active"{{end}}>
        <a href="/project/25003">施工</a>
      </li>
        <form class="navbar-form navbar-left" role="search" method="get" action="/search">
          <div class="form-group">
          <input type="text" class="form-control"  class="search-query span2" placeholder="Search Products" name="keyword" id="keyword"></div>
          <input type="hidden" name="productid" id="productid" value="{{.Category.Id}}">
          <button type="submit" class="btn btn-default" id="search">Submit</button>
        </form>
      <li class="dropdown">
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          规范 <b class="caret"></b>
        </a>
        <ul class="dropdown-menu">
          <li>
            <a href="http://112.74.42.44:8081" target="_blank">查询</a>
          </li>
          <li>
            <a href="http://112.74.42.44:8081/legislation" target="_blank">对标</a>
          </li>
        </ul>
      </li>
      <!-- <li {{if .IsMeetingroomCalendar}}class="active"{{end}}>
        <a href="/meetingroom">会议室</a>
      </li>
      <li {{if .IsCarCalendar}}class="active"{{end}}>
        <a href="/car">车辆</a>
      </li>
      <li {{if .IsOrderCalendar}}class="active"{{end}}>
        <a href="/order">订餐</a>
      </li>
      <li {{if .IsAttendanceCalendar}}class="active"{{end}}>
        <a href="/attendance">考勤</a>
      </li> -->
    </ul>

    <div class="pull-right">
      <ul class="nav navbar-nav">
        {{if or .IsLogin .IsAdmin}}
        <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
              <ul class="dropdown-menu">
                <!-- <li><a href="/admin" title="管理">进入后台</a></li> -->
                <li><a href="/project/25001/gettimeline" title="大事记">大事记</a></li>
                <li><a href="/project/25001/getcalendar" title="项目日历">项目日历</a></li>
                <li><a href="/calendar" title="日程安排">日程安排</a></li>
                <li><a href="/login?exit=true">退出</a></li>
              </ul>
            </li>
        {{else}}
        <li>
          <a href="/login">登录</a>
        </li>
        {{end}}
      </ul>
    </div>
</nav>
{{end}}

