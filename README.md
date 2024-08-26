# ToDo

A simple REST API pet project written in Go, application performs basic API operations with users/lists/tasks etc., includes basic JWT token-based authentication<br><br>
Currently, API is built for a user-specific use, meaning that all endpoints returns information based on the authorization (JWT) token bearer. Viewing other users' profile information, lists and tasks is not implied by this application, so there are no endpoints to do so.<br><br>

<h2>API endpoints:</h2>

<hr>

<h4>Users</h4>
<ul>
    <li>
        <b>GET /api/users/<br></b>
        Returns the token bearer's information
    </li><br>
    <li>
        <b>POST /api/users/<br></b>
        Creates a new user
    </li>
</ul>

<hr>

<h4>Lists</h4>
<ul>
    <li>
        <b>GET /api/lists/</b><br>
        Returns the token bearer's lists
    </li><br>
    <li>
        <b>GET /api/lists/{id}</b><br>
        Returns a list specified by id. If a list doesn't belong to token bearer request will fail
    </li><br>
    <li>
        <b>POST /api/lists/</b><br>
        Creates a new list
    </li>
</ul>

<hr>

<h4>Tasks</h4>
<ul>
    <li>
    <b>GET /api/tasks/</b><br>
    Parameters:
      <ul>
        <li><i>Optional</i> list_id</li>
      </ul><br>
    Returns tasks specified by the <i>list_id</i> parameter, if no <i>list_id</i> parameter is provided, this endpoint will return all the token bearer's tasks. If list specified by <i>list_id</i> doesn't belong to the token bearer request will fail
  </li><br>
  <li>
    <b>GET /api/tasks/{id}</b><br>
    Returns a task specified by id. If a task doesn't belong to token bearer request will fail
  </li><br>
  <li>
    <b>POST /api/tasks/</b><br>
    Creates a new task
  </li>
</ul>

<hr>
  
<h4>Tags</h4>
<ul>
    <li>
    <b>GET /api/tags/</b><br>
    Parameters:
      <ul>
        <li><i>Optional</i> task_id</li>
      </ul><br>
    Returns tags specified by the <i>task_id</i> parameter, if no <i>task_id</i> parameter is provided, this endpoint will return all the token bearer's tags. If task specified by <i>task_id</i> doesn't belong to the token bearer request will fail
  </li><br>
  <li>
    <b>GET /api/tags/{id}</b><br>
    Returns a tag specified by id. If a tag doesn't belong to token bearer request will fail
  </li><br>
  <li>
    <b>POST /api/tags/</b><br>
    Creates a new tag
  </li>
</ul>

<hr>

<h4>Other</h4>
<ul>
    <li>
        <b>POST /api/login/</b><br>
        Used for authentication. A JWT token is returned if credentials are correct
    </li><br>
    <li>
        <b>PATCH /api/emails/verify/</b><br>
        Parameters:
        <ul>
            <li><i>Optional</i> t</li>
        </ul><br>
        Used for email verification. This endpoint will email current user if no <i>t</i> (token) parameter is provided. When user sends another request with <i>t</i> parameter provided, application validates the token and changes the user's email status to "verified" if token is valid
    </li><br>
</ul>

<h2>Used Frameworks/Packages:</h2>
<ul>
  <li>Go Fiber</li>
  <li>jmoiron/sqlx to perform SQL queries (used database is MySQL)</li>
  <li>x/crypto/bcrypt for encryption and hashing passwords</li>
  <li>golang-jwt/jwt/v5 for operations with JWT</li>
  <li>validator/v10 for validating structs fields</li>
  <li>zhashkevych/go-sqlxmock for mocking sql driver behavior</li>
  <li>gopkg.in/gomail.v2 for sending emails</li>
</ul>
