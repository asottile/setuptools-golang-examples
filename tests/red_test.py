from __future__ import annotations

import red


def test_red():
    assert red.red('ohai') == '\x1b[0;31mohai\x1b[0m'
