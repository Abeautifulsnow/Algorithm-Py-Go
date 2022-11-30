/**
 * 检查数据连续性
 * @param nums e.g.[1, 2, 3, 5, 7]
 * @returns e.g.['1-3', 5, 7]
 */
function checkArrContinuity(nums) {
  let start = 0;
  const newArr = [];

  nums.forEach((item, index) => {
    // [1, 2, 3, 5, 7] => 1
    // item不等于前一个值，等于后一个值，那起始值就为item
    if (item - 1 !== nums[index - 1] && item + 1 === nums[index + 1]) {
      start = item;
    } else if (item - 1 === nums[index - 1] && item + 1 !== nums[index + 1]) {
      // [1, 2, 3, 5, 7] => 1, 2, 3 => 1 - 3
      // item等于前一个值，不等于后一个值，那终点值就为item
      newArr.push(start + "-" + item);
    } else if (item - 1 !== nums[index - 1] && item + 1 !== nums[index + 1]) {
      // [1, 2, 3, 5, 7] => 5, 7
      // item与前一个值和后一个值都不想等，属于单独的值
      newArr.push(item);
    }
  });

  return newArr;
}

const demo = [1, 3, 5, 7, 8, 9, 10, 11, 14, 15, 19, 20, 21, 22, 23, 24, 27, 29];
console.log(checkArrContinuity(demo));

// Output:
// [ 1, 3, 5, '7-11', '14-15', '19-24', 27, 29 ]
