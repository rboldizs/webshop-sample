from shop_backend import create_app

def test_uri_root(client):
    '''
        Test case No.5: Flask root endpoint call check
        Application should return greeting html page
        Assert will be thrown if it is not the case
    '''
    resp = client.get('/')
    assert 'Welcome to shop!' in resp.get_data(as_text=True)