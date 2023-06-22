<!-- 具体一个项目的侧栏，右侧为project_products.tpl,显示任意侧栏下的成果 -->
<!DOCTYPE html>
{{template "header"}}
<title>项目详细-EngiCMS</title>
<meta charset="utf-8">
<script src="/static/js/bootstrap-treeview.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
<script src="/static/js/tableExport.js"></script>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
<script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css" />
<!-- <link rel="stylesheet" href="https://unpkg.com/ionicons@4.5.5/dist/css/ionicons.min.css"> -->
<link rel="stylesheet" href="/static/css/jquery.mCustomScrollbar.min.css">
<link rel="stylesheet" href="/static/css/custom.css">
<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
<script src="/static/js/jquery.mCustomScrollbar.concat.min.js"></script>
<script src="/static/js/custom.js"></script>
<!-- <script type="text/javascript" src="/static/js/jquery-ui.min.js"></script> -->
<!-- <style type="text/css">
    #imgmodalDialog .modal-header {
      cursor: move;
    }
  </style> -->
<!-- <style type="text/css">
    @import 'https://unpkg.com/ionicons@4.5.5/dist/css/ionicons.min.css';
    
    body,
    html {
      margin: 0;
      padding: 0;
      height: 100%;
      width: 100%;
      font-family: sans-serif;
    }
    
    .sidebar {
      float: left;
      width: 100px;
      height: 100%;
      margin-left: -100px;
      background: #2c3e50;
      overflow: hidden;
      transition: 0.8s all;
    }
    
    .side {
      margin-left: 0;
    }
    
    .sidebar ul {
      margin: 0;
      padding: 0;
    }
    
    .sidebar ul li {
      list-style: none;
    }
    
    .sidebar ul li a {
      text-decoration: none;
      color: white;
      height: 80px;
      width: 100%;
      font-size: 40px;
      line-height: 80px;
      text-align: center;
      display: block;
      transition: 0.6s all;
    }
    
    .sidebar ul li a:hover {
      background: #34495e;
    }
    
    .btn {
      float: left;
      padding: 0 10px;
      font-size: 40px;
      text-decoration: none;
      color: #2c5e50;
      font-family: ionicons;
      cursor: pointer;
    }
    
    .btn:before {
      content: '\f32a';
    }
    
    .btnc:before {
      content: '\f2c0';
    }
  </style> -->
</head>
<div class="container-fill">{{template "navbar" .}}</div>

<body>
  <div class="page-wrapper toggled">
    <nav id="sidebar" class="sidebar-wrapper">
      <div class="sidebar-content mCustomScrollbar _mCS_1 mCS-autoHide desktop">
        <div id="mCSB_1" class="mCustomScrollBox mCS-light mCSB_vertical mCSB_inside" tabindex="0" style="max-height: none;">
          <div id="mCSB_1_container" class="mCSB_container" dir="ltr">
            <a href="javascript:void(0)" id="toggle-sidebar"> <i class="fa fa-bars"></i>
            </a>
            <div class="sidebar-brand">
              <a href="javascript:void(0)">pro sidebar</a>
            </div>
            <div class="sidebar-menu">
              <ul id="tree"></ul>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <main class="page-content">
      <div class="breadcrumbs">
        <ol class="breadcrumb" style="margin-bottom: 2px;" split="&gt;">
          <li>
            <a href="javascript:gototree({{.Category.Id}})"> <i class="fa fa-home" aria-hidden="true"></i>
              项目编号：{{.Category.Code}}
            </a>
          </li>
        </ol>
      </div>
      <div class="container-fluid">
        <iframe src="/project/{{.Id}}/{{.Id}}" name='iframepage' id="iframepage" frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" onload="this.height=800"></iframe>
      </div>
    </main>
  </div>
  <script type="text/javascript">
  // $('.btn').on('click', function() {
  //   $('.btn').toggleClass('btnc');
  //   $('.sidebar').toggleClass('side');
  // })

  // function onStorageChange(e) {
  //   console.log(e.key);
  // }
  // window.addEventListener('storage', onStorageChange);
  // window.addEventListener('storage',function (e) {
  //   console.log(e.newValue);
  //   window.location
  // })

  //当前页面监听localstorage值得变化
