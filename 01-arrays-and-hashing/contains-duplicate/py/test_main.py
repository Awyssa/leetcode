from main import Solution

def test_contains_duplicate_true():
    solution = Solution()
    assert solution.containsDuplicate([1,2,3,1]) == True

def test_contains_duplicate_false():
    solution = Solution()
    assert solution.containsDuplicate([1,2,3,4]) == False

def test_contains_duplicate_multiple():
    solution = Solution()
    assert solution.containsDuplicate([1,1,1,3,3,4,3,2,4,2]) == True