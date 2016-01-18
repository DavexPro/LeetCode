# Subsets (Medium)

## Link

[Subsets](https://leetcode.com/problems/subsets/)

## Question

> Given a set of distinct integers, nums, return all possible subsets.
> 
> Note:
> 
> Elements in a subset must be in non-descending order.
> 
> The solution set must not contain duplicate subsets.
> 
> For example,
> 
> If nums = [1,2,3], a solution is:
> 
> ```
> [
>   [3],
>   [1],
>   [2],
>   [1,2,3],
>   [1,3],
>   [2,3],
>   [1,2],
>   []
> ]
> ```

## Tips

如果用 Python 写，有个trick函数。

`list(itertools.combinations(['a','b','c'],2))`
