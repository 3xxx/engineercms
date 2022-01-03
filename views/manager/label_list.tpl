<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>标签管理 - Powered by MinDoc</title>

    <!-- Bootstrap -->
    <link href="{{cdncss "/static/mindoc/bootstrap/css/bootstrap.min.css"}}" rel="stylesheet" type="text/css">
    <link href="{{cdncss "/static/mindoc/font-awesome/css/font-awesome.min.css"}}" rel="stylesheet" type="text/css">

    <link href="{{cdncss "/static/mindoc/css/main.css" "version"}}" rel="stylesheet">
    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="{{cdnjs "/static/mindoc/html5shiv/3.7.3/html5shiv.min.js"}}"></script>
    <script src="{{cdnjs "/static/mindoc/respond.js/1.4.2/respond.min.js" }}"></script>
    <![endif]-->
</head>
<body>
<div class="manual-reader">
{{template "widgets/header.tpl" .}}
    <div class="container manual-body">
        <div class="row">
        {{template "manager/widgets.tpl" "label"}}
            <div class="page-right">
                <div class="m-box">
                    <div class="box-head">
                        <strong class="box-title">标签管理</strong>
                    </div>
                </div>
                <div class="box-body">
                    <div class="attach-list" id="labelList">
                        <table class="table">
                            <thead>
                            <tr>
                                <th width="10%">#</th>
                                <th width="55%">标签名称</th>
                                <th width="20%">使用数量</th>
                                <th>操作</th>
                            </tr>
                            </thead>
                            <tbody>
                            {{range $index,$item := .Lists}}
                            <tr>
                                <td>{{$item.LabelId}}</td>
                                <td>{{$item.LabelName}}</td>
                                <td>{{$item.BookNumber}}</td>
                                <td>
                                    <button type="button" data-method="delete" class="btn btn-danger btn-sm" data-id="{{$item.LabelId}}" data-loading-text="删除中...">删除</button>
                                    <a href="{{urlfor "LabelController.Index" ":key" $item.LabelName}}" class="btn btn-success btn-sm" target="_blank">详情</a>

                                </td>
                            </tr>
                            {{else}}
                            <tr><td class="text-center" colspan="6">暂无数据</td></tr>
                            {{end}}
                            </tbody>
                        </table>
                        <nav class="pagination-container">
                        {{.PageHtml}}
                        </nav>
                    </div>
                </div>
            </div>
        </div>
    </div>
{{template "widgets/footer.tpl" .}}
</div>

<script src="{{cdnjs "/static/mindoc/jquery/1.12.4/jquery.min.js"}}"></script>
<script src="{{cdnjs "/static/mindoc/bootstrap/js/bootstrap.min.js"}}"></script>
<script src="{{cdnjs "/static/mindoc/js/jquery.form.js"}}" type="text/javascript"></script>
<script src="{{cdnjs "/static/mindoc/layer/layer.js" }}" type="text/javascript"></script>
<script type="text/javascript">
    $(function () {
        $("#labelList").on("click","button[data-method='delete']",function () {
            var id = $(this).attr("data-id");
            var $this = $(this);
            $(this).button("loading");
            $.ajax({
                url : "{{urlfor "ManagerController.LabelDelete" ":id" ""}}" + id,
                type : "post",
                dataType : "json",
                success : function (res) {
                    if(res.errcode === 0){
                        $this.closest("tr").remove().empty();
                    }else {
                        layer.msg(res.message);
                    }
                },
                error : function () {
                    layer.msg("服务器异常");
                },
                complete : function () {
                    $this.button("reset");
                }
            });
        });
    });
</script>
</body>
</html>