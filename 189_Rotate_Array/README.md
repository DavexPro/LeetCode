# Rotate Array (Easy)

## Link

[Rotate Array](https://leetcode.com/problems/rotate-array/)

## Tips

**You should do the `Rotation` in the list given, otherwise it can't pass.**

And if you want to exchange two variables, then you can try as follow:

1.
```
a, b = b, a
```
2.
```
a ^= b
b ^= a
a ^= b
```
