<!DOCTYPE html>
<html>

<head>
  <meta charset="UTF-8">
  <title>发布数量</title>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <!-- <script type="text/javascript" src="/static/js/bootstrap.min.js"></script> -->
  <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" /> -->
  <!-- <script type="text/javascript" src="/static/js/moment.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/daterangepicker.js"></script> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css" /> -->
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css" /> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script> -->
  <!-- <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script> -->
  <!-- <script src="/static/js/moment-with-locales.min.js"></script> -->
  <script type="text/javascript" src="/static/js/echarts.min.js"></script>
</head>

<body>
  <div class="col-lg-12" id="main" style="width: 800px;height:600px;"></div>
  <pre>
  //取出某个目录（projectid）下用户文章数目
  type Result struct {
    Usernickname string `json:"name"`
    Productid    int64
    Total        int64 `json:"value"`
  }
  
  func GetWxUserArticles(pid int64) (results []*Result, err error) {
    db := GetDB()
    db.Order("total desc").Table("article").Select("product_id as productid, count(*) as total,user.nickname as usernickname").
      Joins("left JOIN product on product.id = article.product_id").
      Joins("left JOIN user on user.id = product.uid").Group("product.uid").
      Joins("left JOIN project on project.id = product.project_id").Where("project.id=?", pid).
     Scan(&results)
    return results, err
  }
  </pre>
  <script type="text/javascript">
  var myChart = echarts.init(document.getElementById('main'));
  $.get('/v1/wx/getwxuserarticles/26159').done(function(data) {
    // 填入数据
    myChart.setOption({
      title: {
        text: '发布文章数量',
        subtext: '统计来自珠三角设代',
        left: 'center'
      },
      tooltip: {
        trigger: 'item',
        formatter: "{a} <br/>{b}: {c} ({d}%)"
      },
      legend: {
        orient: 'vertical',
        x: 'right',
        // data: {{.Select2 }}
      },
      series: [{
        name: '作者文章数量',
        type: 'pie',
        radius: ['10%', '60%'],
        // 根据名字对应到相应的系列
        data: data.data
      }]
    });
  });
  </script>
</body>