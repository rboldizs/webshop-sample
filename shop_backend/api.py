"""Contains all the api blueprints"""

import os
import codecs
from bs4 import BeautifulSoup
from flask import (
    Blueprint, render_template, request, json, Response
)

BP = Blueprint('api', __name__, url_prefix='/api')

@BP.route('/status', methods=('GET', 'POST'))
def status():
    """Status blueprint will create a view with build info"""

    ##Added Code smell as an example, to be detected by the SonarQube##
    if request.method == 'GET':
        if request.user_agent.browser is None:
            return render_json()

    if request.method == 'GET':
        return render_template('api/status.html')

    return 'Not available'


def render_json():
    """Will create a json response with build and version info"""
    try:
        os.chdir(os.path.dirname(__file__))
        tpl_html = codecs.open("templates/api/status.html", 'r', 'utf-8')
        soup = BeautifulSoup(tpl_html, 'html.parser')
        nodes = soup.find_all('bold')
        labels = soup.find_all('label')
        jdict = dict()
        i = 0
        for label in labels:
            jdict[nodes[i].string] = label.string
            i = i + 1
        json_data = json.dumps(jdict)
        resp = Response(json_data, status=200, mimetype='application/json')
        return resp
    except OSError as os_err:
        resp = Response(str(os_err), status=400, mimetype='application/text')
        return resp
    except (ValueError, KeyError, TypeError) as error:
        resp = Response(str(error), status=400, mimetype='application/text')
        return resp
