window.οnlοad=function(){ 
 var button=document.getElementById("button");
  button.onclick = function () {
        var username = document.getElementById("username").value;
        var password = document.getElementById("password").value;
        if(username==""||username==null){
            alert("用户名为空");
            return;
        }
        if(password==""||password==null){
            alert("密码为空");
            return;
        }
   
        fetch("http://101.43.160.254:8080/user",{method:"get"})
        .then(function(response){
          alert("登录成功");
        }).then(function(rejected){
         alert("用户名或密码错误")
        });
    }
} 

