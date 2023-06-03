import matplotlib.pyplot as plt
import numpy as np

xsmall = []
small = []
medium = []
large = []
xlarge = []


def calc_speedup(arr):
    serial_time = arr[0]
    for elem in range(0,6):
        arr[elem] = serial_time/arr[elem]
    return arr

with open('xsmall.txt') as f:
    lines = f.readlines()
    for line in lines:
        elem = line.split()
        print(elem)
        sum = 0
        for e in elem:
            sum += float(e)
        xsmall.append(sum/5)

xsmall = calc_speedup(xsmall)
print(xsmall)

with open('small.txt') as f:
   lines = f.readlines()
   for line in lines:
        elem = line.split()
        sum = 0
        for e in elem:
            sum += float(e)
        small.append(sum/5)
        
small = calc_speedup(small)
print(small)

with open('medium.txt') as f:
   lines = f.readlines()
   for line in lines:
        elem = line.split()
        sum = 0
        for e in elem:
            sum += float(e)
        medium.append(sum/5)
        
medium = calc_speedup(medium)
print(medium)

with open('large.txt') as f:
   lines = f.readlines()
   for line in lines:
        elem = line.split()
        sum = 0
        for e in elem:
            sum += float(e)
        large.append(sum/5) 
        
large = calc_speedup(large)
print(large)

with open('xlarge.txt') as f:
   lines = f.readlines()
   for elem in lines:
        elem = elem.split()
        sum = 0
        for e in elem:
            sum += float(e)
        xlarge.append(sum/5)

print(xlarge)
xlarge = calc_speedup(xlarge)
print(xlarge)

plt.xlabel("Threads")
plt.ylabel("Speedup")

xpoints = [1, 2, 4, 6, 8, 12]

plt.title("SpeedUp Graph for Different Test Size")

plt.plot(xpoints, np.array(xsmall), label ='XSmall', marker = 'o')
plt.plot(xpoints, np.array(small), label ='Small', marker = 'o')
plt.plot(xpoints, np.array(medium), label ='Medium', marker = 'o')
plt.plot(xpoints, np.array(large), label ='Large', marker = 'o')
plt.plot(xpoints, np.array(xlarge), label ='XLarge', marker = 'o')

plt.legend()
plt.savefig('speedup.png')


# plt.savefig('xlarge.png')
