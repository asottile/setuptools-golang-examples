import hello_lib


def test_hello_lib():
    assert hello_lib.ohai('anthony') == 'ohai, anthony'
