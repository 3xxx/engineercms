<!DOCTYPE html>
<!-- saved from url=(0049)https://getbootstrap.com/docs/5.3/examples/album/ -->
<html lang="en" data-bs-theme="auto">

<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <script src="/static/bootstrap5.3/color-modes.js"></script>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="description" content="">
  <meta name="author" content="Mark Otto, Jacob Thornton, and Bootstrap contributors">
  <meta name="generator" content="Hugo 0.118.2">
  <title>FreeCAD 3D Models</title>
  <link rel="canonical" href="https://getbootstrap.com/docs/5.3/examples/album/">

  <link rel="stylesheet" href="/static/bootstrap5.3/css@3">
  <link href="/static/bootstrap5.3/bootstrap.min.css" rel="stylesheet" integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">

  <!-- Favicons -->
  <link rel="apple-touch-icon" href="https://getbootstrap.com/docs/5.3/assets/img/favicons/apple-touch-icon.png" sizes="180x180">
  <link rel="icon" href="https://getbootstrap.com/docs/5.3/assets/img/favicons/favicon-32x32.png" sizes="32x32" type="image/png">
  <link rel="icon" href="https://getbootstrap.com/docs/5.3/assets/img/favicons/favicon-16x16.png" sizes="16x16" type="image/png">
  <link rel="manifest" href="https://getbootstrap.com/docs/5.3/assets/img/favicons/manifest.json">
  <link rel="mask-icon" href="https://getbootstrap.com/docs/5.3/assets/img/favicons/safari-pinned-tab.svg" color="#712cf9">
  <link rel="icon" href="https://getbootstrap.com/docs/5.3/assets/img/favicons/favicon.ico">
  <meta name="theme-color" content="#712cf9">
  <style>
    .bd-placeholder-img {
      font-size: 1.125rem;
      text-anchor: middle;
      -webkit-user-select: none;
      -moz-user-select: none;
      user-select: none;
    }

    @media (min-width: 768px) {
      .bd-placeholder-img-lg {
        font-size: 3.5rem;
      }
    }

    .b-example-divider {
      width: 100%;
      height: 3rem;
      background-color: rgba(0, 0, 0, .1);
      border: solid rgba(0, 0, 0, .15);
      border-width: 1px 0;
      box-shadow: inset 0 .5em 1.5em rgba(0, 0, 0, .1), inset 0 .125em .5em rgba(0, 0, 0, .15);
    }

    .b-example-vr {
      flex-shrink: 0;
      width: 1.5rem;
      height: 100vh;
    }

    .bi {
      vertical-align: -.125em;
      fill: currentColor;
    }

    .nav-scroller {
      position: relative;
      z-index: 2;
      height: 2.75rem;
      overflow-y: hidden;
    }

    .nav-scroller .nav {
      display: flex;
      flex-wrap: nowrap;
      padding-bottom: 1rem;
      margin-top: -1px;
      overflow-x: auto;
      text-align: center;
      white-space: nowrap;
      -webkit-overflow-scrolling: touch;
    }

    .btn-bd-primary {
      --bd-violet-bg: #712cf9;
      --bd-violet-rgb: 112.520718, 44.062154, 249.437846;

      --bs-btn-font-weight: 600;
      --bs-btn-color: var(--bs-white);
      --bs-btn-bg: var(--bd-violet-bg);
      --bs-btn-border-color: var(--bd-violet-bg);
      --bs-btn-hover-color: var(--bs-white);
      --bs-btn-hover-bg: #6528e0;
      --bs-btn-hover-border-color: #6528e0;
      --bs-btn-focus-shadow-rgb: var(--bd-violet-rgb);
      --bs-btn-active-color: var(--bs-btn-hover-color);
      --bs-btn-active-bg: #5a23c8;
      --bs-btn-active-border-color: #5a23c8;
    }

    .bd-mode-toggle {
      z-index: 1500;
    }

    .bd-mode-toggle .dropdown-menu .active .bi {
      display: block !important;
    }


  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  body {
    font-family: Arial, sans-serif;
    /*background: #f5f5f5;*/
  }

  .container {
    max-width: 1400px;
    margin: 0 auto;
    padding: 20px;
  }

  .grid-container {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 20px;
    margin-bottom: 30px;
  }

  /* 调整原有.model-item样式 */
  .model-item {
    position: relative;
    background: white;
    border-radius: 4px;
    overflow: hidden;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    transition: transform 0.2s;
    height: 100%;
    /* 新增高度设置 */
  }

  .model-item:hover {
    transform: translateY(-5px);
  }

  .thumbnail {
    width: 100%;
    height: 200px;
    object-fit: cover;
  }

  .info-overlay {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    background: linear-gradient(to top, rgba(0, 0, 0, 0.8), transparent);
    color: white;
    padding: 15px;
    opacity: 0;
    transition: opacity 0.3s;
  }

  .model-item:hover .info-overlay {
    opacity: 1;
  }

  .model-title {
    font-weight: bold;
    margin-bottom: 8px;
  }

  .model-meta {
    font-size: 0.9em;
    opacity: 0.8;
  }

  .pagination {
    display: flex;
    justify-content: center;
    gap: 10px;
    padding: 20px 0;
  }

  .page-btn {
    padding: 8px 15px;
    border: 1px solid #ddd;
    /*background: white;*/
    cursor: pointer;
    border-radius: 4px;
  }

  .page-btn.active {
    background: #007bff;
    color: white;
    border-color: #007bff;
  }

  .page-btn:hover:not(.active) {
    background: #f8f9fa;
  }

  /* 新增链接样式 */
  .model-link {
    display: block;
    text-decoration: none;
    color: inherit;
  }

  /* 原有样式保持不变，新增以下样式 */
  .loading-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(255, 255, 255, 0.9);
    display: none;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }

  .spinner {
    width: 40px;
    height: 40px;
    border: 4px solid #f3f3f3;
    border-top: 4px solid #3498db;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }

  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }

    100% {
      transform: rotate(360deg);
    }
  }

  .detail-container {
    display: none;
    padding: 20px;
    max-width: 1200px;
    margin: 0 auto;
  }

  .back-button {
    margin-bottom: 20px;
    padding: 10px 20px;
    cursor: pointer;
  }

  .error-message {
    color: red;
    text-align: center;
    padding: 20px;
  }
  </style>
