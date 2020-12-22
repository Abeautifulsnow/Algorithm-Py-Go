#!/usr/bin/env bash
function CreateFolder() {
    echo "Input Python or Go:"
    read program_name
    echo "Input the title, which will be the sub-folder name:"
    read title

    if [[ ${program_name} == "Python" || ${program_name} == "python" ]]
    then
        fileSuffix="py"
        titlePath=./leetcode/Python/code/${title}
    elif [[ ${program_name} == "Go" || ${program_name} == "go" ]] 
    then
        fileSuffix="go"
        titlePath=./leetcode/Go/code/${title}
    else
        echo "Input error, please try again."
        exit 1
    fi

    FileArray=(${title}.${fileSuffix} README.md)
    echo "Create ${title} folder: ${title_path}"
    for i in ${FileArray[@]}
    do
        i_path=${titlePath}/${i}
        echo "filepath: ${i_path}"
        mkdir -p "${titlePath}"
        if [[ ! -e ${i_path} ]]
        then
            touch "${i_path}"
        fi
    done
}

CreateFolder
echo "Task done!!!"