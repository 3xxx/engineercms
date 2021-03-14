<!-- navbar-inverse一个带有黑色背景白色文本的导航栏 
固定在页面的顶部，向 .navbar class 添加 class .navbar-fixed-top
为了防止导航栏与页面主体中的其他内容
的顶部相交错，需要向 <body> 标签添加内边距，内边距的值至少是导航栏的高度。
-->
<style type="text/css">
a.navbar-brand {
  display: none;
}

@media (max-width: 960px) {
  a.navbar-brand {
    display: inline-block;
  }
}
</style>
<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
<nav class="navbar navbar-default navbar-static-top" style="margin-bottom: 5px;" role="navigation">
  <div class="navbar-header">
    <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#target-menu">
      <span class="sr-only">qieh</span>
      <span class="icon-bar"></span>
      <span class="icon-bar"></span>
      <span class="icon-bar"></span>
    </button>
    <a id="11" class="navbar-brand">水利设计</a>
  </div>
  <div class="collapse navbar-collapse" id="target-menu">
    <ul class="nav navbar-nav">
      <li {{if .IsIndex}} class="active" {{end}}>
        <a href="/index">首页</a>
      </li>
      <li {{if .IsProject}} class="active" {{end}}>
        <a href="/project/">项目</a>
      </li>
      <!-- **********定制导航条菜单开始******** -->
      <!-- /project/id——此处id为app.conf里的navigationid1~navigationid9 -->
      <!-- <li {{if .IsNav1}}class="active"{{end}}>
        <a href="/project/25002">设计单位</a>
      </li> -->
      <!-- <li {{if .IsNav3}}class="active"{{end}}>
        <a href="/project/25003">施工专业</a>
      </li> -->
      <!-- <li {{if .IsNav1}} class="dropdown active" {{end}}>
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          试验段 <b class="caret"></b>
        </a>
        <ul class="dropdown-menu">
          <li {{if .IsNav1}} class="active" {{end}}>
            <a href="/project/25002">设计</a>
          </li>
          <li>
            <a href="/project/25005">建设</a>
          </li>
          <li>
            <a href="/project/25004">监理</a>
          </li>
          <li>
            <a href="/project/25003">施工</a>
          </li>
        </ul>
      </li>
      <li {{if .IsNav2}} class="dropdown active" {{end}}>
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          顺德部<b class="caret"></b>
        </a>
        <ul class="dropdown-menu">
          <li {{if .IsNav2}} class="active" {{end}}>
            <a href="/project/26177">设计</a>
          </li>
          <li>
            <a href="/project/26178">建设</a>
          </li>
          <li>
            <a href="/project/26179">监理</a>
          </li>
          <li>
            <a href="/project/26180">施工</a>
          </li>
        </ul>
      </li>
      <li {{if .IsNav3}} class="dropdown active" {{end}}>
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          南沙部<b class="caret"></b>
        </a>
        <ul class="dropdown-menu">
          <li {{if .IsNav3}} class="active" {{end}}>
            <a href="/project/26342">设计</a>
          </li>
          <li>
            <a href="/project/26343">建设</a>
          </li>
          <li>
            <a href="/project/26344">监理</a>
          </li>
          <li>
            <a href="/project/26345">施工</a>
          </li>
        </ul>
      </li>
      <li {{if .IsNav4}} class="dropdown active" {{end}}>
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          东莞部<b class="caret"></b>
        </a>
        <ul class="dropdown-menu">
          <li {{if .IsNav4}} class="active" {{end}}>
            <a href="/project/26507">设计</a>
          </li>
          <li>
            <a href="/project/26508">建设</a>
          </li>
          <li>
            <a href="/project/26509">监理</a>
          </li>
          <li>
            <a href="/project/26510">施工</a>
          </li>
        </ul>
      </li>
      <li {{if .IsNav5}} class="dropdown active" {{end}}>
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          罗田部<b class="caret"></b>
        </a>
        <ul class="dropdown-menu">
          <li {{if .IsNav5}} class="active" {{end}}>
            <a href="/project/26672">设计</a>
          </li>
          <li>
            <a href="/project/26673">建设</a>
          </li>
          <li>
            <a href="/project/26674">监理</a>
          </li>
          <li>
            <a href="/project/26675">施工</a>
          </li>
        </ul>
      </li> -->
      <!-- **********定制导航条菜单结束******** -->
      <li {{if .IsOnlyOffice}} class="active" {{end}}>
        <a href="/onlyoffice">OnlyOffice</a>
      </li>
      <!-- <li {{if or .IsDesignGant .IsConstructGant}}class="dropdown active"{{end}}>
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          进度 <b class="caret"></b>
        </a>
        <ul class="dropdown-menu">
          <li {{if .IsDesignGant}}class="active"{{end}}>
            <a href="/projectgant">设计进度</a>
          </li>
          <li {{if .IsConstructGant}}class="active"{{end}}>
            <a href="/projectgant">施工进度</a>
          </li>
        </ul>
      </li> -->
      <!-- <form class="navbar-form navbar-left" role="search" method="get" action="/search">
        <div class="form-group">
          <input type="text" class="form-control" class="search-query span2" placeholder="Search Products" name="keyword" id="keyword"></div>
        <input type="hidden" name="productid" id="productid" value="{{.Category.Id}}">
        <button type="submit" class="btn btn-default" id="search">Submit</button>
      </form> -->
      <li {{if or .IsStandard .IsLegislation}} class="dropdown active" {{end}} class="dropdown">
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          规范 <b class="caret"></b>
        </a>
        <ul class="dropdown-menu">
          <li {{if .IsStandard}} class="active" {{end}}>
            <a href="/standard" target="_blank">查阅</a>
          </li>
          <li {{if .IsLegislation}} class="active" {{end}}>
            <a href="/legislation" target="_blank">输入框对标</a>
          </li>
          <li {{if .IsLegislation_upfile}} class="active" {{end}}>
            <a href="/legislation/fileinput" target="_blank">excel对标</a>
          </li>
          <li {{if .IsLegislation_wordfile}} class="active" {{end}}>
            <a href="/legislation/wordfileinput" target="_blank">word对标</a>
          </li>
        </ul>
      </li>
      <!-- <li {{if .IsMonthCheck}}class="active"{{end}}>
        <a href="/v1/checkin/monthchecksum">考勤</a>
      </li>
      <li>
        <a href="/diary">日志</a>
      </li> -->
      <!-- <li {{if or .IsMeetingroomCalendar .IsCarCalendar .IsOrderCalendar .IsAttendanceCalendar}} class="dropdown active" {{end}}>
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          预订 <b class="caret"></b>
        </a>
        <ul class="dropdown-menu">
          <li {{if .IsMeetingroomCalendar}} class="active" {{end}}>
            <a href="/meetingroom">会议室&值班安排</a>
          </li>
          <li {{if .IsCarCalendar}} class="active" {{end}}>
            <a href="/car">车辆</a>
          </li>
          <li {{if .IsOrderCalendar}} class="active" {{end}}>
            <a href="/order">订餐</a>
          </li>
          <li {{if .IsAttendanceCalendar}} class="active" {{end}}>
            <a href="/attendance">考勤</a>
          </li>
        </ul>
      </li> -->
      <li {{if .IsBook}} class="active" {{end}}>
        <a href="/mindoc" title="创作你的书籍" target="_blank">BOOK</a>
      </li>
      <li {{if .IsFLV}} class="active" {{end}}>
        <a href="/v1/flv/flvlist" title="视频点播" target="_blank">视频</a>
      </li>
      <li {{if .IsPhotoWipe}} class="active" {{end}}>
        <a href="/v1/wx/photoswipe" title="相册" target="_blank">相册</a>
      </li>
      <li {{if .IsPhotoWipe}} class="active" {{end}}>
        <a href="http://192.168.100.37:8080/merit" title="价值" target="_blank">价值</a>
      </li>
    </ul>
    <div class="pull-right">
      <ul class="nav navbar-nav navbar-right">
        {{if eq true .IsAdmin}}
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
          <ul class="dropdown-menu">
            <li><a href="/admin" title="管理">进入后台</a></li>
            <li><a href="/cms/#/flow/usertobeprocessed" title="邮箱"><i class="fa fa-envelope">&nbsp; 邮箱</i></a></li>
            <li><a href="/v1/cart/getcart" title="购物车"><i class="fa fa-shopping-cart">&nbsp; 购物车</i></a></li>
            <li><a href="javascript:void(0)" id="login">重新登录</a></li>
            <li><a href="/v1/wx/ssologin" title="单点登录">SSO单点登陆</a></li>
            <li><a href="/project/25001/gettimeline" title="大事记">大事记</a></li>
            <li><a href="/project/25001/getcalendar" title="项目日历"><i class="fa fa-calendar">&nbsp; 项目日历</i></a></li>
            <li><a href="/calendar" title="日程安排"><i class="fa fa-calendar-plus-o">&nbsp; 日程安排</i></a></li>
            <li><a href="javascript:void(0)" onclick="logout()">退出</a></li>
          </ul>
        </li>
        {{else if eq true .IsLogin}}
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
          <ul class="dropdown-menu">
            <li><a href="/cms/#/flow/usertobeprocessed" title="邮箱"><i class="fa fa-envelope">&nbsp; 邮箱</i></a></li>
            <li><a href="/v1/cart/getcart" title="购物车"><i class="fa fa-shopping-cart">&nbsp; 购物车</i></a></li>
            <li><a href="/user" title="用户资料">用户资料</a></li>
            <li><a href="javascript:void(0)" id="login">重新登录</a></li>
            <li><a href="/v1/wx/ssologin" title="单点登录">SSO单点登陆</a></li>
            <li><a href="/project/25001/gettimeline" title="大事记">大事记</a></li>
            <li><a href="/project/25001/getcalendar" title="项目日历"><i class="fa fa-calendar">&nbsp; 项目日历</i></a></li>
            <li><a href="/calendar" title="日程安排"><i class="fa fa-calendar-plus-o">&nbsp; 日程安排</i></a></li>
            <li><a href="javascript:void(0)" onclick="logout()">退出</a></li>
          </ul>
        </li>
        {{else}}
        <!-- <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
          <ul class="dropdown-menu">
            <li><a href="/admin" title="管理">进入后台</a></li>
            <li><a href="/cms/#/flow/usertobeprocessed" title="邮箱">邮箱</a></li>
            <li><a href="javascript:void(0)" id="login">重新登录</a></li>
            <li><a href="/v1/wx/ssologin" title="单点登录">SSO单点登陆</a></li>
            <li><a href="/project/25001/gettimeline" title="大事记">大事记</a></li>
            <li><a href="/project/25001/getcalendar" title="项目日历">项目日历</a></li>
            <li><a href="/calendar" title="日程安排">日程安排</a></li>
            <li><a href="javascript:void(0)" onclick="logout()">退出</a></li>
          </ul>
        </li> -->
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
          <ul class="dropdown-menu">
            <li><a href="javascript:void(0)" id="login">登陆</a></li>
            <li><a href="/v1/wx/ssologin" title="单点登录">SSO单点登陆</a></li>
          </ul>
        </li>
        {{end}}
        <li {{if .IsWiki}} class="active" {{end}}>
          <a href="/wiki">Wiki</a>
        </li>
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">
            帮助 <b class="caret"></b>
          </a>
          <ul class="dropdown-menu">
            <li>
              <a href="/doc/ecms" title="工程师知识管理系统">EngineerCMS</a>
            </li>
            <li>
              <a href="/doc/meritms" title="价值和成果管理系统">MeritMS</a>
            </li>
            <li>
              <a href="/doc/hydrows" title="水利供水管线设计工具">HydroWS</a>
            </li>
          </ul>
        </li>
      </ul>
    </div>
  </div>
