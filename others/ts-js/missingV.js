/**
 * 找到数组缺失的所有元素
 * @param nums 有序数组 e.g.[1, 2, 3, 5, 7]
 * @returns [missing data...] e.g.[4, 6]
 */
function getMissingV(nums) {
  const res = [];
  let initV = nums[0];

  nums.forEach((el) => {
    while (initV < el) {
      res.push(initV);
      initV++;
    }
    initV++;
  });

  return res;
}

function getMissingV1(nums) {
  let initV = nums[0];

  const res = nums.reduce((pre, cur) => {
    while (initV < cur) {
      pre.push(initV);
      initV++;
    }
    initV++;

    return pre;
  }, []);

  return res;
}

const demo = [1, 3, 5, 7, 10, 15, 19, 22];
console.log(getMissingV(demo));
console.log(getMissingV1(demo));

// Output:
// [ 2,  4,  6,  8,  9, 11, 12, 13, 14, 16, 17, 18, 20, 21]
