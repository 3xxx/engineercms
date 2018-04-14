<!DOCTYPE html>
<html>
  <head>
  <meta charset="UTF-8">
  <title>角色——oo权限</title>

  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="description" content="A front-end template that helps you build fast, modern mobile web apps.">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0">
    
  <meta name="mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-status-bar-style" content="black">
  <meta name="apple-mobile-web-app-title" content="Material Design Lite">
    
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>

  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css"/>
  <script src="/static/js/bootstrap-treeview.js"></script>
  </head>

<body>

  <div class="col-sm-12 col-md-12 col-lg-12"> 
    
    <!-- 项目表 -->
    <!-- <div class="col-sm-6 col-md-6 col-lg-6">
      <h3>项目列表</h3>
      <table id="table2"  
        data-url="/project/getprojects"
        data-toggle="table"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-click-to-select="true"
        data-side-pagination="client"
        data-pagination="true"
        data-page-size="5"
        data-page-list="[5, 10, All]"
        data-query-params="queryParams"
        >
        <thead>        
          <tr>
          <th data-width="10" data-radio="true" data-select-item-name="proj"></th>
          <th data-formatter="index1">#</th>
          <th data-field="Code">编号</th>
          <th data-field="Title">名称</th>
          <th data-field="Principal">负责人</th>
          <th data-field="Created" data-formatter="localDateFormatter" data-visible="false">建立时间</th>
          </tr>
        </thead>
      </table>
    </div> -->

    <h3>文档列表</h3>
    <div id="toolbar1" class="btn-group">
        <!-- 多文件批量上传 -->
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default" title="批量上传模式"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="editorProdButton" id="editorProdButton" class="btn btn-default"> <i class="fa fa-edit" title="修改成果信息">编辑</i>
        </button>
        <button type="button" data-name="editorAttachButton" id="editorAttachButton" class="btn btn-default"> <i class="fa fa-edit" title="修改成果附件">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
    </div>
    <!--data-click-to-select="true" -->
    <table id="table2" 
        data-toggle="table" 
        data-url="/onlyoffice/data"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="Code"
        data-sort-order="desc"
        data-page-size="15"
        data-page-list="[10,15, 50, 100, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client"
        data-single-select="true"
        data-click-to-select="true"
        data-show-export="true"
        >
      <thead>        
        <tr>
        <!-- radiobox data-checkbox="true" data-formatter="setCode" data-formatter="setTitle"-->
        <th data-width="10" data-radio="true"></th>
        <th data-formatter="index1" data-align="center">#</th>
        <th data-field="Code" data-halign="center">编号</th>
        <th data-field="Title" data-halign="center">名称</th>
        <th data-field="Label" data-formatter="setLable" data-halign="center" data-align="center">关键字</th>
        <th data-field="Principal" data-halign="center" data-align="center">负责人</th>
        <th data-field="Docxlink" data-formatter="setDocx" data-events="actionEvents" data-halign="center" data-align="center">协作</th>
        <!-- <th data-field="Xlsxlink" data-formatter="setXlsx" data-events="actionEvents" data-halign="center" data-align="center">XLSX</th> -->
        <!-- <th data-field="Pptxlink" data-formatter="setPptx" data-events="actionEvents" data-halign="center" data-align="center">PPTX</th> -->
        <th data-field="End" data-formatter="localDateFormatter" data-halign="center" data-align="center">结束时间</th>
        <th data-field="Created" data-formatter="localDateFormatter" data-halign="center" data-visible="false" data-align="center">建立时间</th>
        <th data-field="Updated" data-formatter="localDateFormatter" data-halign="center" data-align="center">更新时间</th>
        </tr>
      </thead>
    </table>

  </div>

  <div class="col-sm-12 col-md-12 col-lg-12">
    <!-- 角色表 -->
    <div id="h-role-info" class="col-sm-4 col-md-4 col-lg-4">

      <h3>Full Access角色表</h3>
      <div id="toolbar2" class="btn-group">
        <button type="button" id="editorButton2" class="btn btn btn-primary btn-sm"> <i class="fa fa-edit">保存修改</i>
        </button>
      </div>
      <table id="table"
        data-url="/admin/role/get/"
        data-toggle="table"
        data-striped="true"
        data-toolbar="#toolbar2"
        data-search="true"
        data-side-pagination="client"
        data-pagination="true"
        data-click-to-select="true"
        data-page-size="5"
        data-page-list="[5, 25, 50, All]"
        data-search="false"
        >
        <thead>
          <tr><!-- data-radio -->
            <th data-field="state" data-checkbox="true" data-select-item-name="role"></th>
            <th data-formatter="index1">#</th>
            <th data-field="Rolenumber">角色编码</th>
            <th data-field="Rolename">角色名称</th>
            <th data-align="center" data-formatter="StatusFormatter">状态</th>
            <th data-align="center" data-field="Createtime" data-formatter="localDateFormatter" data-visible="false">创建时间</th>
            <th data-align="center" data-field="Updated" data-formatter="localDateFormatter" data-visible="false">修改时间</th>
          </tr>
        </thead>
      </table>

    </div>
    <!-- 角色表 -->
    <div id="h-role-info" class="col-sm-4 col-md-4 col-lg-4">

      <h3>Read Only角色表</h3>
      <div id="toolbar3" class="btn-group">
        <button type="button" id="editorButton3" class="btn btn btn-primary btn-sm"> <i class="fa fa-edit">保存修改</i>
        </button>
      </div>
      <table id="table"
        data-url="/admin/role/get/"
        data-toggle="table"
        data-striped="true"
        data-toolbar="#toolbar3"
        data-search="true"
        data-side-pagination="client"
        data-pagination="true"
        data-click-to-select="true"
        data-page-size="5"
        data-page-list="[5, 25, 50, All]"
        data-search="false"
        >
        <thead>
          <tr><!-- data-radio -->
                <th data-field="state" data-checkbox="true" data-select-item-name="role"></th>
                <th data-formatter="index1">#</th>
                <th data-field="Rolenumber">角色编码</th>
                <th data-field="Rolename">角色名称</th>
                <th data-align="center" data-formatter="StatusFormatter">状态</th>
                <th data-align="center" data-field="Createtime" data-formatter="localDateFormatter" data-visible="false">创建时间</th>
                <th data-align="center" data-field="Updated" data-formatter="localDateFormatter" data-visible="false">修改时间</th>
          </tr>
        </thead>
      </table>

    </div>
    <!-- 角色表 -->
    <div id="h-role-info" class="col-sm-4 col-md-4 col-lg-4">

      <h3>Deny Access角色表</h3>
      <div id="toolbar4" class="btn-group">
        <button type="button" id="editorButton" class="btn btn btn-primary btn-sm"> <i class="fa fa-edit">保存修改</i>
        </button>
      </div>
      <table id="table"
        data-url="/admin/role/get/"
        data-toggle="table"
        data-striped="true"
        data-toolbar="#toolbar4"
        data-search="true"
        data-side-pagination="client"
        data-pagination="true"
        data-click-to-select="true"
        data-page-size="5"
        data-page-list="[5, 25, 50, All]"
        data-search="false"
        >
        <thead>
          <tr><!-- data-radio -->
            <th data-field="state" data-checkbox="true" data-select-item-name="role"></th>

            <th data-formatter="index1">#</th>

            <th data-field="Rolenumber">角色编码</th>
            <th data-field="Rolename">角色名称</th>
            <th data-align="center" data-formatter="StatusFormatter">状态</th>
            <th data-align="center" data-field="Createtime" data-formatter="localDateFormatter" data-visible="false">创建时间</th>
            <th data-align="center" data-field="Updated" data-formatter="localDateFormatter" data-visible="false">修改时间</th>
          </tr>
        </thead>
      </table>

    </div>


  </div> 

  <script>
    /*数据json,"PDF":"","DWG":"","DOC":"","XLS":""*/
    // 权限表
    var json =  [{"Id":"1","Title":"添加成果","Action":"POST"},
                {"Id":"2","Title":"编辑成果","Action":"PUT"},
                {"Id":"3","Title":"删除成果","Action":"DELETE"},
                {"Id":"4","Title":"读取成果","Action":"GET"}];
    var json1 = [{"Id":"5","Title":"任意"},
                {"Id":"6","Title":"PDF","checked":true},
                {"Id":"7","Title":"DWG"},
                {"Id":"8","Title":"DOC"},
                {"Id":"8","Title":"XLS"},
                {"Id":"8","Title":"DGN"}];              
    /*初始化table数据*/
    $(function(){
      $("#table3").bootstrapTable({
              data:json,
              // onClickRow: function (row, $element) {
                // alert( "选择了行Id为: " + row.Id );
                // rowid=row.Id//全局变量
                // $('#table1').bootstrapTable('refresh', {url:'/admincategory?pid='+row.Id});
              // }
              // onExpandRow: function (index, row, $Subdetail) {
              //   oInit.InitSubTable(index, row, $Subdetail);
              // }
              onExpandRow: function (index, row, $detail) {
                expandTable(index, row,$detail);
              }
      });
    });

    function expandTable(index, row,$detail) {
        var cur_table = $detail.html('<table id="table4"></table>').find('table');
        if (index==3){
          $(cur_table).bootstrapTable({
            columns: [{
                  checkbox: true,
                  formatter: "stateFormatter"
              }, {
                  field: 'Title',
                  title: '文件扩展名'
              }],
            data:json1,
          })
        }
    }
    //初始化子表格(无限循环)
    // oInit.InitSubTable = function (index, row, $detail) {
    //     var parentid = row.MENU_ID;
    //     var cur_table = $detail.html('<table></table>').find('table');
    //     $(cur_table).bootstrapTable({
    //         url: '/api/MenuApi/GetChildrenMenu',
    //         method: 'get',
    //         queryParams: { strParentID: parentid },
    //         ajaxOptions: { strParentID: parentid },
    //         clickToSelect: true,
    //         detailView: true,//父子表
    //         uniqueId: "MENU_ID",
    //         pageSize: 10,
    //         pageList: [10, 25],
    //         columns: [{
    //             checkbox: true
    //         }, {
    //             field: 'MENU_NAME',
    //             title: '菜单名称'
    //         }, {
    //             field: 'MENU_URL',
    //             title: '菜单URL'
    //         }, {
    //             field: 'PARENT_ID',
    //             title: '父级菜单'
    //         }, {
    //             field: 'MENU_LEVEL',
    //             title: '菜单级别'
    //         }, ],
    //         //无线循环取子表，直到子表里面没有记录
    //         onExpandRow: function (index, row, $Subdetail) {
    //             oInit.InitSubTable(index, row, $Subdetail);
    //         }
    //     });
    // };

    function index1(value,row,index){
      return index+1
    }

    function stateFormatter(value, row, index) {
      if (row.checked === true) {
          return {
              // disabled: true,
              checked: true
          }
      }
      return value;
    }

    function StatusFormatter(value, row, index) {
      // alert(row.Status);
      if (row.Status == "0") {
          return '正常';
      }else{
        return '失效';
      }
    }

    function localDateFormatter(value) {
      return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
    }
    function checkDateFormatter(value) {
      return '<input type="checkbox" name="bike"/>';
    }

    //oo文档列表
    function setDocx(value,row,index){
      if (value){
      if (value.length==1){
        if (value[0].Suffix=="docx"){
          docUrl= '<a href=/onlyoffice/'+value[0].Id+' title="协作" target="_blank"><i class="fa fa-file-word-o fa-lg"></i></a>';
          return docUrl;
        }else if(value[0].Suffix=="xlsx"){
          xlsUrl= '<a href=/onlyoffice/'+value[0].Id+' title="协作" target="_blank"><i class="fa fa-file-excel-o fa-lg" style="color:LimeGreen;"></i></a>';
          return xlsUrl;
        }else if(value[0].Suffix=="pptx"){
          pptUrl= '<a href=/onlyoffice/'+value[0].Id+' title="协作" target="_blank"><i class="fa fa-file-powerpoint-o fa-lg" style="color:Red;"></i></a>';
          return pptUrl;
        }
        
      }else if(value.length==0){
                    
      }else if(value.length>1){
        fileUrl= "<a class='Docx' href='javascript:void(0)' title='查看文档列表'><i class='fa fa-list-ol'></i></a>";
        return fileUrl;
      }
      }
    }

    //关键字
    function setLable(value,row,index){
      // alert(value);
      if (value){//注意这里如果value未定义则出错，一定要加这个判断。
        var array=value.split(",")
        var labelarray = new Array() 
        for (i=0;i<array.length;i++)
        {
          labelarray[i]="<a href='/project/product/keysearch?keyword="+array[i]+"'>" + array[i] + "</a>";
        }
        return labelarray.join(",");
      }
    }

    //角色的添加删除和编辑
    // $(document).ready(function() {
    //   //添加角色
    //   $("#addButton").click(function() {
    //     $('#modalTable').modal({
    //     show:true,
    //     backdrop:'static'
    //     });
    //   })
    //   //删除角色
    //   $("#deleteButton").click(function() {
    //     var selectRow=$('#table').bootstrapTable('getSelections');
    //     if (selectRow.length<=0) {
    //       alert("请先勾选！");
    //       return false;
    //     }
    //     var titles=$.map(selectRow,function(row){
    //       // alert(row.Id);
    //       // return row.Id;
    //       // alert(row.Title);
    //       return row.Title;
    //     })
    //     var ids="";
    //     for(var i=0;i<selectRow.length;i++){
    //       if(i==0){
    //         ids=selectRow[i].Id;
    //       }else{
    //         ids=ids+","+selectRow[i].Id;
    //       }  
    //     }
    //     $.ajax({
    //    type:"post",
    //    url:"/admin/role/delete",
    //    data: {ids:ids},
    //    success:function(data,status){
    //      alert("删除“"+data+"”成功！(status:"+status+".)");
    //      //删除已选数据
    //      $('#table1').bootstrapTable('remove',{
    //        field:'Title',
    //        values:titles
    //      });
    //    }
    //     }); 
    //   })
    //   //编辑角色模态框
    //   $("#editorButton").click(function() {
    //     var selectRow=$('#table').bootstrapTable('getSelections');
    //     if (selectRow.length<1){
    //       alert("请先勾选角色！");
    //       return;
    //     }
    //     if (selectRow.length>1){
    //       alert("请不要勾选一个以上角色！");
    //       return;
    //     }
    //     $("input#cid").remove();
    //     var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
    //     $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
    //     $("#Rolename1").val(selectRow[0].Rolename);
    //     $("#Rolenumber1").val(selectRow[0].Rolenumber);
    //     $("#Status1").val(selectRow[0].Status);
    //     // alert(JSON.stringify(selectRow));
    //     // alert(selectRow[0].Id);
    //     // var title = $('#'+id).attr("value");
    //     // var title = $('#'+id).attr("href");
    //     // var categoryid = $('#categoryid').val();
    //     $('#modalTable1').modal({
    //     show:true,
    //     backdrop:'static'
    //     });
    //   })
    // })

    //新建角色
    // function save(){
    //   // var radio =$("input[type='radio']:checked").val();        
    //   var Rolename = $('#Rolename').val();
    //   var Rolenumber = $('#Rolenumber').val();
    //   // var Status     = $('#Status option:selected').text();
    //   var Status = $('#Status option:selected').val();

    //   if (Rolename){  
    //     $.ajax({
    //             type:"post",
    //             url:"/admin/role/post",
    //             data: {rolename:Rolename,rolenumber:Rolenumber,status:Status},
    //             success:function(data,status){
    //               alert("添加“"+data+"”成功！(status:"+status+".)");
    //              }
    //     });  
    //   } else{
    //     alert("角色名等不能为空！");
    //   }
    //   // $(function(){$('#myModal').modal('hide')}); 
    //   $('#modalTable').modal('hide');
    //   $('#table').bootstrapTable('refresh', {url:'/admin/role/get'});
    //   // "/category/modifyfrm?cid="+cid
    //   // window.location.reload();//刷新页面
    // }

    //更新角色
    // function update(){
    //   // var radio =$("input[type='radio']:checked").val();
    //   var roleid1 = $('#cid').val();
    //   var rolename1 = $('#Rolename1').val();
    //   var rolenumber1 = $('#Rolenumber1').val();
    //   var status1 = $('#Status1').val();
    //   // $('#myModal').on('hide.bs.modal', function () {  
    //   if (rolename1){  
    //       $.ajax({
    //           type:"put",
    //           url:"/admin/role/update",
    //           data: {roleid:roleid1,rolename:rolename1,rolenumber:rolenumber1,status:status1},
    //           success:function(data,status){
    //             alert("修改“"+data+"”成功！(status:"+status+".)");
    //            }
    //       });  
    //   } 
    //   // $(function(){$('#myModal').modal('hide')});
    //   $('#modalTable1').modal('hide');
    //   $('#table').bootstrapTable('refresh', {url:'/admin/role/get/'});
    //   // "/category/modifyfrm?cid="+cid
    //   // window.location.reload();//刷新页面
    // }

    //点击项目显示目录
    $(function(){
      $("#table2").on("check.bs.table",function(e,row,ele){
        $(".info").removeClass("info");
        $(ele).addClass("info");
        rowid=row.Id;//全局变量
        rowtitle=row.Title
        $("#rowtitle").html("项目目录-"+rowtitle);
         
        // $('#iptable').bootstrapTable('refresh', {url:'/admin/project/synchip/'+row.Id});
        $("input#cid").remove();
          var th1="<input id='cid' type='hidden' name='cid' value='" +rowid+"'/>"
        $(".modal-body").append(th1);
        //初始化树   
        $.ajax({  //JQuery的Ajax  
            type: 'POST',    
            dataType : "json",//返回数据类型  
            // async:false, //同步会出现警告：Synchronous XMLHttpRequest on the main thread is deprecated because of its detrimental effects to the end user's experience 
            url: "/admin/project/getprojectcate/"+rowid,//请求的action路径  
            // data: {"flag":true},  //同步请求将锁住浏览器，用户其它操作必须等待请求完成才可以执行  
            error: function () {//请求失败处理函数    
                alert('请求失败');    
            },  
            success:function(data){ //请求成功后处理函数。取到Json对象data
              // alert(data);
              $('#tree').treeview({
                showCheckbox: true,
                data: [data], // data is not optional
                levels: 3,
                enableLinks: true,
                hierarchicalCheck:true,//有效！！
                // propagateCheckEvent:true,
                highlightSearchResults:false,//搜索结果不高亮
                state: {
                  checked: true,
                  // disabled: true,
                  expanded: true,
                  // selected: true
                }
              });
              // $('#tree').treeview('expandAll');//这句为何无效？
              // var singleNode = {
                // text: '设计单位',
                // id:'0.0.0'
                // code:projcatecode3
              // };
              // $('#tree').treeview('checkNode', [ singleNode, { silent: true } ]);
              // $('#tree').treeview('collapseNode', [ singleNode, { silent: true, ignoreChildren: false } ]);
              // $('#tree').treeview('checkAll', { silent: true });
              // $('#tree').treeview('disableNode', [ singleNode, { silent: true, keepState: true } ]);
              $("#details").show();  
            }
        });
      });
    });

    // 保存权限和修改权限
    $("#editorButton3").click(function() {
      // if ({{.role}}!=1){
      //   alert("权限不够！");
      //   return;
      // }
      var selectRow=$('#table').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选角色！");
        return false;
      }
      // var title=$.map(selectRow,function(row){
      //   return row.Title;
      // })
      var roleids="";
      for(var i=0;i<selectRow.length;i++){
        if(i==0){
          roleids=selectRow[i].Id;
        }else{
          roleids=roleids+","+selectRow[i].Id;
        }  
      }

      var selectRow=$('#table3').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选权限！");
        return false;
      }
      // if(confirm("确定删除吗？一旦删除将无法恢复！")){
      // var title=$.map(selectRow,function(row){
      //   return row.Title;
      // })
      var permissionids="";
      for(var i=0;i<selectRow.length;i++){
        if(i==0){
          permissionids=selectRow[i].Title;
        }else{
          permissionids=permissionids+","+selectRow[i].Title;
        }  
      }

      var selectRow=$('#table4').bootstrapTable('getSelections');
      // if (selectRow.length<=0) {
      //   alert("请先勾选权限！");
      //   return false;
      // }
      var sufids="";
      for(var i=0;i<selectRow.length;i++){
        if(i==0){
          sufids=selectRow[i].Title;
        }else{
          sufids=sufids+","+selectRow[i].Title;
        }  
      }

      arr=$('#tree').treeview('getChecked');
        if (arr.length==0){
          alert("请先勾选目录！");
          return;
        }
        // alert(JSON.stringify(arr));
        // if (arr.length>=2){
        //   alert("请不要勾选一个以上！");
        //   return;
        // }
      var treeids="";
      var treenodeids="";
      for(var i=0;i<arr.length;i++){
        if(i==0){
          treeids=arr[i].id;
          treenodeids=arr[i].nodeId;
        }else{
          treeids=treeids+","+arr[i].id;
          treenodeids=treenodeids+","+arr[i].nodeId;
        }  
      }

      $.ajax({
        type:"post",
        url:"/admin/role/permission",
        data: {roleids:roleids,permissionids:permissionids,treeids:treeids,treenodeids:treenodeids,sufids:sufids},
        success:function(data,status){
          alert("添加“"+data+"”成功！(status:"+status+".)");
          // $('#table').bootstrapTable('refresh', {url:'/admin/role/get/'});
        }
      });  
    })
  </script>

  <!-- 显示角色的权限 -->
  <script type="text/javascript">
    // 每次点击角色表table、权限表table2、项目表table3，3个表任何一个点击，都检查是否具备查询条件
    $(function(){
      $("#table").on("check.bs.table",function(e,row,ele){
        //检查table2项目表和table3权限表是否有选择，并且只能单选
        var selectRow2=$('#table2').bootstrapTable('getSelections');
        var selectRow3=$('#table3').bootstrapTable('getSelections');
        if (selectRow2.length>1) {
          // alert("请不要勾选一个以上权限！");
          return false;
        }else if(selectRow2.length==1 && selectRow2.length==1){
          var projectid="";
          for(var i=0;i<selectRow2.length;i++){
            if(i==0){
              projectid=selectRow2[i].Id;
            }else{
              projectid=projectid+","+selectRow2[i].Id;
            }  
          }
          var action="";
          for(var i=0;i<selectRow3.length;i++){
            if(i==0){
              action=selectRow3[i].Action;
            }else{
              action=action+","+selectRow3[i].Action;
            }  
          }
          // alert(JSON.stringify(checkableNodes));
          // $('#btn-check-node.check-node').on('click', function (e) {
          //   $checkableTree.treeview('checkNode', [ checkableNodes, { silent: $('#chk-check-silent').is(':checked') }]);
          // });
          // $('#tree').treeview('checkNode', [ checkableNodes, { silent: true } ]);
          //刷新树   
          $.ajax({  //JQuery的Ajax  
            type: 'GET',    
            dataType : "json",//返回数据类型  
            // async:false, //同步会出现警告：Synchronous XMLHttpRequest on the main thread is deprecated because of its detrimental effects to the end user's experience 
            url: "/admin/role/getpermission",//请求的action路径  
            data: {roleid:row.Id,action:action,projectid:projectid}, 
             //同步请求将锁住浏览器，用户其它操作必须等待请求完成才可以执行  
            error: function () {//请求失败处理函数    
                alert('请求失败');    
            },  
            success:function(data){ //请求成功后处理函数。取到Json对象data
              // var findCheckableNodess = function() {
              //   return $('#tree').treeview('search', [ data, { ignoreCase: false, exactMatch: true } ]);//忽略大小写——这个只支持名称
              // };
              $('#tree').treeview('uncheckAll', { silent: true });
              for(var i=0;i<data.length;i++){
                // alert(data[i]);
                var findCheckableNodess = function() {
                  return $('#tree').treeview('findNodes', [data[i], 'id']);
                }; 
                var checkableNodes = findCheckableNodess();
                // $('#tree').treeview('checkNode', [ checkableNodes, { silent: true } ]);
                $('#tree').treeview('toggleNodeChecked', [ checkableNodes, { silent: true } ]);
              }
            }
          });
        }
        // if (selectRow.length>1) {
        //   alert("请不要勾选一个以上项目！");
        //   return false;
        // }
      });

      $("#table2").on("click-row.bs.table",function(e,row,ele){

      });

      $("#table3").on("click-row.bs.table",function(e,row,ele){

      });
    });
    // onClickRow  click-row.bs.table  row, $element 当用户点击某一行的时候触发，参数包括：
    // row：点击行的数据，
    // $element：tr 元素，
    // field：点击列的 field 名称
  </script>



  <!-- 添加角色 -->
  <div class="container">
    <form class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加角色</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">        
              <div class="form-group must">
                <label class="col-sm-3 control-label">角色名</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Rolename"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">角色编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Rolenumber"></div>
              </div>

              <div class="form-group must">
                <label class="col-sm-3 control-label">状态</label>
                <div class="col-sm-7">
                  <select id="Status" class="form-control">
                    <option value="0">正常</option>
                    <option value="1">失效</option>
                  </select>
                </div>
              </div>

            </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="button" class="btn btn-primary" onclick="save()">保存</button>
        </div>
      </div>
    </div>
      </div>
    </form>
  </div>
  <!-- 修改角色 -->
  <div class="container">
    <form class="form-horizontal">
    <div class="modal fade" id="modalTable1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">修改角色</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">角色名</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Rolename1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">角色编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Rolenumber1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">状态</label>
                <div class="col-sm-7">
                  <select id="Status1" class="form-control">
                    <option value="0">正常</option>
                    <option value="1">失效</option>
                  </select>
                </div>
              </div>
            </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="button" class="btn btn-primary" onclick="update()">修改</button>
        </div>
      </div>
    </div>
      </div>
    </form>
  </div>

</body>
</html>