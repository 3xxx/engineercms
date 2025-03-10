{{define "navbar"}}
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

#NavmodalTable .modal-header {
  cursor: move;
}
</style>
<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
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
        <a href="/project/" id="project">项目</a>
      </li>
      <!-- **********定制导航条菜单开始******** -->
      <!-- **********定制导航条菜单结束******** -->
      <li {{if .IsOnlyOffice}} class="active" {{end}}>
        <a href="/onlyoffice">OnlyOffice</a>
      </li>

      <li {{if .IsPhotoWipe}} class="active" {{end}}>
        <a href="http://192.168.100.37:8080/merit" title="价值" target="_blank">价值</a>
      </li>
    </ul>
    <div class="pull-right">
      <ul class="nav navbar-nav navbar-right">
        <li><a class="navbar-brand" href="javascript:void(0)" onclick="chooseProjectButton()" id="chooseProject">切换项目</a></li>
        {{if eq true .IsAdmin}}
        
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
          <ul class="dropdown-menu">
            <li><a href="/admin" title="管理">进入后台</a></li>
            <li><a href="/cms/#/flow/usertobeprocessed" title="邮箱"><i class="fa fa-envelope">&nbsp; 邮箱</i></a></li>
            <li><a href="/v1/cart/getcart" title="购物车"><i class="fa fa-shopping-cart">&nbsp; 购物车</i></a></li>
            <li><a href="javascript:void(0)" id="login">重新登录</a></li>
            <li><a href="/v1/wx/ssologin" title="单点登录">SSO单点登陆</a></li>

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

            <li><a href="/calendar" title="日程安排"><i class="fa fa-calendar-plus-o">&nbsp; 日程安排</i></a></li>
            <li><a href="javascript:void(0)" onclick="logout()">退出</a></li>
          </ul>
        </li>
        
        {{else}}
        

        
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
          <ul class="dropdown-menu">
            <li><a href="javascript:void(0)" id="login">登陆</a></li>
            <li><a href="/v1/wx/ssologin" title="单点登录">SSO单点登陆</a></li>
          </ul>
        </li>
        
        {{end}}
        <li {{if .IsChat}} class="active" {{end}}>
          <a href="/v1/chat/chat">Chat</a>
        </li>
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
<!-- 项目切换模态框 -->
<div class="form-horizontal">
  <div class="modal fade" id="NavmodalTable">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal">
            <span aria-hidden="true">&times;</span>
          </button>
          <h3 class="modal-title">切换项目</h3>
        </div>
        <div class="modal-body">
          <table id="Navtable2"></table>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="button" class="btn btn-primary" id="saveproj" onclick="setlocalstorage()">保存</button>
        </div>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript">
$(function() {
  // console.log(window.localStorage.getItem('projectid'))
  var projectid = window.localStorage.getItem('projectid')
  // console.log(projectid)
  if (projectid == null) {
    document.getElementById('chooseProject').innerText = '选择项目';
  } else {
    // $('#project').attr('href', '/project/' + projectid);
  }
})

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
      if (!result.islogin) {
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

$(function() {
  var projectid = window.localStorage.getItem('projectid')
  // 初始化【未接受】工作流表格
  $("#Navtable2").bootstrapTable({
    url: '/project/getprojects',
    method: 'get',
    search: 'true',
    showRefresh: 'true',
    showToggle: 'true',
    showColumns: 'true',
    // toolbar:'#toolbar1',
    pagination: 'true',
    sidePagination: "server",
    queryParamsType: '',
    //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
    // limit, offset, search, sort, order 否则, 需要包含:
    // pageSize, pageNumber, searchText, sortName, sortOrder.
    // 返回false将会终止请求。
    pageSize: 15,
    pageNumber: 1,
    pageList: [15, 20, 50, 100],
    singleSelect: "true",
    clickToSelect: "true",
    selectItemName: "project",
    queryParams: function queryParams(params) { //设置查询参数
      var param = {
        limit: params.pageSize, //每页多少条数据
        pageNo: params.pageNumber, // 页码
        searchText: $(".search .form-control").val()
      };
      //搜索框功能
      //当查询条件中包含中文时，get请求默认会使用ISO-8859-1编码请求参数，在服务端需要对其解码
      // if (null != searchText) {
      //   try {
      //     searchText = new String(searchText.getBytes("ISO-8859-1"), "UTF-8");
      //   } catch (Exception e) {
      //     e.printStackTrace();
      //   }
      // }
      return param;
    },
    columns: [{
        title: '选择',
        radio: 'true',
        width: '10',
        align: "center",
        valign: "middle",
        formatter: function(value, row, index) {
          return {checked: row.Id==projectid}
        },
      },
      {
        // field: 'Number',
        title: '序号',
        formatter: function(value, row, index) {
          return index + 1
        },
        align: "center",
        valign: "middle"
      },
      {
        field: 'Code',
        title: '编号',
        // formatter:setCode,
        align: "center",
        valign: "middle"
      },
      {
        field: 'Title',
        title: '名称',
        // formatter:setTitle,
        align: "center",
        valign: "middle"
      },
      // {
      //   field: 'Label',
      //   title: '标签',
      //   formatter:setLable,
      //   align:"center",
      //   valign:"middle"
      // },
      {
        field: 'Principal',
        title: '负责人',
        align: "center",
        valign: "middle"
      },
      // {
      //   field: 'Number',
      //   title: '成果数量',
      //   formatter:setCode,
      //   align:"center",
      //   valign:"middle"
      // },
      // {
      //   field: 'action',
      //   title: '时间轴',
      //   formatter:actionFormatter,
      //   events:actionEvents,
      //   align:"center",
      //   valign:"middle"
      // },
      {
        field: 'Created',
        title: '建立时间',
        formatter: localDateFormatter,
        align: "center",
        valign: "middle"
      }
      // {
      // field: 'dContMainEntity.createTime',
      // title: '发起时间',
      // formatter: function (value, row, index) {
      // return new Date(value).toLocaleString().substring(0,9);
      // }
      // },
      // {
      // field: 'dContMainEntity.operate',
      // title: '操作',
      // formatter: operateFormatter
      // }
    ]
  });
});

// 切换项目
function chooseProjectButton() {
  if (!{{.IsLogin }}) {
    alert("请登录！");
    return;
  }

  // $('#project').$("input[name='types']").attr('checked','true');

  $('#NavmodalTable').modal({
    show: true,
    backdrop: 'static'
  });
}

// 将选择的项目id存入浏览器内存
function setlocalstorage() {
  var selectRow2 = $('#Navtable2').bootstrapTable('getSelections');
  if (selectRow2.length < 1) {
    alert("请先勾选项目！");
    return;
  }
  console.log(selectRow2[0].Id)
  window.localStorage.setItem('projectid', selectRow2[0].Id);
  $('#NavmodalTable').modal('hide');
  // window.location.reload();
}

$(function() {
  $("#NavmodalTable").draggable({ handle: ".modal-header" }); //为模态对话框添加拖拽
  $("#myModal").css("overflow", "hidden"); //禁止模态对话框的半透明背景滚动
})

function localDateFormatter(value) {
  return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
}
</script>
{{end}}
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