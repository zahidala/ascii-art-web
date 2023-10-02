echo "Test 1:"
go run . "" | cat -e
echo "Test 2:"
go run . "\n" | cat -e
echo "Test 3:"
go run . "Hello\n" | cat -e
echo "Test 4:"
 go run . "hello" | cat -e
echo "Test 5:"
 go run . "HeLlO" | cat -e
echo "Test 6:"
 go run . "Hello There" | cat -e
echo "Test 7:"
 go run . "{Hello There}" | cat -e
echo "Test 8:"
 go run . "Hello\nThere" | cat -e
echo "Test 9:"
 go run . "Hello\n\nThere" | cat -e
