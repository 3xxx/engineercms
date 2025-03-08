<!DOCTYPE html>
<html lang="zh-CN">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>微信扫码充值</title>
  <style>
    /*:root {
            --primary-color: #07c160;
            --bg-color: #f5f5f5;
            --card-bg: #fff;
            --text-color: #333;
        }*/

        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }

        .recharge-body {
          --primary-color: #07c160;
            --bg-color: #f5f5f5;
            --card-bg: #fff;
            --text-color: #333;

            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 
                'Helvetica Neue', Arial, sans-serif;
            background: var(--bg-color);
            color: var(--text-color);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
        }

        .recharge-container {
            background: var(--card-bg);
            padding: 2rem;
            border-radius: 12px;
            box-shadow: 0 4px 12px rgba(0,0,0,0.1);
            width: 100%;
            max-width: 400px;
            text-align: center;
        }

        .title {
            font-size: 1.5rem;
            margin-bottom: 1.5rem;
            color: var(--primary-color);
        }

        .input-group {
            margin-bottom: 1.5rem;
        }

        .amount-input {
            width: 100%;
            padding: 0.8rem;
            border: 2px solid #eee;
            border-radius: 8px;
            font-size: 1.1rem;
            text-align: center;
            transition: border-color 0.3s ease;
        }

        .amount-input:focus {
            outline: none;
            border-color: var(--primary-color);
        }

        .submit-btn {
            background: var(--primary-color);
            color: white;
            border: none;
            padding: 0.8rem 1.5rem;
            border-radius: 8px;
            font-size: 1rem;
            cursor: pointer;
            transition: opacity 0.3s ease;
            width: 100%;
        }

        .submit-btn:disabled {
            opacity: 0.7;
            cursor: not-allowed;
        }

        .qrcode-container {
            margin: 1.5rem 0;
            padding: 1rem;
            background: #f8f8f8;
            border-radius: 8px;
            display: none;
        }

        .qrcode-img {
            width: 200px;
            height: 200px;
            margin: 0 auto;
        }

        .status-message {
            margin-top: 3rem;
            color: #666;
            font-size: 0.9rem;
        }

        .loading {
            display: inline-block;
            width: 20px;
            height: 20px;
            border: 3px solid rgba(0,0,0,0.1);
            border-radius: 50%;
            border-top-color: var(--primary-color);
            animation: spin 1s ease-in-out infinite;
        }

        @keyframes spin {
            to { transform: rotate(360deg); }
        }
  </style>
</head>

