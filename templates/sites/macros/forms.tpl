{% macro login_form(action) export %}
       <form method="POST" action="{{ action }}">
               <label for="login">Username</label><input id="login" type="text" name="login" />
               <label for="password">Kennwort</label><input id="password" type="password" name="password" />
               <input type="submit" />
       </form>
{% endmacro %}
