<!-- office下载模式，dwg下载模式 -->
<!DOCTYPE html>

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, height=device-height, initial-scale=1.0, maximum-scale=1.0" />
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <script src="/static/js/tableExport.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />

  <link rel="stylesheet" href="/static/froala/css/froala_editor.css">
  <link rel="stylesheet" href="/static/froala/css/froala_style.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/code_view.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/draggable.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/colors.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/emoticons.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/image_manager.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/image.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/line_breaker.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/table.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/char_counter.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/video.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/fullscreen.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/file.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/quick_insert.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/help.css">
  <link rel="stylesheet" href="/static/froala/css/third_party/spell_checker.css">
  <link rel="stylesheet" href="/static/froala/css/plugins/special_characters.css">
</head>

<body>
  <div>
    <h3 class="modal-title" style="padding:10px">新建项目</h3>
    <div class="form-group must" style="float: left;width: 50%;">
      <label class="col-sm-3 control-label" style="padding-left:10px;">编号</label>
      <div class="col-sm-7">
        <input type="text" class="form-control" id="projcode" name="projcode"></div>
    </div>
    <div class="form-group must" style="float: left;width: 50%;">
      <label class="col-sm-3 control-label" style="padding-left:10px;">名称</label>
      <div class="col-sm-7">
        <input type="tel" class="form-control" id="projname" name="projname"></div>
    </div>
    <div class="form-group must" style="float: left;width: 50%;">
      <label class="col-sm-3 control-label" style="padding-left:10px;">标签</label>
      <div class="col-sm-7">
        <input type="tel" class="form-control" id="projlabel" name="projlabel"></div>
    </div>
    <div class="form-group must" style="float: left;width: 50%;">
      <label class="col-sm-3 control-label" style="padding-left:10px;">负责人</label>
      <div class="col-sm-7">
        <input type="tel" class="form-control" id="projprincipal" name="projprincipal" placeholder="以英文,号分割"></div>
    </div>
    <div class="form-group must" style="float: left;width: 50%;">
      <label class="col-sm-3 control-label" style="padding-left:10px;">类别</label>
      <div class="col-sm-7">
        <select name="admincategory" id="admincategory" class="form-control" required onchange="refreshtable()">
          <option>选择类别：</option>
        </select>
      </div>
    </div>
    <!-- <div> -->
    <div id="details">
      <!--  style="display:none" -->
      <h3>工程目录分级</h3>
      <table id="table1"
      data-page-size="5"
      data-page-list="[5, 25, 50, All]"
      data-unique-id="id"
      data-sort-name="Grade"
      data-pagination="true"
      data-side-pagination="client"
      data-click-to-select="true"
      style="width: 100%;padding: 10px"
      >
        <thead>
          <tr>
            <th data-width="10" data-checkbox="true"></th>
            <th data-formatter="index1">#</th>
            <th data-field="Title">名称</th>
            <th data-field="Code">代码</th>
            <th data-field="Grade" data-sortable="true">级别</th>
          </tr>
        </thead>
      </table>
    </div>
    <div id="details1">
      <!--  style="display:none" -->
      <h3>项目模板</h3>
      <div class="col-sm-3 checkbox">
        <label><input type="checkbox" checked="checked" value="true" id="ispermission">权限继承</label>
      </div>
      <table id="table2"></table>
    </div>

    <label style="padding:10px">文章正文:</label>
    <div id="editor" style="width: 100%;padding: 10px">
      <div id='edit' style="margin-top: 30px;">
      </div>
    </div>

    <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
    <button type="button" class="btn btn-primary" onclick="save2()">保存</button>
  </div>
  <script type="text/javascript" src="/static/froala/js/froala_editor.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/align.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/char_counter.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/code_beautifier.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/code_view.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/colors.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/draggable.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/emoticons.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/entities.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/file.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/font_size.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/font_family.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/fullscreen.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/image.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/image_manager.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/line_breaker.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/inline_style.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/link.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/lists.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/paragraph_format.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/paragraph_style.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/quick_insert.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/quote.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/table.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/save.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/url.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/video.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/help.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/print.min.js"></script>
  <!-- <script type="text/javascript" src="/static/froala/js/third_party/spell_checker.min.js"></script> -->
  <script type="text/javascript" src="/static/froala/js/plugins/special_characters.min.js"></script>
  <script type="text/javascript" src="/static/froala/js/plugins/word_paste.min.js"></script>
  <script src="/static/froala/js/languages/zh_cn.js"></script>
  <script>
  // (function () {
  //   new FroalaEditor("#edit", {
  //     // Set custom buttons with separator between them.
  //     toolbarButtons: [ ['undo', 'redo'], ['bold', 'italic', 'underline', 'strikethrough', 'subscript', 'superscript', 'outdent', 'indent', 'clearFormatting', 'insertTable', 'html'] ],
  //     toolbarButtonsXS: [ ['undo', 'redo'], ['bold', 'italic', 'underline'] ]
  //   })
  // })()
  $(function() {
    //超大屏幕'fullscreen',
    var toolbarButtons = ['bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', 'fontFamily', 'fontSize', '|', 'color', 'emoticons', 'inlineStyle', 'paragraphStyle', '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', 'quote', 'insertHR', '-', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', 'undo', 'redo', 'clearFormatting', 'selectAll', 'html', 'help'];
    //大屏幕
    var toolbarButtonsMD = ['bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', '|', 'fontFamily', 'fontSize', 'color', 'inlineStyle', 'paragraphStyle', '|', 'paragraphFormat', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', '-', 'quote', 'insertLink', 'insertImage', 'insertVideo', 'insertFile', 'insertTable', '|', 'specialCharacters', 'insertHR', 'undo', 'redo', 'clearFormatting', '|', 'html', 'help'];
    //小屏幕'fullscreen',
    var toolbarButtonsSM = ['bold', 'italic', 'underline', 'fontFamily', 'fontSize', 'insertLink', 'insertImage', 'insertVideo', 'insertTable', 'undo', 'redo'];
    //手机
    var toolbarButtonsXS = ['insertImage', 'insertVideo', 'bold', 'italic', 'fontSize', 'undo', 'redo'];
    var pid = $('#pid').val();
    //编辑器初始化并赋值
    // $('#edit').froalaEditor({
    new FroalaEditor("#edit", {
      placeholderText: '请输入内容',
      charCounterCount: true, //默认
      // charCounterMax         : -1,//默认
      saveInterval: 0, //不自动保存，默认10000
      // theme                    : "red",
      height: "300px",
      toolbarBottom: false, //默认
      toolbarButtonsMD: toolbarButtons, //toolbarButtonsMD,
      toolbarButtonsSM: toolbarButtonsMD, //toolbarButtonsSM,
      toolbarButtonsXS: toolbarButtonsXS,
      toolbarInline: false, //true选中设置样式,默认false
      imageUploadMethod: 'POST',
      heightMin: 450,
      charCounterMax: 3000,
      // imageUploadURL: "uploadImgEditor",
      imageParams: { postId: "123" },
      params: {
        acl: '01',
        AWSAccessKeyId: '02',
        policy: '03',
        signature: '04',
      },
      autosave: true,
      autosaveInterval: 2500,
      saveURL: 'hander/FroalaHandler.ashx',
      saveParams: { postId: '1' },
      spellcheck: false,
      imageUploadURL: '/uploadimg', //上传到本地服务器
      imageUploadParams: { pid: '{{.Id}}' },
      imageDeleteURL: 'lib/delete_image.php', //删除图片
      imagesLoadURL: 'lib/load_images.php', //管理图片
      videoUploadURL: '/uploadvideo',
      videoUploadParams: { pid: '{{.Id}}' },
      fileUploadURL: '/uploadimg',
      fileUploadParams: { pid: '{{.Id}}' },
      // enter: $.FroalaEditor.ENTER_BR,
      language: 'zh_cn',
      // toolbarButtons: ['bold', 'italic', 'underline', 'paragraphFormat', 'align','color','fontSize','insertImage','insertTable','undo', 'redo']
    });
  })

  //填充select选项
  // $(document).ready(function() {
  $(function() {
    //   $(array).each(function(index){
    //     alert(this);
    // });
    // $.each(array,function(index){
    //     alert(this);
    // });
    $("#admincategory").append('<option value="a">项目模板</option>');
    if ({{.Select2 }}) { //20171021从meirit修改而来
      $.each({{.Select2 }}, function(i, d) {
        // alert(this);
        // alert(i);
        // alert(d);
        $("#admincategory").append('<option value="' + i + '">' + d + '</option>');
      });
    }
  });

  function index1(value, row, index) {
    // alert( "Data Loaded: " + index );
    return index + 1
  }

/*数据json*/
    var json1 = [{ "Id": "1", "Title": "规划", "Code": "A", "Grade": "1" },
      { "Id": "7", "Title": "可研", "Code": "B", "Grade": "1" },
      { "Id": "2", "Title": "报告", "Code": "B", "Grade": "2" },
      { "Id": "3", "Title": "图纸", "Code": "T", "Grade": "2" },
      { "Id": "4", "Title": "水工", "Code": "5", "Grade": "3" },
      { "Id": "5", "Title": "机电", "Code": "6", "Grade": "3" },
      { "Id": "6", "Title": "施工", "Code": "7", "Grade": "3" }
    ];
    /*初始化table数据*/
    $(function() {
      $("#table1").bootstrapTable({
        data: json1
      });
    });

  //添加文章
  function save2() {
    // var radio =$("input[type='radio']:checked").val();
    var projectid = $('#pid').val();
    var prodcode = $('#prodcode1').val();
    var prodname = $('#prodname1').val();
    var subtext = $('#subtext1').val();
    var prodprincipal = $('#prodprincipal2').val();
    var prodlabel = $('#prodlabel2').val();
    var relevancy = $('#relevancy2').val();
    // var html = $('div#edit').froalaEditor('html.get'); //$('#edit')[0].childNodes[1].innerHTML;
    let editor = new FroalaEditor('#edit', {}, function() {});
    var html = editor.html.get()
    // $('#myModal').on('hide.bs.modal', function () {
    if (prodname && prodcode) {
      $.ajax({
        type: "post",
        url: "/project/product/addarticle",
        data: { pid: projectid, code: prodcode, title: prodname, subtext: subtext, label: prodlabel, content: html, principal: prodprincipal, relevancy: relevancy }, //父级id
        success: function(data, status) {
          alert("添加“" + data + "”成功！(status:" + status + ".)");
          $('#modalTable2').modal('hide');
          $('#table0').bootstrapTable('refresh', { url: '/project/products/' + {{.Id }} });
        },
      });
    } else {
      alert("请填写编号和名称！");
      return;
    }
  }
  </script>

</body>

</html>