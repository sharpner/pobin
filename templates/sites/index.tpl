{% extends "layout/base.tpl" %}
{% block title %}
  Example login page 
{% endblock %}
{% block content %}
    {% import "macros/forms.tpl" login_form %}
    <div>{{ login_form(action) }}</div>
{% endblock %}
