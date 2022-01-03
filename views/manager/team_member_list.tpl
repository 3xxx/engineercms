<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>团队用户管理 - Powered by MinDoc</title>

    <!-- Bootstrap -->
    <link href="{{cdncss "/static/mindoc/bootstrap/css/bootstrap.min.css"}}" rel="stylesheet">
    <link href="{{cdncss "/static/mindoc/font-awesome/css/font-awesome.min.css"}}" rel="stylesheet">
    <link href="{{cdncss "/static/mindoc/select2/4.0.5/css/select2.min.css"}}" rel="stylesheet">
    <link href="{{cdncss "/static/mindoc/css/main.css" "version"}}" rel="stylesheet">
    <style type="text/css">
        .table>tbody>tr>td{vertical-align: middle;}
    </style>
</head>
<body>
<div class="manual-reader">
{{template "widgets/header.tpl" .}}
    <div class="container manual-body">
        <div class="row">
        {{template "manager/widgets.tpl" "team"}}
            <div class="page-right">
                <div class="m-box">
                    <div class="box-head">
                        <strong class="box-title">{{.Model.TeamName}} - 成员管理</strong>
                        <button type="button"  class="btn btn-success btn-sm pull-right" data-toggle="modal" data-target="#addTeamMemberDialogModal"><i class="fa fa-user-plus" aria-hidden="true"></i> 添加成员</button>
                    </div>
                </div>
                <div class="box-body">
                    <div class="users-list" id="teamMemberList">
                        <template v-if="lists.length <= 0">
                            <div class="text-center">暂无数据</div>
                        </template>
                        <template v-else>
                            <table class="table">
                                <thead>
                                <tr>
                                    <th width="80">头像</th>
                                    <th width="100">账号</th>
                                    <th width="100">姓名</th>
                                    <th width="150">角色</th>
                                    <th width="80px">操作</th>
                                </tr>
                                </thead>
                                <tbody>
                                <tr v-for="item in lists">
                                    <td><img :src="item.avatar" onerror="this.src='{{cdnimg "/static/mindoc/images/middle.gif"}}'" class="img-circle" width="34" height="34"></td>
                                    <td>${item.account}</td>
                                    <td>${item.real_name}</td>
                                    <td>
                                        <div class="btn-group">
                                            <button type="button" class="btn btn-default btn-sm"  data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                                角色：${item.role_name}
                                                <span class="caret"></span></button>
                                            <ul class="dropdown-menu">
                                                <li><a href="javascript:;" @click="setTeamMemberRole(item.member_id,1)">管理员 <p class="text">拥有阅读、写作和管理权限</p></a> </li>
                                                <li><a href="javascript:;" @click="setTeamMemberRole(item.member_id,2)">编辑者 <p class="text">拥有阅读和写作权限</p></a> </li>
                                                <li><a href="javascript:;" @click="setTeamMemberRole(item.member_id,3)">观察者 <p class="text">拥有阅读权限</p></a> </li>
                                            </ul>
                                        </div>
                                    </td>
                                    <td>
                                        <button type="button" class="btn btn-danger btn-sm" @click="deleteMember(item.member_id,$event)" data-loading-text="删除中">删除</button>
                                    </td>
                                </tr>
                                </tbody>
                            </table>
                        </template>
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
<!-- Modal -->
<div class="modal fade" id="addTeamMemberDialogModal" tabindex="-1" role="dialog" aria-labelledby="addTeamMemberDialogModalLabel">
    <div class="modal-dialog modal-sm" role="document" style="width: 400px;">
        <form method="post" autocomplete="off" class="form-horizontal" action="{{urlfor "ManagerController.TeamMemberAdd"}}" id="addTeamMemberDialogForm">
            <input type="hidden" name="teamId" value="{{.Model.TeamId}}">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title" id="myModalLabel">添加成员</h4>
                </div>
                <div class="modal-body">
                    <div class="form-group">
                        <label class="col-sm-2 control-label">账号</label>
                        <div class="col-sm-10">
                            <select class="js-data-example-ajax form-control" multiple="multiple" name="memberId" id="memberId"></select>
                        </div>
                    </div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">角色</label>
                        <div class="col-sm-10">
                            <select name="roleId" class="form-control">
                                <option value="1">管理员</option>
                                <option value="2">编辑者</option>
                                <option value="3">观察者</option>
                            </select>
                        </div>
                    </div>

                    <div class="clearfix"></div>
                </div>
                <div class="modal-footer">
                    <span id="form-error-message"></span>
                    <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                    <button type="submit" class="btn btn-success" data-loading-text="保存中..." id="btnAddMember">保存</button>
                </div>
            </div>
        </form>
    </div>
