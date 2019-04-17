from shop_backend import create_app

def test_status(client):
    
    resp = client.get('/api/status')
    assert  200 == resp.status_code 
    assert 'Version:' in resp.get_data(as_text=True)
    assert 'Build id:' in resp.get_data(as_text=True)
    assert 'Build time:' in resp.get_data(as_text=True)
    assert 'Git commit:' in resp.get_data(as_text=True)
