<!-- 具体一个项目侧栏id下所有成果，不含子目录下的成果 -->
<!DOCTYPE html>
<head>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>

  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>

  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>

  <script src="/static/js/tableExport.js"></script>

  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
  <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>

  <script type="text/javascript" src="/static/js/moment.min.js"></script>

  <link rel="stylesheet" type="text/css" href="/static/fex-team-webuploader/css/webuploader.css">
  <script type="text/javascript" src="/static/fex-team-webuploader/dist/webuploader.min.js"></script>
</head>

<body>

<div class="col-lg-12">
  <h3>成果列表</h3>
<div id="toolbar1" class="btn-group">
        <!-- 多文件批量上传 -->
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default" title="批量上传模式"> <i class="fa fa-plus">添加</i>
        </button>
        <!-- 多附件上传 -->
        <button type="button" data-name="addButton1" id="addButton1" class="btn btn-default"> <i class="fa fa-plus-square-o" title="多附件模式">添加</i>
        </button>
        <!-- 添加文章 -->
        <button type="button" data-name="addButton2" id="addButton2" class="btn btn-default"> <i class="fa fa-plus-square" title="文章模式">添加</i>
        </button>
        <!-- <button type="button" data-name="editorProdButton" id="editorProdButton" class="btn btn-default"> <i class="fa fa-edit" title="修改成果信息">编辑</i>
        </button>
        <button type="button" data-name="editorAttachButton" id="editorAttachButton" class="btn btn-default"> <i class="fa fa-edit" title="修改成果附件">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
        <button type="button" data-name="synchIP" id="synchIP" class="btn btn-default">
        <i class="fa fa-refresh">同步</i>
        </button> -->
</div>
<!--data-click-to-select="true" -->
<table id="table0" 
        data-toggle="table" 
        data-url="/project/products/{{.Id}}"
        data-search="true"
        data-show-refresh="true"
        
        
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
        <!-- <th data-field="Created" data-formatter="localDateFormatter">建立时间</th> -->
        <th data-field="Updated" data-visible="false" data-formatter="localDateFormatter">更新时间</th>
        <!-- <th data-field="Created" data-formatter="actionFormatter" events="actionEvents">操作</th> -->
      </tr>
    </thead>
</table>

