<!-- 采用ztree编辑目录、设置同步ip、设置公开私有 -->
<!DOCTYPE html>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
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

<table id="table0"></table>

<SCRIPT type="text/javascript">
  $(function () {
    // 初始化【未接受】工作流表格
    $("#table0").bootstrapTable({
        url : '/project/getprojects',
        method: 'get',
        search:'true',
        showRefresh:'true',
        showToggle:'true',
        showColumns:'true',
        // toolbar:'#toolbar1',
        pagination: 'true',
        sidePagination: "server",
        queryParamsType:'',
        //请求服务器数据时，你可以通过重写参数的方式添加一些额外的参数，例如 toolbar 中的参数 如果 queryParamsType = 'limit' ,返回参数必须包含
        // limit, offset, search, sort, order 否则, 需要包含: 
        // pageSize, pageNumber, searchText, sortName, sortOrder. 
        // 返回false将会终止请求。
        pageSize: 5,
        pageNumber: 1,
        pageList: [15,20, 50, 100],
        singleSelect:"true",
        clickToSelect:"true",
        queryParams:function queryParams(params) {   //设置查询参数
          var param = {
              limit: params.pageSize,   //每页多少条数据
              pageNo: params.pageNumber, // 页码
              searchText: params.searchText // $(".search .form-control").val()
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
        columns: [
          {
            title: '选择',
            radio: 'true',
            width: '10',
            align:"center",
            valign:"middle"
          },
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
              return index+1
            },
            align:"center",
            valign:"middle"
          },
          {
            field: 'Code',
            title: '编号',
            // formatter:setCode,
            align:"center",
            valign:"middle"
          },
          {
            field: 'Title',
            title: '名称',
            // formatter:setTitle,
            align:"center",
            valign:"middle"
          },
          {
            field: 'Principal',
            title: '负责人',
            align:"center",
            valign:"middle"
          },
          {
            field: 'Created',
            title: '建立时间',
            formatter:localDateFormatter,
            align:"center",
            valign:"middle"
          }
        ]
    });
  });

    function removeHoverDom(treeId, treeNode) {
        $("#addBtn_"+treeNode.tId).unbind().remove();
        $("#removeBtn_"+treeNode.tId).unbind().remove();
        $("#editBtn_"+treeNode.tId).unbind().remove();
    };

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
          // showRenameBtn: showRenameBtn,
          drag:{
            isCopy: false,
            isMove: true,
            prev: true,
            next: true,
            inner: true,
            autoOpenTime: 0,
            minMoveSize: 10
          },
          removeTitle: "删除节点",
          renameTitle: "编辑节点名称",
          showRemoveBtn: false,
          showRenameBtn: false, 
        },  
        data: {
          key: {
              children: "nodes",
              name: "text",
              title:""
          },
          simpleData: {
            enable: true,
            idKey: 'nodeId',
            pIdKey: 'parentNodeId'
          },
          keep:{
            leaf:true,
            parent:true,
          }
        },
        callback: {
          beforeRename:beforeRename,
          // onClick: zTreeOnClick, //单击事件  
          onRemove: onRemove, //移除事件  
          onRename: onRename, //修改事件
          // beforeClick: beforeClick,
          beforeDrag:beforeDrag,
          beforeDragOpen:beforeDragOpen,
          beforeDrop:beforeDrop,
          onDrag:onDrag
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

    //拖拽之前调用的函数
  function beforeDrag(treeId,treeNode){
    if(treeNode[0].nodeType == 'GROUP'){
        return false;
    }
    if(treeNode.parentId == null && treeNode.modelType !=null){
        return true;
    }
    var  node =  treeNode[0].getParentNode();
    var modelType = treeNode[0].getParentNode().modelType;
    if(modelType == 'INTERFACE'){
        return false;
    }else {
        return true;
    }
  }

  //预留被拖拽的回调函数
  function onDrag(event, treeId, treeNode){
      //暂时没用到
  }
  
  //拖拽移动到展开父节点之前调用的函数
  function beforeDragOpen(){
      return true;
  }

  //拖拽操作结束之前调用的函数
  function beforeDrop(treeId, treeNode, targetNode, moveType){
    // BRS.fileLoading('show');
    var result = false;
    if(targetNode == null || (moveType != "inner" && !targetNode.parentTId)){
        BRS.fileLoading('hide');
        return false;
    }
    if(targetNode.modelType != null){
        if((targetNode.modelType == 'INTERFACE' && moveType == 'inner') || targetNode.getParentNode().modelType == 'INTERFACE'){
            BRS.fileLoading('hide');
            return false;
        }
    }
    var objDetail = {
        url: '/api/model/' + treeNode[0].id,
        async:false,
    }
    jsonAjax(objDetail,function (detailData) {
        var data = {
            nodeType : detailData.nodeType,
            code : detailData.code,
            name : detailData.name,
            builtIn : detailData.builtIn,
            iconUrl : detailData.iconUrl,
            modelType : detailData.modelType.code,
            interfaceModelId : detailData.interfaceModelId,
        };
        data.id = treeNode[0].id;
        if(moveType != 'inner'){
            data.groupId = targetNode.parentId;
        }else{
            data.groupId = targetNode.id;
        };
        var obj = {
            type:"put",
            showSuccessMsg: false,
            param: {
                params:JSON.stringify(data)
            },
            async:false,
            url: '/api/model',
        };
        jsonAjax(obj,function(updateData){
          if(updateData != null){
            result = true;
            ing('hide');
            return result;
          }
        })
    })

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
  }
  //预留拖拽结束的回调函数
  function onDrop(event, treeId, treeNode, targetNode, moveType){
      befod('hide');
      return result;
  }
  
  //预留拖拽结束的回调函数
  function onDrop(event, treeId, treeNode, targetNode, moveType){
      beforeClick(treeId, treeNode[0]);
  }
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
  
<style type="text/css">
.ztree li span.button.add {margin-left:2px; margin-right: -1px; background-position:-144px 0; vertical-align:top; *vertical-align:middle}
div#rMenu {position:absolute; visibility:hidden; top:0; background-color: #555;text-align: left;padding: 2px;}
div#rMenu ul{list-style:none;margin:0px;}
div#rMenu ul li{
  cursor: pointer;
  margin: 1px 0;
  padding: 0 5px;
  list-style: none inside none; 
  background-color: #DFDFDF; 
} 
</style>
<body>
 
