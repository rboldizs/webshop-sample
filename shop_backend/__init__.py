"""Flask Factory"""
import os

from flask import Flask
from flask_cors import CORS


def create_app(test_config=None):
    """Creates Flask application"""
    # create and configure the app
    app = Flask(__name__, instance_relative_config=True)
    CORS(app)
    app.config.from_mapping(
        SECRET_KEY='valid_key',
    )

    if test_config is None:
        # load the instance config, if it exists, when not testing
        app.config.from_pyfile('config.py', silent=True)
    else:
        # load the test config if passed in
        app.config.from_mapping(test_config)

    # ensure the instance folder exists
    try:
        os.makedirs(app.instance_path)
    except OSError:
        pass

    from . import uri_root
    app.register_blueprint(uri_root.BP)

    from . import api
    app.register_blueprint(api.BP)

    return app
    