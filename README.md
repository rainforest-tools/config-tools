# config-tools
## Overview
![CleanShot 2021-07-17 at 14 59 01](https://user-images.githubusercontent.com/55678466/126028880-fad8e5a9-5fec-47b1-a544-a46c5320a05e.png)

## Template
```bash
#!/bin/bash
#SBATCH --job-name="{{.Name}}"
#SBATCH --partition={{.Partition}}
#SBATCH --ntasks=2
#SBATCH --gres=gpu:{{.GPU}}
#SBATCH --time=3-0:0
#SBATCH --chdir=./
#SBATCH --output=./log/cout.{{.Name}}.$timestamp.txt
#SBATCH --error=./log/cerr.{{.Name}}.$timestamp.txt

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
```
