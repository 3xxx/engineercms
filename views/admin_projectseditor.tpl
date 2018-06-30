<!-- 采用ztree编辑目录、设置同步ip、设置公开私有 -->
<!DOCTYPE html>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/> -->
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/font-awesome.min.css"/> -->
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>

  <link rel="stylesheet" href="/static/css/zTreeStyle/metro.css">
  <script src="/static/js/jquery.ztree.all-3.5.min.js"></script>
</head>

<body>

<div class="col-lg-12">
  <h3>项目列表</h3>
<div id="toolbar1" class="btn-group">
        <button type="button" data-name="editorcateButton" id="editorcateButton" class="btn btn-default"> <i class="fa fa-edit">目录</i>
        </button>
        <button type="button" data-name="editorpubButton" id="editorpubButton" class="btn btn-default"> <i class="fa fa-edit">公开、私有</i>
        </button>
        <button type="button" data-name="editoripButton" id="editoripButton" class="btn btn-default">
        <i class="fa fa-edit">同步IP</i>
        </button>
</div>

<table id="table0"
        data-toggle="table"  
        data-url="/project/getprojects"
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
        <th data-field="Label">标签</th>
        <th data-field="Principal">负责人</th>
        <th data-field="Product">成果数量</th>
        <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
      </tr>
    </thead>
</table>

<SCRIPT type="text/javascript">
    function removeHoverDom(treeId, treeNode) {
        $("#addBtn_"+treeNode.tId).unbind().remove();
        $("#removeBtn_"+treeNode.tId).unbind().remove();
        $("#editBtn_"+treeNode.tId).unbind().remove();
    };
    <!--
    var setting = {
      // async: {
        // enable: true,
        // url:"/admin/project/getprojectcate",
        // autoParam:["id", "name=n", "level=lv"],
        // otherParam:{"otherParam":"zTreeAsyncTest"},
        // dataFilter: filter
      // },
      //页面上的显示效果  
        view: {  
            addHoverDom: addHoverDom, //当鼠标移动到节点上时，显示用户自定义控件  
            removeHoverDom: removeHoverDom, //离开节点时的操作  
            // fontCss: getFontCss //个性化样式  
        }, 
        check: {
                enable: true
            }, 
        edit: {  
            enable: true, //单独设置为true时，可加载修改、删除图标  
            editNameSelectAll: true,  
            // showRemoveBtn: showRemoveBtn,  
            // showRenameBtn: showRenameBtn  
        },  
    
        data: {
                key: {
                    children: "nodes",
                    name: "text",
                    title:""
                }
            },
        callback: {
            beforeRename:beforeRename,
            // onClick: zTreeOnClick, //单击事件  
            onRemove: onRemove, //移除事件  
            onRename: onRename //修改事件  
        } 
    };
      
    function addHoverDom(treeId, treeNode) {  
      var sObj = $("#" + treeNode.tId + "_span"); //获取节点信息  
      if (treeNode.editNameFlag || $("#addBtn_"+treeNode.tId).length>0) return;  
  
      var addStr = "<span class='button add' id='addBtn_" + treeNode.tId + "' title='add node' onfocus='this.blur();'></span>"; //定义添加按钮  
      sObj.after(addStr); //加载添加按钮  
      var btn = $("#addBtn_"+treeNode.tId);  
  
    //绑定添加事件，并定义添加操作  
    if (btn) btn.bind("click", function(){  
        var zTree = $.fn.zTree.getZTreeObj("treeCate");  
        //将新节点添加到数据库中  
        var name='NewNode';  
        $.post('/admin/project/addprojectcate?id='+treeNode.id+'&name='+name,function (data) {
          alert(data);
            var newID = data; //获取新添加的节点Id  
            zTree.addNodes(treeNode, {id:newID, pId:treeNode.id, name:name}); //页面上添加节点  
            var node = zTree.getNodeByParam("id", newID, null); //根据新的id找到新添加的节点  
            zTree.selectNode(node); //让新添加的节点处于选中状态  
        });  
    });
   };