<script type="text/javascript">
  /*数据json使用json数据要删除data-toggle="table"*/
  // var json =
     // 保留***[{"Id":"1","Code":[这个数组也行"SL0001-510-08","SL0001-510-08"],"Title":"水利枢纽布置图","Label":"水电站","Principal":"秦晓川","Product":"8","Created":"2016-11-26"},

        /*初始化table数据*/
        // $(function(){
        //     $("#tabletest").bootstrapTable({
        //         data:json
                // onClickRow: function (row, $element) {
                //   alert( "选择了行Id为: " + row.Id );
                //   rowid=row.Id//全局变量
                //   $('#table1').bootstrapTable('refresh', {url:'/admincategory?pid='+row.Id});
                // }
        //     });
        // });
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
  // }else{
    // return "<a href='/project/product/attachment/"+value+"'>" + value + "</a>";
  // }
    // var x   
    // for (x in array)
    // {
    // array[x];
    // }
          // <a href="/category?op=viewlabel&label=value">
          // <span class="label label-info"></span></a>
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
// var bb;
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
          // $.ajax({//这种同步加载行不通，会混乱。异步又无法传出返回值data
            // type:"get",//这里是否一定要用post？？？
            // url:"/project/product/attachment/"+row.Id,
            // data: {CatalogId:row.Id},
            // dataType:'json',
            // async:false,//必须加这个异步为否定，即同步，否则bb传不出去
            // success:function(data,status){//数据提交成功时返回数据
                // $.each(data,function(i,d){
                //     $("#cars").append('<option value="' + data[i].Username + '"></option>');
                  // });
          //         if (data.length==1){
          //           bb= '<a href="'+data[0].Link+'/'+data[0].Title+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
          //         }else if(data.length==0){
                    
          //         }else if(data.length>1){
          //           bb= '<a class="attachment" href="javascript:void(0)" title="查看附件列表"><i class="fa fa-list-ol"></i></a>';
          //         } 
          //   }
          // });
          // return bb; 
    // return '<a class="attachment" href="javascript:void(0)" title="attachment"><i class="fa fa-paperclip"></i></a>';
  }

  // var pdfUrl;+'&file='+value[0].Link+'/'+value[0].Title
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
    attachUrl= '<a href="'+value+'" title="下载" target="_blank"><i class="fa fa-paperclip"></i></a>';
      return attachUrl;
  }
  //最后面弹出pdf列表中用的'&file='+value+
  function setPdflink(value,row,index){
    pdfUrl= '<a href="/pdf?id='+row.Id+'" title="下载" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
      return pdfUrl;
  }

    // 批量上传
    $("#addButton").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！"+{{.role}});
        return;
      }
      $("input#pid").remove();
      var th1="<input id='pid' type='hidden' name='pid' value='" +{{.Id}}+"'/>"
        $(".modal-body").append(th1);

        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
    })

  $(document).ready(function() {
    $list1 = $('#thelist');
    $btn = $('#ctlBtn');
    state = 'pending';
    // $('#modalTable').on('shown.bs.modal',function(e){
      var uploader=WebUploader.create({
        // 不压缩image
        resize: false,
        // swf文件路径
        swf: '/static/fex-team-webuploader/dist/Uploader.swf',
        // 文件接收服务端。
        server: '/project/product/addattachment',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#picker'
      });

      // 当有文件添加进来的时候
      uploader.on( 'fileQueued', function( file ) {
       $list1.append( '<div id="' + file.id + '" class="item">' +
            '<h4 class="info">' + file.name + '</h4>' +
            '<p class="state">等待上传...</p>' +
        '</div>' );
      }); 

      //传递参数——成果id
      uploader.on( 'startUpload', function() {//uploadBeforeSend——这个居然不行？
      // if (prodlabel){
        var pid = $('#pid').val();
        var prodlabel = $('#prodlabel').val();
        var prodprincipal = $('#prodprincipal').val();
        uploader.option('formData', {
          "pid":pid,
          "prodlabel":prodlabel,
          "prodprincipal":prodprincipal
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

      $('#picker').mouseenter(function(){
        uploader.refresh();
      })

      $('#modalTable').on('hide.bs.modal',function(){
        $list1.text("");
        // uploader.destroy();//销毁uploader
      })
    // });
  })      

    // 多附件模式
    $("#addButton1").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      $("input#pid").remove();
      var th1="<input id='pid' type='hidden' name='pid' value='" +{{.Id}}+"'/>"
        $(".modal-body").append(th1);
        $('#modalTable1').modal({
          show:true,
          backdrop:'static'
        });
    })

  $(document).ready(function() {
      $list = $('#thelist1');
      $btn = $('#ctlBtn1');
      state = 'pending';
    $('#modalTable1').on('shown.bs.modal',function(e){
      var uploader=WebUploader.create({
        // 不压缩image
        resize: false,
        // swf文件路径
        swf: '/static/fex-team-webuploader/dist/Uploader.swf',
        // 文件接收服务端。
        server: '/project/product/addattachment2',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#picker1'
      });

      // 当有文件添加进来的时候
      uploader.on( 'fileQueued', function( file ) {
       $list.append( '<div id="' + file.id + '" class="item">' +
            '<h4 class="info">' + file.name + '</h4>' +
            '<p class="state">等待上传...</p>' +
        '</div>' );
      }); 

          //传递参数——成果id
      uploader.on( 'startUpload', function() {//uploadBeforeSend——这个居然不行？
        var pid = $('#pid').val();
        var prodcode = $('#prodcode').val();
        var prodname = $('#prodname').val();
        var prodlabel = $('#prodlabel1').val();
        var prodprincipal = $('#prodprincipal1').val();
        // var html = ue.getContent();
        // alert(html);
        uploader.option('formData', {
          "pid":pid,
          "prodcode":prodcode,
          "prodname":prodname,
          "prodlabel":prodlabel,
          "prodprincipal":prodprincipal
          // 'content':html,
          // {'tnumber':a,'title':b,'categoryid':c,'category':d,'content':e}
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
        var prodcode = $('#prodcode').val();
        var prodname = $('#prodname').val();
        if (prodcode&&prodname){
          if ( state === 'uploading' ) {
              uploader.stop();
          } else {
              uploader.upload();
          }
        }else{
            alert("编号和名称不能为空"+prodcode+prodname);
        }
      });

      $('#modalTable1').on('hide.bs.modal',function(){
        $list.text("");
        uploader.destroy();//销毁uploader
      })
    });
  }) 

    //****添加文章
    $("#addButton2").click(function() {
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      $("input#pid").remove();
      var th1="<input id='pid' type='hidden' name='pid' value='" +{{.Id}}+"'/>"
        $(".modal-body").append(th1);

        $('#modalTable2').modal({
        show:true,
        backdrop:'static'
        });
    })


    // 编辑成果信息
    $("#editorProdButton").click(function() {
      if ({{.role}}!=1){
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
    })

    // 编辑成果附件——删除附件、文章或追加附件
    var selectrowid;
    $("#editorAttachButton").click(function() {
      if ({{.role}}!=1){
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
        swf: '/static/fex-team-webuploader/dist/Uploader.swf',
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
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选成果！");
        return false;
      }
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
    })


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
  <!-- 批量上传 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">批量添加成果</h3>
            <label>**请选择标准命名的电子文件上传**</label>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
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
              <!-- <div style="display: inline-block;">
                <span id="filePicker" onclick="create()">上传</span>
                <span id="responseText" style="display: inline-block;"></span>
              </div> -->
              <!--SWF在初始化的时候指定，在后面将展示-->
              <div id="uploader" style="position:relative;text-align: center;">
              <!--用来存放文件信息-->
                <div id="thelist"></div>
                <div class="btns">
                  <div id="picker">选择文件</div>
                  <button id="ctlBtn" class="btn btn-default">开始上传</button>
                </div>
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
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            <!-- <button type="button" class="btn btn-primary" onclick="save()">保存</button> -->
          </div>
        </div>
      </div>
    </div>
  </div>

<!-- 多附件 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加成果——多附件模式</h3>
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
                  <input type="tel" class="form-control" id="prodlabel1" name="prodlabel1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">设计</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodprincipal1" name="prodprincipal1"></div>
              </div>
              <!--SWF在初始化的时候指定，在后面将展示-->
              <div id="uploader1" style="position:relative;text-align: center;">
              <!--用来存放文件信息-->
                <div id="thelist1"></div>
                <div class="btns1">
                  <div id="picker1">选择文件</div>
                  <button id="ctlBtn1" class="btn btn-default">开始上传</button>
                </div>
              </div>
            </div>
            <!-- <label>文件简介:</label>
              <div>
                <script id="container" type="text/plain" style="height:200px;"></script>width:1024px;
              </div> -->
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            <!-- <button type="button" class="btn btn-primary" onclick="save1()">保存</button> -->
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- 添加文章 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalTable2">
      <div class="modal-dialog" style="width: 100%">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加文章</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">编号</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="prodcode1" name="prodcode1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">标题</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodname1" name="prodname1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">副标题</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="subtext1" name="subtext1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">关键字</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodlabel2" name="prodlabel2"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">设计</label>
                <div class="col-sm-7">
                  <input type="tel" class="form-control" id="prodprincipal2" name="prodprincipal2"></div>
              </div>
            </div>
            <label>文章正文:</label>
              <div>
                <script id="container" type="text/plain" style="height:200px;"></script><!-- width:1024px; -->
              </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="save2()">保存</button>
          </div>
        </div>
      </div>
    </div>
  </div>
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
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
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
                  <input type="tel" class="form-control" id="prodlabel3" name="prodlabel3"></div>
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
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
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
    var ue = UE.getEditor('container', {
      autoHeightEnabled: true,
      autoFloatEnabled: true
    });
  /* 2.传入参数表,添加到已有参数表里 通过携带参数，实现不同的页面使用不同controllers*/
    ue.ready(function () {
        ue.addListener('focus', function () {//startUpload start-upload startUpload beforeExecCommand是在插入图片之前触发
            var pid = $('#pid').val();
            // var html = ue.getContent();
            ue.execCommand('serverparam', {
              "pid":pid 
            });
        });
    });

  //添加文章
  function save2(){
    // var radio =$("input[type='radio']:checked").val();
    var projectid = $('#pid').val();
    var prodcode = $('#prodcode1').val();
    var prodname = $('#prodname1').val();
    var subtext = $('#subtext1').val();
    var prodprincipal = $('#prodprincipal2').val();
    var prodlabel = $('#prodlabel2').val();
    var html = ue.getContent();
    // $('#myModal').on('hide.bs.modal', function () {  
    if (prodname&&prodcode){  
      $.ajax({
        type:"post",
        url:"/project/product/addarticle",
        data: {pid:projectid,code:prodcode,title:prodname,subtext:subtext,label:prodlabel,content:html,principal:prodprincipal},//父级id
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
      if ({{.role}}!=1){
        alert("权限不够！");
        return;
      }
      var selectRow=$('#attachments').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选！");
        return false;
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
    // $(document).ready(function() {
    //   $('#table0').bootstrapTable({
    //     // onLoadSuccess: function(){
    //       onPostBody: function(){
    //       alert("加载成功");
    //     $('#table0').bootstrapTable('append', randomData());
    //     $('#table0').bootstrapTable('scrollTo', 'bottom'); 
    //          return false;
    //       }
    //     });
    // });
    function randomData() {
        var startId = ~~(Math.random() * 100),
                rows = [];
        for (var i = 0; i < 10; i++) {
            rows.push({
              Id: startId + i,
              Code:startId + i,
              Title: 'test' + (startId + i),
                // id: startId + i,
                // name: 'test' + (startId + i),
                // price: '$' + (startId + i)
            });
        }
        return rows;
    }
</script>

</body>
</html>