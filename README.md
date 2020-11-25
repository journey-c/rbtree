# rbtree

[Doc](https://journey-c.github.io/2020/10/22/red-black-tree/)

# use
```
func main() {
	tree := NewRbTree(cmp, false)

	keyList := []int{10, 3, 17, 7, 1, 0, 10, 12, 4, 5}

	for _, key := range keyList {
		tree.Insert(key, key*2)
	}
	for i := len(keyList) - 1; i >= 0; i-- {
		tree.DeleteByKey(keyList[i])
	}
}

```

# test
```
➜  rbtree git:(main) ✗ go test -v -count=1 
=== RUN   TestRbTree
=========== INSERT: [10 3 17 7 1 0 10 12 4 5] ===========
level:	1	2	3	4	5	6	7	8	...
			17
		12
			10
	10
				7
			5
				4
		3
			1
				0
=========== DELETE: 5 ===========
level:	1	2	3	4	5	6	7	8	...
			17
		12
			10
	10
			7
				4
		3
			1
				0
=========== DELETE: 4 ===========
level:	1	2	3	4	5	6	7	8	...
			17
		12
			10
	10
			7
		3
			1
				0
=========== DELETE: 12 ===========
level:	1	2	3	4	5	6	7	8	...
		17
			10
	10
			7
		3
			1
				0
=========== DELETE: 10 ===========
level:	1	2	3	4	5	6	7	8	...
	17
			7
		3
			1
				0
=========== DELETE: 0 ===========
level:	1	2	3	4	5	6	7	8	...
	17
			7
		3
			1
=========== DELETE: 1 ===========
level:	1	2	3	4	5	6	7	8	...
	17
			7
		3
=========== DELETE: 7 ===========
level:	1	2	3	4	5	6	7	8	...
	17
		3
=========== DELETE: 17 ===========
level:	1	2	3	4	5	6	7	8	...
	3
=========== DELETE: 3 ===========
level:	1	2	3	4	5	6	7	8	...
=========== DELETE: 10 ===========
level:	1	2	3	4	5	6	7	8	...
--- PASS: TestRbTree (0.00s)
PASS
ok  	rbtree	0.441s
```
