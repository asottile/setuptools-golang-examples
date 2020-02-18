import red


def test_red():
    assert red.red('ohai') == '\x1b[31mohai\x1b[0m'
