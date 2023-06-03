echo "Serial Run xsmall"
for n in {1..5}; 
do
    echo "Running serial run :: $n"
    output=$(go run benchmark.go s xsmall)
    echo -n "$output " >> xsmall.txt
done

echo "" >> xsmall.txt

echo "Parallel Run xsmall"

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 2"
    output=$(go run benchmark.go p xsmall 2)
    echo -n "$output " >> xsmall.txt
done

echo ""  >> xsmall.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 4"
    output=$(go run benchmark.go p xsmall 4)
    echo -n "$output " >> xsmall.txt
done

echo ""  >> xsmall.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 6"
    output=$(go run benchmark.go p xsmall 6)
    echo -n "$output " >> xsmall.txt
done

echo ""  >> xsmall.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 8"
    output=$(go run benchmark.go p xsmall 8)
    echo -n "$output " >> xsmall.txt
done

echo ""  >> xsmall.txt

for n in {1..5}; 
do
    echo "Running parallel run :: $n threads :: 12"
    output=$(go run benchmark.go p xsmall 12)
    echo -n "$output " >> xsmall.txt
done

echo ""  >> xsmall.txt

echo "_________________________\n"