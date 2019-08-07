from shop_backend import create_app

def test_config():
    '''
        Test case No.1: Flask Application creation
        Application should be created by the flask app factory, with testing parameter
        Assert will be thrown if it is not the case
    '''
    assert not create_app().testing
    assert create_app({'TESTING' : True}).testing

