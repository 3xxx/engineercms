<!-- 后台主页面，其他为子页面-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>

  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-treeview.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css"/>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>

<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>
<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>

<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
<script src="/static/js/tableExport.js"></script>

</head>
<body>
<div class="col-lg-2">
  <div id="tree"></div>
</div>
<!-- 菜单顶部 -->
<!-- <nav class="navbar navbar-default"> -->
    <!-- <ul class="nav navbar-nav"> -->
<!-- <div class="navbar navba-default navbar-fixed-top"> -->
  <!-- <div class="container-fill"> -->
  <!-- <div class="main-panel"> -->
  <div class="col-lg-10">
      
      <div class="navbar navbar-top">
          <ul class="nav navbar-nav navbar-right">
          <!-- <ul class="nav navbar-nav"> -->
              <li>
                  <a href="/project">项目</a>
              </li>
              <li>
                  <a href="/admin/user/detail">{{.Ip}}</a>
              </li>
              <li class="dropdown">
                  <a href="#" class="dropdown-toggle" data-toggle="dropdown">Dropdown<b class="caret"></b></a>
                  <ul class="dropdown-menu">
                      <li><a href="https://github.com/3xxx" target="_blank">3xxx github</a></li>
                      <li><a href="http://www.sina.com/xc-qin" target="_blank">Weibo</a></li>
                      <li><a href="#">Something</a></li>
                      <li class="divider"></li>
                      <li><a href="http://blog.csdn.net/hotqin888" target="_blank">Blog</a></li>
                  </ul>
              </li>
              <li>
                  <a href="/admin/login/out">Log out</a>
              </li>
          </ul>
      </div>
    <!-- </nav> -->

    <!-- </div> -->
    <!-- 面包屑导航 -->
    <!-- <div class="main-panel-breadcrumb"> -->
      <!-- {include file="./_layout/breadcrumb" /} -->
    <!-- </div> -->
  <!-- </div> -->
