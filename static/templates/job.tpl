#!/bin/bash
#SBATCH --job-name="{{.Name}}"
#SBATCH --partition={{.Partition}}
#SBATCH --ntasks=2
#SBATCH --gres=gpu:{{.GPU}}
#SBATCH --time=3-0:0
#SBATCH --chdir=./
#SBATCH --output=./log/cout.%j.txt
#SBATCH --error=./log/cerr.%j.txt

function timestamp {
        echo $(date +"%Y-%m-%d_%H-%M-%S")
}
# BASENAME=${SLURM_JOB_NAME//"/"/"_"}.$(timestamp)
BASENAME=$SLURM_JOB_NAME/$(timestamp)

function gpus {
        declare -a array=($(echo $CUDA_VISIBLE_DEVICES | tr "," " "))
        echo ${#array[@]}
}

echo
echo "============================ Messages from Goddess ============================"
echo " * Job starting from: "`date`
echo " * Job ID           : "$SLURM_JOBID
echo " * Job name         : "$SLURM_JOB_NAME
echo " * Job partition    : "$SLURM_JOB_PARTITION
echo " * Nodes            : "$SLURM_JOB_NUM_NODES
echo " * Cores            : "$SLURM_NTASKS
echo " * GPUS             : "$(gpus)
echo " * Git Commit ID    : "$(git rev-parse HEAD)
echo " * Working directory: "${SLURM_SUBMIT_DIR}
echo "==============================================================================="
echo

git diff
module load {{range $index, $element := .Modules}}{{$element}} {{end}}
{{.Command}}

echo
echo "============================ Messages from Goddess ============================"
echo " * Jab ended at     : "`date`
echo "==============================================================================="
echo

mkdir -p ./log/$BASENAME
cp ./log/cout.$SLURM_JOBID.txt ./log/$BASENAME/cout.txt
mv ./log/cout.$SLURM_JOBID.txt ./log/cout.txt
cp ./log/cerr.$SLURM_JOBID.txt ./log/$BASENAME/cerr.txt
mv ./log/cerr.$SLURM_JOBID.txt ./log/cerr.txt
