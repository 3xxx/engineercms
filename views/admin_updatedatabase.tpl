<!-- iframe里merit基本信息-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
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
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <!-- <script src="/static/js/jquery.form.js"></script> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/select2.min.css"/> -->
  <!-- <script type="text/javascript" src="/static/js/select2.js"></script> -->
</head>
<body>

<div class="col-lg-12">
<h3>请先备份数据库。——2018年6月30日之前的需要升级数据库，2018年6月30日之后的数据库不用升级。</h3>
<h4>等待时间视数据库数据多少，请耐心等待……直到出现弹框。</h4>
<h4>数据表Project：对parentidpath字段进行了重写，</h4>
<h4>数据表Product：增加TopProjectId字段,删掉content和views字段</h4>

  <!-- <div id="toolbar1" class="btn-group"> -->
  <button type="button" data-name="updateButton" id="updateButton" class="btn btn-default"> <i class="fa fa-plus">点击升级</i>
  </button>
  <!-- <button type="button" data-name="modifyButton" id="modifyButton" class="btn btn-default"> <i class="fa fa-plus">test</i>
  </button> -->
</div>
<script type="text/javascript">
  $(document).ready(function() {
    $("#updateButton").click(function() {
      $.ajax({
        type:"post",
        url:"/updatedatabase",
        success:function(data,status){
          alert("升级“"+data+"”成功！(status:"+status+".)");
        }
      });
    })
  })

  $(document).ready(function() {
    $("#modifyButton").click(function() {
      $.ajax({
        type:"post",
        url:"/modifydatabase",
        success:function(data,status){
          alert("修改“"+data+"”成功！(status:"+status+".)");
        }
      });
    })
  })
</script>
</body>
</html>