echo "Serial Run medium"
for n in {1..5}; 
do
    echo "Running serial run :: $n"
    output=$(go run benchmark.go s medium)
    echo -n "$output " >> medium.txt
done

echo ""  >> medium.txt

echo "Parallel Run medium"

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 2"
    output=$(go run benchmark.go p medium 2)
    echo -n "$output " >> medium.txt  
done

echo ""  >> medium.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 4"
    output=$(go run benchmark.go p medium 4)
    echo -n "$output " >> medium.txt  
done

echo ""  >> medium.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 6"
    output=$(go run benchmark.go p medium 6)
    echo -n "$output " >> medium.txt
done

echo ""  >> medium.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 8"
    output=$(go run benchmark.go p medium 8)
    echo -n "$output " >> medium.txt 
done

echo ""  >> medium.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 12"
    output=$(go run benchmark.go p medium 12)
    echo -n "$output " >> medium.txt
done

echo ""  >> medium.txt