<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>{{i18n .Lang "mgr.config_file"}} - Powered by MinDoc</title>

    <!-- Bootstrap -->
    <link href="{{cdncss "/static/mindoc/bootstrap/css/bootstrap.min.css"}}" rel="stylesheet">
    <link href="{{cdncss "/static/mindoc/font-awesome/css/font-awesome.min.css"}}" rel="stylesheet">
    <link href="{{cdncss "/static/mindoc/editor.md/css/editormd.css"}}" rel="stylesheet">
    <link href="{{cdncss "/static/mindoc/css/main.css" "version"}}" rel="stylesheet">
</head>
<body>
<div class="manual-reader">
{{template "widgets/header.tpl" .}}
    <div class="container manual-body">
        <div class="row">
        {{template "manager/widgets.tpl" .}}
            <div class="page-right">
                <div class="m-box">
                    <div class="box-head">
                        <strong class="box-title"> {{i18n .Lang "mgr.config_mgr"}}</strong>
                    </div>
                </div>
                <div class="box-body">
                    <form method="post" id="configFileContainerForm" action="{{urlfor "ManagerController.Config"}}">
                        <div id="configFileContainer">
                            <textarea style="display:none;" name="configFileTextArea">{{.ConfigContent}}</textarea>
                        </div>

                        <div class="form-group">
                            <button type="submit" id="btnSaveConfigFile" class="btn btn-success" data-loading-text="{{i18n .Lang "message.processing"}}">{{i18n .Lang "common.save"}}</button>
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
<script src="{{cdnjs "/static/mindoc/editor.md/editormd.js" "version"}}" type="text/javascript"></script>
<script src="{{cdnjs "/static/mindoc/js/main.js"}}" type="text/javascript"></script>
<script type="text/javascript">
    $(function() {
        var configEditor = editormd("configFileContainer", {
            width            : "100%",
            height           : 720,
            watch            : false,
            toolbar          : false,
            codeFold         : true,
            searchReplace    : true,
            placeholder      : "",
            mode             : "text/x-properties",
            path             : "{{cdnjs "/static/mindoc/editor.md/lib/"}}"
        });

        $("#configFileContainerForm").ajaxForm({
            beforeSubmit : function () {
                $("#btnSaveBookInfo").button("loading");
            },success : function (res) {
                if(res.errcode === 0) {
                    showSuccess("保存成功", "#form-error-message")
                }else{
                    showError(res.message);
                }
                $("#btnSaveConfigFile").button("reset");
            }
        });
    });
</script>
</body>
</html>