// function removeHoverDom(treeId, treeNode) {  
//     $("#addBtn_"+treeNode.tId).unbind().remove();  
// };
  function onRename(e, treeId, treeNode, isCancel) {  
    //需要对名字做判定的，可以来这里写~~  
    $.post('/admin/project/updateprojectcate?id='+treeNode.id+'&name='+treeNode.name);  
  }
  function beforeRename(treeId, treeNode, newName, isCancel) {  
    if (newName.length == 0) {  
        alert("节点名称不能为空.");  
        return false;  
    }  
    return true;  
  } 
  function onRemove(e, treeId, treeNode) { 
    //需要对删除做判定或者其它操作，在这里写~~
    if(confirm("确认删除 节点 -- " + treeNode.name + " 吗？")){ 
      $.post('/admin/project/delprojectcate?id='+treeNode.id);
    }else{
        return false;
    }  
  }
  function beforeRemove(treeId, treeNode) {  
    var zTree = $.fn.zTree.getZTreeObj("tree");  
    zTree.selectNode(treeNode);  
    return confirm("确认删除 节点 -- " + treeNode.name + " 吗？");  
  }
    // function filter(treeId, parentNode, childNodes) {
    //   if (!childNodes) return null;
    //   for (var i=0, l=childNodes.length; i<l; i++) {
    //     childNodes[i].name = childNodes[i].name.replace(/\.n/g, '.');
    //   }
    //   return childNodes;
    // }
    // $(document).ready(function(){
    //   $.fn.zTree.init($("#treeDemo"), setting);
    // });
    //-->
</SCRIPT>
<script type="text/javascript">
  function index1(value,row,index){
  // alert( "Data Loaded: " + index );
            return index+1
  }
  function localDateFormatter(value) {
     return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }
  // 改变点击行颜色
  $(function(){
     $("#table0").on("click-row.bs.table",function(e,row,ele){
         $(".info").removeClass("info");
         $(ele).addClass("info");
     });
  });

  $(document).ready(function() {
    $("#editorcateButton").click(function() {
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选！");
        return;
      }
      if (selectRow.length>1){
      alert("请不要勾选一个以上！");
      return;
      }
      $("input#cid").remove();
      var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $("#projcatename1").val(selectRow[0].Title);
      $("#projcatecode1").val(selectRow[0].Code);
      //初始化树  
    // function initTree() {  
        $.ajax({  //JQuery的Ajax  
            type: 'POST',    
            dataType : "json",//返回数据类型  
            async:false,  
            url: "/admin/project/getprojectcate/"+selectRow[0].Id,//请求的action路径  
            // data: {"flag":true},  //同步请求将锁住浏览器，用户其它操作必须等待请求完成才可以执行  
            error: function () {//请求失败处理函数    
                alert('请求失败');    
            },  
            success:function(data){ //请求成功后处理函数。  取到Json对象data  
              treeNodes = data;   //把后台封装好的简单Json格式赋给treeNodes      
              $.fn.zTree.init($("#treeCate"), setting, treeNodes);
              //初始化树，传入树的Dom<pre name="code" class="html">  
            }
        });
    // }
    var treeObj = $.fn.zTree.getZTreeObj("treeCate");
      treeObj.expandAll(true);
        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
      })
  })

    //保存
    function save(){
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
            url:"/project/addproject",
            data: {code:projcode,name:projname,label:projlabel,principal:projprincipal,ids:ids},//
            success:function(data,status){
              alert("添加“"+data+"”成功！(status:"+status+".)");
              //按确定后再刷新 
            }
          });  
        }else{
          alert("请填写编号和名称！");
          return;
        } 
        // $(function(){$('#myModal').modal('hide')}); 
          $('#modalTable').modal('hide');
          $('#table0').bootstrapTable('refresh', {url:'/project/getprojects'});
          // "/category/modifyfrm?cid="+cid
          // window.location.reload();//刷新页面
    }

</script>
<!-- 编辑项目目录 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">编辑项目目录</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              
              <ul id="treeCate" class="ztree"></ul>
            </div>

          </div>
          <!-- <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="save()">保存</button>
          </div> -->
        </div>
      </div>
    </div>
  </div>


</div>
</body>
</html>