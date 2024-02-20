## Problem Description:
- Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:

    Input: nums = [1,1,1,4,4,2], k = 2
    Output: [1,2]
    Example 2:

Example 2:
    Input: nums = [1], k = 1
    Output: [1]


Constraints:
```
1 <= nums.length <= 105
-104 <= nums[i] <= 104
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.
```

# Solving the problem:

To solve this problem, I used the Bucket Sort Algorithm. The idea behind this algorithm is to store sets of frequent numbers in corresponding indices.

For example:

```[1,1,1]``` means 3 occurrences of number 1, so it will be stored in index 3.
```[4,4]``` means 2 occurrences of number 4, so it will be stored in index 2.
```[2]``` means 1 occurrence of number 2, so it will be stored in index 1.
The resulting list of lists will be:

the new has slice of slices will be: 

```
bucket = [[],[2],[4],[1],[],[],[]]
```

The final step involves looping through the buckets from the end. If the nested list contains numbers, we append them to a new list called result until its length is not greater than k.
```
result = [1,4]
```
