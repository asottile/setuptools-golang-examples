from __future__ import annotations

import pytest

import sum_go


@pytest.mark.parametrize(('a', 'b', 'c'), ((1, 2, 3), (4, 5, 9)))
def test_sum_go(a, b, c):
    assert sum_go.sum(a, b) == c
