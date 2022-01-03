<!-- iframe里系统基本信息日志页-->
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
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/> -->
  
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script> -->
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script src="/static/js/tableExport.js"></script>
</head>
<body>

<script type="text/javascript">
  function index1(value,row,index){
    return index+1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }
</script>

<div class="col-lg-12">
<h3>基本日志信息</h3>
<div id="toolbar1" class="btn-group">
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button> 
        <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
</div>

<table id="table0"
        data-toggle="table"
        data-url="/v1/adminlog/infolog"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="ProjectName"
        data-sort-order="desc"
        data-page-size="15"
        data-page-list="[15, 25, 50, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client"
        data-click-to-select="true"
        >
    <thead>        
      <tr>
        <!-- radiobox data-checkbox="true"-->
        <th data-width="10" data-checkbox="true"></th>
        <th data-formatter="index1">#</th>
        <th data-field="datetime" data-align="center" data-valign="middle">时间</th>
        <th data-field="tag" data-align="center" data-valign="middle">类别</th>
        <th data-field="file" data-align="center" data-valign="middle">文件位置</th>
        <th data-field="user" data-align="center" data-valign="middle">用户</th>
        <th data-field="action" data-align="center" data-valign="middle">动作</th>
        <th data-field="url" data-align="center" data-valign="middle">资源</th>
      </tr>
    </thead>
</table>

</div>

</body>
</html>