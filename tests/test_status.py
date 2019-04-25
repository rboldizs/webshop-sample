from shop_backend import create_app

def test_status(client):
    client.environ_base['HTTP_USER_AGENT'] = 'Mozilla/5.0 (Macintosh; U; Mac OS X 10.5; en-US; ) Firefox/3.1'
    resp = client.get('/api/status')           
    assert  200 == resp.status_code 
    assert 'Version:' in resp.get_data(as_text=True)
    assert 'Build id:' in resp.get_data(as_text=True)
    assert 'Build time:' in resp.get_data(as_text=True)
    assert 'Git commit:' in resp.get_data(as_text=True)

    resp = client.post('/api/status')
    assert  200 == resp.status_code 
    assert 'Not available' in resp.get_data(as_text=True)

def test_json(client):
    resp = client.get('/api/status')
    assert  200 == resp.status_code 
    assert 'Version:' in resp.get_data(as_text=True)
    assert 'Build id:' in resp.get_data(as_text=True)
    assert 'Build time:' in resp.get_data(as_text=True)
    assert 'Git commit:' in resp.get_data(as_text=True)
