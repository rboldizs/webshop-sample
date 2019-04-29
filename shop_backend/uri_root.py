"""URI Root blueprint with welcome msg"""

from flask import (
    Blueprint, render_template, request
)

BP = Blueprint('uri_root', __name__, url_prefix='/')

@BP.route('/', methods=('GET', 'POST'))
def uri_root():
    """URI Root blueprint will generate a view with html template"""
    if request.method == 'GET':
        return render_template('URI_root/URI_root.html')
    return None
