"""Root blueprint with welcome msg"""

from flask import (
    Blueprint, render_template, request
)

bp = Blueprint('root', __name__, url_prefix='/')

@bp.route('/', methods=('GET', 'POST'))
def root():
    """Root blueprint will generate a view with html template"""
    if request.method == 'GET':
        return render_template('root/root.html')
