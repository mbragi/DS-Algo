function GetIndices(arr, target) {
 //verify the length of the array is greater than 2 n less than 10^4
 if (arr.length < 2 || arr.length > 104) {
  return [];
 }
 const store = [];

  for (let i = 0; i < arr.length; i++) {
   if (arr[i] + arr[i + 1] === target) {
     store.push(i, i + 1);
   }
  }
  return store;
}
console.log("Test 1", GetIndices([2,7,11,15], 9))
console.log("Test 2", GetIndices([3, 2, 4], 6))
console.log("Test 3", GetIndices([3,3], 6))
// console.log("mock 3", GetIndices(["s", 1], 6))