</nav>
<!-- 登录模态框 -->
<div class="form-horizontal">
  <div class="modal fade" id="modalNav">
    <div class="modal-dialog" id="modalDialog">
      <div class="modal-content">
        <div class="modal-header" style="background-color: #8bc34a">
          <button type="button" class="close" data-dismiss="modal">
            <span aria-hidden="true">&times;</span>
          </button>
          <h3 class="modal-title">登录</h3>
          <label id="status"></label>
        </div>
        <div class="modal-body">
          <div class="modal-body-content">
            <div class="form-group" style="width: 100%;">
              <label class="col-sm-3 control-label">用户名 或 邮箱</label>
              <div class="col-sm-7">
                <input id="uname" name="uname" type="text" value="" class="form-control" placeholder="Enter account" list="cars" onkeypress="getKey()"></div>
            </div>
            <div class="form-group" style="width: 100%;">
              <label class="col-sm-3 control-label">密码</label>
              <div class="col-sm-7">
                <input id="pwd" name="pwd" type="password" value="" class="form-control" placeholder="Password" onkeypress="getKey()"></div>
            </div>
            <div class="form-group" style="width: 100%;">
              <label class="col-sm-3 control-label"><input type="checkbox">自动登陆</label>
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          <button type="button" class="btn btn-primary" onclick="login()">登录</button>
        </div>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript">