<body>
  <div class="recharge-body">
    <div class="recharge-container">
      <h1 class="title">微信扫码充值</h1>
      <div class="input-group">
        <input type="number" class="amount-input" placeholder="请输入充值金额（元）" min="1" step="0.01">
      </div>
      <button class="submit-btn" onclick="handleRecharge()">
        生成充值二维码
      </button>
      <div class="qrcode-container">
        <!-- <img class="qrcode-img" id="qrcodeImg"> -->
        <div class="qrcode-img" id="qrcodeImg"></div>
        <p class="status-message" id="statusMessage"></p>
      </div>
    </div>
  </div>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script src="/static/vue.js/qrcode.min.js"></script>
  <!--   <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <script type="text/javascript" src="/static/js/tableExport.js"></script>
  <script type="text/javascript" src="/static/js/moment.min.js"></script> -->
  <script>
  // 模拟服务端接口
  // const mockServer = {
  //   // 获取充值二维码
  //   async getRechargeQrcode(amount) {
  //     // 模拟API延迟
  //     await new Promise(resolve => setTimeout(resolve, 500));

  //     // 返回模拟数据
  //     return {
  //       code: 200,
  //       data: {
  //         qrcodeUrl: `https://dummyimage.com/200x200/07c160/fff&text=￥${amount}`,
  //         orderId: `ORDER_${Date.now()}`
  //       }
  //     };
  //   },

  //   // 检查支付状态
  //   async checkPaymentStatus(orderId) {
  //     // 模拟API延迟
  //     await new Promise(resolve => setTimeout(resolve, 800));

  //     // 随机返回支付结果（模拟）
  //     return Math.random() > 0.5 ? { code: 200, paid: true } : { code: 200, paid: false };
  //   }
  // };

  // DOM元素
  const amountInput = document.querySelector('.amount-input');
  const qrcodeContainer = document.querySelector('.qrcode-container');
  const qrcodeImg = document.getElementById('qrcodeImg');
  const statusMessage = document.getElementById('statusMessage');
  const submitBtn = document.querySelector('.submit-btn');

  let currentOrderId = null;
  let checkInterval = null;

  var out_trade_no = "";

  async function handleRecharge() {
    const amount = parseFloat(amountInput.value);

    // 输入验证
    if (!amount || amount <= 0) {
      alert('请输入有效的充值金额');
      return;
    }

    // 显示加载状态
    submitBtn.disabled = true;
    submitBtn.innerHTML = '<div class="loading"></div>';

    var qrcode = new QRCode(document.getElementById("qrcodeImg"), {
      width: 230,
      height: 230,
      foreground: "#000000",
      background: "#ffffff",
      typeNumber: QRCode.CorrectLevel.L
    });
    var timer = null;


    $.ajax({
      url: '/v1/wx/wxnativepay',
      type: 'post',
      dataType: 'json',
      data: { payamount: parseFloat(amount * 100), out_trade_no: "recharge-pass", description: "userrecharge", attach: "userrecharge" }, //post数据
      success: function(data, textStatus) {
        //从服务器得到数据，显示数据并继续查询
        if (data.CodeUrl && data.CodeUrl !== '') {
          if (qrcode !== null) {
            qrcode.clear();
          }
          qrcode.makeCode(data.CodeUrl);
          // qrcodeImg.src = response.data.qrcodeUrl;
          qrcodeContainer.style.display = 'block';
          // 赋值订单号和生成订单时间
          // document.getElementById("billId").innerHTML = data.out_trade_no;
          out_trade_no = data.out_trade_no;
          // document.getElementById("createTime").innerHTML =data.createtime;
        };

        // loadmsg();
        window.timer = setInterval(() => {
          setTimeout("loadmsg()", 0)
        }, 2000)
      },
      //Ajax请求超时，继续查询
      error: function(XMLHttpRequest, textStatus, errorThrown) {
        if (textStatus == "timeout") {
          // setTimeout("loadmsg()", 1000);
        } else { //异常
          // setTimeout("loadmsg()", 4000);
        }
      },
      complete: function(xhr, status) {
        submitBtn.disabled = false;
        submitBtn.textContent = '生成充值二维码';
      }
      // success(result,status,xhr)
    })

    // try {
    //   // 调用获取二维码接口
    //   const response = await mockServer.getRechargeQrcode(amount);

    //   if (response.code === 200) {
    //     // 显示二维码
    //     qrcodeImg.src = response.data.qrcodeUrl;
    //     qrcodeContainer.style.display = 'block';
    //     currentOrderId = response.data.orderId;

    //     // 开始轮询支付状态
    //     startPaymentPolling();
    //   }
    // } catch (error) {
    //   console.error('充值请求失败:', error);
    //   alert('生成二维码失败，请重试');
    // } finally {
    //   submitBtn.disabled = false;
    //   submitBtn.textContent = '生成充值二维码';
    // }
  }

  function startPaymentPolling() {
    checkInterval = setInterval(async () => {
      try {
        const res = await mockServer.checkPaymentStatus(currentOrderId);

        if (res.paid) {
          clearInterval(checkInterval);
          statusMessage.innerHTML = '✅ 支付成功！页面即将跳转...';
          setTimeout(() => {
            alert('充值成功！');
            location.reload();
          }, 1500);
        } else {
          statusMessage.textContent = '等待扫码支付中...';
        }
      } catch (error) {
        console.error('支付状态查询失败:', error);
        statusMessage.textContent = '状态查询失败，请稍后重试';
      }
    }, 2000);
  }

  // 检查是否支付完成
  function loadmsg() {
    $.ajax({
      type: "GET",
      dataType: "json",
      url: "/v1/wx/queryorderbyouttradeno",
      timeout: 10000, //ajax请求超时时间10s
      data: { type: "wxNativePay", out_trade_no: out_trade_no }, //post数据
      success: function(data, textStatus) {
        //从服务器得到数据，显示数据并继续查询
        if (data.result.trade_state && data.result.trade_state == "SUCCESS") {
          console.log("支付成功！");

          statusMessage.innerHTML = '✅ 支付成功！页面即将跳转...';
          setTimeout(() => {
            alert('充值成功！');
            location.reload();
          }, 1500);

        } else {
          statusMessage.textContent = '等待扫码支付中...';
          //   // setTimeout("loadmsg()", 3000);
          //   window.timer = setInterval(() => {
          //     setTimeout("loadmsg()", 0)
          //   }, 2000)
        }
      },
      //Ajax请求超时，继续查询
      error: function(XMLHttpRequest, textStatus, errorThrown) {
        if (textStatus == "timeout") {
          // setTimeout("loadmsg()", 1000);
          console.error('超时:', textStatus);
          statusMessage.textContent = '超时，请稍后重试';
        } else { //异常
          // setTimeout("loadmsg()", 4000);
          console.error('支付状态查询失败:', XMLHttpRequest.status);
          statusMessage.textContent = '支付状态查询失败，请稍后重试';
        }
      }
    });
  }

  // 输入验证
  amountInput.addEventListener('input', (e) => {
    let value = e.target.value;
    // 限制只能输入数字和一个小数点
    value = value.replace(/[^0-9.]/g, '');
    // 去除多余的小数点
    value = value.replace(/\.{2,}/g, '.');
    // 限制小数点后两位
    if (value.indexOf('.') >= 0) {
      value = value.substring(0, value.indexOf('.') + 3);
    }
    e.target.value = value;
  });
  </script>
</body>

</html>