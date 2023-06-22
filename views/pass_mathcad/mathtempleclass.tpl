<!-- 模板侧栏，右侧为模板列表 -->
<!DOCTYPE html>
<html>
<head>
  <meta name="renderer" content="webkit">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
      <!-- 收藏用logo图标 -->
  <link rel="bookmark"  type="image/x-icon"  href="/static/img/pss.ico"/>
  <!-- 网站显示页logo图标 -->
  <link rel="shortcut icon" href="/static/img/pss.ico">
  <!-- <meta http-equiv="X-UA-Compatible" content="IE=edge,Chrome=1" /> -->
  <!-- <meta http-equiv="X-UA-Compatible" content="IE=9" /> -->
  <!-- <meta name="renderer" content="webkit">
加上这句，360等浏览器就会默认使用google内核，而不是IE内核 。
因为你没加，所以我打开你的那个地址，默认就使用IE内核了。
<meta http-equiv="X-UA-Compatible" content="IE=edge">
加上这一句，如果被用户强行使用IE浏览器，就会使用IE的最高版本渲染内核。
这个也一定要加 -->
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <!-- <meta name="author" content="Jophy" /> -->
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <!-- <link rel="stylesheet" href="/static/css/style.css"> -->
  <!--[if lte IE 9]>兼容ie的方面：先引用bootstrapcss，再引用js
  <script src="bootstrap/js/respond.min.js"></script>
  <script src="bootstrap/js/html5.js"></script>
  <![endif]-->
  <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
  <!--[if lt IE 9]>
  <script src="http://apps.bdimg.com/libs/html5shiv/3.7/html5shiv.min.js"></script>
  <script src="http://apps.bdimg.com/libs/respond.js/1.4.2/respond.min.js"></script>
<![endif]-->
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <title>模板分类列表</title>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css" />
  <meta charset="utf-8">
  <!-- <link rel="stylesheet" href="https://unpkg.com/ionicons@4.5.5/dist/css/ionicons.min.css"> -->
  <link rel="stylesheet" href="/static/css/jquery.mCustomScrollbar.min.css">
  <link rel="stylesheet" href="/static/css/custom.css">
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script src="/static/js/jquery.mCustomScrollbar.concat.min.js"></script>
  <script src="/static/js/custom.js"></script>
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
        <iframe src="/v1/mathcad/getmath/templelist/{{.Id}}" name='iframepage' id="iframepage" frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" onload="this.height=800"></iframe>
      </div>
    </main>
  </div>
  <script type="text/javascript">
  // $('.btn').on('click', function() {
  //   $('.btn').toggleClass('btnc');
  //   $('.sidebar').toggleClass('side');
  // })

  $(function() {
    $('#tree').treeview({
      data: [{{.json }}],
      levels: 2,
      showTags: true,
      loadingIcon: "fa fa-minus",
      lazyLoad: loaddata,
    });
    $('#tree').on('nodeSelected', function(event, data) {
      document.getElementById("iframepage").src = "/v1/mathcad/getmath/templelist/" + data.id;
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
    $("#btn").click(function(e) {
      var arr = $('#tree').treeview('getSelected');
      for (var key in arr) {
        c.innerHTML = c.innerHTML + "," + arr[key].id;
      }
    });
  })

  function loaddata(node, func) {
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

  function gototree(e) {
    // alert(e)
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