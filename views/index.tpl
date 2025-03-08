<!-- 这个是显示左侧栏，右边index_user显示用户的cms情况 -->
<!DOCTYPE html>
{{template "header"}}
<title>EngineerCMS</title>
<meta charset="utf-8">
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
<link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css" />

<script type="text/javascript" src="/static/js/bootstrap-treeview.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
<script type="text/javascript" src="/static/js/tableExport.js"></script>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
<script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>

<style type="text/css">
  .about-link-wrap {
    /* width:100%; */
    /* height:auto; */
    /* clear:both; */
    padding-top: 24px;
}
.about-link-wrap .ali-about-link {
    /* text-decoration:none; */
    /* text-align:left; */
    /* line-height:24px; */
    /* font-size:14px; */
    color: var(--label-color);
    letter-spacing: .5px;
    /* display:inline-block; */
    margin-right: 35px;
}

.ace-homepage-2020-hmod-footer.pc .ali-report-wrap {
    padding-top: 25px;
    /* width:100%; */
    /* height:auto; */
}

</style>
</head>
<div class="container-fill">{{template "navbar" .}}</div>

<body>
  <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
    <iframe src="/index/user" name='main' frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" id="iframepage" onload="this.height=100"></iframe>
  </div>
  <div id="footer">
    <div class="col-lg-12">
      <br>
      <hr />
      <div class="about-link-wrap">
        <a href="/v1/wx/legalnotices" class="ali-about-link" target="_blank">法律声明</a>
        <a href="/v1/wx/cookiespolicy" class="ali-about-link" target="_blank">Cookies政策</a>
        <a href="/v1/wx/integrity" class="ali-about-link" target="_blank">廉正举报</a>
        <a href="/v1/wx/security" class="ali-about-link" target="_blank">安全举报</a>
        <a href="/v1/wx/contact" class="ali-about-link" target="_blank">联系我们</a>
        <a href="https://github.com/3xxx/engineercms" class="ali-about-link" target="_blank"><i class="fa fa-github">源码仓库</i></a>
      </div>
      <div class="about-link-wrap">
        <a target="_blank" href="https://beian.miit.gov.cn/"><img src="/static/img/beiantubiaobianhao.png" class="ali-report-img" alt="粤ICP备2024338992号" loading="lazy"><span class="ali-report-link-text">备案号：粤ICP备2024338992号</span></a>
      </div>
    </div>
  </div>

  <script type="text/javascript">
    function reinitIframe() {
      var iframe = document.getElementById("iframepage");
      try {
        var bHeight = iframe.contentWindow.document.body.scrollHeight;
        var dHeight = iframe.contentWindow.document.documentElement.scrollHeight;
        var height = Math.max(bHeight, dHeight);
        iframe.height = height;
      } catch (ex) {}
    }
    window.setInterval("reinitIframe()", 200);
  
    $(function() {
      $('#treeview').treeview({
        data: {{.json }},
        levels: 2,
        showTags: true,
      });
      $('#treeview').on('nodeSelected', function(event, data) {
        $("#regis").html(data.text);
        $("#regis").css("color", "black");
        window.open('http://' + data.href)
      });
    });
  </script>
</body>
</html>