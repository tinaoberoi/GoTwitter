echo "Serial Run small"
for n in {1..5}; 
do
    echo "Running serial run :: $n"
    output=$(go run benchmark.go s small)
    echo -n "$output " >> small.txt
done

echo ""  >> small.txt

echo "Parallel Run small"

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 2"
    output=$(go run benchmark.go p small 2)
    echo -n "$output " >> small.txt
done

echo ""  >> small.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 4"
    output=$(go run benchmark.go p small 4)
    echo -n "$output " >> small.txt 
done

echo ""  >> small.txt

for n in {1..5};
do
    echo "Running parallel run :: $n threads :: 6"
    output=$(go run benchmark.go p small 6)
    echo -n "$output " >> small.txt  
done

echo ""  >> small.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 8"
    output=$(go run benchmark.go p small 8)
    echo -n "$output " >> small.txt  
done

echo ""  >> small.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 12"
    output=$(go run benchmark.go p small 12)
    echo -n "$output " >> small.txt
done

echo ""  >> small.txt