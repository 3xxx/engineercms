<!-- 这个是显示左侧栏，右边index_user显示用户的cms情况 -->
<!DOCTYPE html>
{{template "header"}}
<title>EngineerCMS</title>
  <script type="text/javascript" src="/static/js/bootstrap-treeview.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css"/>
</head>

<body>
  <div class="navbar navba-default navbar-fixed-top">
    <div class="container-fill">{{template "navbar" .}}</div>
  </div>
  <!-- 侧栏 -->
  <div id="treeview" class="col-xs-2"></div>
  <!-- 右侧frame -->
  <div class="col-lg-10">
    <iframe src="/index/user" name='main' frameborder="0"  width="100%" scrolling="no" marginheight="0" marginwidth="0" id="iframepage" onload="this.height=100"></iframe> 
  </div>  

<div id="footer">
  <div class="col-lg-12">
    <br>
    <hr/>
  </div>
  <div class="col-lg-6">
    <h4>Copyright © 2016-2017 EngineerCMS</h4>
    <p>
      网站由 <i class="user icon"></i>
      <a target="_blank" href="https://github.com/3xxx">@3xxx</a>
      建设，并由
      <a target="_blank" href="http://golang.org">golang</a>
      和
      <a target="_blank" href="http://beego.me">beego</a>
      提供动力。
    </p>

    <p>
      请给 <i class="glyphicon glyphicon-envelope"></i>
      <a class="email" href="mailto:504284@qq.com">我</a>
      发送反馈信息或提交
      <i class="tasks icon"></i>
      <a target="_blank" href="https://github.com/3xxx/engineercms/issues">网站问题</a>
      。
    </p>
  </div>
  <div class="col-lg-6">
    <h4 >更多项目</h4>
    <div >
      <p>
        <a href="https://github.com/3xxx/hydrows">供水管线设计工具</a>
      </p>
      <p>
        <a href="https://github.com/3xxx/merit">价值成果管理系统</a>
      </p>
    </div>
  </div>
</div> 

  <script type="text/javascript">
    function reinitIframe(){//http://caibaojian.com/frame-adjust-content-height.html
      var iframe = document.getElementById("iframepage");
      try{
        var bHeight = iframe.contentWindow.document.body.scrollHeight;
        var dHeight = iframe.contentWindow.document.documentElement.scrollHeight;
        var height = Math.max(bHeight, dHeight); iframe.height = height;
        // console.log(height);//这个显示老是在变化
      }catch (ex){
      } 
    } 
    window.setInterval("reinitIframe()", 200);
  </script>

  <script type="text/javascript">
    $(function() {
          $('#treeview').treeview({
          data: {{.json}},//[{{.json}}]——有时候加个中括号就行了。
          // data:alternateData,
          levels: 2,// expanded to 5 levels
          // enableLinks:true,
          showTags:true,
        });
        $('#treeview').on('nodeSelected', function(event, data) {
          // alert("名称："+data.text);
          // alert("节点id："+data.nodeId);
          // alert("部门id："+data.Id);  
          // alert("部门级别："+data.Level);
          $("#regis").html(data.text);//显示部门名称
          $("#regis").css("color","black");
          //点击菜单，右侧显示那个人的主页
          // document.getElementById("iframepage").src="/index/main/"+data.Id;
          //点击菜单，打开新标签页个人主页
          window.open('http://'+data.href)
          //当前窗口打开这个个人主页：
          // parent.location.href='/index/main/'+data.Id
        });   
    });
  </script>
</body>
</html>
<!-- <button id="directNextpage" onclick="window.location.reload('/topic/add?id={{.Category.Id}}&mid=1')">Direct Next Page</button> 这个方法跳不出去iframe
  onclick="window.open('/topic/add?id={{.Category.Id}}&mid=1')"新标签页中打开
  onclick="parent.location.href='/topic/add?id={{.Category.Id}}&mid=2'跳出iframe重新打开-->