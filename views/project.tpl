<!-- 具体一个项目的侧栏，右侧为project_products.tpl,显示任意侧栏下的成果 -->
<!DOCTYPE html>
{{template "header"}}
<title>项目详细-EngiCMS</title>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css"/>
</head>
<!-- <div class="navbar navbar-default navbar-static-top"> -->
  <div class="container-fill">{{template "navbar" .}}</div>
<!-- </div> -->
<body>
<div class="col-xs-3 col-sm-3 col-md-3 col-lg-3">
  <div id="tree"></div>
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
            var data = 
            [
              {
                text: "系统设置",
                // icon: "fa fa-tachometer icon",
                // selectedIcon: "glyphicon glyphicon-stop",
                href: "#node-1",
                selectable: true,
                id: '00',
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
                    // icon: "fa fa-cog",
                    text: "目录设置",
                    id: '160929',
                    lazyLoad:true,
                  }, 
                  { 
                    // icon: "fa fa-bug",
                    text: "爬虫设置",
                    id: '02',
                    lazyLoad:true,
                  }, 
                  { 
                    // icon: "fa fa-th-list",
                    text: "项目权限",
                    id: '03'
                  }, 
                  { 
                    // icon: "fa fa-user",
                    text: "账号管理",
                    id: '04',
                    selectable: false,
                    // nodes: 
                    // [
                    //   { icon: "fa fa-users",
                    //     text: '用户组',
                    //     id: '05'
                    //   },
                    //   { icon: "fa fa-user",
                    //     text: 'IP权限',
                    //     id: '06'
                    //   }
                    // ]
                  }
                ]
              }
            ]
            // return data;
          $('#tree').treeview({
            // data: data,         // data is not optional
            data:[{{.json}}],
            levels: 2,
            showTags:true,
            loadingIcon:"fa fa-hourglass",
            lazyLoad:loaddata
              // var $Tree = $('#tv').treeview({
              // data: defaultData,
              // lazyLoad: function (node, display) {
              //     data = defaultData2;
              //     display(data);
              // }
              // });
            // enableLinks: true,
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
            //点击任何一级，都是显示这一级下的成果
          document.getElementById("iframepage").src="/project/{{.Id}}/"+data.id;
          // $("#iframepage").contents().find("#table0").bootstrapTable('refresh', {url:'/project/products/'+data.id});
          // $('#table0').bootstrapTable('refresh', {url:'/project/products/'+data.id});
          //?secid="+data.Id+"&level="+data.Level;
          //这里用刷新右侧表格中的数据refresh行不通！！！！
          $.ajax({
            type:"get",
            url:"/project/navbar/"+data.id,
            // data: { uname: $("#uname").val()},
            // dataType:'json',//dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
            success:function(data,status){
              // $.each(data,function(i,d){
                $(".breadcrumb #nav").remove();
                for (i=0;i<data.length;i++)
                  {
                    // $(".breadcrumb").append('<li><a href="javascript:void(0)"><i class="fa fa-home">项目编号：' + {{.Category.Code}}+ '</a></li>');onclick="SomeJavaScriptCode"
                    $(".breadcrumb").append('<li id="nav"><a href="javascript:gototree('+data[i].Id+')">' + data[i].Title + '</a></li>');
                  }
              // });
            }
          });
        }); 

        // var obj = {};
        // obj.text = "123";
        
        $("#btn").click(function (e) {
            var arr = $('#tree').treeview('getSelected');
            for (var key in arr) {
                c.innerHTML = c.innerHTML + "," + arr[key].id;
            }
        })
    })

    function loaddata(node,func){//这个技巧真高，即能返回参数，又能把参数通过函数发回去
      // alert(node.id);
      // alert(func);
      $.ajax({
        type:"get",
        url:"/project/getprojcate",
        data: {id:node.id},
        success:function(data,status){
          if (data){
            func(data);
          }
        }
      });
      // var singleNode = {
      //   text: "projcatename2",
      //   id:"08",
      // };
      //var _this = this;
      // this._options.lazyLoad(node, function (nodes) {
        // Adding the node will expand its parent automatically
        // _this.addNode(nodes, node);
      // });
      // func(singleNode);//这样明显优雅很多
      // $("#tree").treeview("addNode", [singleNode,node]);这一句和上面一句等同
      // if (projcatename2)
      //   {  
      //       $.ajax({
      //           type:"post",
      //           url:"/admin/project/addprojectcate",
      //           data: {id:arr[0].id,name:projcatename2,code:projcatecode2},
      //           success:function(data,status){
      //             alert("添加“"+data+"”成功！(status:"+status+".)");
      //             var singleNode = {
      //               text: projcatename2,
      //               id:data,
      //               code:projcatecode2
      //             };
      //             $("#tree").treeview("addNode", [singleNode,arr]);
      //             $('#modalTable2').modal('hide');
      //            }
      //       });  
      //   } 
    }

    function gototree(e){
      document.getElementById("iframepage").src="/project/{{.Id}}/"+e;
      var findCheckableNodess = function() {
        return $('#tree').treeview('findNodes', [e, 'id']);
      }; 
      var checkableNodes = findCheckableNodess();
        $('#tree').treeview('toggleNodeSelected', [ checkableNodes, { silent: true } ]);
        $('#tree').treeview('toggleNodeExpanded', [ checkableNodes, { silent: true } ]);
        $('#tree').treeview('revealNode', [ checkableNodes, { silent: true } ]);
    }
        // Select/unselect/toggle nodes
        // $('#input-select-node').on('keyup', function (e) {
        //   selectableNodes = findSelectableNodes();
        //   $('.select-node').prop('disabled', !(selectableNodes.length >= 1));
        // });
        // $('#btn-select-node.select-node').on('click', function (e) {
        //   $selectableTree.treeview('selectNode', [ selectableNodes, { silent: $('#chk-select-silent').is(':checked') }]);
        // });
        // $('#btn-unselect-node.select-node').on('click', function (e) {
        //   $selectableTree.treeview('unselectNode', [ selectableNodes, { silent: $('#chk-select-silent').is(':checked') }]);
        // });
        // $('#btn-toggle-selected.select-node').on('click', function (e) {
        //   $selectableTree.treeview('toggleNodeSelected', [ selectableNodes, { silent: $('#chk-select-silent').is(':checked') }]);
        // });
    function index1(value,row,index){
    // alert( "Data Loaded: " + index );
      return index+1
    }