<!-- </div>     -->
<!-- if (!empty($parents)) { -->
<!-- $html  = ""; -->
<!-- <div class="main-panel-breadcrumb"> -->
  <div class="breadcrumbs">
    <ol class="breadcrumb" split="&gt;">
      <li>
        <a href="javascript:void(0)"> <i class="fa fa-home" aria-hidden="true"></i>
          后台
        </a>
      </li>
      <li>
        <a href="javascript:void(0)"> <i class="fa '. $parents['picon'] .' " aria-hidden="true"></i>
          日历
        </a>
      </li>
<!--       <li>
        <a href="javascript:void(0)">
          <i class="fa '. $parents['icon'] .' " aria-hidden="true"></i>
          分级目录
        </a>
      </li> -->
      <!-- <li>
        <a href="javascript:void(0)">
          <i class="fa '. $parents['act_icon'] .' " aria-hidden="true"></i>
          parents
        </a>
      </li> -->
    </ol>
  </div>
<!-- </div> -->
</div>



<script type="text/javascript">
    $(function () {
        // function getTree() {
          // text: "Node 1",
          // icon: "glyphicon glyphicon-stop",
          // selectedIcon: "glyphicon glyphicon-stop",
          // color: "#000000",
          // backColor: "#FFFFFF",
          // href: "#node-1",
          // selectable: true,
          // state: {
          //   checked: true,
          //   disabled: true,
          //   expanded: true,
          //   selected: true
          // },
          // tags: ['available'],
          // Some logic to retrieve, or generate tree structure
            // <div class="hello-user">欢迎您~'. $_authUser['username'] .'  </div>
            var data = 
            [
              {
                text: "欢迎您~{{.Ip}}", 
                text1: "欢迎您~{{.Ip}}",
                selectable: true,
                id: '010',
              },
              {
                text: "系统设置",
                text1: "欢迎您~{{.Ip}}",
                icon: "fa fa-tachometer icon",
                // selectedIcon: "glyphicon glyphicon-stop",
                href: "#node-1",
                // selectable: true,
                id: '01',
                selectable: false,
                // state: {
                  // checked: true,
                  // disabled: true,
                  // expanded: true,
                  // selected: true
                // },
                tags: ['available'],
                nodes: 
                [
                  { 
                    icon: "fa fa-cog",
                    text: "基本设置",
                    text1: "欢迎您~{{.Ip}}",
                    id: '011',
                    nodeId: '011'
                  },
                  { 
                    icon: "fa fa-align-right",
                    text: "分级目录",
                    id: '012',
                    nodeId: '012'
                  }, 
                  { 
                    icon: "fa fa-bug",
                    text: "搜索IP",
                    id: '013'
                  }
                ] 
              },
              {
                text: "权限管理",
                icon: "fa fa-balance-scale",
                // selectedIcon: "glyphicon glyphicon-stop",
                href: "#node-1",
                // selectable: true,
                id: '02',
                selectable: false,
                // state: {
                  // checked: true,
                  // disabled: true,
                  // expanded: true,
                  // selected: true
                // },
                tags: ['available'],
                nodes: 
                [
                  { icon: "fa fa-safari",
                    text: '系统权限',
                    id: '021',
                    state: {
                      // checked: true,
                      disabled: true,
                      // expanded: true,
                      // selected: true
                    }
                  },
                  { icon: "fa fa-navicon",
                    text: '项目权限',
                    id: '022',
                    state: {
                      // checked: true,
                      disabled: true,
                      // expanded: true,
                      // selected: true
                    }
                  }
                ]
              },
              {
                text: "账号管理",
                icon: "fa fa-users icon",
                // selectedIcon: "glyphicon glyphicon-stop",
                href: "#node-1",
                // selectable: true,
                id: '03',
                selectable: false,
                // state: {
                  // checked: true,
                  // disabled: true,
                  // expanded: true,
                  // selected: true
                // },
                tags: ['available'],
                nodes: 
                [
                  { 
                    icon: "fa fa-align-right",
                    text: "组织结构",
                    id: '030'
                  },
                  { icon: "fa fa-users",
                    text: '用户',
                    id: '031'
                  },
                  { icon: "fa fa-th",
                    text: 'IP地址段',
                    id: '032'
                  },
                  { icon: "fa fa-group",
                    text: '用户组',
                    id: '033'
                  },
                ]
              }, 
              {
                text: "项目设置",
                icon: "fa fa-list-alt icon",
                // selectedIcon: "glyphicon glyphicon-stop",
                href: "#node-1",
                // selectable: true,
                id: '04',
                selectable: false,
                showTags:false,
                tags: ['available'],
                nodes: 
                [
                  { 
                    icon: "fa fa-edit",
                    text: "编辑目录",
                    id: '041',
                    tags: [''],
                  },
                  { 
                    icon: "fa fa-edit",
                    text: "同步IP",
                    id: '042'
                  },
                  { 
                    icon: "fa fa-lock",
                    text: "项目权限",
                    id: '043'
                  },
                  { 
                    icon: "fa fa-copy",
                    text: "快捷编辑",
                    id: '044',
                  }
                ]
              } 
            ]
            // return data;

        $('#tree').treeview({
            data: data,         // data is not optional
            levels: 2,
            enableLinks: true,
            showTags:false,
            // showCheckbox: true,
            state: {
              checked: true,
              disabled: true,
              expanded: true,
              selected: true
            }
            // multiSelect: true
          });  
        // }
          // alert(JSON.stringify({{.json}}));
         // $('#treeview').treeview('collapseAll', { silent: true });
          // $('#tree').treeview({
          // data: [{{.json}}],//defaultData,
          // data:alternateData,
          // levels: 3,// expanded to 5 levels
          // enableLinks:true,
          // showTags:true,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
        // });

        $('#tree').on('nodeSelected', function(event, data) {
            // alert("名称："+data.text);
            // alert("节点id："+data.nodeId);
            // alert("部门id："+data.id);  
            // alert("部门级别："+data.Level);
            // $("#regis").html(data.text);//显示部门名称
            // $("#regis").css("color","black");
          document.getElementById("iframepage").src="/admin/"+data.id;
          if (data.id=="010"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;日历")
          }else if (data.id=="011"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;系统设置&gt;"+data.text)
          }else if(data.id=="012"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;系统设置&gt;"+data.text)
          }else if(data.id=="013"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;系统设置&gt;"+data.text)
          }else if(data.id=="021"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;权限管理&gt;"+data.text)
          }else if(data.id=="022"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;权限管理&gt;"+data.text)
          }else if(data.id=="030"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;账号管理&gt;"+data.text)
          }else if(data.id=="031"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;账号管理&gt;"+data.text)
          }else if(data.id=="032"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;账号管理&gt;"+data.text)
          }else if(data.id=="033"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;账号管理&gt;"+data.text)
          }else if(data.id=="041"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;项目设置&gt;"+data.text)
          }else if(data.id=="042"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;项目设置&gt;"+data.text)
          }else if(data.id=="043"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;项目设置&gt;"+data.text)
          }else if(data.id=="044"){
            $(".breadcrumb").html("<i class='fa fa-home'></i>后台&gt;项目设置&gt;"+data.text)
          }
          //?secid="+data.Id+"&level="+data.Level;
        });

        // var obj = {};
        // obj.text = "123";
        // $("#btn").click(function (e) {
        //     var arr = $('#tree').treeview('getSelected');
        //     for (var key in arr) {
        //         c.innerHTML = c.innerHTML + "," + arr[key].id;
        //     }
        // })
    }) 

    function index1(value,row,index){
    // alert( "Data Loaded: " + index );
      return index+1
    }
</script>

<div class="col-lg-10">
    <!-- <div class="form-group"> -->
        <!-- <label class="control-label" id="regis" for="LoginForm-UserName"></label> 显示部门名称  -->
    <!-- </div> -->
        <!-- <iframe src="/secofficeshow" name='main' id="iframepage" frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" onLoad="iFrameHeight()"></iframe> -->
        <!-- <iframe src="/secofficeshow" name='main' id="iframepage" frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" onload="changeFrameHeight()"></iframe> -->
       <iframe src="/admin/01" name='main' frameborder="0"  width="100%" scrolling="no" marginheight="0" marginwidth="0" id="iframepage" onload="this.height=100"></iframe> 
</div>  


<script type="text/javascript">
  function reinitIframe(){//http://caibaojian.com/frame-adjust-content-height.html
    var iframe = document.getElementById("iframepage");
    try{
      var bHeight = iframe.contentWindow.document.body.scrollHeight;
      var dHeight = iframe.contentWindow.document.documentElement.scrollHeight;
      var height = Math.max(bHeight, dHeight,800);
      iframe.height = height;
       // console.log(height);//这个显示老是在变化
     }catch (ex){
     } 
  } 
  window.setInterval("reinitIframe()", 200);
 </script>

</body>
</html>