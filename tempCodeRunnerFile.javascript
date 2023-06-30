function birthday(s, d, m) {
  // Write your code here
  if (m === 1 && s.length === 1 && s[0] === d) return 1;
  else if (m > 1 && s.length > 1) {
    let count = 0;

    for (let i = 0; i <= s.length - m; i++) {
      let tempCount = s[i];
      for (let j = i + 1; j < i + m; j++) {
        tempCount += s[j];
      }
    console.log(tempCount)
      if (tempCount === d) {
        count += 1;
      }
    }

    return count;
  }

  return 0;
}

const result = birthday([2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1], 18, 7);

console.log(result);
