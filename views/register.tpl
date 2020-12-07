<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html charset=utf-8">
</head>

<body>
   <h2>Register</h2>
<form action="/register" method="POST">
    Username: <input type="text" name="username">
    Email: <input type="email" name="email">
    Password: <input type="password" name="password">
    Password Confirm: <input type="password" name="passwordConfirm">
     <h2>
    {{.Email}}
     
     <h2>
    <button type="submit" >SignUp</button>
</form>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
