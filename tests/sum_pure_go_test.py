import pytest

import sum_pure_go


@pytest.mark.parametrize(('a', 'b', 'c'), ((1, 2, 3), (4, 5, 9)))
def test_sum_pure_go(a, b, c):
    assert sum_pure_go.sum(a, b) == c