</head>

<body>
  <svg xmlns="http://www.w3.org/2000/svg" class="d-none">
    <symbol id="check2" viewBox="0 0 16 16">
      <path d="M13.854 3.646a.5.5 0 0 1 0 .708l-7 7a.5.5 0 0 1-.708 0l-3.5-3.5a.5.5 0 1 1 .708-.708L6.5 10.293l6.646-6.647a.5.5 0 0 1 .708 0z"></path>
    </symbol>
    <symbol id="circle-half" viewBox="0 0 16 16">
      <path d="M8 15A7 7 0 1 0 8 1v14zm0 1A8 8 0 1 1 8 0a8 8 0 0 1 0 16z"></path>
    </symbol>
    <symbol id="moon-stars-fill" viewBox="0 0 16 16">
      <path d="M6 .278a.768.768 0 0 1 .08.858 7.208 7.208 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277.527 0 1.04-.055 1.533-.16a.787.787 0 0 1 .81.316.733.733 0 0 1-.031.893A8.349 8.349 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.752.752 0 0 1 6 .278z"></path>
      <path d="M10.794 3.148a.217.217 0 0 1 .412 0l.387 1.162c.173.518.579.924 1.097 1.097l1.162.387a.217.217 0 0 1 0 .412l-1.162.387a1.734 1.734 0 0 0-1.097 1.097l-.387 1.162a.217.217 0 0 1-.412 0l-.387-1.162A1.734 1.734 0 0 0 9.31 6.593l-1.162-.387a.217.217 0 0 1 0-.412l1.162-.387a1.734 1.734 0 0 0 1.097-1.097l.387-1.162zM13.863.099a.145.145 0 0 1 .274 0l.258.774c.115.346.386.617.732.732l.774.258a.145.145 0 0 1 0 .274l-.774.258a1.156 1.156 0 0 0-.732.732l-.258.774a.145.145 0 0 1-.274 0l-.258-.774a1.156 1.156 0 0 0-.732-.732l-.774-.258a.145.145 0 0 1 0-.274l.774-.258c.346-.115.617-.386.732-.732L13.863.1z"></path>
    </symbol>
    <symbol id="sun-fill" viewBox="0 0 16 16">
      <path d="M8 12a4 4 0 1 0 0-8 4 4 0 0 0 0 8zM8 0a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 0zm0 13a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 13zm8-5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 .5.5zM3 8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2A.5.5 0 0 1 3 8zm10.657-5.657a.5.5 0 0 1 0 .707l-1.414 1.415a.5.5 0 1 1-.707-.708l1.414-1.414a.5.5 0 0 1 .707 0zm-9.193 9.193a.5.5 0 0 1 0 .707L3.05 13.657a.5.5 0 0 1-.707-.707l1.414-1.414a.5.5 0 0 1 .707 0zm9.193 2.121a.5.5 0 0 1-.707 0l-1.414-1.414a.5.5 0 0 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .707zM4.464 4.465a.5.5 0 0 1-.707 0L2.343 3.05a.5.5 0 1 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .708z"></path>
    </symbol>
  </svg>
  <div class="dropdown position-fixed bottom-0 end-0 mb-3 me-3 bd-mode-toggle">
    <button class="btn btn-bd-primary py-2 dropdown-toggle d-flex align-items-center" id="bd-theme" type="button" aria-expanded="false" data-bs-toggle="dropdown" aria-label="Toggle theme (auto)">
      <svg class="bi my-1 theme-icon-active" width="1em" height="1em">
        <use href="#circle-half"></use>
      </svg>
      <span class="visually-hidden" id="bd-theme-text">Toggle theme</span>
    </button>
    <ul class="dropdown-menu dropdown-menu-end shadow" aria-labelledby="bd-theme-text">
      <li>
        <button type="button" class="dropdown-item d-flex align-items-center" data-bs-theme-value="light" aria-pressed="false">
          <svg class="bi me-2 opacity-50 theme-icon" width="1em" height="1em">
            <use href="#sun-fill"></use>
          </svg>
          Light
          <svg class="bi ms-auto d-none" width="1em" height="1em">
            <use href="#check2"></use>
          </svg>
        </button>
      </li>
      <li>
        <button type="button" class="dropdown-item d-flex align-items-center" data-bs-theme-value="dark" aria-pressed="false">
          <svg class="bi me-2 opacity-50 theme-icon" width="1em" height="1em">
            <use href="#moon-stars-fill"></use>
          </svg>
          Dark
          <svg class="bi ms-auto d-none" width="1em" height="1em">
            <use href="#check2"></use>
          </svg>
        </button>
      </li>
      <li>
        <button type="button" class="dropdown-item d-flex align-items-center active" data-bs-theme-value="auto" aria-pressed="true">
          <svg class="bi me-2 opacity-50 theme-icon" width="1em" height="1em">
            <use href="#circle-half"></use>
          </svg>
          Auto
          <svg class="bi ms-auto d-none" width="1em" height="1em">
            <use href="#check2"></use>
          </svg>
        </button>
      </li>
    </ul>
  </div>
  <header data-bs-theme="dark">
    <div class="collapse text-bg-dark" id="navbarHeader">
      <div class="container">
        <div class="row">
          <div class="col-sm-8 col-md-7 py-4">
            <h4>关于</h4>
            <p class="text-body-secondary">FreeCAD模型库展示页面。全参数化模型，可通过调整参数驱动模型。</p>
          </div>
          <div class="col-sm-4 offset-md-1 py-4">
            <h4>Contact</h4>
            <ul class="list-unstyled">
              <li><a href="/docs/freecad" class="text-white">建模文档</a></li>
              <li><a href="/docs/freecad/freecad-020-004" class="text-white">接口文档</a></li>
              {{if eq true .IsLogin}}
              <li><a href="/user" class="text-white">{{.Username}}</a></li>
              {{else}}
              <li><a href="/login?url=/v1/freecad/modellist" class="text-white">用户登录</a></li>
              {{end}}
            </ul>
          </div>
        </div>
      </div>
    </div>
    <div class="navbar navbar-dark bg-dark shadow-sm">
      <div class="container">
        <a href="#" class="navbar-brand d-flex align-items-center">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" aria-hidden="true" class="me-2" viewBox="0 0 24 24">
            <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"></path>
            <circle cx="12" cy="13" r="4"></circle>
          </svg>
          <strong>模型库展示</strong>
        </a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarHeader" aria-controls="navbarHeader" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
      </div>
    </div>
  </header>
  <main>
    <section class="py-5 text-center container">
      <div class="row py-lg-5">
        <div class="col-lg-6 col-md-8 mx-auto">
          <h1 class="fw-light">FreeCAD参数化三维模型库</h1>
          <p class="lead text-body-secondary">一次建模，全球受用。</p>
          <p>
            <a href="/docs/freecad" class="btn btn-primary my-2">建模文档</a>
            <a href="/v1/freecad/uploadmodel" class="btn btn-secondary my-2">上传模型</a>
          </p>
        </div>
      </div>
    </section>
    <div class="album py-5 bg-body-tertiary">
      <!-- <div class="container"> -->
        <!-- <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 g-3"> -->
  <!-- 保持之前的HTML结构不变 -->
      <div class="loading-overlay" id="loading">
        <div class="spinner"></div>
      </div>
      <div class="container" id="mainContainer">
        <div class="grid-container" id="modelGrid"></div>
        <div class="pagination" id="pagination"></div>
        <div class="error-message" id="errorContainer"></div>
      </div>
      <!-- 保持详情容器不变 -->
      <div class="detail-container" id="detailContainer">
        <button class="back-button" id="backButton">返回列表</button>
        <div id="detailContent"></div>
      </div>          
        <!-- </div> -->
      <!-- </div> -->
    </div>
  </main>
  <footer class="text-body-secondary py-5">
    <div class="container">
      <p class="float-end mb-1">
        <a href="/v1/freecad/modellist/#">Back to top</a>
      </p>
      <p class="mb-1">Album is © Bootstrap v5.3 !</p>
      <p class="mb-0">New to Bootstrap? <a href="https://getbootstrap.com/">Visit the homepage</a> or read our <a href="https://getbootstrap.com/docs/5.3/getting-started/introduction/">getting started guide</a>.</p>
    </div>
  </footer>

  <script src="/static/bootstrap5.3/jquery.min.js"></script>
  <script src="/static/bootstrap5.3/bootstrap.bundle.min.js" integrity="sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL" crossorigin="anonymous"></script>

  <script>
  // 状态管理对象
  const state = {
    currentPage: 1,
    itemsPerPage: 20,
    currentModel: null,
    models: [] // 改为从服务器获取
  };

  // 新增API配置
  const API = {
    MODELS: '/v1/freecad/modellistdata', // 根据实际API地址修改
    MODEL_DETAIL: id => `/v1/freecad/model/${id}`
  };

  // GET /api/models 响应示例
