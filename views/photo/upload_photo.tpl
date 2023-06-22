<!DOCTYPE html>
<html>

<head>
  <meta name="renderer" content="webkit">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <!-- 收藏用logo图标 -->
  <link rel="bookmark" type="image/x-icon" href="/static/img/pss.ico" />
  <!-- 网站显示页logo图标 -->
  <link rel="shortcut icon" href="/static/img/pss.ico">
  <!-- KRAJEE EXPLORER THEME (ADVANCED) -->
  <title>照片上传</title>
  <!-- <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script> -->
  <link href="/static/bootstrap-fileinput/css/all.css" rel="stylesheet">
  <link href="/static/bootstrap-fileinput/assets/prod/all-krajee.min.css?ver=201903112143" rel="stylesheet">
  <link href="/static/bootstrap-fileinput/assets/1d958cec/css/fileinput.css?ver=201909132002" rel="stylesheet">
  <link href="/static/bootstrap-fileinput/assets/1d958cec/themes/explorer/theme.min.css?ver=201908311938" rel="stylesheet">
  <link href="/static/bootstrap-fileinput/assets/fc69cbca/css/dropdown.min.css" rel="stylesheet">
  <script src="/static/bootstrap-fileinput/js/buttons.js" async></script>
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <script src="/static/bootstrap-fileinput/assets/1d958cec/js/plugins/purify.min.js"></script>
  <script src="/static/bootstrap-fileinput/js/jquery-3.1.1.min.js" crossorigin="anonymous"></script>
  <!-- <script type="text/javascript" src="/static/js/bootstrap.min.js"></script> -->
  <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
  <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.bundle.min.js" crossorigin="anonymous"></script>
  <script src="/static/bootstrap-fileinput/assets/prod/all-krajee.min.js?ver=201903112143"></script>
  <script src="/static/bootstrap-fileinput/assets/1d958cec/js/plugins/sortable.min.js"></script>
  <script src="/static/bootstrap-fileinput/assets/1d958cec/js/plugins/piexif.min.js"></script>
  <script src="/static/bootstrap-fileinput/assets/1d958cec/js/fileinput.js?ver=201909132002"></script>
  <script src="/static/bootstrap-fileinput/assets/1d958cec/themes/explorer/theme.min.js?ver=201908311938"></script>
  <script src="/static/bootstrap-fileinput/assets/fc69cbca/js/dropdown.min.js"></script>
</head>
<div class="file-loading">
  <input id="input-ke-2" name="input-ke-2[]" type="file" multiple data-browse-on-zone-click="true">
  <!-- 注意：服务的用上面的name值来获取文件！！！！ -->
  <!-- <input id="input-b1" name="input-b1" type="file" class="file" > -->
