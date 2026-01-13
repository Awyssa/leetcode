const groupAnagrams = (words) => {
	let returnObject = {};
	for (let word of words) {
		const sortedWord = word.split("").sort().join("");

		if (!returnObject[sortedWord])
			returnObject[sortedWord] = [word];
		else
			returnObject[sortedWord].push(word);
	}
	return Object.values(returnObject);
};