<!-- 首页右侧的frame -->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>EngineerCMS</title>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
</head>

<div class="container-fill">{{template "navbar" .}}</div>
<body>
<div class="bs-example">

<div id="details">
<h3 id="rowtitle">搜索结果</h3>

<table id="table1" 
        data-toggle="table" 
        data-search="true"
        data-url="/project/product/search?keyword={{.Key}}&productid={{.Pid}}"
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
        <th data-formatter="index1">#</th>
        <!-- <th data-field="Id">编号</th> -->
        <th data-field="Code">编号</th>
        <th data-field="Title">名称</th>
        <th data-field="Label" data-formatter="setLable">关键字</th>
        <th data-field="Principal">设计</th>
        <th data-field="Articlecontent" data-formatter="setArticle" data-events="actionEvents">文章</th>
        <th data-field="Attachmentlink" data-formatter="setAttachment" data-events="actionEvents">附件</th>
        <th data-field="Pdflink" data-formatter="setPdf" data-events="actionEvents">PDF</th>
        <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
        <!-- <th data-field="Created" data-formatter="actionFormatter" events="actionEvents">操作</th> -->
      </tr>
    </thead>
</table>
</div>

</div>  

<script type="text/javascript">
// 改变点击行颜色
  $(function(){
     $("#table0").on("click-row.bs.table",function(e,row,ele){
         $(".info").removeClass("info");
         $(ele).addClass("info");
         rowid=row.Id;//全局变量
     });
  });

  $(document).ready(function(){
    $("#search").click(function(){//这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
      var radio =$("input[type='radio']:checked").val();
      $.ajax({
        type:"post",//这里是否一定要用post，是的，因为get会缓存？？
        url:"/index/searchproduct",
        data: {keyword: $("#keyword").val(),radiostring:radio},
        success:function(data,status){//数据提交成功时返回数据
          //显示结果表
          $("#rowtitle").html("搜寻结果");
          $("#details").show();
          $('#table1').bootstrapTable('append', data);
          $('#table1').bootstrapTable('scrollTo', 'bottom');
        }       
      });              
    });
  });

  function getKey(){  
    if(event.keyCode==13){  
     var radio =$("input[type='radio']:checked").val();
      $.ajax({
        type:"post",//这里是否一定要用post，是的，因为get会缓存？？
        url:"/index/searchproduct",
        data: {keyword: $("#keyword").val(),radiostring:radio},
        success:function(data,status){//数据提交成功时返回数据

          //显示结果表
          $("#rowtitle").html("搜寻结果");
          $("#details").show();
          $('#table1').bootstrapTable('append', data);
          $('#table1').bootstrapTable('scrollTo', 'bottom');
        }       
      });
    }     
  } 

  function index1(value,row,index){
  // alert( "Data Loaded: " + index );
    return index+1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }
  function setCode(value,row,index){
    return "<a href='/project/product/attachment/"+row.Id+"'>" + value + "</a>";
  }
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
  function setTitle(value,row,index){
    return "<a href='/project/product/"+row.Id+"'>" + value + "</a>";
  }
  function setArticle(value,row,index){
    // return '<a class="article" href="javascript:void(0)" title="article"><i class="fa fa-file-text-o"></i></a>';
    if (value){
      if (value.length==1){//'<a href="/project/product/article/'
        articleUrl= '<a href="'+value[0].Link+'/'+value[0].Id+'" title="查看" target="_blank"><i class="fa fa-file-text-o"></i></a>';
        return articleUrl;
      }else if(value.length==0){
                    
      }else if(value.length>1){
        articleUrl= "<a class='article' href='javascript:void(0)' title='查看文章列表'><i class='fa fa-list-ol'></i></a>";
        return articleUrl;
      }
    }
  }

  function setAttachment(value,row,index){
    if (value){
      if (value.length==1){
        attachUrl= '<a href="'+value[0].Link+'/'+value[0].Title+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
        return attachUrl;
      }else if(value.length==0){
                    
      }else if(value.length>1){
        attachUrl= "<a class='attachment' href='javascript:void(0)' title='查看附件列表'><i class='fa fa-list-ol'></i></a>";
        return attachUrl;
      }
    }
  }
  function setPdf(value,row,index){
    if (value){
      if (value.length==1){
        pdfUrl= '<a href="'+value[0].Link+'/'+value[0].Title+'" title="打开pdf" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
        return pdfUrl;
      }else if(value.length==0){
                    
      }else if(value.length>1){
        pdfUrl= "<a class='pdf' href='javascript:void(0)' title='查看pdf列表'><i class='fa fa-list-ol'></i></a>";
        return pdfUrl;
      }
    }
  }

  window.actionEvents = {
    'click .article': function (e, value, row, index) {
      var site=/http:\/\/.*?\//.exec(value[1].Link);//非贪婪模式 
      if (site){
        $('#articles').bootstrapTable('refresh', {url:'/project/product/syncharticles?site='+site+'&id='+row.Id});
      }else{
        $('#articles').bootstrapTable('refresh', {url:'/project/product/articles/'+row.Id});
      }
      $('#modalarticle').modal({
        show:true,
        backdrop:'static'
      }); 
    },
    'click .attachment': function (e, value, row, index) {
      // for(var i=0;i<value.length;i++)
      // alert(value[i].Link);
      // var ret=/http:(.*)\:/.exec(value[i].Link);//http://127.0.0.1:
      var site=/http:\/\/.*?\//.exec(value[1].Link);//非贪婪模式 
      if (site){//跨域
        // alert("1");
        // $.getJSON(ret+'project/product/attachment/'+row.Id,function(){
          // $('#attachs').bootstrapTable('load', randomData());
        // })
        $('#attachs').bootstrapTable('refresh', {url:'/project/product/synchattachment?site='+site+'&id='+row.Id});
        // $('#attachs').bootstrapTable('refresh', {url:site+'project/product/attachment/'+row.Id});
      }else{
        // alert("2");
        $('#attachs').bootstrapTable('refresh', {url:'/project/product/attachment/'+row.Id});
        }
        $('#modalattach').modal({
          show:true,
          backdrop:'static'
        });
    },

    'click .pdf': function (e, value, row, index) {
      var site=/http:\/\/.*?\//.exec(value[1].Link);//非贪婪模式 
      if (site){//跨域
        $('#pdfs').bootstrapTable('refresh', {url:'/project/product/synchpdf?site='+site+'&id='+row.Id});
      }else{
        $('#pdfs').bootstrapTable('refresh', {url:'/project/product/pdf/'+row.Id});
      }
      $('#modalpdf').modal({
        show:true,
        backdrop:'static'
      }); 
    },
  };

  //最后面弹出文章列表中用的_根据上面的click，弹出模态框，给模态框中的链接赋值
  function setArticlecontent(value,row,index){
    articleUrl= '<a href="'+value+'" title="下载" target="_blank"><i class="fa fa-file-text-o"></i></a>';
      return articleUrl;
  }
  //最后面弹出附件列表中用的
  function setAttachlink(value,row,index){
    attachUrl= '<a href="'+value+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
      return attachUrl;
  }
  //最后面弹出pdf列表中用的
  function setPdflink(value,row,index){
    pdfUrl= '<a href="'+value+'" title="下载" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
      return pdfUrl;
  }  
