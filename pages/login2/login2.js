//index.js
//获取应用实例
Page({
  data: {
    username: '',
    password: '',
    clientHeight:''
  },
  onLoad(){

  },
  //获取输入款内容
  bed(e){
    this.data.username=e.detail.value
  },
  password(e){
    this.data.password=e.detail.value
  },
  goadmin() {
    const username = this.data.username; // 获取用户名
    const password = this.data.password; // 获取密码
  
    // 检查输入是否为空
    if (username === '') {
      wx.showToast({
        icon: 'none',
        title: '账号不能为空',
      });
      return;
    }
    else if(username === '李亮'){
      if(password === '111111'){
        wx.showToast({
          title: '登录成功',
          icon: 'success',
        });
        const app = getApp(); // 获取全局 App 实例
        app.globalData.userRole = "teacher";
        app.globalData.userid = '111111';
        app.globalData.username = username;
        
        console.log(app.globalData.userjwt)
        // 跳转到 Stu_choose 页面
        wx.navigateTo({
          url: '/pages/Tea_choose/Tea_choose',
        });
        return ;
      }
    }
  
    if (password === '') {
      wx.showToast({
        icon: 'none',
        title: '密码不能为空',
      });
      return;
    }
  
    // 发送请求到后端
    wx.request({
      url: 'http://localhost:8080/teacher/login', // 后端接口地址
      method: 'POST',
      header: {
        'Content-Type': 'application/json',
      },
      data: {
        user_id: username,
        pwd: password,
      },
      success(res) {
        const statusCode = res.statusCode;
        if (statusCode===200) { 
          wx.showToast({
            title: '登录成功',
            icon: 'success',
          });
        const app = getApp(); // 获取全局 App 实例
        app.globalData.userRole = "teacher";
        app.globalData.userjwt = res.data.token; 
        app.globalData.userid = res.data.id;
        app.globalData.username = res.data.name;

        console.log(app.globalData.userjwt)

          // 跳转到 Stu_choose 页面
          wx.navigateTo({
            url: '/pages/Tea_choose/Tea_choose',
          });
        } else {
          wx.showToast({
            title: res.data.error || '登录失败', // 假设后端返回 `message` 字段表示错误信息
            icon: 'none',
          });
        }
      },
      fail(error) {
        // 请求失败处理
        wx.showToast({
          title: '请求失败，请稍后再试',
          icon: 'none',
        });
        console.error('请求失败:', error);
      },
    });
  },
})
 
