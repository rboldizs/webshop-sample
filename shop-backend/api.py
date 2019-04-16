import functools

from flask import (
    Blueprint, render_template, request
)

bp = Blueprint('api', __name__, url_prefix='/api')

@bp.route('/status', methods=('GET', 'POST'))
def status():
    if request.method == 'GET':
        return render_template('api/status.html')