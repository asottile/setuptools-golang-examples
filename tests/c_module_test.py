import c_module


def test_c_module():
    assert c_module.hello_world() == 'hello world'
