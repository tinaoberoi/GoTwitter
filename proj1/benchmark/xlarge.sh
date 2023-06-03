echo "Serial Run xlarge"
for n in {1..5}; 
do
    echo "Running serial run :: $n"
    output=$(go run benchmark.go s xlarge)
    echo -n "$output " >> xlarge.txt
done

echo ""  >> xlarge.txt

echo "Parallel Run xlarge"

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 2"
    output=$(go run benchmark.go p xlarge 2)
    echo -n "$output " >> xlarge.txt
done

echo ""  >> xlarge.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 4"
    output=$(go run benchmark.go p xlarge 4)
    echo -n "$output " >> xlarge.txt
done

echo ""  >> xlarge.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 6"
    output=$(go run benchmark.go p xlarge 6)
    echo -n "$output " >> xlarge.txt
done

echo ""  >> xlarge.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 8"
    output=$(go run benchmark.go p xlarge 8)
    echo -n "$output " >> xlarge.txt
done

echo ""  >> xlarge.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 12"
    output=$(go run benchmark.go p xlarge 12)
    echo -n "$output " >> xlarge.txt
done

echo ""  >> xlarge.txt

echo "_________________________\n"
