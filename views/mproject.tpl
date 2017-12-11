<!-- 具体一个项目的侧栏，右侧为project_products.tpl,显示任意侧栏下的成果 -->
<!DOCTYPE html>
{{template "mheader"}}
<title>项目详细-EngiCMS</title>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css"/>
<style type="text/css">
  .navbar-default {
  background-color: #e74c3c;
  border-color: #c0392b;
}
.navbar-default .navbar-brand {
  color: #ecf0f1;
}
.navbar-default .navbar-brand:hover,
.navbar-default .navbar-brand:focus {
  color: #ffbbbc;
}
.navbar-default .navbar-text {
  color: #ecf0f1;
}
.navbar-default .navbar-nav > li > a {
  color: #ecf0f1;
}
.navbar-default .navbar-nav > li > a:hover,
.navbar-default .navbar-nav > li > a:focus {
  color: #ffbbbc;
}
.navbar-default .navbar-nav > .active > a,
.navbar-default .navbar-nav > .active > a:hover,
.navbar-default .navbar-nav > .active > a:focus {
  color: #ffbbbc;
  background-color: #c0392b;
}
.navbar-default .navbar-nav > .open > a,
.navbar-default .navbar-nav > .open > a:hover,
.navbar-default .navbar-nav > .open > a:focus {
  color: #ffbbbc;
  background-color: #c0392b;
}
.navbar-default .navbar-toggle {
  border-color: #c0392b;
}
.navbar-default .navbar-toggle:hover,
.navbar-default .navbar-toggle:focus {
  background-color: #c0392b;
}
.navbar-default .navbar-toggle .icon-bar {
  background-color: #ecf0f1;
}
.navbar-default .navbar-collapse,
.navbar-default .navbar-form {
  border-color: #ecf0f1;
}
.navbar-default .navbar-link {
  color: #ecf0f1;
}
.navbar-default .navbar-link:hover {
  color: #ffbbbc;
}

@media (max-width: 767px) {
  .navbar-default .navbar-nav .open .dropdown-menu > li > a {
    color: #ecf0f1;
  }
  .navbar-default .navbar-nav .open .dropdown-menu > li > a:hover,
  .navbar-default .navbar-nav .open .dropdown-menu > li > a:focus {
    color: #ffbbbc;
  }
  .navbar-default .navbar-nav .open .dropdown-menu > .active > a,
  .navbar-default .navbar-nav .open .dropdown-menu > .active > a:hover,
  .navbar-default .navbar-nav .open .dropdown-menu > .active > a:focus {
    color: #ffbbbc;
    background-color: #c0392b;
  }
}
</style> 
</head>

<body style="height: auto; overflow-y: auto;">

<!-- <div class="navbar navbar-default navbar-static-top"> -->
  <div class="container-fill">{{template "mnavbar" .}}</div>
<!-- </div> -->

<!-- <div class="col-xs-3 col-sm-3 col-md-3 col-lg-3">
        <div id="tree"></div>
        </div> -->






<!-- <body> -->
<!-- <div class="col-xs-3 col-sm-3 col-md-3 col-lg-3">
  <div id="tree"></div>
</div> -->

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
                icon: "fa fa-tachometer icon",
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
                    icon: "fa fa-cog",
                    text: "目录设置",
                    id: '01',
                    nodeId: '01'
                  }, 
                  { 
                    icon: "fa fa-bug",
                    text: "爬虫设置",
                    id: '02'
                  }, 
                  { 
                    icon: "fa fa-th-list",
                    text: "项目权限",
                    id: '03'
                  }, 
                  { 
                    icon: "fa fa-user",
                    text: "账号管理",
                    id: '04',
                    selectable: false,
                    nodes: 
                    [
                      { icon: "fa fa-users",
                        text: '用户组',
                        id: '05'
                      },
                      { icon: "fa fa-user",
                        text: 'IP权限',
                        id: '06'
                      }
                    ]
                  }
                ]
              }
            ]
            // return data;

          $('#tree').treeview({
            // data: data,         // data is not optional
            data:[{{.json}}],
            levels: 2,
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
                // $(".breadcrumb").append('<li><a href="javascript:void(0)"><i class="fa fa-home">项目编号：' + {{.Category.Code}}+ '</a></li>');
                    $(".breadcrumb").append('<li id="nav"><a href="javascript:void(0)">' + data[i].Title + '</a></li>');
                  }
              // });
            }
          });
        }); 


        var obj = {};
        obj.text = "123";
        

        $("#btn").click(function (e) {
            var arr = $('#tree').treeview('getSelected');
            for (var key in arr) {
                c.innerHTML = c.innerHTML + "," + arr[key].id;
            }
        })
    }) 

    function index1(value,row,index){
    // alert( "Data Loaded: " + index );
      return index+1
    }
  </script>

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