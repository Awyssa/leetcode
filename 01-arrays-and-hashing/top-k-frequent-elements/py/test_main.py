import pytest
from main import Solution


class TestTopKFrequent:
    def setup_method(self):
        self.solution = Solution()

    def test_example_1(self):
        """Test Example 1: [1,1,1,2,2,3], k=2 should return [1,2]"""
        result = self.solution.topKFrequent([1, 1, 1, 2, 2, 3], 2)
        assert sorted(result) == [1, 2]

    def test_example_2(self):
        """Test Example 2: [1], k=1 should return [1]"""
        result = self.solution.topKFrequent([1], 1)
        assert result == [1]

    def test_same_frequency(self):
        """Test all elements have same frequency"""
        result = self.solution.topKFrequent([1, 2, 3], 2)
        assert len(result) == 2

    def test_larger_array_with_negatives(self):
        """Test larger array with negative numbers"""
        result = self.solution.topKFrequent([4, 1, -1, 2, -1, 2, 3], 2)
        assert sorted(result) == [-1, 2]
