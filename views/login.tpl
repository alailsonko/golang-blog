<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html charset=utf-8">
</head>

<body>
   <h2>Login</h2>
<form action="/login" method="POST">
    Email: <input type="email" name="email">
    Password: <input type="password" name="password">
     <h2>
    {{.Email}}
     
     <h2>
     {{.flash.error}}
    <button type="submit" >SignIn</button>
</form>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
