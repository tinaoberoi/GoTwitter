echo "Serial Run large"
for n in {1..5}; 
do
    echo "Running serial run :: $n"
    output=$(go run benchmark.go s large)
    echo -n "$output " >> large.txt
done

echo ""  >> large.txt

echo "Parallel Run large"

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 2"
    output=$(go run benchmark.go p large 2)
    echo -n "$output " >> large.txt
done

echo ""  >> large.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 4"
    output=$(go run benchmark.go p large 4)
    echo -n "$output " >> large.txt
done

echo ""  >> large.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 6"
    output=$(go run benchmark.go p large 6)
    echo -n "$output " >> large.txt 
done

echo ""  >> large.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 8"
    output=$(go run benchmark.go p large 8)
    echo -n "$output " >> large.txt
done

echo ""  >> large.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 12"
    output=$(go run benchmark.go p large 12)
    echo -n "$output " >> large.txt
done

echo ""  >> large.txt

echo "_________________________\n"