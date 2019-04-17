"""Root blueprint with welcome msg"""

from flask import (
    Blueprint, render_template, request
)

BP = Blueprint('root', __name__, url_prefix='/')

@BP.route('/', methods=('GET', 'POST'))
def root():
    """Root blueprint will generate a view with html template"""
    if request.method == 'GET':
        return render_template('root/root.html')
    return None