<div>${tips}</div>
<div class="modal-header">
      <h3 id="myModalLabel">系统模块维护 <span style="margin-left: 200px"> 
        <shiro:hasPermission name="resource:add">
        <a class="btn btn-primary" type="button" href="${ctx}/goAddResource">增加资源</a>
      </shiro:hasPermission></span>
    </h3> 
</div>
 
<div>
    <ul id="treeDemo" class="ztree"></ul>
</div>
<div id="rMenu">
  <ul>
    <li id="add_res" onclick="addRes();">增加资源</li>
    <li id="edit_res" onclick="editRes();">编辑资源</li>
    <li id="del_res" onclick="delRes();">删除资源</li>
    <li id="add_fun" onclick="addFun();">增加功能</li>
    <li id="edit_fun" onclick="editFun();">修改功能</li>
    <li id="del_fun" onclick="delFun();">删除功能</li>
    <li id="tree_reset" onclick="resetTree();">恢复列表</li>
  </ul>
</div>
</body>

<script type="text/javascript">
      var setting = {
    edit: {//设置编辑/删除 按钮
      enable: true,
      showRemoveBtn: true,
      showRenameBtn: true
    },
    view: {
      expandSpeed: "fast",
      nameIsHTML: true,
      fontCss: getFont,
      addHoverDom: addHoverDom,//自定义添加按钮
      removeHoverDom: removeHoverDom,//自定义隐藏按钮
    },
    data: {
      simpleData: {
        enable: true
      }
    },
    async: {//初始化异步加载资源节点
      enable: true,
      url: "${ctx}/initResourceNodes?time"+new Date().getTime(),
      type : "post",
      autoParam : ["id","name"],
      dataFilter : ajaxDataFilter //异步加载回调函数
    },
    callback: {//核心回调函数
      beforeDrag: beforeDrag,
      beforeDrop: beforeDrop,
      onDrop: zTreeOnDrop,
      beforeEditName: beforeEditName,
      beforeRemove: beforeRemove,
      onRemove: onRemove,
      onRightClick: OnRightClick
    }
  };
  
  function getFont(treeId, node) {
    return node.font ? node.font : {};
  }
  
  function addHoverDom(treeId, treeNode) {
    var sObj = $("#" + treeNode.tId + "_span");
    if (treeNode.editNameFlag || $("#addBtn_"+treeNode.tId).length>0) return;
    if(treeNode.level == 0){
      var addStr = "<shiro:hasPermission name='resource:add'>"
                +"<span class='button add' id='addBtn_" + treeNode.tId+ "' title='增加资源' onfocus='this.blur();'></span>"
            +"</shiro:hasPermission>";
      sObj.after(addStr);
    }
    if(treeNode.level == 1){
      var addStr = "<shiro:hasPermission name='resource:add'>"
                +"<span class='button add' id='addBtn_" + treeNode.tId+ "' title='增加功能' onfocus='this.blur();'></span>"
            +"</shiro:hasPermission>";
      sObj.after(addStr);
    }
    var btn = $("#addBtn_"+treeNode.tId);
    if (btn) btn.bind("click", function(){
      if(treeNode.level == 0){
        window.location.href="${ctx}/goAddResource?treeNodeId="+treeNode.id+"&time=" +new Date().getTime();
      }
      if(treeNode.level == 1){
        window.location.href="${ctx}/goAddFunction/"+treeNode.id;
      } 
    });
  };
  
  function removeHoverDom(treeId, treeNode) {
    $("#addBtn_"+treeNode.tId).unbind().remove();
  };
  
  function beforeEditName(treeId, treeNode) {
    if(treeNode.level == 0 || treeNode.level == 1){
      window.location.href="${ctx}/goUpdateResource/"+treeNode.id;
    }
    if(treeNode.level == 2){
      window.location.href="${ctx}/goUpdateFunc/"+treeNode.id;
    } 
    return false;
  }
  
  function beforeRemove(treeId, treeNode) {
    var zTree = $.fn.zTree.getZTreeObj(treeId);
    zTree.selectNode(treeNode);
    if(treeNode.type == 1){
      return confirm("确认删除资源  " + treeNode.name + "  吗？"); 
    }
    return confirm("确认删除功能  " + treeNode.name + "  吗？"); 
  }
  
  function onRemove(e, treeId, treeNode) {
    window.location.href="${ctx}/deleteResource/"+treeNode.id;
  }
  
  function beforeDrag(treeId, treeNodes) {
    for (var i=0,l=treeNodes.length; i<l; i++) {
      if (treeNodes[i].drag === false) {
        return false;
      }
    }
    return true;
  }
  
  function beforeDrop(treeId, treeNodes, targetNode, moveType) {
    var multiTreeLevel = treeNodes[0].level;
    if(targetNode == null){//将某个节点拖拽为根节点
      return false;
    }else{
      if(moveType == 'inner' && targetNode.level >= 2){//三层节点{0=跟资源节点，1=资源节点, 2=功能节点}
        return false;
      }
    }
    for(var i=0; i<treeNodes.length; i++){//多选拖拽为同级节点
      if(treeNodes[i].level != multiTreeLevel){
        return false;
      }
      if(moveType == 'inner' && treeNodes[i].isParent){
        return false;
      }
    }
    
    if(treeNodes.length >1){
      return confirm("确定对   " + treeNodes[0].name + "   等资源进行移动 ?");
    }
    return confirm("确定对   " + treeNodes[0].name + "   资源进行移动 ?");
  }
  
  //获取当前节点的同级兄弟节点
  function getPeerNodes(targetNode){
    if(targetNode == null){
      return null
    }else{
      if(targetNode.getParentNode() != null){
        return targetNode.getParentNode().children;
      }
      return null;
    }
  }
  //拖拽结束后Ajax相应操作，更新资源表
  function zTreeOnDrop(event, treeId, treeNodes, targetNode, moveType, isCopy) { 
    console.info(moveType +":"+isCopy);
    if(moveType != null){ // moveType == null 表示拖拽无效
      var afterDragPeerNodes=null;
      if(moveType == 'inner'){//目标节点成为父节点
        afterDragPeerNodes = targetNode.children;
      }else{
        if(getPeerNodes(targetNode) == null){//根节点
          afterDragPeerNodes = $.fn.zTree.getZTreeObj("treeDemo").getNodes();
        }else{//普通next/prev拖拽
          afterDragPeerNodes=getPeerNodes(targetNode);  
        }
      }
      $.ajax({  
        url : "${ctx}/updateResourceNodes?time"+new Date().getTime(),
         type : "post",
         dataType : "json",
         data : {  "moveType" : moveType,
            "treeNodes" : JSON.stringify(treeNodes),
           "targetNode" : JSON.stringify(targetNode),
          "targetParentNode" : JSON.stringify(targetNode.getParentNode()),
          "afterDragPeerNodes" : JSON.stringify(afterDragPeerNodes)
              }
      });
    }
  };
  
  //右键单击显示功能列表
  function OnRightClick(event, treeId, treeNode) {
    if (!treeNode && event.target.tagName.toLowerCase() != "button" && $(event.target).parents("a").length == 0) {
      $.fn.zTree.getZTreeObj(treeId).cancelSelectedNode();
      showRMenu("root", event.clientX, event.clientY);
    } else if (treeNode && !treeNode.noR) {
      $.fn.zTree.getZTreeObj(treeId).selectNode(treeNode);
      showRMenu(treeNode, event.clientX, event.clientY);
    }
  }
  
  function showRMenu(type, x, y) {
    $("#rMenu ul").show();
     if (type == "root") {
      $("#add_res").show();
      $("#edit_res").hide();
      $("#del_res").hide();
      $("#add_fun").hide();
      $("#edit_fun").hide();
      $("#del_fun").hide();
      $("#tree_reset").show();
    } else {
      if(type.level == 0){
        $("#add_res").show();
        $("#edit_res").show();
        $("#del_res").show();
        $("#add_fun").hide();
        $("#edit_fun").hide();
        $("#del_fun").hide();
        $("#tree_reset").show();
      }
      if(type.level == 1){
        $("#add_res").hide();
        $("#edit_res").show();
        $("#del_res").show();
        $("#add_fun").show();
        $("#edit_fun").hide();
        $("#del_fun").hide();
        $("#tree_reset").show();
      }
      if(type.level == 2){
        $("#add_res").hide();
        $("#edit_res").hide();
        $("#del_res").hide();
        $("#add_fun").hide();
        $("#edit_fun").show();
        $("#del_fun").show();
        $("#tree_reset").show();
      }
    } 
    rMenu.css({"top":y+"px", "left":x+"px", "visibility":"visible"});
 
    $("body").bind("mousedown", onBodyMouseDown);
  }
  
  function hideRMenu() {
    if (rMenu) rMenu.css({"visibility": "hidden"});
    $("body").unbind("mousedown", onBodyMouseDown);
  }
  function onBodyMouseDown(event){
    if (!(event.target.id == "rMenu" || $(event.target).parents("#rMenu").length>0)) {
      rMenu.css({"visibility" : "hidden"});
      $.fn.zTree.getZTreeObj("treeDemo").cancelSelectedNode();
    }
  }
  
  function addRes(){
    var nodes = $.fn.zTree.getZTreeObj('treeDemo').getSelectedNodes();
    if(nodes && nodes.length>0){
      window.location.href="${ctx}/goAddResource?treeNodeId="+nodes[0].id+"&time=" +new Date().getTime();
    }else{
      window.location.href="${ctx}/goAddResource?time=" +new Date().getTime();
    }
    hideRMenu();
  }
  
  function editRes(){
    var nodes = $.fn.zTree.getZTreeObj('treeDemo').getSelectedNodes();
    window.location.href="${ctx}/goUpdateResource/"+nodes[0].id;
    hideRMenu();
  }
  
  function delRes(){
    var nodes = $.fn.zTree.getZTreeObj('treeDemo').getSelectedNodes();
    if(confirm("确认删除资源  " + nodes[0].name + "  吗？")){
      window.location.href="${ctx}/deleteResource/"+nodes[0].id;
    }
    hideRMenu();
  }
  
  function addFun(){
    var nodes = $.fn.zTree.getZTreeObj('treeDemo').getSelectedNodes();
    window.location.href="${ctx}/goAddFunction/"+nodes[0].id;
    hideRMenu();
  }
  
  function editFun(){
    var nodes = $.fn.zTree.getZTreeObj('treeDemo').getSelectedNodes();
    window.location.href="${ctx}/goUpdateFunc/"+nodes[0].id;
    hideRMenu();
  }
  
  function delFun(){
    var nodes = $.fn.zTree.getZTreeObj('treeDemo').getSelectedNodes();
    window.location.href="${ctx}/deleteResource/"+nodes[0].id;
    hideRMenu();
  }
  function resetTree() {
    hideRMenu();
    $.fn.zTree.init($("#treeDemo"), setting, null);
  }
  
  function ajaxDataFilter(treeId, parentNode, responseData) {
    for(var i=0;i<responseData.length;i++){
      responseData[i].font = {'font-weight':'bold'};
      if(responseData[i].type == 1){
        if(!responseData[i].hasOwnProperty('parentid')){
          responseData[i].iconOpen = '${ctx}/static/ztree/css/zTreeStyle/img/diy/1_open.png';
          responseData[i].iconClose = '${ctx}/static/ztree/css/zTreeStyle/img/diy/1_close.png'; 
        }else{
          if(!responseData[i].isParent){
            responseData[i].icon = '${ctx}/static/ztree/css/zTreeStyle/img/diy/8.png';
          }
        }
      }
       if(responseData[i].type == 2){
           responseData[i].icon = '${ctx}/static/ztree/css/zTreeStyle/img/diy/9.png';
      }  
    }
    return responseData;
  };
  
  function setCheck() {
    var zTree = $.fn.zTree.getZTreeObj("treeDemo");
    zTree.setting.edit.drag.isCopy = false;
    zTree.setting.edit.drag.isMove = true;
    zTree.setting.edit.drag.prev = true;
    zTree.setting.edit.drag.inner = true;
    zTree.setting.edit.drag.next = true;
    zTree.setting.edit.removeTitle = "删除";
    zTree.setting.edit.renameTitle = "修改";
  }
  
  var rMenu;
  $(document).ready(function(){
    $.fn.zTree.init($("#treeDemo"), setting, null);
    setCheck();
    rMenu = $("#rMenu");
  });
</script>


</div>
</body>
</html>