// 弹出登录模态框
$("#login").click(function() {
  $('#modalNav').modal({
    show: true,
    backdrop: 'static'
  });
})

//登陆功能
function login() {
  var uname = document.getElementById("uname");
  if (uname.value.length == 0) {
    alert("请输入账号");
    return
  }
  var pwd = document.getElementById("pwd");
  if (pwd.value.length == 0) {
    alert("请输入密码");
    return
  }

  $.ajax({
    type: 'post',
    url: '/loginpost',
    data: {
      "uname": $("#uname").val(),
      "pwd": $("#pwd").val()
    },
    success: function(result) {
      if (result.islogin == 0) {
        $("#status").html("登陆成功");
        $('#modalNav').modal('hide');
        window.location.reload();
      } else if (result.islogin == 1) {
        $("#status").html("用户名或密码错误！")
      } else if (result.islogin == 2) {
        $("#status").html("密码错误")
      }
    }
  })
}
//登出功能
function logout() {
  $.ajax({
    type: 'get',
    url: '/logout',
    data: {},
    success: function(result) {
      if (result.islogin) {
        // $("#status").html("登出成功");
        alert("登出成功");
        window.location.reload();
      } else {
        // $("#status").html("登出失败");
        alert("登出失败")
      }
    }
  })
}

function getKey() {
  if (event.keyCode == 13) {
    login()
  }
}
</script>

