# Information retrieval systems

### Description:
Course on information retrieval systems

* 1: Show the number of comparisons and number of moves of Quicksort, InsertionSort, and BubbleSort on arrays of size 10, 20, 40
* 2: Implement Balanced merging, Oscillated Sorting and Polyphase Sorting algorithms*
* 3: The program that can search for a substring in a text using the Naive method, Boyer-Moore method, Knuth-Morris-Pratt method.


### Run 1, 3 point
```
go run cmd/main.go
```

### Run 2 point
```
cmake -B ./internal/lab2 -S ./internal/lab2
make -C ./internal/lab2/Makefile
./internal/lab2/main
```

### Generate test files
input - input path,<br> output - sorted result path,<br> records - number to generate and calc,<br> block - size of buffer
```
python3 internal/lab2/utils/generate_input.py --output test.txt --records 1000
```

*made in C++ due to the peculiarities of the language (I just wanted to make my life difficult.)