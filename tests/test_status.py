from shop_backend import create_app

def test_status(client):
    '''
        Test case No.2: Flask Status endpoint check
        Application should return version info with build parameters
        Assert will be thrown if it is not the case
        Note: This test case is mimic the browser header, therefore the rendered html is expected
    '''
    client.environ_base['HTTP_USER_AGENT'] = 'Mozilla/5.0 (Macintosh; U; Mac OS X 10.5; en-US; ) Firefox/3.1'
    resp = client.get('/api/status')           
    assert  200 == resp.status_code 
    assert 'Version:' in resp.get_data(as_text=True)
    assert 'Build id:' in resp.get_data(as_text=True)
    assert 'Build time:' in resp.get_data(as_text=True)
    assert 'Git commit:' in resp.get_data(as_text=True)

    '''
        Test case No.3: Post methods
        Application should return not available if the request method is POST
        Assert will be thrown if it is not the case
    '''
    resp = client.post('/api/status')
    assert  200 == resp.status_code 
    assert 'Not available' in resp.get_data(as_text=True)

def test_json(client):
    '''
        Test case No.4: Flask Status endpoint API call check
        Application should return version info with build parameters
        Assert will be thrown if it is not the case
        Note: This test case is mimic the non-browser call, therefore no rendered html is expected, but json response
    '''
    resp = client.get('/api/status')
    assert  200 == resp.status_code 
    assert 'Version:' in resp.get_data(as_text=True)
    assert 'Build id:' in resp.get_data(as_text=True)
    assert 'Build time:' in resp.get_data(as_text=True)
    assert 'Git commit:' in resp.get_data(as_text=True)
