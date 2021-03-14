<!-- onlyoffice左侧导航栏浮动验证 -->
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
            // return data;
          $('#tree').treeview({
            // data: data,         // data is not optional
            data:[{{.json}}],
            levels: 3,
            showTags:true,
            loadingIcon:"fa fa-hourglass",
            // lazyLoad:loaddata,
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

        $("#btn").click(function (e) {
            var arr = $('#tree').treeview('getSelected');
            for (var key in arr) {
                c.innerHTML = c.innerHTML + "," + arr[key].id;
            }
        });
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

 </script>
</body>
</html>