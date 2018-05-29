<!-- 具体一个项目侧栏id下所有成果，不含子目录下的成果 -->
<!DOCTYPE html>
{{template "header"}}
<title>项目成果-EngiCMS</title>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
    <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
  <style type="text/css">
    /*模态框效果*/
    #modalDialog .modal-header {cursor: move;}
    #modalDialog1 .modal-header {cursor: move;}
    #modalDialog2 .modal-header {cursor: move;}
    #modalDialog3 .modal-header {cursor: move;}
    #modalDialog4 .modal-header {cursor: move;}
    #modalDialog5 .modal-header {cursor: move;}
    #modalDialog6 .modal-header {cursor: move;}
    #modalDialog7 .modal-header {cursor: move;}
  </style>
</head>
  <div class="container-fill">{{template "navbar" .}}</div>

<body>
<div class="col-lg-12">
  <!-- 面包屑导航 -->
  <!-- <div class="breadcrumbs">
    <ol class="breadcrumb" split="&gt;">
      <li>
        <i class="fa fa-home" aria-hidden="true">项目编号-名称：</i><a href="javascript:void(0)"> 
          {{.Category.Code}}-{{.Category.Title}}
        </a>/<a href="javascript:void(0)">成果列表</a>
      </li>
    </ol>
  </div> -->

  <div id="toolbar1" class="btn-group">
        <button type="button" data-name="editorProdButton" id="editorProdButton" class="btn btn-default"> <i class="fa fa-edit" title="修改成果信息">编辑</i>
        </button>
        <button type="button" data-name="editorAttachButton" id="editorAttachButton" class="btn btn-default"> <i class="fa fa-edit" title="修改成果附件">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
  </div>
  <!--data-click-to-select="true" -->
  <table id="table0" 
        data-toggle="table" 
        data-url="/project/products/all/{{.Id}}"
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
        <!-- <th data-field="Id">编号</th> data-visible="false" -->
        <th data-field="Code" data-halign="center">编号</th>
        <th data-field="Title" data-halign="center">名称</th>
        <th data-field="Label" data-formatter="setLable" data-halign="center" data-align="center">关键字</th>
        <th data-field="Principal" data-halign="center" data-align="center">设计</th>
        <th data-field="Articlecontent" data-formatter="setArticle" data-events="actionEvents" data-halign="center" data-align="center">文章</th>
        <th data-field="Attachmentlink" data-formatter="setAttachment" data-events="actionEvents" data-halign="center" data-align="center">附件</th>
        <th data-field="Pdflink" data-formatter="setPdf" data-events="actionEvents" data-halign="center" data-align="center">PDF</th>
        <th data-field="Created" data-formatter="localDateFormatter" data-halign="center" data-visible="false" data-align="center">建立时间</th>
        <th data-field="Updated" data-formatter="localDateFormatter" data-halign="center" data-align="center">更新时间</th>
        <th data-field="Relevancy" data-formatter="RelevFormatter" events="actionRelevancy" data-halign="center">关联</th>
      </tr>
    </thead>
  </table>

  <script type="text/javascript">
  function index1(value,row,index){
    return index+1
  }

  function localDateFormatter(value) {
    return moment(value, 'YYYY-MM-DD').format('YYYY-MM-DD');
  }

  function RelevFormatter(value) {
    if (value){
      if (value.length==1){//'<a href="/project/product/article/'
        var array=value[0].Relevancy.split(",")
        var relevarray = new Array() 
        for (i=0;i<array.length;i++)
        {
          relevarray[i]=array[i];
        }
        return relevarray.join(",");
        // articleUrl= '<a href="'+value[0].Link+'/'+value[0].Id+'" title="查看" target="_blank"><i class="fa fa-file-text-o"></i></a>';
        // return articleUrl;
      }else if(value.length==0){
                    
      }else if(value.length>1){
        var relevarray = new Array()
        for (i=0;i<value.length;i++)
          {
            relevarray[i]=value[i].Relevancy;
          }
        return relevarray.join(",");
        // articleUrl= "<a class='article' href='javascript:void(0)' title='查看文章列表'><i class='fa fa-list-ol'></i></a>";
        // return articleUrl;
      }
    }
  }

  function setCode(value,row,index){
    return "<a href='/project/product/attachment/"+row.Id+"'>" + value + "</a>";
  }
  function setLable(value,row,index){
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
  function setCodetest(value,row,index){
    //保留，数组和字符串以及循环的处理
    // array=value.split(",")
    // var labelarray = new Array() 
    // for (i=0;i<value.length;i++)//value是数组"Code":[数组"SL0001-510-08","SL0001-510-08"],
    // {
    //   labelarray[i]="<a href='/project/product/attachment/"+value[i]+"'>" + value[i] + "</a>";
    // }
    // if (value.match(",")!=null){
    if (value){
      array=value.split(",")
      var labelarray = new Array() 
      for (i=0;i<array.length;i++)
      {
        labelarray[i]="<a href='/project/product/attachment/"+array[i]+"'>" + array[i] + "</a>";
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
        attachUrl= '<a href="/attachment?id='+value[0].Id+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
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
        pdfUrl= '<a href="/pdf?id='+value[0].Id+'" title="打开pdf" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
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
    attachUrl= '<a href="/attachment?id='+row.Id+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
      return attachUrl;
  }
  //最后面弹出pdf列表中用的
  function setPdflink(value,row,index){
    pdfUrl= '<a href="/pdf?id='+row.Id+'" title="下载" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
      return pdfUrl;
  }

  // 编辑成果信息
  $("#editorProdButton").click(function() {
      // if ({{.role}}!=1){
      //   alert("权限不够！");
      //   return;
      // }
      if ({{.RoleAdd}}!="true"){
        alert("权限不够！");
        return;
      }
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选成果！");
        return;
      }
      if (selectRow.length>1){
        alert("请不要勾选一个以上成果！");
        return;
      }
      // alert(selectRow[0].Uid=={{.Uid}});
      // alert({{.Uid}});
      if (selectRow[0].Uid==={{.Uid}}||{{.RoleUpdate}}=="true"){
      
        if (selectRow[0].Attachmentlink[0]){//||selectRow[0].Pdflink[0].Link||selectRow[0].Articlecontent[0].Link)
        var site=/http:\/\/.*?\//.exec(selectRow[0].Attachmentlink[0].Link);//非贪婪模式 
        }
        if (selectRow[0].Articlecontent[0]){
        var site=/http:\/\/.*?\//.exec(selectRow[0].Articlecontent[0].Link);//非贪婪模式 
        }
        if (selectRow[0].Pdflink[0]){
        var site=/http:\/\/.*?\//.exec(selectRow[0].Pdflink[0].Link);//非贪婪模式 
        }
        if (site){
          alert("同步成果不允许！");
          return;
        }

        $("input#cid").remove();
        var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
        $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
        $("#prodcode3").val(selectRow[0].Code);
        $("#prodname3").val(selectRow[0].Title);
        $("#prodlabel3").val(selectRow[0].Label);
        $("#prodprincipal3").val(selectRow[0].Principal);

        $('#modalProdEditor').modal({
        show:true,
        backdrop:'static'
        });

      }else{
        alert("权限不够！"+selectRow[0].Uid);
        return;
      }
  })

  // 编辑成果附件——删除附件、文章或追加附件
  var selectrowid;
  $("#editorAttachButton").click(function() {
      // if ({{.role}}!=1){
      //   alert("权限不够！");
      //   return;
      // }
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选成果！");
        return;
      }
      if (selectRow.length>1){
      alert("请不要勾选一个以上成果！");
      return;
      }

      if (selectRow[0].Uid==={{.Uid}}||{{.RoleDelete}}=="true"){

      if (selectRow[0].Attachmentlink[0]){//||selectRow[0].Pdflink[0].Link||selectRow[0].Articlecontent[0].Link)
      var site=/http:\/\/.*?\//.exec(selectRow[0].Attachmentlink[0].Link);//非贪婪模式 
      }
      if (selectRow[0].Articlecontent[0]){
      var site=/http:\/\/.*?\//.exec(selectRow[0].Articlecontent[0].Link);//非贪婪模式 
      }
      if (selectRow[0].Pdflink[0]){
      var site=/http:\/\/.*?\//.exec(selectRow[0].Pdflink[0].Link);//非贪婪模式 
      }
      if (site){
        alert("同步成果不允许！");
        return;
      }
      selectrowid=selectRow[0].Id;
      $("input#pid").remove();
      var th1="<input id='pid' type='hidden' name='pid' value='" +selectRow[0].Id+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $('#attachments').bootstrapTable('refresh', {url:'/project/product/allattachments/'+selectRow[0].Id});//取得所有附件列表和文章列表
      $('#modalAttachEditor').modal({
      show:true,
      backdrop:'static'
      });

      }else{
        alert("权限不够！"+selectRow[0].Uid);
        return;
      }
  })

  $(document).ready(function() {
    var uploader;
    $('#modalAttachEditor').on('shown.bs.modal',function(){
      // var $ = jQuery,
      $list2 = $('#thelist2'),
      $btn = $('#ctlBtn2'),
      state = 'pending',
      // uploader;
      uploader = WebUploader.create({
        // 不压缩image
        resize: false,
        // swf文件路径
        swf: '/static/js/Uploader.swf',
        // 文件接收服务端。
        server: '/project/product/updateattachment',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#picker2'
      });
      // 当有文件添加进来的时候
      uploader.on( 'fileQueued', function( file ) {
        $list2.append( '<div id="' + file.id + '" class="item">' +
              '<h4 class="info">' + file.name + '</h4>' +
              '<p class="state">等待上传...</p>' +
          '</div>' );
      }); 

      //传递参数——成果id
      uploader.on( 'startUpload', function() {//uploadBeforeSend——这个居然不行？
        var pid = $('#pid').val();
        uploader.option('formData', {
          "pid":pid,
        });        
      });

      // 文件上传过程中创建进度条实时显示。
      uploader.on( 'uploadProgress', function( file, percentage ) {
        var $li = $( '#'+file.id ),
        $percent = $li.find('.progress .progress-bar');
        // 避免重复创建
        if ( !$percent.length ) {
          $percent = $('<div class="progress progress-striped active">' +
                '<div class="progress-bar" role="progressbar" style="width: 0%">' +
              '</div>' +
              '</div>').appendTo( $li ).find('.progress-bar');
        }
        $li.find('p.state').text('上传中');
        $percent.css( 'width', percentage * 100 + '%' );
      });

      uploader.on( 'uploadSuccess', function( file ) {
          $( '#'+file.id ).find('p.state').text('已上传');
      });

      uploader.on( 'uploadError', function( file ) {
          $( '#'+file.id ).find('p.state').text('上传出错');
      });

      uploader.on( 'uploadComplete', function( file ) {
          $( '#'+file.id ).find('.progress').fadeOut();
          $('#attachments').bootstrapTable('refresh', {url:'/project/product/allattachments/'+selectrowid});
          $('#table0').bootstrapTable('refresh', {url:'/project/products/'+{{.Id}}});
      });

      uploader.on( 'all', function( type ) {
        if ( type === 'startUpload' ) {
            state = 'uploading';
        } else if ( type === 'stopUpload' ) {
            state = 'paused';
        } else if ( type === 'uploadFinished' ) {
            state = 'done';
        }
        if ( state === 'uploading' ) {
            $btn.text('暂停上传');
        } else {
            $btn.text('开始上传');
        }
      });
    
      $btn.on( 'click', function() {
          if ( state === 'uploading' ) {
              uploader.stop();
          } else {
              uploader.upload();
          }
      });
    })

    $('#modalAttachEditor').on('hide.bs.modal',function(){
      $list2.text("");
      uploader.destroy();//销毁uploader
    })
  })
    
  // 删除成果
  $("#deleteButton").click(function() {
      // if ({{.role}}!=1){
      //   alert("权限不够！");
      //   return;
      // }
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选成果！");
        return false;
      }
     //问题：如果多选，而其中有自己的，也有自己不具备权限的********
      if (selectRow[0].Uid==={{.Uid}}||{{.RoleDelete}}=="true"){
        

      if (selectRow[0].Attachmentlink[0]){//||selectRow[0].Pdflink[0].Link||selectRow[0].Articlecontent[0].Link)
      var site=/http:\/\/.*?\//.exec(selectRow[0].Attachmentlink[0].Link);//非贪婪模式 
      }
      if (selectRow[0].Articlecontent[0]){
      var site=/http:\/\/.*?\//.exec(selectRow[0].Articlecontent[0].Link);//非贪婪模式 
      }
      if (selectRow[0].Pdflink[0]){
      var site=/http:\/\/.*?\//.exec(selectRow[0].Pdflink[0].Link);//非贪婪模式 
      }
      if (site){
        alert("同步成果不允许！");
        return;
      }
      if(confirm("确定删除成果吗？一旦删除将无法恢复！")){
        var title=$.map(selectRow,function(row){
          return row.Title;
        })
        var ids="";
        for(var i=0;i<selectRow.length;i++){
          if(i==0){
            ids=selectRow[i].Id;
          }else{
            ids=ids+","+selectRow[i].Id;
          }  
        }
        $.ajax({
          type:"post",
          url:"/project/product/deleteproduct",
          data: {ids:ids},
          success:function(data,status){
            alert("删除“"+data+"”成功！(status:"+status+".)");
            //删除已选数据
            $('#table0').bootstrapTable('remove',{
              field:'Title',
              values:title
            });
          }
        });
      }

      }else{
        alert("权限不够！"+selectRow[0].Uid);
        return;
      }  
  })

  </script>
  <!-- 文章列表 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalarticle">
      <div class="modal-dialog" id="modalDialog3">
        <div class="modal-content">
          <div class="modal-header" style="background-color: #8bc34a">
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
      <div class="modal-dialog" id="modalDialog4">
        <div class="modal-content">
          <div class="modal-header" style="background-color: #8bc34a">
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
      <div class="modal-dialog" id="modalDialog5">
        <div class="modal-content">
          <div class="modal-header" style="background-color: #8bc34a">
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
                      <th data-field="Link" data-formatter="setPdflink">查看</th>
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
  <!-- 编辑成果名称等信息 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalProdEditor">
      <div class="modal-dialog" id="modalDialog6">
        <div class="modal-content">
          <div class="modal-header" style="background-color: #8bc34a">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">编辑成果信息</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="prodcode3" name="prodcode3"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">标题</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodname3" name="prodname3"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">关键字</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodlabel3" name="prodlabel3" placeholder="以英文,号分割"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">设计</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodprincipal3" name="prodprincipal3"></div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="updateprod()">保存</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- 编辑成果附件 删除附件或追加附件-->
  <div class="form-horizontal">
    <div class="modal fade" id="modalAttachEditor">
      <div class="modal-dialog" id="modalDialog7">
        <div class="modal-content">
          <div class="modal-header" style="background-color: #8bc34a">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">编辑成果附件</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div id="attachtoolbar" class="btn-group">
                <button type="button" data-name="deleteAttachButton" id="deleteAttachButton" class="btn btn-default">
                <i class="fa fa-trash">删除</i>
                </button>
              </div>
              <table id="attachments"
                    data-toggle="table"
                    data-toolbar="#attachtoolbar"
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
              <!--SWF在初始化的时候指定，在后面将展示-->
              <div id="uploader1" style="position:relative;text-align: center;">
              <!--用来存放文件信息-->
                <div id="thelist2"></div>
                <div class="btns">
                  <div id="picker2">选择文件</div>
                  <button id="ctlBtn2" class="btn btn-default">开始上传</button>
                </div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>

