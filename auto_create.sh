#!/usr/bin/env bash
function CreateFolder() {
    echo "Input Python or Go(Case-Insensitive):"
    read program_name
    echo "Input the title, which will be the sub-folder name:"
    read title

    case ${program_name} in
    [Pp][Yy][Tt][Hh][Oo][Nn])
        fileSuffix="py"
        titlePath=./leetcode/Python/code/${title}
        ;;

    [Gg][Oo])
        fileSuffix="go"
        titlePath=./leetcode/Go/code/${title}
        ;;
    *)
        echo "Input error, please try again."
        exit 1
        ;;
    esac

    FileArray=(${title}.${fileSuffix} README.md)
    echo "Create ${title} folder: ${title_path}"
    for i in ${FileArray[@]}; do
        i_path=${titlePath}/${i}
        echo "filepath: ${i_path}"
        mkdir -p "${titlePath}"
        if [[ ! -e ${i_path} ]]; then
            touch "${i_path}"
        fi
    done
}

CreateFolder
echo "Task done!!!"