<!--前端递归解析json 数据
 $(function () {
        var showlist = $("<ul></ul>");
        showall(menulist.menulist, showlist); 
        $("#div_menu").append(showlist);
});
/**
 * parent为要组合成html的容器
* menu_list为后台json数据
 */
function showall(menu_list, parent) {
    for (var menu in menu_list) {
        //如果有子节点，则遍历该子节点
        if (menu_list[menu].menulist.length > 0) {
            //创建一个子节点li
            var li = $("<li></li>");
            //将li的文本设置好，并马上添加一个空白的ul子节点，并且将这个li添加到父亲节点中
            $(li).append(menu_list[menu].MName).append("<ul></ul>").appendTo(parent);
            //将空白的ul作为下一个递归遍历的父亲节点传入
            showall(menu_list[menu].menulist, $(li).children().eq(0));
        }
        //如果该节点没有子节点，则直接将该节点li以及文本创建好直接添加到父亲节点中
        else {
           $("<li></li>").append(menu_list[menu].MName).appendTo(parent);
        }
    }
} 

/**
 * Created on 2017/6/27.
 */
$(function () {
    $.getJSON({
        type: "get",
        url: "dist/json/nav.json",
        success: function (data) {
            var showList = $("<ul class='" + data.ulClass + "'><li class='header'>主导航</li></ul>");
            showAll(data, showList);
            $(".sidebar").append(showList);
        }
    });
    //data为json数据
    //parent为要组合成html的容器
    function showAll(data, parent) {
        $.each(data.children, function (index, fatherLi) {//遍历数据集
            var li1 = $("<li class='" + fatherLi.liClass + "'><a href='" + fatherLi.link + "'><i class=" + fatherLi.iClass + "></i>" + fatherLi.label + "</a></li>");//没有children的初始li结构
            var li2 = $("<li class='" + fatherLi.liClass + "'><a href='" + fatherLi.link + "'><i class=" + fatherLi.iClass + "></i>" + fatherLi.label + "<span class='" + fatherLi.spanClass + "'><i class='" + fatherLi.spanChildIClass + "'></i></span>" + "</a></li>");//有children的初始li结构
            //console.log($(li1).html());
            //console.log($(li2).html());
            if (fatherLi.children.length > 0) { //如果有子节点，则遍历该子节点
                var ul = $("<ul class='" + fatherLi.children[0].ulClass + "'></ul>");
                $(li2).append(ul).appendTo(parent);//将li的初始化选择好，并马上添加带类名的ul子节点，并且将这个li添加到父亲节点中
                showAll(fatherLi.children[0], $(li2).children().eq(1));//将空白的ul作为下一个递归遍历的父亲节点传入，递归调用showAll函数
            } else {
                $(li1).appendTo(parent);//如果该节点没有子节点，则直接将该节点li以及文本创建好直接添加到父亲节点中
            }
        });
    }
});

js代码



-->