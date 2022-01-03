<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>配置管理 - Powered by MinDoc</title>

    <!-- Bootstrap -->
    <link href="{{cdncss "/static/mindoc/bootstrap/css/bootstrap.min.css"}}" rel="stylesheet">
    <link href="{{cdncss "/static/mindoc/font-awesome/css/font-awesome.min.css"}}" rel="stylesheet">
    <link href="{{cdncss "/static/mindoc/css/main.css" "version"}}" rel="stylesheet">
    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="/static/mindoc/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="/static/mindoc/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
</head>
<body>
<div class="manual-reader">
    {{template "widgets/header.tpl" .}}
    <div class="container manual-body">
        <div class="row">
        {{template "manager/widgets.tpl" "setting"}}
            <div class="page-right">
                <div class="m-box">
                    <div class="box-head">
                        <strong class="box-title"> 配置管理</strong>
                    </div>
                </div>
                <div class="box-body">
                    <form method="post" id="gloablEditForm" action="{{urlfor "ManagerController.Setting"}}">
                        <div class="form-group">
                            <label>网站标题</label>
                            <input type="text" class="form-control" name="SITE_NAME" id="siteName" placeholder="网站标题" value="{{.SITE_NAME}}">
                        </div>
                        <div class="form-group">
                            <label>域名备案</label>
                            <input type="text" class="form-control" name="site_beian" id="siteName" placeholder="域名备案" value="{{.site_beian}}" maxlength="50">
                        </div>
                        <div class="form-group">
                            <label>网站描述</label>
                            <textarea rows="3" class="form-control" name="site_description" style="height: 90px" placeholder="网站描述">{{.site_description}}</textarea>
                            <p class="text">描述信息不超过500个字符</p>
                        </div>
                            <div class="form-group">
                                <label>启用匿名访问</label>
                                <div class="radio">
                                    <label class="radio-inline">
                                        <input type="radio" {{if eq .ENABLE_ANONYMOUS "true"}}checked{{end}} name="ENABLE_ANONYMOUS" value="true">开启<span class="text"></span>
                                    </label>
                                    <label class="radio-inline">
                                        <input type="radio" {{if eq .ENABLE_ANONYMOUS "false"}}checked{{end}} name="ENABLE_ANONYMOUS" value="false">关闭<span class="text"></span>
                                    </label>
                                </div>
                            </div>
                        <div class="form-group">
                            <label>启用注册</label>
                            <div class="radio">
                                <label class="radio-inline">
                                    <input type="radio" {{if eq .ENABLED_REGISTER "true"}}checked{{end}} name="ENABLED_REGISTER" value="true">开启<span class="text"></span>
                                </label>
                                <label class="radio-inline">
                                    <input type="radio" {{if eq .ENABLED_REGISTER "false"}}checked{{end}} name="ENABLED_REGISTER" value="false">关闭<span class="text"></span>
                                </label>
                            </div>
                        </div>
                        <div class="form-group">
                            <label>启用验证码</label>
                            <div class="radio">
                                <label class="radio-inline">
                                    <input type="radio" {{if eq .ENABLED_CAPTCHA "true"}}checked{{end}} name="ENABLED_CAPTCHA" value="true">开启<span class="text"></span>
                                </label>
                                <label class="radio-inline">
                                    <input type="radio" {{if eq .ENABLED_CAPTCHA "false"}}checked{{end}} name="ENABLED_CAPTCHA" value="false">关闭<span class="text"></span>
                                </label>
                            </div>
                        </div>
                        <div class="form-group">
                            <label>启用文档历史</label>
                            <div class="radio">
                                <label class="radio-inline">
                                    <input type="radio" {{if eq .ENABLE_DOCUMENT_HISTORY "true"}}checked{{end}} name="ENABLE_DOCUMENT_HISTORY" value="true">开启<span class="text"></span>
                                </label>
                                <label class="radio-inline">
                                    <input type="radio" {{if eq .ENABLE_DOCUMENT_HISTORY "false"}}checked{{end}} name="ENABLE_DOCUMENT_HISTORY" value="false">关闭<span class="text"></span>
                                </label>
                            </div>
                        </div>

                        <div class="form-group">
                            <button type="submit" id="btnSaveBookInfo" class="btn btn-success" data-loading-text="保存中...">保存修改</button>
                            <span id="form-error-message" class="error-message"></span>
                        </div>
                        </form>

                    <div class="clearfix"></div>

                </div>
            </div>
        </div>
    </div>
    {{template "widgets/footer.tpl" .}}
</div>


<script src="{{cdnjs "/static/mindoc/jquery/1.12.4/jquery.min.js"}}" type="text/javascript"></script>
<script src="{{cdnjs "/static/mindoc/bootstrap/js/bootstrap.min.js"}}" type="text/javascript"></script>
<script src="{{cdnjs "/static/mindoc/js/jquery.form.js"}}" type="text/javascript"></script>
<script src="{{cdnjs "/static/mindoc/js/main.js"}}" type="text/javascript"></script>
<script type="text/javascript">
    $(function () {
        $("#gloablEditForm").ajaxForm({
            beforeSubmit : function () {
                var title = $.trim($("#siteName").val());

                if (title === ""){
                    return showError("网站标题不能为空");
                }
                $("#btnSaveBookInfo").button("loading");
            },success : function (res) {
                if(res.errcode === 0) {
                    showSuccess("保存成功")
                }else{
                    showError(res.message);
                }
                $("#btnSaveBookInfo").button("reset");
            }
        });
    });
</script>
</body>
</html>