var orignalSetItem = localStorage.setItem;
localStorage.setItem = function(key, newValue) {
  var setItemEvent = new Event("setItemEvent");
  setItemEvent.key = key;
  setItemEvent.newValue = newValue;
  window.dispatchEvent(setItemEvent);
  orignalSetItem.apply(this, arguments);
};

window.addEventListener("setItemEvent", function(e) {
  if (e.key == 'projectid') {
    var _this = localStorage.getItem("projectid")
    if (_this != e.newValue) {
      alert('key值为projectid的值发生改变,之前是' + _this + '当前为' + e.newValue)
      window.open("/project/"+e.newValue, "_self" )
    } else {
      alert('key值为projectid值' + e.newValue);
    }
  }
});

var orignalremoveItem = localStorage.removeItem;
localStorage.removeItem = function(key, newValue) {
  var removeItemEvent = new Event("removeItemEvent");
  removeItemEvent.key = key;
  removeItemEvent.newValue = newValue;
  window.dispatchEvent(removeItemEvent);
  orignalremoveItem.apply(this, arguments);
};

window.addEventListener("removeItemEvent", function(e) {
  if (localStorage.getItem("projectid")) {
    if (e.key == 'projectid') {
      alert("key值为demo，删除成功");
    }
  } else {
    alert("本地数据无key值为demo")
  }
});

  $(function() {
    $('#tree').treeview({
      data: [{{.json }}],
      levels: 2,
      showTags: true,
      loadingIcon: "fa fa-minus",
      lazyLoad: loaddata,
      emptyIcon:""//没有子节点的节点图标
    });

    $('#tree').on('nodeSelected', function(event, data) {
      document.getElementById("iframepage").src = "/project/{{.Id}}/" + data.id;
      $.ajax({
        type: "get",
        url: "/project/navbar/" + data.id,
        success: function(data, status) {
          $(".breadcrumb #nav").remove();
          for (i = 0; i < data.length; i++) {
            $(".breadcrumb").append('<li id="nav"><a href="javascript:gototree(' + data[i].Id + ')">' + data[i].Title + '</a></li>');
          }
        }
      });
    });
  })

  $("#btn").click(function(e) {
    var arr = $('#tree').treeview('getSelected');
    for (var key in arr) {
      c.innerHTML = c.innerHTML + "," + arr[key].id;
    }
  });

  function loaddata(node, func) {//这个技巧真高，即能返回参数，又能把参数通过函数发回去
    $.ajax({
      type: "get",
      url: "/project/getprojcate",
      data: { id: node.id },
      success: function(data, status) {
        if (data) {
          func(data);
        }
      }
    });
  }
  // 20220407添加跳转
  $(function() {
    // alert({{.Gototree}})
    // alert({{.Node}})
    if ({{.Gototree }}) {
      $.ajax({
        type: "get",
        url: "/project/navbar/" + {{.Node }},
        success: function(data, status) {
          $(".breadcrumb #nav").remove();
          for (i = 0; i < data.length; i++) {
            $(".breadcrumb").append('<li id="nav"><a href="javascript:gototree(' + data[i].Id + ')">' + data[i].Title + '</a></li>');
          }
        }
      });
      gototree({{.Node }})
    }
  })

  function gototree(e) {
    document.getElementById("iframepage").src = "/project/{{.Id}}/" + e;
    var findCheckableNodess = function() {
      return $('#tree').treeview('findNodes', [e, 'id']);
    };
    var checkableNodes = findCheckableNodess();
    $('#tree').treeview('toggleNodeSelected', [checkableNodes, { silent: true }]);
    $('#tree').treeview('toggleNodeExpanded', [checkableNodes, { silent: true }]);
    $('#tree').treeview('revealNode', [checkableNodes, { silent: true }]);
  }

  function reinitIframe() {
    var iframe = document.getElementById("iframepage");
    try {
      var bHeight = iframe.contentWindow.document.body.scrollHeight;
      var dHeight = iframe.contentWindow.document.documentElement.scrollHeight;
      var height = Math.max(bHeight, dHeight, 800);
      iframe.height = height;
    } catch (ex) {}
  }
  window.setInterval("reinitIframe()", 200);

  // $(document).ready(function() {
  //   $("#imgmodalDialog").draggable({ handle: ".modal-header" });
  // })
  </script>
</body>

</html>