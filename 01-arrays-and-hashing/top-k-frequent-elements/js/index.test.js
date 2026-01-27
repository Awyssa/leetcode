const topKFrequent = require('./index.js');

// Helper function to compare arrays (order doesn't matter)
const arraysEqualUnordered = (arr1, arr2) => {
	if (arr1.length !== arr2.length) return false;
	const sorted1 = [...arr1].sort((a, b) => a - b);
	const sorted2 = [...arr2].sort((a, b) => a - b);
	return sorted1.every((val, idx) => val === sorted2[idx]);
};

describe('topKFrequent', () => {
	test('Example 1: [1,1,1,2,2,3], k=2 should return [1,2]', () => {
		const result = topKFrequent([1, 1, 1, 2, 2, 3], 2);
		expect(arraysEqualUnordered(result, [1, 2])).toBe(true);
	});

	test('Example 2: [1], k=1 should return [1]', () => {
		const result = topKFrequent([1], 1);
		expect(result).toEqual([1]);
	});

	test('All elements have same frequency', () => {
		const result = topKFrequent([1, 2, 3], 2);
		expect(result).toHaveLength(2);
	});

	test('Larger array with negative numbers', () => {
		const result = topKFrequent([4, 1, -1, 2, -1, 2, 3], 2);
		expect(arraysEqualUnordered(result, [-1, 2])).toBe(true);
	});
});