</script>

<!-- 文章列表 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalarticle">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">文章列表</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!-- <div id="pdfs" style="display:none"> -->
                <!-- <h3>工程目录分级</h3> -->
                <table id="articles"
                      data-toggle="table"
                      data-page-size="5"
                      data-page-list="[5, 25, 50, All]"
                      data-unique-id="id"
                      data-pagination="true"
                      data-side-pagination="client"
                      data-click-to-select="true">
                    <thead>     
                    <tr>
                      <th data-width="10" data-checkbox="true"></th>
                      <th data-formatter="index1">#</th>
                      <th data-field="Title">名称</th>
                      <th data-field="Subtext">副标题</th>
                      <th data-field="Link" data-formatter="setArticlecontent">查看</th>
                      <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
                      <th data-field="Updated" data-formatter="localDateFormatter">修改时间</th>
                    </tr>
                  </thead>
                </table>
              <!-- </div> -->
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- 除了**pdf**之外的附件列表 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalattach">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">非PDF附件列表</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!-- <div id="pdfs" style="display:none"> -->
                <!-- <h3>工程目录分级</h3> -->
                <table id="attachs"
                      data-toggle="table"
                      data-page-size="5"
                      data-page-list="[5, 25, 50, All]"
                      data-unique-id="id"
                      data-pagination="true"
                      data-side-pagination="client"
                      data-click-to-select="true">
                    <thead>     
                    <tr>
                      <th data-width="10" data-checkbox="true"></th>
                      <th data-formatter="index1">#</th>
                      <th data-field="Title">名称</th>
                      <th data-field="FileSize">大小</th>
                      <th data-field="Link" data-formatter="setAttachlink">下载</th>
                      <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
                      <th data-field="Updated" data-formatter="localDateFormatter">修改时间</th>
                    </tr>
                  </thead>
                </table>
              <!-- </div> -->
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- pdf附件列表 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalpdf">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">pdf附件列表</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <!-- <div id="pdfs" style="display:none"> -->
                <!-- <h3>工程目录分级</h3> -->
                <table id="pdfs"
                      data-toggle="table"
                      data-page-size="5"
                      data-page-list="[5, 25, 50, All]"
                      data-unique-id="id"
                      data-pagination="true"
                      data-side-pagination="client"
                      data-click-to-select="true">
                    <thead>     
                    <tr>
                      <th data-width="10" data-checkbox="true"></th>
                      <th data-formatter="index1">#</th>
                      <th data-field="Title">名称</th>
                      <th data-field="FileSize">大小</th>
                      <th data-field="Link" data-formatter="setPdflink">下载</th>
                      <th data-field="Created" data-formatter="localDateFormatter">建立时间</th>
                      <th data-field="Updated" data-formatter="localDateFormatter">修改时间</th>
                    </tr>
                  </thead>
                </table>
              <!-- </div> -->
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>

</body>
</html>