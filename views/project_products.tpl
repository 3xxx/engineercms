<!-- 具体一个项目侧栏id下所有成果，不含子目录下的成果 -->
<!DOCTYPE html>
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
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
  <link rel="stylesheet" type="text/css" href="/static/css/font-awesome.min.css"/>
  <script src="/static/js/tableExport.js"></script>

  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>

  <script type="text/javascript" src="/static/js/moment.min.js"></script>
</head>

<body>
<div class="col-lg-12">
  <h3>成果列表{{.Id}}</h3>
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
        data-url="/project/products/{{.Id}}"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="ProjectName"
        data-sort-order="desc"
        data-page-size="5"
        data-page-list="[5,20, 50, 100, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client"
        data-single-select="true"
        data-click-to-select="true"
        >
    <thead>        
      <tr>
        <!-- radiobox data-checkbox="true"-->
        <th data-width="10" data-radio="true"></th>
        <th data-formatter="index1">#</th>
        <th data-field="Code">编号</th>
        <th data-field="Title">名称</th>
        <th data-field="Label">关键字</th>
        <th data-field="Principal">设计</th>
        <th data-field="Product">附件数量</th>
        <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
      </tr>
    </thead>
</table>
<!-- <div class="gridview2"></div> -->

<script type="text/javascript">
  /*数据json*/
  var json =
     [{"Id":"1","Code":"SL0001-510-08","Title":"水利枢纽布置图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      {"Id":"2","Code":"SL0002-530-10","Title":"电力布置图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      {"Id":"3","Code":"SL0003-650-20","Title":"市政布置图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      {"Id":"4","Code":"SL0004-750-60","Title":"建筑平面图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      {"Id":"5","Code":"SL0005-870-20","Title":"交通纵面图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},
      {"Id":"6","Code":"SL0006-230-25","Title":"境外鸟瞰图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"}];
        /*初始化table数据*/
        $(function(){
            $("#table0").bootstrapTable({
                data:json
                // onClickRow: function (row, $element) {
                  // alert( "选择了行Id为: " + row.Id );
                  // rowid=row.Id//全局变量
                  // $('#table1').bootstrapTable('refresh', {url:'/admincategory?pid='+row.Id});
                // }
            });
        });
  function index1(value,row,index){
  // alert( "Data Loaded: " + index );
            return index+1
  }

  function localDateFormatter(value) {
     return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }
  // 改变点击行颜色
  $(function(){
     // $("#table").bootstrapTable('destroy').bootstrapTable({
     //     columns:columns,
     //     data:json
     // });
     $("#table0").on("click-row.bs.table",function(e,row,ele){
         $(".info").removeClass("info");
         $(ele).addClass("info");
         // rowid=row.Id//全局变量
         // $('#table1').bootstrapTable('refresh', {url:'/admin/category/'+row.Id});
     });
     // $("#get").click(function(){
     //     alert("商品名称：" + getContent().TuanGouName);
     // })
  });
  // $(function () {
  //     var $result = $('#eventsResult');
  //     var selectRow=$('#table').bootstrapTable('getSelections');
  
  //     $('#table').on('all.bs.table', function (e, name, args) {
  //         console.log('Event:', name, ', data:', args);
  //     })
  //     .on('click-row.bs.table', function (e, row, $element) {
  //       alert("选择！"+row.Id);
  //       if (selectRow.length<1){
  //         selectRow=$('#table').bootstrapTable('getSelections');
  //         alert("请选择"+selectRow.length);
  //         // return;
  //         }
  //         $result.text('Event: click-row.bs.table');
  //     })
  //     .on('dbl-click-row.bs.table', function (e, row, $element) {
  //         $result.text('Event: dbl-click-row.bs.table');
  //     })
  //     .on('sort.bs.table', function (e, name, order) {
  //         $result.text('Event: sort.bs.table');
  //     })
  //     .on('check.bs.table', function (e, row) {
  //         $result.text('Event: check.bs.table');
  //     })
  //     .on('uncheck.bs.table', function (e, row) {
  //         $result.text('Event: uncheck.bs.table');
  //     })
  //     .on('check-all.bs.table', function (e) {
  //         $result.text('Event: check-all.bs.table');
  //     })
  //     .on('uncheck-all.bs.table', function (e) {
  //         $result.text('Event: uncheck-all.bs.table');
  //     })
  //     .on('load-success.bs.table', function (e, data) {
  //         $result.text('Event: load-success.bs.table');
  //     })
  //     .on('load-error.bs.table', function (e, status) {
  //         $result.text('Event: load-error.bs.table');
  //     })
  //     .on('column-switch.bs.table', function (e, field, checked) {
  //         $result.text('Event: column-switch.bs.table');
  //     })
  //     .on('page-change.bs.table', function (e, number, size) {
  //         $result.text('Event: page-change.bs.table');
  //     })
  //     .on('search.bs.table', function (e, text) {
  //         $result.text('Event: search.bs.table');
  //     });
  // });


  // $(document).ready(function() {
  // $("#addButton").click(function() {
    
  // var selectRow=$('#table').bootstrapTable('getSelections');  
  // if (selectRow.length<1){
  // selectRow=$('#table').bootstrapTable('getSelections');
  // alert("请选择"+selectRow.length);
  // return;
  // }
        // $('#modalTable').modal({
        // show:true,
        // backdrop:'static'
        // });
    // })
  // })

  $(document).ready(function() {
    $("#addButton").click(function() {
      $("input#pid").remove();
      var th1="<input id='pid' type='hidden' name='pid' value='" +{{.Id}}+"'/>"
        $(".modal-body").append(th1);

        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#editorButton").click(function() {
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选类别！");
        return;
      }
      if (selectRow.length>1){
      alert("请不要勾选一个以上目录！");
      return;
      }
      $("input#cid").remove();
      var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $("#projcatename1").val(selectRow[0].Title);
      $("#projcatecode1").val(selectRow[0].Code);
      // alert(JSON.stringify(selectRow));
      // alert(selectRow[0].Id);
      // var title = $('#'+id).attr("value");
      // var title = $('#'+id).attr("href");
      // var categoryid = $('#categoryid').val();
        $('#modalTable1').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#deleteButton").click(function() {
      var selectRow=$('#table0').bootstrapTable('getSelections');
      // if (selectRow.length<1){
      //   alert("请先勾选类别！");
      //   return;
      // }
      if (selectRow.length<=0) {
        alert("请先勾选类别！");
        return false;
      }
      var ids=$.map(selectRow,function(row){
        return row.id;
      })
      //删除已选数据
      $('$table0').bootstrapTable('remove',{
        field:'id',
        values:ids
      });
    })
  })

   /*数据json*/
  var json1 = [{"Id":"1","Title":"规划","Code":"A","Grade":"1"},
              {"Id":"7","Title":"可研","Code":"B","Grade":"1"},
              {"Id":"2","Title":"报告","Code":"B","Grade":"2"},
              {"Id":"3","Title":"图纸","Code":"T","Grade":"2"},
              {"Id":"4","Title":"水工","Code":"5","Grade":"3"},
              {"Id":"5","Title":"机电","Code":"6","Grade":"3"},
              {"Id":"6","Title":"施工","Code":"7","Grade":"3"}];
  /*初始化table数据*/
  $(function(){
      $("#table1").bootstrapTable({
          data:json1
      });
  });

//填充select选项
$(document).ready(function(){
//   $(array).each(function(index){
//     alert(this);
// });
 
// $.each(array,function(index){
//     alert(this);
// });
  $.each({{.Select2}},function(i,d){
  // alert(this);
  // alert(i);
  // alert(d);
   $("#admincategory").append('<option value="' + i + '">'+d+'</option>');
   });
});

//根据选择，刷新表格
function refreshtable(){
  var newcategory = $("#admincategory option:selected").text();
  // alert("你选的是"+newcategory);
  //根据名字，查到id，或者另外写个categoryname的方法
  $('#table1').bootstrapTable('refresh', {url:'/admin/categorytitle?title='+newcategory});
}

//保存
function save111(){
      // var radio =$("input[type='radio']:checked").val();
      var projcode = $('#projcode').val();
      var projname = $('#projname').val();
      var projlabel = $('#projlabel').val();
      var projprincipal = $('#projprincipal').val();

      var selectRow3=$('#table1').bootstrapTable('getSelections');
      if (selectRow3.length<1){
        alert("请先勾选目录！");
        return;
      }
      var ids="";
      for(var i=0;i<selectRow3.length;i++){
        if(i==0){
          ids=selectRow3[i].Id;
        }else{
          ids=ids+","+selectRow3[i].Id;
        }  
      }

      // $('#myModal').on('hide.bs.modal', function () {  
      if (projname&&projcode)
        {  
          $.ajax({
            type:"post",
            url:"/project/category/addcategory",
            data: {code:projcode,name:projname,label:projlabel,principal:projprincipal,ids:ids},//
            success:function(data,status){
              alert("添加“"+data+"”成功！(status:"+status+".)");
            }
          });  
        }else{
          alert("请填写编号和名称！");
          return;
        } 
        // $(function(){$('#myModal').modal('hide')}); 
          $('#modalTable').modal('hide');
          // "/category/modifyfrm?cid="+cid
          // window.location.reload();//刷新页面
  }

  //模态框可拖曳—要引入ui-jquery.js
  // $("#modalTable").draggable({
  //   handle:".modal-header",
  //   cusor:"move",
  //   refreshPositions:false
  // });
  // 来自群，保留，批量
  // var rows= $('#account-table').bootstrapTable('getSelections');
  //       if(rows.length==0) {
  //           layer.alert('请您选择要删除的子账号！', {
  //               title:'提示信息',
  //               closeBtn: 0,
  //               icon: 0,
  //               skin: 'layui-layer-lan',
  //               shift:0 //动画类型
  //           });
  //           return false;
  //       }
  //           var ids="";
  //           for(var i=0;i<rows.length;i++){
  //               if(i==0){
  //                   ids=rows[i].frontUserId;
  //               }else{
  //                   ids=ids+","+rows[i].frontUserId;
  //               }
  //           }
</script>
<!-- 添加成果 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加成果</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="prodcode" name="prodcode"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">名称</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodname" name="prodname"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">关键字</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodlabel" name="prodlabel"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">设计</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodprincipal" name="prodprincipal"></div>
              </div>
              <!-- <div class="form-group">
                <label class="col-sm-3 control-label">标签</label>
                <div class="col-sm-7">
                  <input type="number" class="form-control digits" name="label" maxlength="20" placeholder="至多20个字符" required></div>
              </div> -->
              <!-- <div class="form-group must">
                <label class="col-sm-3 control-label">负责人</label>
                  <div class="col-sm-7">
                    <input type="password" class="form-control" name="password" id="password" maxlength="32" placeholder="至多32个字符" required></div>
              </div> -->

            <!-- <div class="form-group must">
              <label class="col-sm-3 control-label">确认密码</label>
              <div class="col-sm-7">
                <input type="password" class="form-control equalto" name="password2" maxlength="32" placeholder="至多32个字符" required data-rule-equalto="#password" data-msg-equalto="密码不一致"></div>
            </div> -->
            </div>

            
            <label>文件简介:</label>
              <div>
                <script id="container" type="text/plain" style="height:300px;"></script><!-- width:1024px; -->
              </div>

          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="save()">保存</button>
          </div>
        </div>
      </div>
    </div>
  </div>


</div>

<script type="text/javascript">
//实例化编辑器
    //议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.getEditor('editor')就能拿到相关的实例
    // var ue = UE.getEditor('container');
    var ue = UE.getEditor('container', {
    // toolbars: [
    //     ['fullscreen', 'source', 'undo', 'redo', 'bold']
    // ],
    toolbars: [[
            'fullscreen', 'source', '|', 'undo', 'redo', '|',
            'bold', 'italic', 'underline', 'fontborder', 'strikethrough', 'superscript', 'subscript', 'removeformat', 'formatmatch', 'autotypeset', 'blockquote', 'pasteplain', '|', 'forecolor', 'backcolor', 'insertorderedlist', 'insertunorderedlist', 'selectall', 'cleardoc', '|',
            'rowspacingtop', 'rowspacingbottom', 'lineheight', '|',
            'customstyle', 'paragraph', 'fontfamily', 'fontsize', '|',
            'directionalityltr', 'directionalityrtl', 'indent', '|',
            'justifyleft', 'justifycenter', 'justifyright', 'justifyjustify', '|', 'touppercase', 'tolowercase', '|',
            'link', 'unlink', 'anchor', '|', 'imagenone', 'imageleft', 'imageright', 'imagecenter', '|',
            'emotion', 'map', 'gmap', 'insertframe', 'insertcode', 'webapp', 'pagebreak', 'template', 'background', '|',
            'horizontal', 'date', 'time', 'spechars', 'wordimage', '|',
            'inserttable', 'deletetable', 'insertparagraphbeforetable', 'insertrow', 'deleterow', 'insertcol', 'deletecol', 'mergecells', 'mergeright', 'mergedown', 'splittocells', 'splittorows', 'splittocols', 'charts', '|',
            'print', 'preview', 'searchreplace', 'help', 'drafts'
        ]],
    autoHeightEnabled: true,
    autoFloatEnabled: true
});

function save(){
      // var radio =$("input[type='radio']:checked").val();
      var prodname = $('#prodname').val();
      var prodcode = $('#prodcode').val();
      var projectid = $('#pid').val();
      var prodprincipal = $('#prodprincipal').val();
      var prodlabel = $('#prodlabel').val();
      var html = ue.getContent();
      // $('#myModal').on('hide.bs.modal', function () {  
      if (prodname&&prodcode)
        {  
            $.ajax({
                type:"post",
                url:"/project/addproduct",
                data: {pid:projectid,title:prodname,code:prodcode,label:prodlabel,content:html,principal:prodprincipal},//父级id
                success:function(data,status){
                  alert("添加“"+data+"”成功！(status:"+status+".)");
                 }
            });  
        }else{
          alert("请填写编号和名称！");
          return;
        } 
        // $(function(){$('#myModal').modal('hide')}); 
          $('#modalTable').modal('hide');
          $('#table0').bootstrapTable('refresh', {url:'/project/products/'+{{.Id}}});
          // "/category/modifyfrm?cid="+cid
          // window.location.reload();//刷新页面
  }

</script>

</body>
</html>