</div>

<script type="text/javascript">
  //实例化编辑器
    // var ue = UE.getEditor('container', {
    //   autoHeightEnabled: true,
    //   autoFloatEnabled: false,
    //   // topOffset:100,
    //   initialFrameWidth:'100%'
    // });
  /* 2.传入参数表,添加到已有参数表里 通过携带参数，实现不同的页面使用不同controllers*/
    // ue.ready(function () {
    //     ue.addListener('focus', function () {//startUpload start-upload startUpload beforeExecCommand是在插入图片之前触发
    //         var pid = $('#pid').val();
    //         // var html = ue.getContent();
    //         ue.execCommand('serverparam', {
    //           "pid":pid 
    //         });
    //     });
    // });

  //添加文章
  function save2(){
    // var radio =$("input[type='radio']:checked").val();
    var projectid = $('#pid').val();
    var prodcode = $('#prodcode1').val();
    var prodname = $('#prodname1').val();
    var subtext = $('#subtext1').val();
    var prodprincipal = $('#prodprincipal2').val();
    var prodlabel = $('#prodlabel2').val();
    var relevancy = $('#relevancy').val();
    var html = ue.getContent();
    // $('#myModal').on('hide.bs.modal', function () {  
    if (prodname&&prodcode){  
      $.ajax({
        type:"post",
        url:"/project/product/addarticle",
        data: {pid:projectid,code:prodcode,title:prodname,subtext:subtext,label:prodlabel,content:html,principal:prodprincipal,relevancy:relevancy},//父级id
        success:function(data,status){
          alert("添加“"+data+"”成功！(status:"+status+".)");
          $('#modalTable2').modal('hide');
          $('#table0').bootstrapTable('refresh', {url:'/project/products/'+{{.Id}}});
        },
        
      });
    }else{
      alert("请填写编号和名称！");
      return;
    }
  }
  // 编辑成果信息
  function updateprod(){
    // var radio =$("input[type='radio']:checked").val();
    var projectid = $('#cid').val();
    var prodcode = $('#prodcode3').val();
    var prodname = $('#prodname3').val();
    var prodlabel = $('#prodlabel3').val();
    var prodprincipal = $('#prodprincipal3').val();
 
    if (prodname&&prodcode){  
      $.ajax({
        type:"post",
        url:"/project/product/updateproduct",
        data: {pid:projectid,code:prodcode,title:prodname,label:prodlabel,principal:prodprincipal},//父级id
        success:function(data,status){
          alert("添加“"+data+"”成功！(status:"+status+".)");
          $('#modalProdEditor').modal('hide');
          $('#table0').bootstrapTable('refresh', {url:'/project/products/'+{{.Id}}});
        },
        
      });
    }else{
      alert("请填写编号和名称！");
      return;
    }
  }
  // 删除附件
  $("#deleteAttachButton").click(function() {
      // if ({{.role}}!=1){
      //   alert("权限不够！");
      //   return;
      // }
      if ({{.RoleDelete}}!="true"){
        alert("权限不够！");
        return;
      }
      var selectRow=$('#attachments').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选！");
        return false;
      }

      if ({{.RoleDelete}}!="true"){
        alert("权限不够！"+selectRow[0].Uid);
        return;
      }

      if(confirm("确定删除吗？一旦删除将无法恢复！")){
        var title=$.map(selectRow,function(row){
          return row.Title;
        })
        var ids="";
        for(var i=0;i<selectRow.length;i++){
          if(i==0){
            ids=selectRow[i].Id;
          }else{
            ids=ids+","+selectRow[i].Id;
          }  
        }
        $.ajax({
          type:"post",
          url:"/project/product/deleteattachment",
          data: {ids:ids},
          success:function(data,status){
            alert("删除“"+data+"”成功！(status:"+status+".)");
            //删除已选数据
            $('#attachments').bootstrapTable('remove',{
              field:'Title',
              values:title
            });
          }
        });
      }  
  })

    //******表格追加项目同步ip中的数据*******
    $(function () {
        $('#synchIP').click(function () {
          // alert("ha ");
          $.ajax({
            type:"get",
            url:"/project/synchproducts/"+{{.Id}},
            // data: {ids:ids},
            success:function(data,status){
              alert("同步成功！(status:"+status+".)");
              //追加数据
              $('#table0').bootstrapTable('append', data);
              $('#table0').bootstrapTable('scrollTo', 'bottom');
            }
          });
        });
    });

    function randomData() {
        var startId = ~~(Math.random() * 100),
                rows = [];
        for (var i = 0; i < 10; i++) {
            rows.push({
              Id: startId + i,
              Code:startId + i,
              Title: 'test' + (startId + i),
            });
        }
        return rows;
    }

    //勾选后输入框可用

    function station_select(){ 
      if(box.checked){ 
        document.getElementById("relevancy").disabled=false; 
      } else{ 
        document.getElementById("relevancy").disabled=true; 
      } 
    }

    $(document).ready(function(){
        $("#modalDialog").draggable({ handle: ".modal-header" });//为模态对话框添加拖拽
        $("#modalDialog1").draggable({ handle: ".modal-header" });
        $("#modalDialog2").draggable({ handle: ".modal-header" });
        $("#modalDialog3").draggable({ handle: ".modal-header" });
        $("#modalDialog4").draggable({ handle: ".modal-header" });
        $("#modalDialog5").draggable({ handle: ".modal-header" });
        $("#modalDialog6").draggable({ handle: ".modal-header" });
        $("#modalDialog7").draggable({ handle: ".modal-header" });
        $("#myModal").css("overflow", "hidden");//禁止模态对话框的半透明背景滚动
    })
</script>

</body>
</html>