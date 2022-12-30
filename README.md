<h3> GIN PAYMENT SERVICE </h3>
<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" height="100px" style="display: block" alt="gin gonic logo"/>

How to run ?
In the project directory, Add .env file first. </br>
env file contain this value:
JWT_SECRET = your jwt secret </br>
TOKEN_EXPIRE = your jwt token lifetime </br>

DB_HOST = your database host </br>
PORT = your database port </br>
DB_USER = your database username</br>
DB_PASSWORD = your database password</br>
DB_NAME = your database name </br>

API Reference:
<table>
<thead>
    <td>Api End point</td> <td> Description</td>
</thead>
<tbody>
    <tr>
        <td>POST /api/auth/register</td>
        <td>user registration</td>
    </tr>
    <tr>
        <td>POST /api/auth/login</td>
        <td>user login and get token</td>
    </tr>
    <tr>
        <td>POST /api/customer/register</td>
        <td>customer registration (must register for user first)</td>
    </tr>
</tbody>
</table>
<h3>Api request Example</h3>
This is example request for api endpoint <br>
POST /api/auth/register -> <br>
{ <br>
    "username":"your username",
<br>
    "password":"your password",
<br>
    "passwordVerify":"verify password"
<br>
}<br> <br>
POST /api/auth/login -> <br>{ <br>
    "username":"your username",<br>
    "password":"your password"<br>
} <br> 
POST /api/customer/register -> <br>
{ <br>
    "customer_name":"iqbal"
<br>
}
<br>
*note: for customer endpoint, you need to logged in first and add jwt token on Authorization in header request<br>
