#!/bin/bash
#
#SBATCH --mail-user=toberoi@cs.uchicago.edu
#SBATCH --mail-type=ALL
#SBATCH --job-name=proj1_benchmark 
#SBATCH --output=./slurm/out/%j.%N.stdout
#SBATCH --error=./slurm/out/%j.%N.stderr
#SBATCH --chdir=/home/toberoi/project-1-tinaoberoi/proj1/benchmark
#SBATCH --partition=debug 
#SBATCH --nodes=1
#SBATCH --ntasks=1
#SBATCH --cpus-per-task=16
#SBATCH --mem-per-cpu=900
#SBATCH --exclusive
#SBATCH --time=120:00


module load golang/1.16.2

./xsmall.sh

./small.sh

./medium.sh

./large.sh

./xlarge.sh

python3 plot.py