</script>

<div class="col-xs-9 col-sm-9 col-md-9 col-lg-9">
<!-- 面包屑导航 -->
  <div class="breadcrumbs">
    <ol class="breadcrumb" split="&gt;">
      <li>
        <a href="javascript:gototree({{.Category.Id}})"> <i class="fa fa-home" aria-hidden="true"></i>
          项目编号：{{.Category.Code}}
        </a>
      </li>
      <!-- {{range $index, $elem :=.jishu}}
      <li>
        <a href="javascript:void(0)">
          jibie.title
        </a>
      </li>
      {{end}}-->
    </ol>
  </div>
    <!-- <div class="form-group"> -->
        <!-- <label class="control-label" id="regis" for="LoginForm-UserName"></label> 显示部门名称  -->
    <!-- </div> -->
        <!-- <iframe src="/secofficeshow" name='main' id="iframepage" frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" onLoad="iFrameHeight()"></iframe> -->
        <!-- <iframe src="/secofficeshow" name='main' id="iframepage" frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" onload="changeFrameHeight()"></iframe> -->
        <!-- 默认显示所有成果？还是项目简介？当为项目id时，判断是一级，显示里面的成果 -->
       <iframe src="/project/{{.Id}}/{{.Id}}" name='iframepage' id="iframepage" frameborder="0"  width="100%" scrolling="no" marginheight="0" marginwidth="0"  onload="this.height=800"></iframe> 
</div>  


<script type="text/javascript">
 function reinitIframe(){//http://caibaojian.com/frame-adjust-content-height.html
  var iframe = document.getElementById("iframepage");
    try{
    var bHeight = iframe.contentWindow.document.body.scrollHeight;
    var dHeight = iframe.contentWindow.document.documentElement.scrollHeight; var height = Math.max(bHeight, dHeight,800); iframe.height = height;
      // console.log(height);//这个显示老是在变化
    }catch (ex){
    } 
    } 
    window.setInterval("reinitIframe()", 200);


    // $(function () {
    //     $('#search').click(function () {
    //       var productid=$("#productid").val();
    //       var keyword=$("#keyword").val();
    //       if (productid){
    //         var url="/projects/search"
    //       }else{
    //         var url="/project/product/search"
    //       }
    //       $.ajax({
    //         type:"post",//这里是否一定要用post？？？
    //         url:url,
    //         data: {productid:productid,keyword:keyword},
    //         success:function(data,status){//数据提交成功时返回数据
    //           $("#attachtitle").html(rowtitle+'—附件列表');
    //           $('#attachs').bootstrapTable('refresh', {url:'/achievement/catalog/attachment/'+row.id});
    //           $('#modalattach').modal({
    //             show:true,
    //             backdrop:'static'
    //           });
    //         }
    //       });

    //     });
    // });
 </script>
</body>
</html>