// [
//   {
//     "id": 1,
//     "title": "机械手臂",
//     "author": "张三",
//     "createdAt": "2023-08-15",
//     "views": 1234,
//     "thumbnailUrl": "/thumbnails/1.jpg",
//     "detailUrl": "/models/1"
//   }
//   // ...
// ]

// GET /api/models/1 响应示例
// {
//   "id": 1,
//   "title": "机械手臂",
//   "author": "张三",
//   "createdAt": "2023-08-15",
//   "views": 1234,
//   "previewUrl": "/previews/1.jpg",
//   "description": "详细描述...",
//   "downloadUrl": "/downloads/1.stl",
//   "tags": ["机械", "工业"]
// }

  // 修改后的初始化函数
  async function init() {
    try {
      showLoading();
      await fetchModels();
      // parseURLParameters();
      setupEventListeners();
      renderPage();
    } catch (error) {
      showError(error.message);
    } finally {
      hideLoading();
    }
  }

  // 修改fetchModels函数
  async function fetchModels(page = 1) {
    const response = await fetch(`${API.MODELS}?page=${page}&perPage=${state.itemsPerPage}`);
    const data = await response.json();
    
    state.models = data.items;
    state.total = data.total;
    state.perPage = data.perPage;
  }

  // 修改分页处理逻辑
  function getPaginatedModels() {
    return state.models; // 直接使用服务端返回的分页数据
  }

  // 修改后的renderModels函数
  function renderModels() {
    const currentModels = getPaginatedModels();
    elements.modelGrid.innerHTML = currentModels.map(model => `
        <a href="${model.modelurl}&id=${model.ID}" class="model-link">
            <div class="model-item">
                <img src="${model.faceimgurl}" class="thumbnail" alt="${model.title}">
                <div class="info-overlay">
                    <div class="model-title">${model.titleb}</div>
                    <div class="model-meta">
                        <div>作者：${model.author}</div>
                        <div>日期：${model.CreatedAt}</div>
                        <div>浏览：${model.views}</div>
                    </div>
                </div>
            </div>
        </a>
    `).join('');
  }

  // 修改后的分页渲染
  function renderPagination() {
    const totalPages = Math.ceil(state.total / state.perPage) // Math.ceil(state.models.length / state.itemsPerPage);
    const pageBtns = [];

    for (let i = 1; i <= totalPages; i++) {
      pageBtns.push(
        `<button class="page-btn ${i === state.currentPage ? 'active' : ''}">${i}</button>`
      );
    }

    elements.pagination.innerHTML = `
        <button class="page-btn prev">上一页</button>
        ${pageBtns.join('')}
        <button class="page-btn next">下一页</button>
    `;
  }

  // 新增错误显示函数
  function showError(message) {
    elements.errorContainer.innerHTML = `
        <div class="error-message">⚠️ ${message}</div>
    `;
  }

  // 修改后的模型点击处理
  async function handleModelClick(e) {
    const link = e.target.closest('.model-link');
    if (!link) return;

    e.preventDefault();
    // showLoading();

    if (link) {
      e.preventDefault();
      const url = link.getAttribute('href');
      // 这里可以添加页面跳转逻辑
      console.log('跳转到：', url);
      var glbpath = url.replace(".FCStd",".glb")
      window.location.href = "/v1/freecad/online3deditor?model="+glbpath;
    }
  }

  // 修改后的显示详情函数
  function showModelDetail(model) {
    state.currentModel = model;

    elements.mainContainer.style.display = 'none';
    elements.detailContainer.style.display = 'block';

    elements.detailContent.innerHTML = `
        <h2>${model.title}</h2>
        <div class="meta-info">
            <p>作者：${model.author}</p>
            <p>发布日期：${model.CreatedAt}</p>
            <p>浏览量：${model.views}</p>
        </div>
        <img src="${model.faceimgurl}" style="max-width: 100%">
        <p>${model.description}</p>
    `;
  }


  // 增强版模型数据
  // const models = Array.from({ length: 50 }, (_, i) => ({
  //   id: i + 1,
  //   title: `模型 ${i + 1}`,
  //   author: `设计师 ${i % 10 + 1}`,
  //   date: `2023-0${(i % 12) + 1}-${(i % 28) + 1}`,
  //   views: Math.floor(Math.random() * 1000),
  //   img: `https://picsum.photos/300/200?random=${i}`,
  //   detailUrl: `/models/${i+1}`,
  //   description: `这是模型 ${i+1} 的详细描述，包含详细的技术参数和设计说明。`
  // }));

  // DOM元素缓存
  const elements = {
    loading: document.getElementById('loading'),
    mainContainer: document.getElementById('mainContainer'),
    detailContainer: document.getElementById('detailContainer'),
    backButton: document.getElementById('backButton'),
    modelGrid: document.getElementById('modelGrid'),
    pagination: document.getElementById('pagination'),
    errorContainer: document.getElementById('errorContainer'),
    detailContent: document.getElementById('detailContent')
  };

  // URL参数解析
  function parseURLParameters() {
    const params = new URLSearchParams(window.location.search);
    state.currentPage = parseInt(params.get('page')) || 1;

    const pathParts = window.location.pathname.split('/');
    if (pathParts[1] === 'models' && pathParts[2]) {
      showModelDetail(parseInt(pathParts[2]));
    }
  }

  // 事件监听
  function setupEventListeners() {
    // 分页点击
    elements.pagination.addEventListener('click', handlePaginationClick);

    // 模型点击
    elements.modelGrid.addEventListener('click', handleModelClick);

    // 返回按钮
    elements.backButton.addEventListener('click', () => {
      history.back();
    });

    // 历史记录变化
    window.addEventListener('popstate', handlePopState);
  }

  // 处理分页点击
  async function handlePaginationClick(e) {
    if (!e.target.classList.contains('page-btn')) return;

    e.preventDefault();
    const btn = e.target;

    if (btn.classList.contains('prev')) {
      state.currentPage = Math.max(1, state.currentPage - 1);
    } else if (btn.classList.contains('next')) {
      state.currentPage = Math.min(
        Math.ceil(state.total / state.perPage),
        state.currentPage + 1
      );
    } else if (!isNaN(btn.textContent)) {
      state.currentPage = parseInt(btn.textContent);
    }

    // updateURL({ page: state.currentPage });

    try {
      showLoading();
      await fetchModels(state.currentPage);
      // parseURLParameters();
      // setupEventListeners();
      renderPage();
    } catch (error) {
      showError(error.message);
    } finally {
      hideLoading();
    }

    // renderPage();
  }

  // 显示加载动画
  function showLoading() {
    elements.loading.style.display = 'flex';
  }

  // 隐藏加载动画
  function hideLoading() {
    elements.loading.style.display = 'none';
  }

  // 更新URL
  function updateURL(path, replace = false) {
    const method = replace ? 'replaceState' : 'pushState';
    history[method]({
      page: state.currentPage,
      model: state.currentModel?.ID
    }, '', path);
  }

  // 处理历史记录变化
  function handlePopState(e) {
    if (window.location.pathname === '/') {
      elements.mainContainer.style.display = 'block';
      elements.detailContainer.style.display = 'none';
      parseURLParameters();
      renderPage();
    } else {
      parseURLParameters();
    }
  }

  // 渲染主页面
  function renderPage() {
    renderModels();
    renderPagination();
  }

  // 初始化应用
  init();
  </script>

  <script type="text/javascript">
  // var i = 1;
  // var flag = false; // 防止多次调用下拉触发事件
  // var key = "";

  // $(function() {
    // var loadData = function(ii) {
    //   $.getJSON("/v1/freecad/freecaddata?page=" + ii, function(data) {
    //     if (data.length == 0) {
    //       $("#loadMore").removeClass('hidden').text('已加载全部数据！');
    //       key = "over"
    //     } else if (data.length > 0) {
    //       $.each(data, function(i, fcmodellist) {
    //         var glbpath = fcmodellist.path.replace("./","/").replace(".FCStd",".glb")
    //         var facepath = fcmodellist.faceimgurl.replace("./","/")
    //         var _card_ = ""
    //         _card_ = _card_ + '<div class="card shadow-sm"><img src="'+facepath+'" class="bd-placeholder-img card-img-top" width="100%" height="225" href="coffee.html"><title>'+fcmodellist.titleb+'</title><div class="card-body"><p class="card-text">'+fcmodellist.titleb+'，'+fcmodellist.description+'<br>'+"特征参数："+fcmodellist.indicators+'<br>作者：'+fcmodellist.User.Nickname+'</p><div class="d-flex justify-content-between align-items-center"><div class="btn-group"><a href="/v1/freecad/online3dview#model='+glbpath+'" class="btn btn-sm btn-outline-secondary">View</a><a href="/v1/freecad/online3deditor?model='+glbpath+'&id='+fcmodellist.ID+'" class="btn btn-sm btn-outline-secondary">Edit</a><a href="/v1/freecad/update?id='+fcmodellist.ID+'" class="btn btn-sm btn-outline-secondary">Update</a></div><small class="text-body-secondary">'+"00 下载"+'</small><small class="text-body-secondary">'+"00 RMB"+'</small></div></div></div>'
    //         $(".row-cols-md-3").append('<div class="col">' + _card_ + '</div>')
    //       });
    //     }
    //     if (data.length >= 9) {
    //       i = i + 1;
    //     } else if (data.length < 9) {
    //       key = "over"
    //     }
    //   })
    // };

    // loadData(1);

    // var tcScroll = function() {
    //   $(window).on('scroll', function() {
    //     // var scrollTop = document.documentElement.scrollTop
    //     var clientHeight = document.documentElement.clientHeight //整个网页的高度（包括未显示的部分）
    //     var scrollTop = $(this).scrollTop(); //滚动条滚动上去的高度
    //     var bodyHeight = document.body.clientHeight //设备可见区域高度
    //     // console.log(scrollTop,clientHeight,bodyHeight)//1706 2205 499   631 760 1466   705 760 1466   706 760 1466
    //     if (bodyHeight - scrollTop - clientHeight  <= 0 && flag === false) {
    //       //此处是滚动条到底部时候触发的事件，在这里写要加载的数据，或者是拉动滚动条的操作
    //       flag = true;
    //       console.log("触底了!!!!");
    //       //此处是滚动条到底部时候触发的事件，在这里写要加载的数据，或者是拉动滚动条的操作
    //       if (key != "over"){
    //         loadData(i);
    //       }
    //       setTimeout(() => {
    //         flag = false;
    //       }, 500);
    //     }
    //   })
    // }
    // tcScroll();
  // });
  </script>
</body>
</html>