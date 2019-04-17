from shop_backend import create_app

def test_root(client):
    resp = client.get('/')
    assert 'Welcome to shop!' in resp.get_data(as_text=True)