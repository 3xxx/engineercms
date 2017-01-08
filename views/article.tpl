<!DOCTYPE html>
<title>CMS系统</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/ueditor/ueditor.parse.js"></script>
<style type="text/css">
img{max-width:100%}
nav {
  background: #4682B4;
}
h3 {line-height: 150%;
  /*color: #DC143C;#000000*/
  color:#FFFFFF;
  background: #4682B4;/*#fff*/
  margin: 0;
  padding: 0;
  font-family: Georgia, Palatino, serif;
  }
  h3 a {color: inherit;}/*h3标签里面的a标签继承父级颜色*/

   h4{
      color: #DC143C;
  font-family: georgia, palatino, "New Century Schoolbook",
  times, serif;
  font-weight: normal;
  /*font-size: 2em;*/
  margin-top: 1em;
  margin-bottom: 0;
  }

</style>
</head>

<body>
<div class="col-lg-8 col-md-offset-2">
<div class="page-header">
  <h2>{{.product.Code}}-{{.product.Title}} <small>{{.article.Subtext}}</small></h2>
  <label>作者：{{.product.Principal}}</label>
  <small>发表于：{{dateformat .product.Created "2006-01-02"}}</small>
  <span class="label label-info">{{.product.Label}}</span> 
</div>

  <div class="content">
    {{str2html .article.Content}}
  </div>
</div>
<script>

 //实例化编辑器
    //议使用工厂方法getEditor创建和引用编辑器实例，如果在某个闭包下引用该编辑器，直接调用UE.getEditor('editor')就能拿到相关的实例
    // var ue = UE.getEditor('container');
    var ue = UE.getEditor('container', {
      autoHeightEnabled: true,
      autoFloatEnabled: true
    });
        // 语法
        // uParse(selector,[option])
        /*
         selector支持
         id,class,tagName
         */
        /*
         目前支持的参数
         option:
         highlightJsUrl 代码高亮相关js的路径 如果展示有代码高亮，必须给定该属性
         highlightCssUrl 代码高亮相关css的路径 如果展示有代码高亮，必须给定该属性
         liiconpath 自定义列表样式的图标路径，可以不给定，默认'http://bs.baidu.com/listicon/',
         listDefaultPaddingLeft : 自定义列表样式的左边宽度 默认'20',
         customRule 可以传入你自己的处理规则函数，函数第一个参数是容器节点
         */
        uParse('.content',{
            rootPath : '/static/ueditor/'
        })
</script>
</body>
</html>