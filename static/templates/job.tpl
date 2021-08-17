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
BASENAME={{.Name}}.$(timestamp)

echo
echo "============================ Messages from Goddess ============================"
echo " * Job starting from: "`date`
echo " * Job ID           : "$SLURM_JOBID
echo " * Job name         : "$SLURM_JOB_NAME
echo " * Job partition    : "$SLURM_JOB_PARTITION
echo " * Nodes            : "$SLURM_JOB_NUM_NODES
echo " * Cores            : "$SLURM_NTASKS
echo " * Working directory: "${SLURM_SUBMIT_DIR}
echo "==============================================================================="
echo

module load {{range $index, $element := .Modules}}{{$element}} {{end}}
{{.Command}}

echo
echo "============================ Messages from Goddess ============================"
echo " * Jab ended at     : "`date`
echo "==============================================================================="
echo

mv ./log/cout.$SLURM_JOBID.txt ./log/cout.$BASENAME.txt
mv ./log/cerr.$SLURM_JOBID.txt ./log/cerr.$BASENAME.txt
