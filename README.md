# ToDo

A simple REST API pet project written in Go, application performs basic API operations with users/lists/tasks etc., includes basic JWT token-based authentication<br><br>
Currently, API is built for a user-specific use, meaning that all endpoints returns information based on the authorization (JWT) token bearer. Viewing other users' profile information, lists and tasks is not implied by this application, so there are no endpoints to do so.<br><br>

<h2>API endpoints:</h2>
<ul>
  <li>
    <b>POST /api/login/</b><br>
    Used for authentication. A JWT token is returned if credentials are correct
  </li><br>
  <li>
    <b>POST /api/users/<br></b>
    Creates a new user
  </li><br>
  <li>
    <b>GET /api/users/<br></b>
    Returns the token bearer's information
  </li><br>
  <li>
    <b>GET /api/lists/</b><br>
    Returns the token bearer's lists
  </li><br>
  <li>
    <b>GET /api/tasks/</b><br>
    Parameters:
      <ul>
        <li><i>Optional</i> list_id</li>
      </ul><br>
    Returns tasks specified by the <i>list_id</i> parameter, if no <i>list_id</i> parameter is provided, this endpoing will return all the token bearer's tasks. If list specified by <i>list_id</i> doesn't belong to the token bearer request will fail
  </li>
</ul>

<h2>Used Frameworks/Packages:</h2>
<ul>
  <li>Go Fiber</li>
  <li>jmoiron/sqlx to perform SQL queries (used database is MySQL)</li>
  <li>x/crypto/bcrypt for encryption and hashing passwords</li>
  <li>golang-jwt/jwt/v5 for operations with JWT</li>
  <li>validator/v10 for validating structs fields</li>
</ul>
