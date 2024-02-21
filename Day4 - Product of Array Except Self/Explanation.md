## Problem Description:

- Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
  The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
  You must write an algorithm that runs in O(n) time and without using the division operation.

## Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

## Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Constraints:

```2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
```

# Solving the problem:

To resolve the issue, we adopt a systematic approach involving both prefix and suffix computations. Let's illustrate this with an example using the input array [1, 2, 3, 4].

- Prefix Computation:

  - Initialize the prefix variable to 1.
  - Traverse the array and at each step:
  - Multiply the current prefix with the current number.
  - Append the result to the res slice.
  - Update the prefix.
  
  Example:

    ```
    res = [1, 1], prefix = 1 * n = 1 * 2, prefix = 2
    res = [1, 1, 2], prefix = 6
    res = [1, 1, 2, 6], prefix = 24
    ```

- Suffix Computation:

  - Initialize the postfix variable to 1.
  - Traverse the array in reverse order and at each step:
  - Multiply the current postfix with the corresponding number from the res slice.
  - Update the res slice.
  - Update the postfix.

  Example:

   ```
   res = [1, 1, 2, 6], postfix = 1 * n = 1 * 4, prefix = 4
   res = [1, 1, 8, 6], postfix = 4 * 3, postfix = 12
   res = [1,12,8,6], postfix = 12 * 2, postfix = 24
   res = [24, 12, 8, 6], postfix = 24
    ```

By performing these steps, we efficiently compute both the prefix and suffix products, resulting in the optimized solution.

