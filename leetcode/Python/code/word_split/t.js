const wordBreak = function (text, wordDict) {
  const len = text.length;
  const dp = new Array(len).fill(false);
  const map = {};
  wordDict.forEach((item) => (map[item] = true));

  for (let i = 0; i < len; i++) {
    let str = "";
    for (let j = i; j < len; j++) {
      str += text[j];
			if (map[str] && (i === 0 || dp[i - 1])) {
        dp[j] = [].concat(dp[i - 1] || [], str);
        if (j === len - 1) {
          return dp[len - 1];
        }
      }
    }
  }
  return null;
};

const res = wordBreak("appleddesktop", [
	"led",
	"a",
	"desk",
  "app",
  "apple",
  "top",
]);
console.log(res);