</div><!--END Modal-->

<script src="{{cdnjs "/static/mindoc/jquery/1.12.4/jquery.min.js"}}"></script>
<script src="{{cdnjs "/static/mindoc/bootstrap/js/bootstrap.min.js"}}"></script>
<script src="{{cdnjs "/static/mindoc/vuejs/vue.min.js"}}"></script>
<script src="{{cdnjs "/static/mindoc/js/jquery.form.js"}}" type="text/javascript"></script>
<script src="{{cdnjs "/static/mindoc/select2/4.0.5/js/select2.full.min.js"}}"></script>
<script src="{{cdnjs "/static/mindoc/select2/4.0.5/js/i18n/zh-CN.js"}}"></script>
<script src="{{cdnjs "/static/mindoc/js/main.js"}}" type="text/javascript"></script>
<script type="text/javascript">
    $(function () {
        var modalCache = $("#addTeamMemberDialogModal form").html();

        /**
         * 添加用户
         */
        $("#addTeamMemberDialogForm").ajaxForm({
            beforeSubmit : function () {
                var memberId = $.trim($("#memberId").val());
                if(memberId === ""){
                    return showError("账号不能为空");
                }
                $("#btnAddMember").button("loading");
            },
            success : function (res) {
                if(res.errcode === 0){
                    app.lists.splice(0,0,res.data);
                    $("#addTeamMemberDialogModal").modal("hide");
                }else{
                    showError(res.message);
                }
                $("#btnAddMember").button("reset");
            }
        });
        $("#addTeamMemberDialogModal").on("hidden.bs.modal",function () {
            $(this).find("form").html(modalCache);
        }).on("show.bs.modal",function () {
            $('.js-data-example-ajax').select2({
                language: "zh-CN",
                minimumInputLength : 1,
                minimumResultsForSearch: Infinity,
                maximumSelectionLength:1,
                width : "100%",
                ajax: {
                    url: '{{urlfor "ManagerController.TeamSearchMember" "teamId" .Model.TeamId}}',
                    dataType: 'json',
                    data: function (params) {
                        return {
                            q: params.term, // search term
                            page: params.page
                        };
                    },
                    processResults: function (data, params) {
                        return {
                            results : data.data.results
                        }
                    }
                }
            });
        });

        var app = new Vue({
            el : "#teamMemberList",
            data : {
                lists : {{.Result}}
            },
            delimiters : ['${','}'],
            methods : {
                setTeamMemberRole : function (member_id, role) {
                    var $this = this;
                    $.ajax({
                        url :"{{urlfor "ManagerController.TeamChangeMemberRole"}}",
                        dataType :"json",
                        type :"post",
                        data : { "memberId" : member_id,"roleId" : role ,"teamId":{{.Model.TeamId}}},
                        success : function (res) {
                            if(res.errcode === 0){
                                for (var index in $this.lists) {
                                    var item = $this.lists[index];

                                    if (item.member_id === member_id) {

                                        $this.lists.splice(index,1,res.data);
                                        break;
                                    }
                                }
                            }else{
                                alert("操作失败：" + res.message);
                            }
                        }
                    })
                },
                deleteMember : function (id, e) {
                    var $this = this;
                    $.ajax({
                        url : "{{urlfor "ManagerController.TeamMemberDelete"}}",
                        type : "post",
                        data : { "memberId":id ,"teamId":{{.Model.TeamId}}},
                        dataType : "json",
                        success : function (res) {
                            if (res.errcode === 0) {

                                for (var index in $this.lists) {
                                    var item = $this.lists[index];
                                    if (item.member_id == id) {
                                        $this.lists.splice(index,1);
                                        break;
                                    }
                                }
                            } else {
                                alert("操作失败：" + res.message);
                            }
                        }
                    });
                }
            }
        });
        Vue.nextTick(function () {
            $("[data-toggle='tooltip']").tooltip();
        });
    });
</script>
</body>
</html>