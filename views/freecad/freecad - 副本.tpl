<!DOCTYPE html>
<html>

<head>
  <title>Bootstrap Tree View</title>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <link href="/static/css/bootstrap-treeview.css" rel="stylesheet">
  <link rel="stylesheet" href="/static/css/jquery.mCustomScrollbar.min.css">
  <link rel="stylesheet" href="/static/css/custom.css">
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <style type="text/css">
  .page-wrapper .sidebar-wrapper {
    position: fixed;
    /*top: 50px;*/
    left: 5;
    /*bottom: 0;*/
    /*width: 200px;*/
    /*background-color: #23262E;*/
  }

  .page-wrapper .container-fluid {
    /*position: fixed;*/
    /*top: 50px;*/
    right: 0;
    /*bottom: 0;*/
    left: 200px;
    /*background-color: #F8F9FA;*/
    overflow: hidden;
  }
  </style>
</head>

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
              <div id="form"></div>
              <ul id="tree"></ul>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <main class="page-content">
      <!--       <div class="breadcrumbs">
        <ol class="breadcrumb" style="margin-bottom: 2px;" split="&gt;">
          <li>
            <a href="javascript:gototree({{.Category.Id}})"> <i class="fa fa-home" aria-hidden="true"></i>
              项目编号：{{.Category.Code}}
            </a>
          </li>
        </ol>
      </div> -->
      <div class="container-fluid">
        <iframe src="/v1/freecad/getfreecad/{{.Id}}" name='iframepage' id="iframepage" frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" onload="this.height=800"></iframe>
      </div>
    </main>
  </div>
  <!--   <div class="container">
    <h1>Bootstrap Tree View - DOM Tree</h1>
    <br />
    <div class="row">
      <div class="col-sm-12">
        <label for="treeview"></label>
        <div id="tree" />
      </div>
    </div>
  </div> -->
  <!-- 省市区地域查询展示 -->
  <!--   <div class="container">
    <div class="row">
      <div class="col-sm-4">
        <h4>快捷搜索</h4>
        <div class="form-group">
          <label for="input-search" class="sr-only">快捷搜索:</label>
          <input type="input" class="form-control" id="input-search" placeholder="请输入要查询的省市区名称信息..." value="">
        </div>
        <button type="button" class="btn btn-success" id="btn-search">搜索</button>
        <button type="button" class="btn btn-default" id="btn-clear-search">清除</button>
      </div>
      <div class="col-sm-4">
        <h4>省市区名称层级树</h4>
        <div id="treeview-searchable" class=""></div>
      </div>
      <div class="col-sm-4">
        <h4>查询结果展示</h4>
        <div id="search-output"></div>
      </div>
    </div>
  </div> -->
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <script src="/static/js/jquery.mCustomScrollbar.concat.min.js"></script>
  <script src="/static/js/custom.js"></script>
  <script type="text/javascript">
  $(function() {
    $('#tree').treeview({
      data: {{.json }},// 数据结构供bootstrap treeview用
      // data: getTree(),// 数据结构供antd tree测试用
      levels: 1,
      showTags: true,
      loadingIcon: "fa fa-minus",
      lazyLoad: loaddata,
    });

    $('#tree').on('nodeSelected', function(event, data) {
      document.getElementById("iframepage").src = "/v1/freecad/getfreecad/" + data.id;
    });
  })

  function getTree() {
    // Some logic to retrieve, or generate tree structure
    $.ajax({
      type: "get",
      url: "/v1/freecad/freecadmenu",
      success: function(data, status) {
        return data;
      }
    });
  }
  

  function loaddata(node, func) {
    $.ajax({
      type: "get",
      url: "/v1/freecad/freecadlist",
      data: { id: node.id },
      success: function(data, status) {
        if (data) {
          func(data);
        }
      }
    });
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




  // function lazyLoad(data, func) {
  //   func([{
  //     text: "Parent 2-1",
  //     lazyLoad: true
  //   }, ])
  //   console.log(data, func)
  // }

  // $(function() {
  //   var options = {
  //     bootstrap2: false,
  //     showTags: true,
  //     levels: 5,
  //     data: [{
  //         icon: "glyphicon glyphicon-stop",
  //         text: "Parent 2",
  //         lazyLoad: true
  //       },
  //       {
  //         text: "Parent 3"
  //       },
  //     ],
  //     lazyLoad: function(data, func) {
  //       lazyLoad(data, func);
  //     }
  //   };

  //   $('#tree').treeview(options);
  //   $('#tree').on('nodeSelected', function(event, data) {
  //     console.log(data);
  //   });
  // });
  </script>
  <script type="text/javascript">
  $(function() {
    var defaultData;
    // $.ajax({
    //   type: "post",
    //   url: "",
    //   dataType: "json",
    //   success: function(result) {
    //     defaultData = result;
    //     var initSearchableTree = function() {
    //       return $('#treeview-searchable').treeview({
    //         data: defaultData,
    //         nodeIcon: 'glyphicon glyphicon-globe',
    //         emptyIcon: '', //没有子节点的节点图标
    //         //collapsed: true,
    //       });
    //     };
    //     var $searchableTree = initSearchableTree();
    //     $('#treeview-searchable').treeview('collapseAll', {
    //       silent: false //设置初始化节点关闭
    //     });
    //     var findSearchableNodes = function() {
    //       return $searchableTree.treeview('search', [$.trim($('#input-search').val()), { ignoreCase: false, exactMatch: false }]);
    //     };
    //     var searchableNodes = findSearchableNodes();
    //     // Select/unselect/toggle nodes
    //     $('#input-search').on('keyup', function(e) {
    //       var str = $('#input-search').val();
    //       if ($.trim(str).length > 0) {
    //         searchableNodes = findSearchableNodes();
    //       } else {
    //         $('#treeview-searchable').treeview('collapseAll', {
    //           silent: false //设置初始化节点关闭
    //         });
    //       }
    //       //$('.select-node').prop('disabled', !(searchableNodes.length >= 1));
    //     });
    //     var search = function(e) {
    //       var pattern = $.trim($('#input-search').val());
    //       var options = {
    //         ignoreCase: $('#chk-ignore-case').is(':checked'),
    //         exactMatch: $('#chk-exact-match').is(':checked'),
    //         revealResults: $('#chk-reveal-results').is(':checked')
    //       };
    //       var results = $searchableTree.treeview('search', [pattern, options]);

    //       var output = '<p>' + results.length + ' 匹配的搜索结果</p>';
    //       $.each(results, function(index, result) {
    //         output += '<p>- <span style="color:red;">' + result.text + '</span></p>';
    //       });
    //       $('#search-output').html(output);
    //     }
    //     $('#btn-search').on('click', search);
    //     $('#input-search').on('keyup', search);
    //     $('#btn-clear-search').on('click', function(e) {
    //       $searchableTree.treeview('clearSearch');
    //       $('#input-search').val('');
    //       $('#search-output').html('');
    //       $('#treeview-searchable').treeview('collapseAll', {
    //         silent: false //设置初始化节点关闭
    //       });
    //     });

    //   },
    //   error: function() {
    //     alert("省市区地域信息加载失败！")
    //   }
    // });
  });
  </script>
</body>