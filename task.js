function GetIndices(arr, target) {
  //verify the length of the array is greater than 2 n less than 104
  if (arr.length < 2 || arr.length > 104) {
    return "verify the length of the array is greater than 2 n less than 104";
  }

  for (let i = 0; i < arr.length; i++) {
    if (arr[i] + arr[i + 1] === target) {
      return [i, i + 1];
    }
  }
}
console.log("Test 1", GetIndices([2, 7, 11, 15], 9));
console.log("Test 2", GetIndices([3, 2, 4], 6));
console.log("Test 3", GetIndices([3, 3], 6));