</div>
<script>
// must load the font-awesome.css for this example
$("#input-ke-2").fileinput({
  language: 'zh',
  theme: "explorer",
  // uploadAsync: true, //异步上传，一个文件请求一次post
  uploadAsync: false, //同步上传，多文件只请求一次post
  uploadUrl: "/v1/wx/uploadphotodata/{{.Id}}",
  minFileCount: 1,
  maxFileCount: 20,
  maxFileSize: 200000,
  removeFromPreviewOnError: true,
  overwriteInitial: false,
  allowedFileExtensions: ["jpg", "png", "gif", "pdf"], //接收的文件后缀
  previewFileIcon: '<i class="fas fa-file"></i>',
  initialPreview: [
    // IMAGE DATA
    // 'https://picsum.photos/800/560?image=1071',
    // IMAGE RAW MARKUP
    // '<img src="https://picsum.photos/800/560?image=1075" class="kv-preview-data file-preview-image">',
    // TEXT DATA
    // "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ut mauris ut libero fermentum feugiat eu et dui. Mauris condimentum rhoncus enim, sed semper neque vestibulum id. Nulla semper, turpis ut consequat imperdiet, enim turpis aliquet orci, eget venenatis elit sapien non ante. Aliquam neque ipsum, rhoncus id ipsum et, volutpat tincidunt augue. Maecenas dolor libero, gravida nec est at, commodo tempor massa. Sed id feugiat massa. Pellentesque at est eu ante aliquam viverra ac sed est.",
    // PDF DATA
    // 'http://kartik-v.github.io/bootstrap-fileinput-samples/samples/pdf-sample.pdf',
    // VIDEO DATA
    // "http://kartik-v.github.io/bootstrap-fileinput-samples/samples/small.mp4"
  ],
  initialPreviewAsData: true, // defaults markup  
  initialPreviewConfig: [
    // {caption: "Image 11.jpg", size: 762980, url: "/site/file-delete", downloadUrl: 'https://picsum.photos/800/460?image=11', key: 11},
    // {previewAsData: false, size: 823782, caption: "Image 13.jpg", url: "/site/file-delete", downloadUrl: 'https://picsum.photos/800/460?image=13', key: 13}, 
    // {caption: "Lorem Ipsum.txt", type: "text", size: 1430, url: "/site/file-delete", key: 12}, 
    // {type: "pdf", size: 8000, caption: "PDF Sample.pdf", url: "/site/file-delete", key: 14}, 
    // {type: "video", size: 375000, filetype: "video/mp4", caption: "Krajee Sample.mp4", url: "/site/file-delete", key: 15} 
  ],
  // initialPreviewDownloadUrl:'https://picsum.photos/1920/1080?image={key}', // the key will be dynamically replaced 
  uploadExtraData: {
    img_key: "1000",
    img_keywords: "happy, nature"
  },
  preferIconicPreview: true, // this will force thumbnails to display icons for following file extensions
  previewFileIconSettings: { // configure your icon file extensions
    'doc': '<i class="fas fa-file-word text-primary"></i>',
    'xls': '<i class="fas fa-file-excel text-success"></i>',
    'ppt': '<i class="fas fa-file-powerpoint text-danger"></i>',
    'pdf': '<i class="fas fa-file-pdf text-danger"></i>',
    'zip': '<i class="fas fa-file-archive text-muted"></i>',
    'htm': '<i class="fas fa-file-code text-info"></i>',
    'txt': '<i class="fas fa-file-text text-info"></i>',
    'mov': '<i class="fas fa-file-video text-warning"></i>',
    'mp3': '<i class="fas fa-file-audio text-warning"></i>',
    // note for these file types below no extension determination logic 
    // has been configured (the keys itself will be used as extensions)
    'jpg': '<i class="fas fa-file-image text-danger"></i>',
    'gif': '<i class="fas fa-file-image text-muted"></i>',
    'png': '<i class="fas fa-file-image text-primary"></i>'
  },
  previewFileExtSettings: { // configure the logic for determining icon file extensions
    'doc': function(ext) {
      return ext.match(/(doc|docx)$/i);
    },
    'xls': function(ext) {
      return ext.match(/(xls|xlsx)$/i);
    },
    'ppt': function(ext) {
      return ext.match(/(ppt|pptx)$/i);
    },
    'zip': function(ext) {
      return ext.match(/(zip|rar|tar|gzip|gz|7z)$/i);
    },
    'htm': function(ext) {
      return ext.match(/(htm|html)$/i);
    },
    'txt': function(ext) {
      return ext.match(/(txt|ini|csv|java|php|js|css)$/i);
    },
    'mov': function(ext) {
      return ext.match(/(avi|mpg|mkv|mov|mp4|3gp|webm|wmv)$/i);
    },
    'mp3': function(ext) {
      return ext.match(/(mp3|wav)$/i);
    },
    'mov': function(ext) {
      return ext.match(/(avi|mpg|mkv|mov|mp4|3gp|webm|wmv)$/i);
    },
    'mcdx': function(ext) {
      return ext.match(/(mcdx|mctx)$/i);
    }
  }
}).on("filebatchselected", function(event, data) {
  console.log("同步上传选择完成！")
  if (data.length == 0) {
    return;
  }
}).on("fileuploaded", function(event, data, previewId, index) { //一个文件上传成功
  // console.log('文件上传成功！');
  // console.log('文件上传成功！'+data);
  if (data.response.info == "ERROR") {
    console.log("上传失败回调：" + data.response.data + ",请重新上传")
  }
  // else{
  //   alert("上传附件成功")
  // }
  // console.log(data);
  // console.log(data.response.data)
}).on('fileerror', function(event, data, msg) {
  // alert(data.data);
  console.log(data);
}).on("filebatchuploadsuccess", function(event, data, previewId, index) {
  console.log("同步上传附件成功！");
  console.log(data.response.data);
}).on("filebatchuploadcomplete", function(event, data, previewId, index) { //event, data, previewId, index
  console.log("上传附件完成");
  console.log(data);
});
</script>

</html>